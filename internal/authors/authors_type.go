package authors

type Author struct {
	Id   string `json:"id" example:"28487cdf-6694-4efb-9916-c0d5f596ed06"`
	Name string `json:"name" example:"Jon Snow"`
	Bio  string `json:"bio" example:"998th Lord Commander of the Night's Watch"`
}

type CreateAuthorDto struct {
	Name string `json:"name" binding:"required,min=2" example:"Jon Snow"`
	Bio  string `json:"bio" example:"998th Lord Commander of the Night's Watch"`
}

type UpdateAuthorDto struct {
	Name string `json:"name" binding:"min=2" example:"Aegon Targaryen"`
	Bio  string `json:"bio" example:"King Aegon of Houses Targaryen and Stark"`
}
