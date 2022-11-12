package evaluator

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func (h Handler) Handle(ctx *fiber.Ctx) error {
	var payload domains.AuthenticationData
	if err := ctx.BodyParser(&payload); err != nil {
		h.logger.Error("Impossible to read request payload", zap.Error(err))
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	mockedUserStatistics := domains.UserProfile{
		UserID: "test",
		Countries: []domains.UserProfileStatistics{
			{Entry: "BRAZIL", Total: 10},
		},
		Regions: []domains.UserProfileStatistics{
			{Entry: "RJ", Total: 10},
		},
	}

	if risk, err := h.service.Evaluate(mockedUserStatistics, payload); err != nil {
		h.logger.Error("Impossible to evaluate", zap.Error(err))
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	} else {
		return ctx.Status(http.StatusOK).JSON(risk)
	}
}
