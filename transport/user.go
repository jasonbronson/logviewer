package transport

import (
	"net/http"
	"time"
	"os"

	"github.com/jasonbronson/logreader/config"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(g *gin.Context) {

	userParams := UserLogin{}
	if err := g.BindJSON(&userParams); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userEmail := os.Getenv("EMAIL")
	userPassword := os.Getenv("PASSWORD")

	if userEmail ==  userParams.Email && userPassword == userParams.Password{
		claims := CustomClaims{
			jwt.StandardClaims{
				Audience: config.Cfg.JWTAudience,
				IssuedAt: time.Now().Unix(),
				Issuer:   config.Cfg.JWTIssuer,
				Subject:  userEmail,
			},
			userEmail,
			nil,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedString, err := token.SignedString([]byte(config.Cfg.JWTSecret))
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		g.Writer.Header().Set("Content-Type", "text/plain")
		g.Writer.WriteHeader(http.StatusOK)
		g.Writer.Write([]byte(signedString))
		return
	} else {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email/password"})
		return
	}

}

type CustomClaims struct {
	jwt.StandardClaims
	Email             string     `json:"email"`
	Expiration        *time.Time `json:"expiration"`
}