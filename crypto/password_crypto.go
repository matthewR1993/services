package crypto

import (
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha256"
	"math/rand"
	"time"
	"reflect"
	b64 "encoding/base64"
)

func random_bytes(strlen int) []byte {
	var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz123456789"  // Allowed symbols
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return result
}

/*
  PBKDF2 algorithm is used here with SHA-256 hashing algorhitm;
  key length is 256; hashed for 4096 times.
  Return format: <algorithm>$<iterations>$<salt>$<hash>
*/
func GeneratePDKDF2key(password []byte) []byte {
	// generate random salt
	salt := random_bytes(32)

	// calculate hash
	key := pbkdf2.Key(password, salt, 4096, sha256.Size, sha256.New)
	s := b64.StdEncoding.EncodeToString(key) // base64 encoded
	hash := []byte(s)

	dlr := byte(0x24)  // '$' - symbol 
	algo := []byte("SHA256")
	iter := []byte("4096")
	res := append(algo, dlr)
	res = append(res, iter...)
	res = append(res, dlr)
	res = append(res, salt...)
	res = append(res, dlr)
	res = append(res, hash...)
	// res format: <algorithm>$<iterations>$<salt>$<hash>
	return res
}

func GetHash(password []byte, salt []byte) []byte {
	key := pbkdf2.Key(password, salt, 4096, sha256.Size, sha256.New)
	s := b64.StdEncoding.EncodeToString(key)
	hash := []byte(s)
	return hash
}

func CompareHashes(a []byte, b []byte) bool {
	res := false
	res = reflect.DeepEqual(a, b)
	return res
}



