package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sayed-imran/go-design-pattern/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.APIUser
	_ = json.NewDecoder(r.Body).Decode(&user)

}

