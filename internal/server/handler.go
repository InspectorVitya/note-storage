package httpserver

import (
	"encoding/json"
	"github.com/inspectorvitya/note-storage/internal/model"
	"net/http"
	"strconv"
)

func (s *Server) Main(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.GetAll(w, r)
	case http.MethodPost:
		s.CreateNote(w, r)
	case http.MethodDelete:
		s.DeleteNote(w, r)
	default:
		w.WriteHeader(404)
	}

}

func (s *Server) CreateNote(w http.ResponseWriter, r *http.Request) {
	note := model.Note{}
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err := s.App.CreateNote(r.Context(), note)
	w.WriteHeader(201)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
	}
}

func (s *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := s.App.GetAllNotes(r.Context())
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(notes)
}

func (s *Server) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = s.App.DeleteNote(r.Context(), model.IDNote(id))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}
func (s *Server) GetLast(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		note, err := s.App.GetLastNote(r.Context())
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(note)
		return
	}
	w.WriteHeader(404)
}
func (s *Server) GetFirs(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		note, err := s.App.GetFirstNote(r.Context())
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(note)
		return
	}
	w.WriteHeader(404)
}
