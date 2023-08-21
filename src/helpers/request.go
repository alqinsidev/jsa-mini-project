package helpers

import (
	"alqinsidev/jsa-mini-project/aduan/domain"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRequestQuery(c *fiber.Ctx) *domain.RequestQuery {
	startDate := GetThisMonthStartDate()
	startDateQuery := c.Query("start_date")
	if startDateQuery != "" {
		startDate = startDateQuery
	}

	endDate := GetThisMonthEndDate()
	endDateQuery := c.Query("end_date")
	if endDateQuery != "" {
		endDate = endDateQuery
	}

	var currentPage int64 = 1
	currentPageQuery := c.Query("page")
	if currentPageQuery != "" {
		currentPage, _ = strconv.ParseInt(currentPageQuery, 10, 64)
	}

	var pageSize int64 = 10
	pageSizeQuery := c.Query("limit")
	if pageSizeQuery != "" {
		pageSize, _ = strconv.ParseInt(pageSizeQuery, 10, 64)
	}

	sortBy := "created_at"
	sortByQuery := c.Query("sort_by")
	if sortByQuery != "" {
		sortBy = sortByQuery
	}

	sortDirection := "DESC"
	sortDirectionQuery := c.Query("sort_dir")
	if sortDirectionQuery != "" {
		if sortDirectionQuery == "ASC" || sortDirectionQuery == "DESC" {
			sortDirection = sortDirectionQuery
		}
	}

	Filters := make(map[string]interface{})

	statusQuery := c.Query("status")
	if statusQuery != "" {
		Filters["status"], _ = strconv.ParseInt(statusQuery, 10, 64)
	}

	Keyword := ""
	KeywordQuery := c.Query("q")
	if KeywordQuery != "" {
		Keyword = strings.ToLower(KeywordQuery)
	}

	return &domain.RequestQuery{
		Keyword:       Keyword,
		Page:          currentPage,
		Limit:         pageSize,
		SortBy:        sortBy,
		SortDirection: sortDirection,
		StartDate:     startDate,
		EndDate:       endDate,
		Filters:       Filters,
	}

}
