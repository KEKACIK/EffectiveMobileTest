package api

import (
	"TestTask/pkg/logging"
	"TestTask/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPagination(
	items []any,
	total int,
	page int,
	limit int,
) PaginationResponse {
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

func debug(l *logging.Logger, url, method, host string, code int) {
	// Выводим в консоль метод, ссылку и айпи запроса, а также код ответа
	l.Debug(fmt.Sprintf("%s %s %s - %d", method, url, host, code))
}

func ErrorNotFoundResponse(w http.ResponseWriter) {
	// Функция возврата ошибки: 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{})
}

func ErrorResponse(w http.ResponseWriter, code int, err error) {
	// Универсальная функция возврата ошибки
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"message": err.Error(),
	})
}
