package domains

type AuthenticationData struct {
	UserID                string `json:"user_id"`
	AuthenticationSuccess bool   `json:"authentication_success"`
	UserAgent             string `json:"user_agent"`
	IP                    string `json:"ip"`
	Country               string `json:"country"`
	Region                string `json:"region"`
	OS                    string `json:"os"`
	Browser               string `json:"browser"`
	DeviceType            string `json:"device_type"`
}
