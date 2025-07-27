package dto

type Paging struct {
	PageNumber  int16 `form:"pageNumber" binding:"required,min=1"`
	RowsPerPage int16 `form:"rowsPerPage" binding:"required,min=1"`
}
