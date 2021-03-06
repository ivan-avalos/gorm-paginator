package pagination

import (
	"math"

	"github.com/jinzhu/gorm"
)

// Param 分页参数
type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

// Paginator 分页返回
type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	FirstPage   int         `json:"first_page"`
	LastPage    int         `json:"last_page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
	Error       error       `json:"-"`
}

// Paging 分页
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	var err error
	go countRecords(db, result, done, &err, &count)
	if err != nil {
		paginator.Error = err
		return &paginator
	}

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	err = db.Limit(p.Limit).Offset(offset).Find(result).Error
	<-done
	if err != nil {
		paginator.Error = err
		return &paginator
	}

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))
	paginator.FirstPage = 1

	if paginator.TotalPage > 0 {
		paginator.LastPage = paginator.TotalPage
	} else {
		paginator.LastPage = paginator.FirstPage
	}

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, err *error, count *int) {
	*err = db.Model(anyType).Count(count).Error
	done <- true
}
