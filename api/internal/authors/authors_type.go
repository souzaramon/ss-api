package authors

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type CreateAuthorDto struct {
	Name string `json:"name" binding:"required,min=2"`
	Bio  string `json:"bio"`
}

type UpdateAuthorDto struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
