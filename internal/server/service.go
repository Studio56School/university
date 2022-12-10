package server

import (
	"github.com/Studio56School/university/internal/config"
	"github.com/Studio56School/university/internal/logger"
	"github.com/Studio56School/university/internal/service"
	"github.com/Studio56School/university/internal/storage"
)

// Все сервисы
type ServerServices struct {
	Srv service.IService
}

// Создать все сервисы
func newServices(conf *config.Config, logger *logger.Logger, universityStorage *storage.Repo) (*ServerServices, error) {
	// Создаем репозитории

	// Создаем сервисы
	srv := service.NewService(conf, logger, universityStorage)

	return &ServerServices{
		Srv: srv,
	}, nil
}