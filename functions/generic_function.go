package functions

func PrepareResponse[T any, U any](result []T, mapper func(T) U) []U {
	var response []U
	for _, rowValue := range result {
		mapDto := mapper(rowValue)
		response = append(response, mapDto)
	}
	return response
}
