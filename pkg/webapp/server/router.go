package server

import (
	"net/http"

	"github.com/bakurits/fileshare/pkg/auth"
	"github.com/bakurits/fileshare/pkg/webapp/db"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Server dependencies for endpoints
type Server struct {
	r *gin.Engine

	AuthConfig *auth.Config

	Repository    db.Repository
	StaticFileDir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

// Init initializes server
func (s *Server) Init() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("fileshare", store))

	router.GET("/", s.userExtractorMiddleware(s.homepageHandler()))
	router.GET("/login", s.loginHandler())
	router.GET("/logout", s.logoutHandler())
	router.GET("/auth", s.authHandler())
	router.POST("/get-user-token", s.getUserTokenHandler())
	s.r = router
}
