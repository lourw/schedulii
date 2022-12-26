package middleware

import (
	"fmt"
	"os"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var engine *gin.Engine
var session sessions.Session

func setup() {
	engine = gin.Default()
	store := cookie.NewStore([]byte("secret"))

	engine.Use(sessions.Sessions("schedulii", store))

	engine.GET("/", func(c *gin.Context) {
		session = sessions.Default(c)
		session.Set("user", "hello")
		err := session.Save()
		if err != nil {
			fmt.Print("Session save error")
		}
	})

	auth := engine.Group("/auth")
	auth.Use(CheckAuthenticated)
	{
		auth.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World")
		})
	}
}

func TestCheckAuthenticated_ValidUser(t *testing.T) {
	w := httptest.NewRecorder()
	session_req, _ := http.NewRequest("GET", "/", nil)
	auth_req, _ := http.NewRequest("GET", "/auth", nil)
	engine.ServeHTTP(w, session_req)
	engine.ServeHTTP(w, auth_req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuthenticated_InvalidUser(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/auth", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
