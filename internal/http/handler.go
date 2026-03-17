package http

import (
	"fmt"
	"net/http"
	"time"

	client "github.com/L1irik259/TestForOzonHTTPService/internal/client"
)

func Handler(client *client.ItemServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dateStr := r.URL.Query().Get("date_req")
		if dateStr == "" {
			http.Error(w, "Параметр date_req обязателен", http.StatusBadRequest)
			return
		}

		date, err := time.Parse("02/01/2006", dateStr)
		if err != nil {
			http.Error(w, "Неверный формат даты, используйте 02/01/2006", http.StatusBadRequest)
			return
		}

		items, err := client.FindAllItemsByDate(date)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка gRPC: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/xml")

		fmt.Fprintln(w, "<Valutes>")

		for _, item := range items {
			fmt.Fprintf(w, `<Valute ID="%s">
					<NumCode>%s</NumCode>
					<CharCode>%s</CharCode>
					<Nominal>%d</Nominal>
					<Name>%s</Name>
					<Value>%s</Value>
					<VunitRate>%s</VunitRate>
				</Valute>
				`,
				item.Id,
				item.NumCode,
				item.CharCode,
				item.Nominal,
				item.Name,
				item.Value,
				item.VunitRate,
			)
		}

		fmt.Fprintln(w, "</Valutes>")
	}
}
