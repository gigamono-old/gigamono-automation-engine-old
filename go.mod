module github.com/sageflow/sageengine

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.3
	github.com/mitchellh/mapstructure v1.4.1
	github.com/sageflow/sagedb v0.0.0-20210117193554-834c3eaadd58
	github.com/sageflow/sageflow v0.0.0-20210108233356-e663f3625227
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/sageflow/sageflow v0.0.0-20210108233356-e663f3625227 => ../sageflow
