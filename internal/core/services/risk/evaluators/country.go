package evaluators

import "github.com/pedrolopesme/open-rba/internal/core/domains"

type CountryEvaluator struct{}

func (c CountryEvaluator) Evaluate(riskProfile domains.RiskProfile, userProfile domains.UserProfile, attempt domains.AuthenticationData) float32 {
	totalWeight := riskProfile.TotalWeight()
	countryWeight := userProfile.GetCountryPercentage(attempt.Country)

	if totalWeight == 0 {
		return 0.0
	}

	return countryWeight * float32(riskProfile.CountryFactor) / float32(totalWeight)
}
