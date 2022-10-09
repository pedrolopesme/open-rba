package domains

type (
	RiskClassification string

	Risk struct {
		Score          float32
		Classification RiskClassification
	}
)

const (
	HIGH   RiskClassification = "HIGH"
	MEDIUM RiskClassification = "MEDIUM"
	LOW    RiskClassification = "LOW"
)

func (r *Risk) AddScore(score float32) {
	r.Score += score
}
