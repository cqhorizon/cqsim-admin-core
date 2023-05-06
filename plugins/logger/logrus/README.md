# logrus

[logrus](https://github.com/sirupsen/logrus) logger implementation for __cqsim-resource-service__ [meta logger](https://github.com/cqsim-resource-service-team/cqsim-resource-service-core/tree/master/logger).

## Usage

```go
import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/cqsim-resource-service-team/cqsim-resource-service-core/logger"
)

func ExampleWithOutput() {
	logger.DefaultLogger = NewLogger(logger.WithOutput(os.Stdout))
	logger.Infof("testing: %s", "Infof")
}

func ExampleWithLogger() {
	l := logrus.New() // *logrus.Logger
	logger.DefaultLogger = NewLogger(WithLogger(l))
	logger.Infof("testing: %s", "Infof")
}
```

