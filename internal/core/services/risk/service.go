package risk

import (
	"github.com/pedrolopesme/open-rba/internal/core/domains"
	"github.com/pedrolopesme/open-rba/internal/core/ports"
	"go.uber.org/zap"
)

type RiskService struct {
	risk_profile domains.RiskProfile
	persistence  ports.Repository
	logger       zap.Logger
}

func NewRiskService(logger zap.Logger, persistence ports.Repository) *RiskService {
	return &RiskService{
		risk_profile: *domains.NewDefaultRisckWeight(),
		persistence:  persistence,
		logger:       logger,
	}
}

func (r *RiskService) Evaluate(userProfile domains.UserProfile, attempt domains.AuthenticationData) (domains.Risk, error) {
	risk := domains.Risk{}
	risk.AddScore(r.evaluateCountry(userProfile, attempt.Country))

	if risk.Score >= 80 {
		risk.Classification = domains.LOW
	} else if risk.Score >= 50 {
		risk.Classification = domains.LOW
	} else {
		risk.Classification = domains.HIGH
	}

	return risk, nil
}

func (r *RiskService) evaluateCountry(userProfile domains.UserProfile, attempt string) float32 {
	totalWeight := r.risk_profile.TotalWeight()
	countryWeight := userProfile.GetCountryPercentage(attempt)

	if totalWeight == 0 {
		return 0.0
	}

	return countryWeight * 100 / float32(totalWeight)
}
