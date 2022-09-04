package domains

type (
	RiskClassification string

	Risk struct {
		Score          float64
		Classification RiskClassification
	}
)

const (
	HIGH   RiskClassification = "HIGH"
	MEDIUM RiskClassification = "MEDIUM"
	LOW    RiskClassification = "LOW"
)
