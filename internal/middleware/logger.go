package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger 日志中间件
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// 创建响应写入器包装器来捕获状态码
		wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		// 处理请求
		next.ServeHTTP(wrapper, r)
		
		// 记录日志
		duration := time.Since(start)
		log.Printf(
			"[%s] %s %s %d %v %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			wrapper.statusCode,
			duration,
			r.UserAgent(),
		)
	})
}

// responseWriter 包装器用于捕获状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}