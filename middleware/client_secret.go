package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var clientSecret string

func init() {
	clientSecret = os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		clientSecret = "minha-senha-secreta" // cuidado: não use isso em produção
	}
}

// ClientSecretMiddleware verifica o header "Client-Secret"
func ClientSecretMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerSecret := c.GetHeader("Client-Secret")

		if headerSecret == "" || headerSecret != clientSecret {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"erro": "Client secret inválido ou ausente",
			})
			return
		}

		c.Next()
	}
}
