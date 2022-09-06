package api

import (
	"github.com/pedrolopesme/open-rba/internal/core/ports"
	"github.com/pedrolopesme/open-rba/internal/core/services/collector"
	"github.com/pedrolopesme/open-rba/internal/persistence/redisclient"
	"go.uber.org/zap"
)

type API struct {
	logger           zap.Logger
	collectorService ports.Collector
}

func NewAPI() *API {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	redis := redisclient.NewRedisClient()

	return &API{
		logger:           *logger,
		collectorService: collector.NewCollectorService(*logger, redis),
	}
}
