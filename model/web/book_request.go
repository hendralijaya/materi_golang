package web

type BookCreateRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type BookUpdateRequest struct {
	Id     uint64
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
