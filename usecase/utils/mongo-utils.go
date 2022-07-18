package utils

func ConvertSortToMongo(typeSort string) int {
	constantSort := make(map[string]interface{})

	constantSort["desc"] = -1
	constantSort["asc"] = 1

	return constantSort[typeSort].(int)
}
