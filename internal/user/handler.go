package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xyersh/examle-REST-app/internal/handlers"
)

type handler struct {
}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartialUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("This is the users' list"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Got user by UUID"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("User's creating"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Full update of the user"))
}

func (h *handler) PartialUpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Partial update of the user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("DElete the user"))
}
