package models

const (
	pageDefault     int = 1
	pageSizeDefault int = 20
)

type Meta struct {
	PageSize     int  `json:"pageSize"`
	Offset       int  `json:"offSet"`
	TotalRecords uint `json:"totalRecords"`
}

func NewMeta(pageSize int, offset int, totalRecords uint) *Meta {
	return &Meta{
		PageSize:     pageSize,
		Offset:       offset,
		TotalRecords: totalRecords,
	}
}
