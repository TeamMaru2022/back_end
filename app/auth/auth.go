package auth

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserJwt struct {
	Id   uint16
	Mail string
	jwt.StandardClaims
}

func CreateTokenString(id uint16, mail string) string {

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &UserJwt{
		Id:   id,
		Mail: mail,
	})

	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}

func ValidateTokenString(tokenstring string) UserJwt {
	log.Println(tokenstring)

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})

	log.Println(token.Claims, err)
	userJwt := UserJwt{}
	token, err = jwt.ParseWithClaims(tokenstring, &userJwt, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	return userJwt

}
