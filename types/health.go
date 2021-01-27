package types

// Health health object
// swagger:response foobarResponse
type Health struct {
	StartTime string `json:"startTime"`
	UpTime    string `json:"uptime"`
}
