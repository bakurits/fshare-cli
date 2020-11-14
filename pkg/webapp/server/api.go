package server

import (
	"net/http"

	"github.com/bakurits/ph"
	"github.com/gin-gonic/gin"
)

func (s *Server) getUserTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		user, err := s.Repository.GetUser(email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		if !ph.Compare(user.Password, password) {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		c.JSON(http.StatusOK, user.Token)
	}
}
