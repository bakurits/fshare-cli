package server

import (
	"fmt"
	"github.com/bakurits/fileshare/pkg/auth"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/bakurits/fileshare/pkg/webapp/db"
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

	router.GET("/", s.homepageHandler())
	router.GET("/login", s.loginHandler())
	router.GET("/logout", s.logoutHandler())
	router.GET("/auth", s.authHandler())

	s.r = router
}

func (s *Server) homepageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("email") == nil {
			c.Redirect(http.StatusSeeOther, "/login")
		}
		s.executeTemplate(c.Writer, struct{}{}, true, "homepage")
	}

}

func (s *Server) loginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("email") != nil {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
		state := randToken()
		session.Set("state", state)
		_ = session.Save()
		_, _ = c.Writer.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + s.getLoginURL(state) + "'><button>Login with Google!</button> </a> </body></html>"))
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
			_ = c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
			return
		}

		client, err := s.AuthConfig.ClientFromCode(c.Query("code"))
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
		}
		session.Set("email", client.Email)
		_ = session.Save()
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func (s *Server) getLoginURL(state string) string {
	return s.AuthConfig.AuthCodeURL(state)
}

func (s *Server) executeTemplate(w http.ResponseWriter, data interface{}, withLayout bool, fileNames ...string) {
	var files []string
	if withLayout {
		files = append(files, s.StaticFileDir+"/layout.gohtml")
	}

	for _, file := range fileNames {
		files = append(files, s.StaticFileDir+"/"+file+".gohtml")
	}

	err := template.Must(template.ParseFiles(files...)).Execute(w, data)
	if err != nil {
		log.Println(err)
	}

}
