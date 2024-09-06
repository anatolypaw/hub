package authservice

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("неверный логин/пароль")
	ErrNotAuthorized      = errors.New("not authorized")
	ErrNoPermission       = errors.New("no permission")
)

type user struct {
	passHash   string
	permission string
}

type Auth struct {
	userStore map[string]user

	// Ключ - токен, значение - имя пользователя, которому токен выдан
	authTokenStore map[string]string
}

// Создает сервис авторизации и управления пользователями
func New() Auth {
	users := make(map[string]user)
	session := make(map[string]string)

	auth := Auth{
		userStore:      users,
		authTokenStore: session,
	}

	return auth
}

// Добавляет пользователя
func (a *Auth) AddUser(username, password, permission string) error {
	if username == "" || password == "" || permission == "" {
		return ErrInvalidCredentials
	}

	a.userStore[username] = user{
		passHash:   password,
		permission: permission,
	}

	return nil
}

// Аутентификация пользовтеля, возвращает session id
// fixme: токены не удаляются по таймауту, можно всю память закончить
func (a *Auth) Login(username, password string) (string, error) {
	user, ok := a.userStore[username]
	if !ok {
		return "", ErrInvalidCredentials
	}

	if user.passHash != password {
		return "", ErrInvalidCredentials
	}

	s, err := generateSessionID()
	if err != nil {
		return "", err
	}

	a.authTokenStore[s] = username
	return s, nil
}

// Авторизация пользователя и проверка наличия у пользрователя требуемых прав
func (a *Auth) Authorize(session string, permission []string) error {
	username, ok := a.authTokenStore[session]
	if !ok {
		return ErrNotAuthorized
	}

	user, ok := a.userStore[username]
	if !ok {
		return ErrNotAuthorized
	}

	// проверка прав пользователя
	for _, p := range permission {
		if user.permission == p {
			return nil
		}
	}
	return ErrNoPermission
}

// Получить имя пользователя по
func (a *Auth) GetUsernameByAuthToken(authToken string) string {
	username, ok := a.authTokenStore[authToken]
	if !ok {
		return "not found"
	}
	return username
}

// generateSessionID generates a secure random session ID
func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
