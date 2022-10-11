package evaluators

import "github.com/pedrolopesme/open-rba/internal/core/domains"

type RegionEvaluator struct{}

func (c RegionEvaluator) Evaluate(riskProfile domains.RiskProfile, userProfile domains.UserProfile, attempt domains.AuthenticationData) float32 {
	totalWeight := riskProfile.TotalWeight()
	regionWeight := userProfile.GetRegionPercentage(attempt.Region)

	if totalWeight == 0 {
		return 0.0
	}

	return regionWeight * float32(riskProfile.RegionFactor) / float32(totalWeight)
}
