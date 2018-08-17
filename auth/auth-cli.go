package auth

import (
	log "github.com/sirupsen/logrus"
)

type AuthCLI struct {
	Log *log.Entry
}

func NewAuthCLI(log *log.Entry) *AuthCLI {

	client := &AuthCLI{
		Log: log,
	}

	return client
}

func (s *AuthCLI) Login(username string, password string) (AuthData, error) {

	data := AuthData{Status: 200, Token: "my-token", UserName: username}

	return data, nil
}

func (s *AuthCLI) Me(token string) (AuthData, error) {

	data := AuthData{Name: "doggo", Title: "sir"}

	return data, nil
}
