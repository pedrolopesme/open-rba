package domains

type RiskProfile struct {
	Name          string
	CountryFactor int
	RegionFactor  int
}

func NewDefaultRisckWeight() *RiskProfile {
	return &RiskProfile{
		Name:          "DEFAULT",
		CountryFactor: 10,
		RegionFactor:  1,
	}
}

func (rp *RiskProfile) TotalWeight() int {
	return rp.CountryFactor + rp.RegionFactor
}
