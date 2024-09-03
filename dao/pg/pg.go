package pg

import (
	"fmt"
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var pgDB *sqlx.DB

func Init(cfg *settings.PostgresConfig) (err error) {
	//"user=bob password=secret host=1.2.3.4 port=5432 dbname=mydb sslmode=verify-full"

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	pgDB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return
	}

	if err = pgDB.Ping(); err != nil {
		zap.L().Error("ping postgres failed", zap.Error(err))
		return err
	}

	pgDB.SetMaxOpenConns(cfg.MaxOpenConns)
	pgDB.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}
