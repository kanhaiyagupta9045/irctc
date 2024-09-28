package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type SignedDetails struct {
	Email string
	jwt.StandardClaims
}

func GenerateAccessToken(email string) (string, error) {
	ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
	if ACCESS_TOKEN_SECRET == "" {
		return "", errors.New("access token secret is not set")
	}
	accessClaims := &SignedDetails{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(ACCESS_TOKEN_SECRET))
	if err != nil {
		return "", err
	}

	return token, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	check := true
	msg := ""

	if err != nil {
		fmt.Println("error", err)
		msg = fmt.Sprintln("password is incorrect")
		check = false
	}
	return check, msg
}

func ValidateToken(signedToken string) (*SignedDetails, error) {
    token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        secretKey := os.Getenv("ACCESS_TOKEN_SECRET")
        if secretKey == "" {
            return nil, errors.New("SECRET_KEY not set in environment")
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        return nil, fmt.Errorf("failed to parse token: %w", err)
    }
    claims, ok := token.Claims.(*SignedDetails)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token claims or token is not valid")
    }
    if claims.ExpiresAt < time.Now().Unix() {
        return nil, errors.New("token expired")
    }

    return claims, nil
}

