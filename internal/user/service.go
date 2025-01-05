package user

import (
	"fmt"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/auth"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/encryption"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user/model"
	"github.com/google/uuid"
)

type Service struct {
	repo       *Repository
	auth       *auth.Service
	encryption *encryption.Service
}

func NewService(repo *Repository, auth *auth.Service, encryption *encryption.Service) *Service {
	return &Service{
		repo:       repo,
		auth:       auth,
		encryption: encryption,
	}
}

func (s *Service) Register(user *model.User) (string, error) {
	existingUser, err := s.repo.GetUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return "", fmt.Errorf("username already taken")
	}

	existingUser, err = s.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return "", fmt.Errorf("email already registered")
	}

	if len(user.PasswordHash) < 60 { // bcrypt hashes are 60 characters
		return "", fmt.Errorf("password must be hashed before sending to server")
	}

	user.ID = uuid.New().String()
	user.ActiveStatus = true
	user.IsAdmin = false

	if err := s.repo.CreateUser(user); err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}

	token, err := s.auth.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return token, nil
}

func (s *Service) Login(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if len(password) < 60 { // bcrypt hashes are 60 characters
		return "", fmt.Errorf("password must be hashed before sending to server")
	}

	if user.PasswordHash != password {
		return "", fmt.Errorf("invalid credentials")
	}

	if !user.ActiveStatus {
		return "", fmt.Errorf("account is inactive")
	}

	token, err := s.auth.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return token, nil
}

func (s *Service) GetUserById(userID string) (*model.User, error) {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	return user, nil
}

func (s *Service) GetUserByUsername(username string) (*model.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	return user, nil
}

func (s *Service) UpdateUser(user *model.User) error {
	existingUser, err := s.repo.GetUserById(user.ID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Preserved fields from existing user
	user.CreatedAt = existingUser.CreatedAt
	user.LastLogin = existingUser.LastLogin

	if err := s.repo.UpdateUser(user); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (s *Service) DeleteUser(userID string) error {
	if err := s.repo.DeleteUser(userID); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (s *Service) GetDecryptedAPIKeys(
	userID string,
	requestedTextProvider textgen.ProviderType,
	requestedImageProvider imagegen.ProviderType,
) (*encryption.DecryptAPIKeys, error) {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	textKey, exists := user.TextModelApiKeys[requestedTextProvider]
	if !exists {
		return nil, fmt.Errorf("no API key saved for %s", requestedTextProvider)
	}

	imageKey, exists := user.ImageModelApiKeys[requestedImageProvider]
	if !exists {
		return nil, fmt.Errorf("no API key saved for %s", requestedImageProvider)
	}

	decryptedTextKey, err := s.encryption.Decrypt(textKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt text API key: %w", err)
	}

	decryptedImageKey, err := s.encryption.Decrypt(imageKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt image API key: %w", err)
	}

	return &encryption.DecryptAPIKeys{
		TextProvider:  requestedTextProvider,
		TextAPIKey:    decryptedTextKey,
		ImageProvider: requestedImageProvider,
		ImageAPIKey:   decryptedImageKey,
	}, nil
}

func (s *Service) GetAvailableLLMProviders(userID string) (map[string][]string, error) {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	availableProviders := map[string][]string{
		"text":  {},
		"image": {},
	}

	if user.TextModelApiKeys != nil {
		for provider := range user.TextModelApiKeys {
			availableProviders["text"] = append(availableProviders["text"], string(provider))
		}
	}
	if user.ImageModelApiKeys != nil {
		for provider := range user.ImageModelApiKeys {
			availableProviders["image"] = append(availableProviders["image"], string(provider))
		}
	}

	return availableProviders, nil
}

func (s *Service) UpdateLLMProviderAPIKeys(userID string, keys *encryption.APIKeys) error {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if user.TextModelApiKeys == nil {
		user.TextModelApiKeys = make(map[textgen.ProviderType]string)
	}
	if user.ImageModelApiKeys == nil {
		user.ImageModelApiKeys = make(map[imagegen.ProviderType]string)
	}

	for provider := range keys.TextProviders {
		textProvider, err := textgen.NewProviderType(provider)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("%v", err.Error()))
		}

		encryptedTextKey, err := s.encryption.Encrypt(keys.TextProviders[string(textProvider)])
		if err != nil {
			return fmt.Errorf("failed to encrypt text API key: %w", err)
		}

		user.TextModelApiKeys[textProvider] = encryptedTextKey
	}

	for provider := range keys.ImageProviders {
		imageProvider, err := imagegen.NewProviderType(provider)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("%v", err.Error()))
		}

		encryptedImageKey, err := s.encryption.Encrypt(keys.ImageProviders[string(imageProvider)])
		if err != nil {
			return fmt.Errorf("failed to encrypt image API key: %w", err)
		}

		user.ImageModelApiKeys[imageProvider] = encryptedImageKey
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
