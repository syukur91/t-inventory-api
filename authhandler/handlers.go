package authhandler

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/syukur91/t-inventory-api/auth"
)

type Handler struct {
	AuthClient auth.Auth
	Log        *log.Entry
}

// SetRoutes is
func (h *Handler) SetRoutes(r *echo.Group) {

	r.POST("/login", h.login)
	r.POST("/me", h.me)

}

func (h *Handler) login(c echo.Context) error {

	authData := new(auth.AuthData)
	err := c.Bind(authData)

	auth, err := h.AuthClient.Login(authData.UserName, authData.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, auth)
}

func (h *Handler) me(c echo.Context) error {

	authData := new(auth.AuthData)
	err := c.Bind(authData)

	auth, err := h.AuthClient.Me(authData.Token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, auth)
}
