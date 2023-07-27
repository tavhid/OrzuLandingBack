package databases

import (
	"cl/pkg/config"
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// RedisModule ...
var RedisModule = fx.Provide(NewRedisConn)

// PostgresModule ...
var PostgresModule = fx.Provide(NewPostgresConn)

// CFTModule ...
var CFTModule = fx.Provide(NewCFTConn)

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

// NewCFTConn ...
func NewCFTConn(params Dependencies) *sql.DB {
	return Cft(params)
}
