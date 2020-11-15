package server

import (
	"fmt"
	"net/http"

	"github.com/bakurits/fileshare/pkg/webapp/db"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Server) homePageHandler() handlerWithUser {

	type Homepage struct {
		Email         string
		IsPasswordSet bool
	}

	return func(user db.User, c *gin.Context) {
		s.executeTemplate(c.Writer, Homepage{Email: user.Email, IsPasswordSet: user.Password != ""}, true, "homepage")
	}

}

func (s *Server) loginPageHandler() gin.HandlerFunc {

	type LoginResponse struct {
		AuthLink string
	}

	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("email") != nil {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
		state := randToken()
		session.Set("state", state)
		_ = session.Save()

		s.executeTemplate(c.Writer, LoginResponse{AuthLink: s.getLoginURL(state)}, true, "login")
	}
}

func (s *Server) logoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("email")
		_ = session.Save()
		c.Redirect(http.StatusSeeOther, "/")
	}

}

func (s *Server) authHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		retrievedState := session.Get("state")
		if retrievedState != c.Query("state") {
			_ = c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", retrievedState))
			return
		}

		client, err := s.AuthConfig.ClientFromCode(c.Query("code"))
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		_ = s.Repository.AddUser(db.User{
			Email:    client.Email,
			Password: "",
			Token:    db.TokenStore(*client.Token),
		})

		session.Set("email", client.Email)
		_ = session.Save()
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func (s *Server) getEmailFromPasswordRecoveryRequest(c *gin.Context) string {
	session := sessions.Default(c)
	email := ""
	if token := c.Query("token"); token != "" {
		info, err := s.Repository.GetPasswordRestoreInfo(token)
		if err == nil {
			email = info.Email
		}
	} else {
		email = session.Get("email").(string)
	}
	return email
}

func (s *Server) setPasswordHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := s.getEmailFromPasswordRecoveryRequest(c)
		if email == "" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		session.Set("emil", email)
		_ = session.Save()

		s.executeTemplate(c.Writer, struct{}{}, true, "password_recovery")
	}
}
