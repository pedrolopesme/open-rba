package risk

import (
	"github.com/pedrolopesme/open-rba/internal/core/domains"
	"github.com/pedrolopesme/open-rba/internal/core/ports"
	"github.com/pedrolopesme/open-rba/internal/core/services/risk/evaluators"
	"go.uber.org/zap"
)

type RiskService struct {
	riskProfile      domains.RiskProfile
	persistence      ports.Repository
	logger           zap.Logger
	countryEvaluator evaluators.Evaluator
}

func NewRiskService(logger zap.Logger, persistence ports.Repository) *RiskService {
	return &RiskService{
		riskProfile:      *domains.NewDefaultRisckWeight(),
		persistence:      persistence,
		logger:           logger,
		countryEvaluator: evaluators.CountryEvaluator{},
	}
}

func (r *RiskService) Evaluate(userProfile domains.UserProfile, attempt domains.AuthenticationData) (domains.Risk, error) {
	risk := domains.Risk{}
	risk.AddScore(r.countryEvaluator.Evaluate(r.riskProfile, userProfile, attempt))
	risk = calculateClassification(risk)
	return risk, nil
}

func calculateClassification(risk domains.Risk) domains.Risk {
	if risk.Score >= 80 {
		risk.Classification = domains.LOW
	} else if risk.Score >= 50 {
		risk.Classification = domains.LOW
	} else {
		risk.Classification = domains.HIGH
	}

	return risk
}
