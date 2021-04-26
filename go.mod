module github.com/gigamono/gigamono-workflow-engine

go 1.15

require (
	github.com/gigamono/gigamono v0.0.0-20210426004714-d66ce0dbccc4
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/gigamono/gigamono v0.0.0-20210426004714-d66ce0dbccc4 => ../gigamono
