package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

func Encrypt(data, passphrase []byte) ([]byte, error) {
	cypher, err := buildCypher(passphrase)
	if err != nil {
		return nil, err
	}

	// make the nonce array as iv
	nonce := make([]byte, cypher.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	// seal the data (prepend the nonce to the encrypted data)
	return cypher.Seal(nonce, nonce, data, nil), nil
}

func Decrypt(encrypted, passphrase []byte) ([]byte, error) {
	cypher, err := buildCypher(passphrase)
	if err != nil {
		return nil, err
	}

	// split nonce from data payload
	nonce, payload := encrypted[:cypher.NonceSize()], encrypted[cypher.NonceSize():]

	// decrypt data
	data, err := cypher.Open(nil, nonce, payload, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildCypher(passphrase []byte) (cipher.AEAD, error) {
	// key stretching with a KDF
	derivedKey := keyStretching(passphrase)

	// compute the hash of the derived key
	hash := sha256.Sum256(derivedKey)

	// build AES cypher based on the hash
	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return nil, err
	}

	// wrap the block cypher in GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm, nil
}
