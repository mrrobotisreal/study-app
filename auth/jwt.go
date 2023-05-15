package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"study-app/types"
)

func GenerateJWT(username string, email string) string {
	claims := &types.JWTClaim{
		Username: username,
		Email:    email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("WinTer"))
	if err != nil {
		log.Println("Error Performing token.SignedString() Function Call..\n\n", err.Error())
		return ""
	}
	return tokenString
}

func ValidateJWT(signedToken string, username string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken, &types.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("WinTer"), nil
	})
	if err != nil {
		log.Println("Error Performing jwt.ParseWithClaims() Function Call..\n\n", err.Error())
		return err
	}
	claims, ok := token.Claims.(*types.JWTClaim)
	if !ok {
		log.Println("Error Performing token.Claims.(*JWTClaim) Read/Assignment Operation..\n\n", ok)
		err = errors.New("Error Performing token.Claims.(*JWTClaim) Read/Assignment Operation..")
		return err
	}
	if claims.Username != username {
		log.Println("Supplied username does not match claims.Username..\nJWT is not valid; returning error..\n\n")
		err = errors.New("Error: Incorrect Claims\n" +
			"One or more claims do not match the supplied information; token is not valid")
		return err
	}
	return nil
}
