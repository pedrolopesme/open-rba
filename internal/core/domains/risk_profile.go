package domains

type RiskProfile struct {
	Name          string
	CountryFactor int
}

func NewDefaultRisckWeight() *RiskProfile {
	return &RiskProfile{
		Name:          "DEFAULT",
		CountryFactor: 10,
	}
}

func (rp *RiskProfile) TotalWeight() int {
	return rp.CountryFactor
}
