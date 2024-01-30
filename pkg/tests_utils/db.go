package tests_utils

import (
	"FinanceTODO/pkg/systems"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

func GetDB() *sqlx.DB {
	systems.SetupLogger()

	cfg := systems.AppConfig{
		Database: systems.DBConfig{
			Host: "localhost",
			Port: "5432",
			User: "postgres",
			Pass: "postgres",
			Name: "FinanceTodo",
		},
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Pass, cfg.Database.Name))
	if err != nil {
		log.Fatal().Err(err).Msg("ошибка при подключении к базе данных")
	}

	return db
}

func ClearTestDatabase(db *sqlx.DB) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal().Err(err).Msg("ошибка при удалении данных в конце теста")
	}
}
