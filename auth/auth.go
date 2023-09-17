package auth

import (
	"errors"
	"expense-tracker/pkg/logger"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// adding fields like jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UUID      string
	jwt.RegisteredClaims
}

func GenerateToken(validTimeInHour uint32, email, firstName, lastName, uuid, JWTSecret string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(validTimeInHour) * time.Hour)
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UUID:      uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// In JWT, the expiry time is expressed as unix milliseconds
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SignedToken, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		logger.Error("error in sigining token")
		return "", err
	}
	return SignedToken, nil

}

func ValidateToken(token, secret string) (*SignedDetails, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("unable to parse token")
	}
	// claims, ok := jwtToken.Claims.(*SignedDetails) {
	// 	//return error
	// }
	if claims, ok := jwtToken.Claims.(*SignedDetails); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
	// if !ok {
	//     msg = fmt.Sprintf("the token is invalid")
	//     msg = err.Error()
	//     return
	// }
	// _=claims
	// if claims.ExpiresAt < time.Now().Local().Unix() {
	// 	//token expired
	// }

}
