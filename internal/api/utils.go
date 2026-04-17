package api

import (
	"TestTask/pkg/utils"
	"encoding/json"
	"net/http"
)

func GetPagination(items []any, total int, page int, limit int) PaginationResponse {
	// Универсальная функция пагинации
	// Все структуры необходимо привести в вид Any
	result := PaginationResponse{Total: total, Page: page, Limit: limit}

	maxPages := utils.RoundUp(total, limit)
	if page > maxPages {
		result.Items = []any{}
		return result
	}

	startI := (page - 1) * limit
	endI := utils.Min(total, startI+limit)

	result.Items = items[startI:endI]
	return result
}

func SendErrorNotFoundResponse(w http.ResponseWriter) {
	// Функция возврата ошибки: 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{})
}

func SendErrorResponse(w http.ResponseWriter, code int, err error) {
	// Универсальная функция возврата ошибки
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"message": err.Error(),
	})
}
