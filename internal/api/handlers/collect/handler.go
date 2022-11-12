package collect

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pedrolopesme/open-rba/internal/core/domains"
	"github.com/pedrolopesme/open-rba/internal/core/ports"
	"go.uber.org/zap"
)

type Handler struct {
	logger  zap.Logger
	service ports.Collector
}

func NewHandler(logger zap.Logger, service ports.Collector) *Handler {
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

	if err := h.service.Collect(payload); err != nil {
		h.logger.Error("Impossible collect data", zap.Error(err))
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(http.StatusOK)
}
