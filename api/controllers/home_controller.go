package controllers

import (
	"net/http"

	"github.com/dhafinkawakibi/iot_platform/api/responses"
)

// Home Controller
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
