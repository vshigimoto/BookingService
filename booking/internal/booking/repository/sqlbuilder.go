package repository

func SqlBuilder(sortKey, sortBy string) string {
	queryBase := "SELECT * from hotel"
	if sortKey == "" {
		sortKey = "id"
	}
	queryBase = queryBase + " ORDER BY " + sortKey + " "
	queryBase = queryBase + sortBy + " "
	return queryBase
}
