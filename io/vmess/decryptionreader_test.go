package vmess

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	mrand "math/rand"
	"testing"
)

func randomBytes(p []byte, t *testing.T) {
	nBytes, err := rand.Read(p)
	if err != nil {
		t.Fatal(err)
	}
	if nBytes != len(p) {
		t.Errorf("Unable to generate %d bytes of random buffer", len(p))
	}
}

func TestNormalReading(t *testing.T) {
	testSize := 256
	plaintext := make([]byte, testSize)
	randomBytes(plaintext, t)

	keySize := 16
	key := make([]byte, keySize)
	randomBytes(key, t)
	iv := make([]byte, keySize)
	randomBytes(iv, t)

	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	aesMode := cipher.NewCBCEncrypter(aesBlock, iv)

	ciphertext := make([]byte, testSize)
	aesMode.CryptBlocks(ciphertext, plaintext)

	ciphertextcopy := make([]byte, testSize)
	copy(ciphertextcopy, ciphertext)

	reader, err := NewDecryptionReader(bytes.NewReader(ciphertextcopy), key, iv)
	if err != nil {
		t.Fatal(err)
	}

	readtext := make([]byte, testSize)
	readSize := 0
	for readSize < testSize {
		nBytes := mrand.Intn(16) + 1
		if nBytes > testSize-readSize {
			nBytes = testSize - readSize
		}
		bytesRead, err := reader.Read(readtext[readSize : readSize+nBytes])
		if err != nil {
			t.Fatal(err)
		}
		if bytesRead != nBytes {
			t.Errorf("Expected to read %d bytes, but only read %d bytes", nBytes, bytesRead)
		}
		readSize += nBytes
	}
	if !bytes.Equal(readtext, plaintext) {
		t.Errorf("Expected plaintext %v, but got %v", plaintext, readtext)
	}
}
