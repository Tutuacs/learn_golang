package todo

type Todo struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Name   string `json:"name"`
	Done   bool   `json:"done"`
}
