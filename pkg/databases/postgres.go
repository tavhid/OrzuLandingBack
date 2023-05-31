package databases

import (
	"cl/internal/structs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres ...
func Postgres(params Dependencies) (pdb *gorm.DB) {
	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",

		params.Config.Postgres.Host,
		fmt.Sprint(params.Config.Postgres.Port),
		params.Config.Postgres.DatabaseName,
		params.Config.Postgres.Username,
		params.Config.Postgres.Password,
		params.Config.Postgres.SSLMode,
	)

	pdb, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  connStr,
				PreferSimpleProtocol: true,
			},
		),
		&gorm.Config{
			DisableAutomaticPing:                     true,
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		panic(err)
	}

	conn, err := pdb.DB()
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}
	pdb.AutoMigrate(&structs.Application{})
	pdb.AutoMigrate(&structs.Merchant{})
	params.Logger.Info("Postgres pong!")

	return
}
