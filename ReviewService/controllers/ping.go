package controllers

import (
	"ReviewService/utils"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Server is running", map[string]string{"status": "pong"}) // No error expected here
}
