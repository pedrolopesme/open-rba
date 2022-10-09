package ports

import "github.com/pedrolopesme/open-rba/internal/core/domains"

type Collector interface {
	Collect(data domains.AuthenticationData) error
}

type RiskEvaluator interface {
	Evaluate(userProfile domains.UserProfile, attempt domains.AuthenticationData) (domains.Risk, error)
}
