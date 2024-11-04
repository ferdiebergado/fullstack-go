package request

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	idPath  = "id"
	intBase = 10
	bitSize = 64

	queryParamPage      = "page"
	queryParamLimit     = "limit"
	queryParamSortCol   = "sortCol"
	queryParamSortDir   = "sortDir"
	queryParamSearch    = "search"
	queryParamSearchCol = "searchCol"

	recordsPerPage = 5
	sortColumn     = "updated_at"
	sortDir        = "DESC"
)

type QueryParams struct {
	urlValues url.Values
	Page      int64
	Offset    int64
	Limit     int64
	SortCol   string
	SortDir   string
	Search    string
	SearchCol string
}

func NewQueryParams(r *http.Request) *QueryParams {

	q := &QueryParams{
		urlValues: r.URL.Query(),
	}

	q.Page = q.GetPage()
	q.Limit = q.GetLimit()
	q.SortCol = q.GetSortCol()
	q.SortDir = q.GetSortDir()
	q.Search = q.GetSearch()
	q.SearchCol = q.GetSearchCol()
	q.Offset = (q.Page - 1) * q.Limit

	return q
}

func (q *QueryParams) GetPage() int64 {
	// TODO: Validate query params
	s := q.urlValues.Get(queryParamPage)

	page, err := strconv.ParseInt(s, intBase, bitSize)

	if err != nil || page < 1 {
		return 1
	}

	return page
}

func (q *QueryParams) GetLimit() int64 {
	// TODO: Validate query params
	s := q.urlValues.Get(queryParamLimit)

	limit, err := strconv.ParseInt(s, 0, 64)

	if err != nil || limit < 1 {
		return recordsPerPage
	}

	return limit
}

func (q *QueryParams) GetSortCol() string {
	// TODO: Validate query params
	sortCol := q.urlValues.Get(queryParamSortCol)

	if sortCol == "" {
		return sortColumn
	}

	return sortCol
}

func (q *QueryParams) GetSortDir() string {
	// TODO: Validate query params
	s := q.urlValues.Get(queryParamSortDir)

	if s == "1" {
		return "ASC"
	} else if s == "-1" {
		return "DESC"
	}

	return sortDir
}

func (q *QueryParams) GetSearch() string {
	// TODO: Validate query params
	return q.urlValues.Get(queryParamSearch)
}

func (q *QueryParams) GetSearchCol() string {
	// TODO: Validate query params
	return q.urlValues.Get(queryParamSearchCol)
}

func ParseResourceId(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue(idPath), intBase, bitSize)
}
