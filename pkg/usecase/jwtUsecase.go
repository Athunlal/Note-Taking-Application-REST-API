package usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	"github.com/athunlal/Note-Taking-Application/pkg/usecase/interfaces"
	"github.com/golang-jwt/jwt"
)

type jwtUseCase struct {
	SecretKey string
}

// ValidateJwtUser implements interfaces.JwtUseCase.
func (*jwtUseCase) ValidateJwtUser(userid uint) (domain.User, error) {
	panic("unimplemented")
}

func (u *jwtUseCase) GenerateRefreshToken(userid int, email string, role string) (string, error) {

	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 7).Unix() // Refresh token expires in 7 days
	refreshTokenClaims := jwt.MapClaims{
		"exp":  refreshTokenExpiresAt,
		"sub":  userid,
		"type": "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(u.SecretKey))
	return refreshTokenString, err
}

func (u *jwtUseCase) GenerateAccessToken(userid int, email string, role string) (string, error) {
	claims := domain.JwtClaims{
		Userid: uint(userid),
		Email:  email,
		Source: "AccessToken",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(500)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(u.SecretKey))

	return accessToken, err

}

func (u *jwtUseCase) VerifyToken(token string) (bool, *domain.JwtClaims) {
	claims := &domain.JwtClaims{}
	tkn, err := u.GetTokenFromString(token, claims)
	if err != nil {
		return false, claims
	}
	if tkn.Valid {
		if err := claims.Valid(); err != nil {
			return false, claims
		}
	}
	return true, claims
}

func (u *jwtUseCase) GetTokenFromString(signedToken string, claims *domain.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(u.SecretKey), nil
	})
}

func (use *userUseCase) ValidateJwtUser(userId uint) (*domain.User, error) {
	user, err := use.Repo.FindUserById(int(userId))
	if err != nil {
		return user, errors.New("Unauthorized User")
	}
	return user, nil
}

func NewJWTUseCase() interfaces.JwtUseCase {
	return &jwtUseCase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
