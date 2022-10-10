package evaluators

import "github.com/pedrolopesme/open-rba/internal/core/domains"

type Evaluator interface {
	Evaluate(riskProfile domains.RiskProfile, userProfile domains.UserProfile, attempt domains.AuthenticationData) float32
}
