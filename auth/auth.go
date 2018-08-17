package auth

type AuthData struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
	Status   int    `json:"status,omitempty"`
	Name     string `json:"name,omitempty"`
	Title    string `json:"title,omitempty"`
}

type Auth interface {
	Login(username string, password string) (AuthData, error)
	Me(token string) (AuthData, error)
}
