module github.com/gigamono/gigamono-automation-engine

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/gigamono/gigamono v0.0.0
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-pg/pg/v10 v10.9.1 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/soheilhy/cmux v0.1.4
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.1.0 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/examples v0.0.0-20210430185331-b418de839e73 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/gigamono/gigamono v0.0.0 => ../gigamono
