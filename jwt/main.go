package main

import (
	"github.com/robbert229/jwt"
	"devx/iferr"
	"fmt"
	"time"
)

func main() {
	key := "LF7uXyDnuuKoUYonzboPgCQ0h33Nc7P8UWZ9Y5nW"
	algorithm := jwt.HmacSha256(key)
	token := createClaims(algorithm, 0, 10)
	readToken(algorithm, token)
}

func createClaims(algorithm jwt.Algorithm, daysNotBefore, daysExpire int) string {
	// Set claims
	claims := jwt.NewClaim()
	claims.Set("Role", "Admin")
	claims.SetTime("nbf", time.Now().AddDate(0, 0, daysNotBefore)) // Not before
	claims.SetTime("exp", time.Now().AddDate(0, 0, daysExpire))    // Expire

	// Sign the claims
	token, err := algorithm.Encode(claims)
	iferr.Panic(err)

	fmt.Println(token)
	return token
}

func readToken(algorithm jwt.Algorithm, token string) {
	// Validate token
	if algorithm.Validate(token) != nil {
		panic("validation error")
	}
	fmt.Println("Authenticated")

	// Read the claims
	claims, err := algorithm.Decode(token)
	iferr.Panic(err)
	role, err := claims.Get("Role")
	iferr.Panic(err)
	fmt.Println("Role:", role.(string))
}
