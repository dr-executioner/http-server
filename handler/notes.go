package handler

import (
	"encoding/json"
	"fmt"
	"http-server/model"
	"http-server/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var notes []model.Note

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriterError(w, http.StatusMethodNotAllowed, "Only POST allowed on this route")
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Title == "" || input.Content == "" {
		utils.WriterError(w, http.StatusBadRequest, "Invalid Input")
		return
	}

	note := model.Note{
		ID:      uuid.New().String(),
		Title:   input.Title,
		Content: input.Content,
	}

	notes = append(notes, note)
	utils.WriteJSON(w, http.StatusCreated, note)
}

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriterError(w, http.StatusMethodNotAllowed, "Only GET allowed on this route")
		return
	}

	w.Header().Set("Content-Type", "text/html")

	var html strings.Builder
	for _, note := range notes {
		html.WriteString(fmt.Sprintf("<div><h3>%s</h3><p>%s</p></div><hr>", note.Title, note.Content))
	}
	fmt.Fprintf(w, html.String())
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriterError(w, http.StatusMethodNotAllowed, "MEthod not allowed")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/notes/")
	if id == "" {
		utils.WriterError(w, http.StatusBadRequest, "Missing ID")
		return
	}

	for _, note := range notes {
		if note.ID == id {
			utils.WriteJSON(w, http.StatusOK, note)
			return
		}
	}

	utils.WriterError(w, http.StatusNotFound, "Note not found")
}

func DeleteNoteById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.WriterError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/notes/")
	if id == "" {
		utils.WriterError(w, http.StatusBadRequest, "Missing ID")
		return
	}

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			utils.WriteJSON(w, http.StatusOK, map[string]string{
				"message": "Note Deleted",
			})

			return
		}
		utils.WriterError(w, http.StatusNotFound, "Note not found")
	}
}
