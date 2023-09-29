package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(signKey string, refreshKey string, userID string) map[string]any {
	var res = map[string]any{}
	accessToken := generateToken(signKey, userID)
	if accessToken == "" {
		return nil
	}
	refreshToken := generateRefreshToken(refreshKey, accessToken)
	if refreshToken == "" {
		return nil
	}
	res["access_token"] = accessToken
	res["refresh_token"] = refreshToken
	return res
}

func generateToken(signKey string, id string) string {
	var claims = jwt.MapClaims{}
	claims["id"] = 1
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := signature.SignedString([]byte(signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func generateRefreshToken(signKey string, accessToken string) string {
	var claims = jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := signature.SignedString([]byte(signKey))

	if err != nil {
		return ""
	}

	return refreshToken
}

func ExtractToken(token *jwt.Token) any {
	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil
		}

		expiration := int64(claims["exp"].(float64))
		currentTime := time.Now().Unix()
		if expiration < currentTime {
			return nil
		}

		return claims
	}
	return nil
}
