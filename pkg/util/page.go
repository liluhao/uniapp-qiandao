package util

type PageRequest struct {
	Offset int    `json:"offset" form:"offset"`
	Limit  int    `json:"limit" form:"limit"`
	Logo   string `json:"logo" form:"logo" binding:"required"`
}
