package models

type Role string

const (
	User      Role = "user"
	Assistant Role = "assistant"
)

type TextHistory struct {
	Role    Role
	Content string
}
