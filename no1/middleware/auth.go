package middleware

import (
	"majoo-be-test/config"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	jwtv3 "github.com/golang-jwt/jwt"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

const (
	accessTokenCookieName  = "access-token"
	refreshTokenCookieName = "refresh-token"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwtv3.StandardClaims
}

func GenerateTokensAndSetCookies(userID int, cfg *config.Config, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(userID, cfg.JwtSecretKey, cfg.JwtTokenExpiryDuration)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)

	return nil
}

func Logout(c echo.Context) {
	timeNow := time.Now()
	setTokenCookie(accessTokenCookieName, "", timeNow, c)
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, "Your token is expired, please login")
}

func JWTmiddleware(JwtSecretKey string) echo.MiddlewareFunc {
	return echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
		Claims:                  &Claims{},
		SigningKey:              []byte(JwtSecretKey),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: JWTErrorChecker,
	})
}

func generateAccessToken(userID int, JwtSecretKey string, JwtTokenExpiryDuration int) (string, time.Time, error) {
	TokenExpiryTime := time.Now().Add(time.Duration(JwtTokenExpiryDuration) * time.Second)
	return generateToken(userID, TokenExpiryTime, []byte(JwtSecretKey))
}

func generateToken(userID int, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwtv3.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwtv3.NewWithClaims(jwtv3.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}
