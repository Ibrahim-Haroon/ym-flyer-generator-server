package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user/model"
)

// Repository errors
var (
	ErrUserNotFound = fmt.Errorf("user not found")
	ErrUserExists   = fmt.Errorf("user already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *model.User) error {
	exists, err := r.userExists(user.Username, user.Email)
	if err != nil {
		return fmt.Errorf("error checking user existence: %w", err)
	}
	if exists {
		return ErrUserExists
	}

	textModelKeysJSON, err := json.Marshal(user.TextModelApiKeys)
	if err != nil {
		return fmt.Errorf("error marshaling text model keys: %w", err)
	}

	imageModelKeysJSON, err := json.Marshal(user.ImageModelApiKeys)
	if err != nil {
		return fmt.Errorf("error marshaling image model keys: %w", err)
	}

	now := time.Now().UTC()

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        INSERT INTO users (
            id, username, password_hash, email, created_at, updated_at,
            last_login, active_status, is_admin, text_model_api_keys,
            image_model_api_keys
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `

	_, err = tx.Exec(
		query,
		user.ID,
		user.Username,
		user.PasswordHash,
		user.Email,
		now, // created_at
		now, // updated_at
		now, // last_login
		user.ActiveStatus,
		user.IsAdmin,
		textModelKeysJSON,
		imageModelKeysJSON,
	)

	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	user.CreatedAt = now.Format(time.RFC3339)
	user.UpdatedAt = now.Format(time.RFC3339)
	user.LastLogin = now.Format(time.RFC3339)

	return nil
}

func (r *Repository) userExists(username, email string) (bool, error) {
	var exists bool
	query := `
        SELECT EXISTS(
            SELECT 1
            FROM users
            WHERE username = $1 OR email = $2
        )
    `
	err := r.db.QueryRow(query, username, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %w", err)
	}
	return exists, nil
}

func (r *Repository) scanUser(row *sql.Row) (*model.User, error) {
	var user model.User
	var textModelKeysJSON, imageModelKeysJSON []byte
	var createdAt, updatedAt, lastLogin time.Time

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.Email,
		&createdAt,
		&updatedAt,
		&lastLogin,
		&user.ActiveStatus,
		&user.IsAdmin,
		&textModelKeysJSON,
		&imageModelKeysJSON,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}

	user.CreatedAt = createdAt.Format(time.RFC3339)
	user.UpdatedAt = updatedAt.Format(time.RFC3339)
	user.LastLogin = lastLogin.Format(time.RFC3339)

	if err := json.Unmarshal(textModelKeysJSON, &user.TextModelApiKeys); err != nil {
		return nil, fmt.Errorf("error unmarshaling text model keys: %w", err)
	}
	if err := json.Unmarshal(imageModelKeysJSON, &user.ImageModelApiKeys); err != nil {
		return nil, fmt.Errorf("error unmarshaling image model keys: %w", err)
	}

	return &user, nil
}

func (r *Repository) GetUserByUsername(username string) (*model.User, error) {
	query := `
        SELECT id, username, password_hash, email, created_at, updated_at,
               last_login, active_status, is_admin, text_model_api_keys,
               image_model_api_keys
        FROM users
        WHERE username = $1
    `
	return r.scanUser(r.db.QueryRow(query, username))
}

func (r *Repository) GetUserById(userID string) (*model.User, error) {
	query := `
        SELECT id, username, password_hash, email, created_at, updated_at,
               last_login, active_status, is_admin, text_model_api_keys,
               image_model_api_keys
        FROM users
        WHERE id = $1
    `
	return r.scanUser(r.db.QueryRow(query, userID))
}

func (r *Repository) GetUserByEmail(email string) (*model.User, error) {
	query := `
        SELECT id, username, password_hash, email, created_at, updated_at,
               last_login, active_status, is_admin, text_model_api_keys,
               image_model_api_keys
        FROM users
        WHERE email = $1
    `
	return r.scanUser(r.db.QueryRow(query, email))
}

func (r *Repository) UpdateUser(user *model.User) error {
	textModelKeysJSON, err := json.Marshal(user.TextModelApiKeys)
	if err != nil {
		return fmt.Errorf("error marshaling text model keys: %w", err)
	}

	imageModelKeysJSON, err := json.Marshal(user.ImageModelApiKeys)
	if err != nil {
		return fmt.Errorf("error marshaling image model keys: %w", err)
	}

	now := time.Now().UTC()

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        UPDATE users
        SET username = $1,
            password_hash = $2,
            email = $3,
            updated_at = $4,
            active_status = $5,
            is_admin = $6,
            text_model_api_keys = $7,
            image_model_api_keys = $8
        WHERE id = $9
    `

	result, err := tx.Exec(
		query,
		user.Username,
		user.PasswordHash,
		user.Email,
		now,
		user.ActiveStatus,
		user.IsAdmin,
		textModelKeysJSON,
		imageModelKeysJSON,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rows == 0 {
		return ErrUserNotFound
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	user.UpdatedAt = now.Format(time.RFC3339)
	return nil
}

func (r *Repository) DeleteUser(userID string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
		DELETE FROM users
	 	WHERE id = $1
 	`

	result, err := tx.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rows == 0 {
		return ErrUserNotFound
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (r *Repository) UpdateLastLogin(userID string) error {
	now := time.Now().UTC()
	query := `
		UPDATE users
		SET last_login = $1
		WHERE id = $2
	`

	result, err := r.db.Exec(query, now, userID)
	if err != nil {
		return fmt.Errorf("error updating last login: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}
