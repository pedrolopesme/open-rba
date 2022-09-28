package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/open-rba/internal/api/handlers/collect"
	"github.com/pedrolopesme/open-rba/internal/core/services/collector"
	"github.com/pedrolopesme/open-rba/internal/persistence/memory"
	"go.uber.org/zap"
)

const (
	PORT = 8080
)

type API struct {
	logger zap.Logger
	echo   echo.Echo
}

func NewAPI() *API {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &API{
		logger: *logger,
		echo:   *echo.New(),
	}
}

func (a *API) Setup() {
	e := a.echo

	persistence := memory.NewMemory()
	collectorService := collector.NewCollectorService(a.logger, persistence)

	e.POST("/collect", collect.NewHandler(a.logger, collectorService).Handle)
}

func (a *API) Run() {
	a.logger.Info(fmt.Sprintf("Starting up application at port %d", PORT))
	a.echo.Logger.Fatal(a.echo.Start(fmt.Sprintf(":%d", PORT)))
}
