module github.com/cherya/memezis

go 1.18

require (
	github.com/aws/aws-sdk-go v1.44.82
	github.com/azr/phash v0.2.0
	github.com/cherya/memezis/pkg/memezis v0.0.0-20220413080315-e1205989ab4b
	github.com/cherya/memezis/pkg/queue v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.9
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.6
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/sirupsen/logrus v1.9.0
	github.com/utrack/clay/v2 v2.4.9
	google.golang.org/grpc v1.48.0
)

require (
	github.com/azr/gift v1.1.2 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/peterbourgon/mergemap v0.0.1 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	golang.org/x/image v0.0.0-20220722155232-062f8c9fd539 // indirect
	golang.org/x/net v0.0.0-20220822230855-b0a4917ee28c // indirect
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
	golang.org/x/text v0.3.8 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/cherya/memezis/pkg/queue => ./pkg/queue

replace github.com/cherya/memezis/pkg/memezis => ./pkg/memezis
