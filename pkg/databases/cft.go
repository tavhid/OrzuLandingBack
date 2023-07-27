package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

// We can have many types of sql.DB
func Cft(params Dependencies) *sql.DB {
	urlDB := fmt.Sprintf("%s:%s/%s", params.Config.CFT.Host, params.Config.CFT.Port, params.Config.CFT.ServiceName)
	params.Logger.Info(urlDB)
	db, err := sql.Open("godror", fmt.Sprintf(`user="%s", password="%s", connectString="%s"`, params.Config.CFT.Username, params.Config.CFT.Password, urlDB))
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(0)
	db.SetConnMaxLifetime(0)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	params.Logger.Info("Cft pong!")

	return db
}
