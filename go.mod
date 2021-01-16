module github.com/sageflow/sageengine

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.3
	github.com/mitchellh/mapstructure v1.1.2
	github.com/sageflow/sagedb v0.0.0-20210113125433-7f36eda586cf
	github.com/sageflow/sageutils v0.0.0-20210105150335-9448053ab68b
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/sageflow/sagedb v0.0.0-20210108233746-64884d2812f4 => ../sagedb

replace github.com/sageflow/sageutils v0.0.0-20210105150335-9448053ab68b => ../sageutils
