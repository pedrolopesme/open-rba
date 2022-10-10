package collector

import (
	"github.com/pedrolopesme/open-rba/internal/core/domains"
	"github.com/pedrolopesme/open-rba/internal/core/ports"
	"go.uber.org/zap"
)

type CollectorService struct {
	persistence ports.Repository
	logger      zap.Logger
}

func NewCollectorService(logger zap.Logger, persistence ports.Repository) *CollectorService {
	return &CollectorService{
		persistence: persistence,
		logger:      logger,
	}
}

func (c CollectorService) Collect(data domains.AuthenticationData) error {
	c.logger.Info("Collecting authentication data",
		zap.String("user-id", data.UserID),
	)

	if err := validate(&data); err != nil {
		c.logger.Error("Invalid authentication data",
			zap.String("user-id", data.UserID),
			zap.Error(err),
		)
		return err
	}

	if err := c.persistence.Insert(data.UserID, data); err != nil {
		c.logger.Error("Cannot save authentication",
			zap.String("user-id", data.UserID),
			zap.Error(err),
		)
	}

	return nil
}
