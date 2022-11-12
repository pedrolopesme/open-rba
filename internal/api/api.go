package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pedrolopesme/open-rba/internal/api/handlers/collect"
	"github.com/pedrolopesme/open-rba/internal/api/handlers/evaluator"
	"github.com/pedrolopesme/open-rba/internal/core/services/collector"
	"github.com/pedrolopesme/open-rba/internal/core/services/risk"
	"github.com/pedrolopesme/open-rba/internal/persistence/memory"
	"go.uber.org/zap"
)

const (
	PORT = 8080
)

type API struct {
	logger zap.Logger
	app    *fiber.App
}

func NewAPI() *API {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app := fiber.New()
	app.Use(cors.New())

	return &API{
		logger: *logger,
		app:    app,
	}
}

func (a *API) Setup() {
	app := a.app

	persistence := memory.NewMemory()
	collectorService := collector.NewCollectorService(a.logger, persistence)
	evaluatorService := risk.NewRiskService(a.logger, persistence)

	app.Post("/collect", collect.NewHandler(a.logger, collectorService).Handle)
	app.Post("/evaluate", evaluator.NewHandler(a.logger, evaluatorService).Handle)
}

func (a *API) Run() {
	a.logger.Info(fmt.Sprintf("Starting up application at port %d", PORT))
	app := a.app
	app.Listen(fmt.Sprintf(":%d", PORT))
}
