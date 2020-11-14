package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getUserTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		email, _, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		info, err := s.Repository.GetUser(email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		token := info.Token
		c.JSON(http.StatusOK, token)
	}
}
