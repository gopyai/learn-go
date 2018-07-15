package mdw

import (
	"time"

	"devx/iferr"

	"github.com/pkg/errors"
	"github.com/robbert229/jwt"
)

var algorithm = jwt.HmacSha256("LF7uXyDnuuKoUYonzboPgCQ0h33Nc7P8UWZ9Y5nW")

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrUnexpected   = errors.New("unexpected error")
)

func CreateJWT(daysNotBefore, daysExpire int, email string) string {
	// Set claims
	claims := jwt.NewClaim()
	claims.SetTime("nbf", time.Now().AddDate(0, 0, daysNotBefore)) // Not before
	claims.SetTime("exp", time.Now().AddDate(0, 0, daysExpire))    // Expire
	claims.Set("ema", email)

	// Sign the claims
	token, err := algorithm.Encode(claims)
	iferr.Panic(err)
	return token
}

func verifyAndReadToken(token string) (email string, err error) {
	// Validate token
	if algorithm.Validate(token) != nil {
		return "", ErrInvalidToken
	}

	// Read the claims
	claims, err := algorithm.Decode(token)
	if err != nil {
		return "", err
	}
	ema, err := claims.Get("ema")
	if err != nil {
		return "", err
	}

	switch ema.(type) {
	case string:
		return ema.(string), nil
	default:
		return "", ErrUnexpected
	}
}
