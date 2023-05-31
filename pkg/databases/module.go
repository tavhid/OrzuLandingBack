package databases

import (
	"cl/pkg/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// RedisModule ...
var RedisModule = fx.Provide(NewRedisConn)

// PostgresModule ...
var PostgresModule = fx.Provide(NewPostgresConn)

// Dependencies ...
type Dependencies struct {
	fx.In

	Logger *logrus.Logger
	Config *config.Config
}

// NewRedisConn ...
func NewRedisConn(params Dependencies) *redis.Client {
	return Redis(params)
}

// NewPostgresConn ...
func NewPostgresConn(params Dependencies) *gorm.DB {
	return Postgres(params)
}
