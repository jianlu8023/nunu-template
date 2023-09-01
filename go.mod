module nunu-template

go 1.16

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/go-co-op/gocron v1.28.2
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/pkg/errors v0.9.1
	github.com/sony/sonyflake v1.1.0
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.9.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.1
)

require (
	github.com/go-redis/redis/v8 v8.11.5
	go.uber.org/atomic v1.11.0 // indirect
)

replace (
	github.com/gin-gonic/gin v1.9.1 => github.com/gin-gonic/gin v1.8.2
	go.uber.org/atomic v1.11.0 => go.uber.org/atomic v1.9.0
	golang.org/x/sys => golang.org/x/sys v0.0.0-20201204225414-ed752295db88
)
