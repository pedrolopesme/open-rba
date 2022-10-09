package evaluator

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/open-rba/internal/core/domains"
	"github.com/pedrolopesme/open-rba/internal/core/ports"
	"go.uber.org/zap"
)

type Handler struct {
	logger  zap.Logger
	service ports.RiskEvaluator
}

func NewHandler(logger zap.Logger, service ports.RiskEvaluator) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h Handler) Handle(ctx echo.Context) error {
	var payload domains.AuthenticationData
	if err := ctx.Bind(&payload); err != nil {
		h.logger.Error("Impossible to read request payload", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	mockedUserStatistics := domains.UserProfile{
		UserID: "test",
		Countries: []domains.UserProfileStatistics{
			{Entry: "BRAZIL", Total: 10},
		},
	}

	if risk, err := h.service.Evaluate(mockedUserStatistics, payload); err != nil {
		h.logger.Error("Impossible to evaluate", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		return ctx.JSON(http.StatusOK, risk)
	}
}
