package app

import (
	"ragnarok/config"
	"ragnarok/internal/pkg/util"
)

func InitApp() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	_, err = util.ConnectDB(cfg)
	if err != nil {
		return err
	}

	err = StartServer()
	if err != nil {
		return err
	}

	return nil
}
