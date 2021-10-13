package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

const (
	PASSPHARSE = "thisisthepassphrase"
)

func DoExample() {

	str := "string to encrypt"
	fmt.Printf("String before encrypting: %s\n", str)

	encryptedString := encrypt(str)
	fmt.Printf("String after encrypting: %s\n", encryptedString)

	decryptedString := decrypt(encryptedString)
	fmt.Printf("String after decrypting: %s\n", decryptedString)

}

func createHash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}
func decrypt(input string) string {

	data, err := hex.DecodeString(input)
	if err != nil {
		failed(err)
	}
	key := createHash(PASSPHARSE)
	block, err := aes.NewCipher(key)
	if err != nil {

		failed(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		failed(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		failed(err)
	}
	return string(plaintext[:])
}

func encrypt(input string) string {
	data := []byte(input)
	key := createHash(PASSPHARSE)
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		failed(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		failed(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	encryptedString := hex.EncodeToString(ciphertext)
	return encryptedString
}

func failed(err error) {
	log.Fatalf("Failure: %s", err)
}
