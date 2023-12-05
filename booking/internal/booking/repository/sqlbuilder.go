package repository

func SqlBuilder(sortKey, sortBy string) string {
	queryBase := "SELECT * from hotel"
	if sortKey == "" {
		sortKey = "id"
	}
	if sortBy != "ASC" && sortBy != "DESC" {
		return queryBase
	}
	queryBase = queryBase + " ORDER BY " + sortKey + " "
	queryBase = queryBase + sortBy + " "
	return queryBase
}
