module github.com/cherya/memezis

go 1.14

require (
	github.com/aws/aws-sdk-go v1.32.13
	github.com/azr/gift v1.1.2 // indirect
	github.com/azr/phash v0.0.0-20180422161411-60fe3536b55c
	github.com/cherya/memezis/pkg/memezis v0.0.0-00010101000000-000000000000
	github.com/cherya/memezis/pkg/queue v0.0.0-20200617154535-92e0758b2834
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.7.0
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/sirupsen/logrus v1.6.0
	github.com/utrack/clay/v2 v2.4.9
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/grpc/examples v0.0.0-20200630190442-3de8449f8555 // indirect
)

replace github.com/cherya/memezis/pkg/queue => ./pkg/queue

replace github.com/cherya/memezis/pkg/memezis => ./pkg/memezis
