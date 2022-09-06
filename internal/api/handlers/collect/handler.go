package collect

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/open-rba/internal/core/domains"
	"github.com/pedrolopesme/open-rba/internal/core/ports"
	"go.uber.org/zap"
)

type Handler struct {
	logger  zap.Logger
	service ports.Collector
}

func (h Handler) Handle(ctx echo.Context) error {
	var payload domains.AuthenticationData
	if err := ctx.Bind(payload); err != nil {
		h.logger.Error("Impossible to read request payload", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.Collect(payload); err != nil {
		h.logger.Error("Impossible collect data", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
