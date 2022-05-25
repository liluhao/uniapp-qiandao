package viewmodel

type CreateClassRequest struct {
	ClassName string `json:"class_name" form:"class_name"`
	Capacity  int    `json:"capacity" form:"capacity"`
}

type ClassInfo struct {
	ClassId   string `json:"class_id"`
	ClassName string `json:"class_name"`
}

type GetClassListResponse struct {
	TotalCount uint64      `json:"totalCount"`
	Class      []ClassInfo `json:"class"`
}
