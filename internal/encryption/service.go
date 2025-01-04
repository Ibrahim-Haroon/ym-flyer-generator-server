package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

var (
	ErrInvalidKeyLength  = errors.New("encryption key must be 32 bytes (256 bits)")
	ErrInvalidCiphertext = errors.New("invalid ciphertext")
	ErrDecryptionFailed  = errors.New("decryption failed")
	ErrEncryptionFailed  = errors.New("encryption failed")
)

type APIKeys struct {
	TextProvider  textgen.ProviderType
	TextAPIKey    string
	ImageProvider imagegen.ProviderType
	ImageAPIKey   string
}

type Service struct {
	gcm cipher.AEAD
}

func NewService(key []byte) (*Service, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes (256 bits)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	return &Service{gcm: gcm}, nil
}

func (s *Service) Encrypt(plaintext string) (string, error) {
	nonce := make([]byte, s.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := s.gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

func (s *Service) Decrypt(encodedCiphertext string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	if len(ciphertext) < s.gcm.NonceSize() {
		return "", fmt.Errorf("ciphertext too short")
	}
	nonce := ciphertext[:s.gcm.NonceSize()]
	encryptedData := ciphertext[s.gcm.NonceSize():]

	plaintext, err := s.gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %w", err)
	}

	return string(plaintext), nil
}

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	return key, nil
}
