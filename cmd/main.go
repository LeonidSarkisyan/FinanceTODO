package main

import (
	"FinanceTODO/internal/handlers"
	"FinanceTODO/internal/repositories"
	"FinanceTODO/internal/services"
	"FinanceTODO/pkg/server"
	"FinanceTODO/pkg/systems"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func main() {
	systems.SetupLogger()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msgf("ошибка при подключении env - файла: %s", err.Error())
	}

	cfg, err := systems.GetAndSetupConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("ошибка при конфигурировании конфига")
	}
	log.Info().Msg("конфиг загружен")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("ошибка при подключении к базе данных")
	}

	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)
	server_ := new(server.Server)

	go func() {
		if err = server_.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			log.Fatal().Err(err).Msg("ошибка при запуске сервера")
		}
	}()

	log.Printf("FinanceTodo started on port %s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	log.Print("Server shutdown")

	if err = server_.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("ошибка при остановке сервера")
	}

	if err = db.Close(); err != nil {
		log.Error().Err(err).Msg("ошибка при закрытии соединения с БД")
	}
}

func executeInitScript(db *sqlx.DB, script string) error {
	_, err := db.Exec(script)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении SQL-скрипта: %v", err)
	}
	return nil
}
