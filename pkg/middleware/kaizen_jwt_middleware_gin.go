package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdaptHTTPMiddleware converte um middleware http.Handler -> http.Handler
// para o formato gin.HandlerFunc
func AdaptHTTPMiddleware(mw func(http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cria um handler final que simplesmente continua a cadeia do Gin
		finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Atualiza request no gin.Context
			c.Request = r
			c.Next()
		})

		// Aplica o middleware externo
		mw(finalHandler).ServeHTTP(c.Writer, c.Request)

		// Se o middleware já escreveu resposta (ex: http.Error), aborta o Gin
		if c.Writer.Written() {
			c.Abort()
		}
	}
}
