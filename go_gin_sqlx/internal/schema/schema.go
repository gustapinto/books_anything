package schema

type NewAuthorPayload struct {
	Name string `json:"name"`
}

type NewBookPayload struct {
	ISBN     string `json:"isbn"`
	Name     string `json:"name"`
	AuthorId uint   `json:"author_id"`
}
