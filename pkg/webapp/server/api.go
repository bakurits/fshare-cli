package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getUserTokenHandler() gin.HandlerFunc {
	type LoginParameters struct {
		Email    string
		Password string
	}

	var user LoginParameters
	return func(c *gin.Context) {
		if err := c.BindJSON(&user); err == nil {
			info, err := s.Repository.GetUser(user.Email)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{})
				return
			}
			token := info.Token
			c.JSON(http.StatusOK, gin.H{"user": user.Email, "password": user.Password, "token": token})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{})
	}
}
