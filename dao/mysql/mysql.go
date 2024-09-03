package mysql

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	//"user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.user")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	dbname := viper.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}

	maxOpenConns := viper.GetInt("mysql.max_open_conns")
	maxIdleConns := viper.GetInt("mysql.max_idle_conns")

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	return
}
func Close() {
	_ = db.Close()
}
