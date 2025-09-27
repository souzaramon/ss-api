package titles

type Title struct {
	Id       string `json:"id" example:"28487cdf-6694-4efb-9916-c0d5f596ed06"`
	IdAuthor string `json:"id_author" example:"28487cdf-6694-4efb-9916-c0d5f596ed06"`
	Name     string `json:"name" example:"Lord of Mysteries"`
}

type CreateTitleDto struct {
	Name     string `json:"name" binding:"required,min=2" example:"Lord of Mysteries"`
	IdAuthor string `json:"id_author" example:"28487cdf-6694-4efb-9916-c0d5f596ed06"`
}

type UpdateTitleDto struct {
	Name     string `json:"name" binding:"min=2" example:"Circle of Inevitably"`
	IdAuthor string `json:"id_author" example:"28487cdf-6694-4efb-9916-c0d5f596ed06"`
}
