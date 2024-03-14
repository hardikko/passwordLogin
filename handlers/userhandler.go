package handlers

import (
	"encoding/json"
	"fmt"
	"learngo/helpers"
	"learngo/models"
	"learngo/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// user create handler
func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.User{}

	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	obj, err := service.UserCreateService(ctx, req)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	render.JSON(w, r, obj)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := helpers.StringToInt64(idStr)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	obj, err := service.UserService(ctx, id)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	render.JSON(w, r, obj)
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := service.UserListService(ctx)
	if err != nil {
		render.JSON(w, r, err.Error())
	}

	// This will also work
	// result, err := userListStore(ctx)
	// if err != nil {
	// 	render.JSON(w, r, err.Error())
	// }

	render.JSON(w, r, result)
}
