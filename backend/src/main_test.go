package main

import (
	"net/http"
	"net/http/httptest"
	"schedulii/src/models"
	router "schedulii/src/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	ginEngine := gin.Default()
	router.SetupRoutes(ginEngine, &models.Env{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	ginEngine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello World", w.Body.String())
}
