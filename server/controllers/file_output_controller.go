package controllers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/serpedious/automatic-trading-system/server/bitflyer/usecase"
)

func GetCsvFile(w http.ResponseWriter, r *http.Request) {
	apiClient := usecase.CreateClient()
	params := map[string]string{
		"currency_code": "JPY",
		"count":         "10",
		"before":        "0",
		"after":         "0",
	}
	balance_history_data, _ := apiClient.GetBalanceHistory(params)
	w.Header().Set("Content-Disposition", "attachment; filename=test.csv")
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", string(len(balance_history_data)))
	for _, record := range balance_history_data {
		rtCst := reflect.TypeOf(record)
		rvCst := reflect.ValueOf(record)
		for i := 0; i < rtCst.NumField(); i++ {
			f := rtCst.Field(i)
			v := rvCst.FieldByName(f.Name).Interface()
			value := ""
			if _, ok := v.(float64); ok {
				value = strconv.FormatFloat(v.(float64), 'f', 2, 64)
			}

			if _, ok := v.(int); ok {
				value = strconv.Itoa(v.(int))
			}

			if _, ok := v.(string); ok {
				value = v.(string)
			}
			w.Write([]byte(value + ","))
		}
		w.Write([]byte("\n"))
	}
}
