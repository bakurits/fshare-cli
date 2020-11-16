package server

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/bakurits/fileshare/pkg/webapp/db"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func randToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// handlerWithUser gin handler function with user
type handlerWithUser func(user db.User, c *gin.Context)

// userExtractorMiddleware extracts user from session and passes it to handlerWithUser
func (s *Server) userExtractorMiddleware(handler handlerWithUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("email")
		if email == nil {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		u, err := s.Repository.GetUser(email.(string))
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		handler(u, c)
	}
}

func (s *Server) getLoginURL(state string) string {
	return s.AuthConfig.AuthCodeURL(state)
}

// executeTemplate executes templates with given filenames
// if withLayout is true than template executes with layout file
func (s *Server) executeTemplate(w http.ResponseWriter, data interface{}, withLayout bool, fileNames ...string) {
	tplRoot := s.StaticFileDir + "/tpls"
	var files []string
	if withLayout {
		files = append(files, tplRoot+"/layout.gohtml")
	}
	for _, file := range fileNames {
		files = append(files, tplRoot+"/"+file+".gohtml")
	}

	err := template.Must(template.ParseFiles(files...)).Execute(w, data)
	if err != nil {
		log.Println(err)
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
