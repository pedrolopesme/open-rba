package domains

type UserProfile struct {
	UserID    string                  `json:"user_id"`
	Countries []UserProfileStatistics `json:"countries"`
}

type UserProfileStatistics struct {
	Entry string `json:"entry"`
	Total int    `json:"total"`
}

func (up *UserProfile) GetCountryPercentage(entry string) float32 {
	total := 0
	totalItem := 0
	for _, country := range up.Countries {
		total += country.Total

		if country.Entry == entry {
			totalItem += country.Total
		}
	}

	if total == 0 {
		return 0
	}

	return float32(totalItem) * 100 / float32(total)
}
