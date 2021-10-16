package security

import (
	"crypto/rsa"
	"fmt"
	//"os"
	"io/ioutil"
	"log"
	//"net/http"
    "github.com/beego/beego/v2/core/logs"
	"app/models"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "var/certs/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "var/certs/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type AccessClaims struct {
	UserInfo models.UserInfo
	StandardClaims jwt.StandardClaims
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

type JWTClaims struct {
    username string
    exp int
    iat int
}

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	if err != nil {
		logs.Error(err)
	}
}

func GenerateToken() string {
	// set our claims
	claims := AccessClaims {
		UserInfo: models.UserInfo {"user", "human", "fi", "women"},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 1645620008,
			Issuer:    "beego",
		},
	}

	// create a signer for rsa 256
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims.StandardClaims)

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprintln(w, "Sorry, error while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return ""
	}

	return tokenString
}

func Valid(tokenString string) (map[string]interface{}, bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return verifyKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return claims, false
	}
}