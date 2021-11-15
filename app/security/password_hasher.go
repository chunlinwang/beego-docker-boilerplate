package security

import (
"crypto/rand"
// "encoding/base64"
// "fmt"
// "log"

"golang.org/x/crypto/argon2"
)

type argon2Param struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func init() {
	// b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	// b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	// fmt.Println(b64Salt)
	// fmt.Println(b64Hash)
}

func GetArgon2Param() (*argon2Param) {
	return &argon2Param{
			memory:      64 * 1024,
			iterations:  3,
			parallelism: 2,
			saltLength:  32,
			keyLength:   64,
		}
}

func GenerateFromPassword(password string, salt []byte) (hash []byte, err error) {
	c := GetArgon2Param();

	hash = argon2.IDKey([]byte(password), salt, c.iterations, c.memory, c.parallelism, c.keyLength)
 
	return hash, nil
}

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}