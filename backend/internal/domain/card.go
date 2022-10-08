package domain

type Card struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Body        string  `json:"body"`
	Price       float32 `json:"price"`
	Thumbnail   string  `json:"thumbnail"`
}
