package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Gunakan kunci rahasia yang sama dengan yang ada di Login handler
var jwtSecretKey = []byte("rahasia_negara_kkp_corp")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		// Token biasanya formatnya "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau kadaluwarsa"})
			c.Abort()
			return
		}

		// Mengambil claims dari token dan menyimpannya ke context Gin
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"]) // <--- INI BAGIAN PENTINGNYA
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Gagal memproses klaim token"})
			c.Abort()
			return
		}
	}
}
