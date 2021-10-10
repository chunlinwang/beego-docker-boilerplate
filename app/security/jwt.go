package security

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    "github.com/beego/beego/v2/core/logs"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "certs/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "certs/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type JWTClaims struct {
    username string
    exp int
    iat int
}

// read the key files before starting http handlers
func init() {
    //	now := time.Now().UTC()

	signBytes, err := ioutil.ReadFile(privKeyPath)
	logs.Error("signBytes error", err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	logs.Error("signKey error", err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	logs.Error("verifyBytes error", err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	logs.Error("verifyKey error", err)
}

func getToken() {
	// set our claims
	claims := AccessClaims{
		Access:   "level1",
		UserInfo: UserInfo{user, "human"},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "asym-example",
		},
	}


	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
}