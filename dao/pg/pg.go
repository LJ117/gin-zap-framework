package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var pgDB *sqlx.DB

func Init() (err error) {
	//"user=bob password=secret host=1.2.3.4 port=5432 dbname=mydb sslmode=verify-full"
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.user")
	host := viper.GetString("postgres.host")
	port := viper.GetInt("postgres.port")
	dbname := viper.GetString("postgres.dbname")
	sslmode := viper.GetString("postgres.sslmode")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", user, password, host, port, dbname, sslmode)

	pgDB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return
	}

	maxOpenConns := viper.GetInt("postgres.max_open_conns")
	maxIdleConns := viper.GetInt("postgres.max_idle_Conns")

	pgDB.SetMaxOpenConns(maxOpenConns)
	pgDB.SetMaxIdleConns(maxIdleConns)
	return
}
