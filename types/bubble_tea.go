package types

// BubbleTeaRequest represent a bubble tea drink
type BubbleTeaRequest struct {
	Size    string   `json:"size"`
	Name    string   `json:"name"`
	Flavors []string `json:"flavors"`
}

// Flavor represent a flavor
// type Flavor struct {
// 	Flavor string `json:"flavor"`
// }

// BubbleTeaResponse represents a BubbleTea request response
type BubbleTeaResponse struct {
	Success  bool   `json:"success"`
	StatusCd int    `json:"statusCd"`
	Message  string `json:"message"`
}
