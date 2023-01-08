package book

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
