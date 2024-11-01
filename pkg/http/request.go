package http

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	idPath  = "id"
	intBase = 10
	bitSize = 64

	queryParamPage    = "page"
	queryParamLimit   = "limit"
	queryParamSortCol = "sortCol"
	queryParamSortDir = "sortDir"
	queryParamSearch  = "search"
	recordsPerPage    = 5
	sortColumn        = "start_date"
	sortDir           = 1
)

type QueryParams struct {
	urlValues url.Values
	Page      int64
	Offset    int64
	Limit     int64
	SortCol   string
	SortDir   int
	Search    string
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

func (q *QueryParams) GetSortDir() int {
	// TODO: Validate query params
	s := q.urlValues.Get(queryParamSortDir)

	sortDir, err := strconv.Atoi(s)

	if err != nil {
		return sortDir
	}

	if sortDir != 1 && sortDir != -1 {
		return sortDir
	}

	return sortDir
}

func (q *QueryParams) GetSearch() string {
	// TODO: Validate query params
	searchText := q.urlValues.Get(queryParamSearch)

	if searchText == "" {
		return searchText
	}

	return "%" + searchText + "%"
}

func ParseResourceId(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue(idPath), intBase, bitSize)
}
