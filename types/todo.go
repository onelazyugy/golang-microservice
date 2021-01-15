package types

// Todo is a type to hold a todo item
type Todo struct {
	ItemID      int    `json:"itemId"`
	ItemName    string `json:"itemName"`
	Description string `json:"description"`
}
