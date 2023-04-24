package request

type TagsReq struct {
	Name        string `validate:"required, min-1, max=100" json:"name"`
	Description string `validate:"required, min-1, max=255" json:"description"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}
