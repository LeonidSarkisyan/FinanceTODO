package systems

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port    string
	Address string
	Debug   bool

	Database DBConfig
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func GetAndSetupConfig() (*AppConfig, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return getConfig()
}

func getConfig() (*AppConfig, error) {
	cfg := AppConfig{
		Address: viper.GetString("address"),
		Port:    viper.GetString("port"),
		Debug:   viper.GetBool("debug"),
	}

	if cfg.Debug {
		cfg.Database = DBConfig{
			Host: viper.GetString("test_db.host"),
			Port: viper.GetString("test_db.port"),
			User: viper.GetString("test_db.user"),
			Pass: viper.GetString("test_db.pass"),
			Name: viper.GetString("test_db.name"),
		}
	} else {
		cfg.Database = DBConfig{
			Host: viper.GetString("prod_db.host"),
			Port: viper.GetString("prod_db.port"),
			User: viper.GetString("prod_db.user"),
			Pass: viper.GetString("prod_db.pass"),
			Name: viper.GetString("prod_db.name"),
		}
	}

	return &cfg, nil
}
