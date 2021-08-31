package vmess

import (
	"bytes"
	"crypto/aes"
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

	cipher, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext := make([]byte, testSize)
	for encryptSize := 0; encryptSize < testSize; encryptSize += blockSize {
		cipher.Encrypt(ciphertext[encryptSize:], plaintext[encryptSize:])
	}

	ciphertextCopy := make([]byte, testSize)
	copy(ciphertextCopy, ciphertext)

	reader, err := NewDecryptionReader(bytes.NewReader(ciphertextCopy), key)
	if err != nil {
		t.Fatal(err)
	}

	readText := make([]byte, testSize)
	readSize := 0
	for readSize < testSize {
		nBytes := mrand.Intn(16) + 1
		if nBytes > testSize-readSize {
			nBytes = testSize - readSize
		}
		bytesRead, err := reader.Read(readText[readSize : readSize+nBytes])
		if err != nil {
			t.Fatal(err)
		}
		if bytesRead != nBytes {
			t.Errorf("Expected to read %d bytes, but only read %d bytes", nBytes, bytesRead)
		}
		readSize += nBytes
	}
	if !bytes.Equal(readText, plaintext) {
		t.Errorf("Expected plaintext %v, but got %v", plaintext, readText)
	}
}
