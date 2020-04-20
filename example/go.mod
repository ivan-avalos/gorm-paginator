module example

require (
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/ivan-avalos/gorm-paginator/pagination v0.0.0-20200420171620-44b491c88153
	github.com/jinzhu/gorm v1.9.2
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/ugorji/go/codec v0.0.0-20181209151446-772ced7fd4c2 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace github.com/ivan-avalos/gorm-pagination => ../

go 1.13
