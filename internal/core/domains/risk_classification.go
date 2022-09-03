package domains

type (
	RiskClassification string

	Risk struct {
		Score         float64
		Classifiction RiskClassification
	}
)

const (
	HIGH   RiskClassification = "HIGH"
	MEDIUM RiskClassification = "MEDIUM"
	LOW    RiskClassification = "LOW"
)
