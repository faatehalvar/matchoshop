package domain

import (
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/dgrijalva/jwt-go"
)

const ACCESS_TOKEN_DURATION = time.Hour
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 30

const HMAC_SAMPLE_SECRET = "machoshop-secret"

type User struct {
	UserID    int64  `db:"user_id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	RoleID    int64  `db:"role_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type AccessTokenClaims struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	TokenType string `json:"token_type"`
	UserID    int64  `json:"user_id"`
	RoleID    int64  `json:"role_id"`
	jwt.StandardClaims
}

type AuthToken struct {
	token *jwt.Token
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func (r User) AccessTokenClaims() AccessTokenClaims {
	return AccessTokenClaims{
		UserID: r.UserID,
		RoleID: r.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}

func (r AccessTokenClaims) RefreshTokenClaims() RefreshTokenClaims {
	return RefreshTokenClaims{
		TokenType: "refresh_token",
		UserID:    r.UserID,
		RoleID:    r.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_DURATION).Unix(),
		},
	}
}

func (r AuthToken) NewAccessToken() (string, *errs.AppError) {
	signedString, err := r.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		logger.Error("Failed while signing access token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}

func (r AuthToken) NewRefreshToken() (string, *errs.AppError) {

	claims := r.token.Claims.(AccessTokenClaims)
	refreshClaims := claims.RefreshTokenClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))

	if err != nil {
		logger.Error("Failed while signing refresh token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate refresh token")
	}

	return signedString, nil
}
