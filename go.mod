module github.com/sageflow/sageengine

go 1.15

require (
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/sageflow/sageflow v0.0.0-20210213140648-cdbbdd211183
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/sageflow/sageflow v0.0.0-20210213140648-cdbbdd211183 => ../sageflow
