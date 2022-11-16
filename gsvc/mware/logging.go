package mware

import (
	"net/http"
	"time"

	"gsvc/pkg/util"
)

// responseObserver is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseObserver struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(log *util.Logger, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		obs := &responseObserver{ResponseWriter: w}
		next.ServeHTTP(obs, r)
		blank := " "
		hyphen := "-"
		// LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-agent}i\"" combined
		log.Sugar().Info(
			r.RemoteAddr, blank,
			hyphen, blank,
			r.URL.User, blank,
			start.Format(time.RFC3339), blank,
			r.Method, blank, r.URL.EscapedPath(), blank, r.Proto, blank,
			obs.status, blank,
			obs.written, blank,
			time.Since(start), blank,
			r.Referer(), blank,
			r.UserAgent(),
		)
	}
	return http.HandlerFunc(fn)
}
