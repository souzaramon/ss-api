package authors

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateAuthorDto struct {
	Name string `json:"name" binding:"required,min=2"`
}
