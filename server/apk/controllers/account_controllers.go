package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aliamerj/aircup/server/apk/modules"
	resType "github.com/aliamerj/aircup/server/internal/constant"
	"github.com/aliamerj/aircup/server/internal/utility"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"
)

var validate = validator.New()

func Register(c echo.Context, db *gorm.DB) error {
	var account modules.Account

	if err := c.Bind(&account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format in request body. Please ensure the request body is a valid JSON object.")
	}

	// Validate the account data
	if err := validate.Struct(account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := hashPassword(account.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not hash password")
	}
	account.Password = hashedPassword

	// Save the account to the database
	if result := db.Create(&account); result.Error != nil {
		if result.Error.Error() == "UNIQUE constraint failed: accounts.email" {
			return echo.NewHTTPError(http.StatusBadRequest, "Email is already exists")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	// login
	// Expiration settings
	accessTokenExp := time.Now().Add(15 * time.Minute)    // Access token expires in 15 minutes
	refreshTokenExp := time.Now().Add(7 * 24 * time.Hour) // Refresh token expires in 7 days

	accessToken, refreshToken, err := generateTokens(account, accessTokenExp, refreshTokenExp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate tokens")
	}

	// Update or create session record
	var session modules.Session
	upsertData := modules.Session{RefreshToken: refreshToken, ExpiresAt: refreshTokenExp.Unix()}
	result := db.Where(modules.Session{AccountID: account.ID}).Assign(upsertData).FirstOrCreate(&session)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not update or create refresh token session")
	}

	// Set access token in an HttpOnly cookie
	c.SetCookie(&http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Expires:  accessTokenExp,
		HttpOnly: true,
		Secure:   true, // Set to true in production
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	// Set refresh token in a separate HttpOnly cookie
	c.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  refreshTokenExp,
		HttpOnly: true,
		Secure:   true, // Set to true in production
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
	return c.JSON(http.StatusCreated, utility.SuccesRespnse("Account successfully created", resType.AccountCreated, map[string]string{
		"id":        fmt.Sprintf("%v", account.ID),
		"firstName": account.FirstName,
		"lastName":  account.LastName,
		"email":     account.Email,
	},
	))

}

func Login(c echo.Context, db *gorm.DB) error {
	var credentials modules.LoginCredentials
	var account modules.Account
	if err := c.Bind(&credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format in request body. Please ensure the request body is a valid JSON object.")
	}
	// Validate the credentials
	if err := validate.Struct(credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Input")
	}
	// Retrieve the user by email
	if tx := db.First(&account, modules.Account{Email: credentials.Email}); tx.Error != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(credentials.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or passwordx")
	}

	// Expiration settings
	accessTokenExp := time.Now().Add(15 * time.Minute)    // Access token expires in 15 minutes
	refreshTokenExp := time.Now().Add(7 * 24 * time.Hour) // Refresh token expires in 7 days

	accessToken, refreshToken, err := generateTokens(account, accessTokenExp, refreshTokenExp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate tokens")
	}

	// Update or create session record
	var session modules.Session
	upsertData := modules.Session{RefreshToken: refreshToken, ExpiresAt: refreshTokenExp.Unix()}
	result := db.Where(modules.Session{AccountID: account.ID}).Assign(upsertData).FirstOrCreate(&session)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not update or create refresh token session")
	}

	// Set access token in an HttpOnly cookie
	c.SetCookie(&http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Expires:  accessTokenExp,
		HttpOnly: true,
		Secure:   true, // Set to true in production
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	// Set refresh token in a separate HttpOnly cookie
	c.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  refreshTokenExp,
		HttpOnly: true,
		Secure:   true, // Set to true in production
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
	return c.JSON(http.StatusOK, utility.SuccesRespnse("Login success", resType.Login, map[string]string{
		"id":        fmt.Sprintf("%v", account.ID),
		"firstName": account.FirstName,
		"lastName":  account.LastName,
		"email":     account.Email,
	},
	))
}

func VerifySession(c echo.Context, db *gorm.DB) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse user token")
	}

	claims, ok := user.Claims.(*utility.JwtAuthClaims)
	if !ok || !user.Valid {
		// Access token is invalid or claims are not correct, try to refresh it
		return refreshAccessToken(c, db)
	}

	// If the access token is still valid, return the user info
	return c.JSON(http.StatusOK, utility.SuccesRespnse("Token valid", resType.Refresh, map[string]string{
		"id":        fmt.Sprintf("%v", claims.AccountID),
		"email":     claims.Email,
		"firstName": claims.FirstName,
		"lastName":  claims.LastName,
	}))
}

func refreshAccessToken(c echo.Context, db *gorm.DB) error {
	refreshTokenCookie, err := c.Cookie("refreshToken")
	if err != nil || refreshTokenCookie.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "No refresh token provided")
	}

	// Validate refresh token from the database
	var session modules.Session
	if err := db.Where("refresh_token = ?", refreshTokenCookie.Value).First(&session).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid refresh token")
	}

	if session.ExpiresAt < time.Now().Unix() {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh token expired")
	}

	// Generate a new access token
	newClaims := &utility.JwtAuthClaims{
		AccountID: session.AccountID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return errors.New("JWT_SECRET in .env is missing")
	}

	newAccessToken, err := generateTokenWithClaims(newClaims, jwtSecret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate new access token")
	}

	// Set new access token in response header
	c.SetCookie(&http.Cookie{
		Name:     "accessToken",
		Value:    newAccessToken,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
		Secure:   true, // Set to true in production
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusOK, utility.SuccesRespnse("Refresh Token success", resType.Refresh, map[string]string{
		"id":        fmt.Sprintf("%v", session.AccountID),
		"email":     newClaims.Email,
		"firstName": newClaims.FirstName,
		"lastName":  newClaims.LastName,
	}))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Generates both access and refresh tokens
func generateTokens(account modules.Account, accessTokenExp, refreshTokenExp time.Time) (accessToken string, refreshToken string, err error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if len(jwtSecret) == 0 {
		return "", "", errors.New("JWT_SECRET in .env is missing")
	}

	// Access token
	accessClaims := &utility.JwtAuthClaims{
		AccountID: account.ID,
		Email:     account.Email,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExp),
		},
	}
	accessToken, err = generateTokenWithClaims(accessClaims, jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh token
	refreshClaims := &utility.JwtAuthClaims{
		AccountID: account.ID,
		Email:     account.Email,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExp),
		},
	}
	refreshToken, err = generateTokenWithClaims(refreshClaims, jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateTokenWithClaims(claims *utility.JwtAuthClaims, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
