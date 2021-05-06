package middleware

import (
	"2021_1_Noskool_team/internal/pkg/monitoring"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func PanicMiddleware(metrics *monitoring.PromMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqTime := time.Now()
			defer func() {
				if err := recover(); err != nil {
					respTime := time.Since(reqTime)
					url := "/api/v1/" + re.ReplaceAllString(r.URL.String()[8:], ":num")
					metrics.Hits.WithLabelValues(
						strconv.Itoa(http.StatusInternalServerError), url, r.Method).Inc()

					metrics.Timings.WithLabelValues(
						strconv.Itoa(http.StatusInternalServerError), url,
						r.Method).Observe(respTime.Seconds())

					logrus.Error(fmt.Sprintf("panic catched: %s", err))
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
