package pagination

type Pagination struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
}
