package middleware

import (
	"2021_1_Noskool_team/internal/pkg/monitoring"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	re = regexp.MustCompile("[0-9]+")
)

func LoggingMiddleware(metrics *monitoring.PromMetrics) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			msg := fmt.Sprintf("URL: %s, METHOD: %s, REMOTE_ADDR%s",
				r.RequestURI, r.Method, r.RemoteAddr)
			logrus.Info(msg)

			reqTime := time.Now()
			next.ServeHTTP(w, r)
			respTime := time.Since(reqTime)
			if r.URL.Path != "/metrics" {
				url := "/api/v1/" + re.ReplaceAllString(r.URL.String()[8:], ":num")
				metrics.Hits.WithLabelValues(strconv.Itoa(http.StatusOK), url, r.Method).Inc()
				metrics.Timings.WithLabelValues(
					strconv.Itoa(http.StatusOK), url, r.Method).Observe(respTime.Seconds())
			}
		})
	}
}
