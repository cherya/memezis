module github.com/cherya/memezis

go 1.13

require (
	github.com/aws/aws-sdk-go v1.30.24
	github.com/cherya/memezis/pkg/queue v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gocraft/work v0.5.1
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/mux v1.7.3
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/joho/godotenv v1.3.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lib/pq v1.5.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/pressly/goose v2.6.0+incompatible // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	google.golang.org/appengine v1.6.5
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/cherya/memezis/pkg/queue => ./pkg/queue
