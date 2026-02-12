package middleware

import (
	"net/http"
	"time"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func TelemetriMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		tracer := otel.Tracer("http-server-middleware")

		ctx,span := tracer.Start(r.Context(), r.URL.Path)
		defer span.End()

		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))

		duration := time.Since(start)
		span.SetAttributes(
			attribute.String("http.method", r.Method),
			attribute.String("http.url", r.URL.String()),
			attribute.Float64("http.duration_ms", float64(duration.Milliseconds())),
		)
	})
}