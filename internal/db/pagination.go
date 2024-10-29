package db

type QueryParams struct {
	Offset  int64
	Limit   int64
	SortCol string
	SortDir int
	Search  string
}
