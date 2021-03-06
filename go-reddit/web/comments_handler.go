package web

import (
	"net/http"

	goreddit "devisions.org/go-reddit"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type CommentsHandler struct {
	store    goreddit.Store
	sessions *scs.SessionManager
}

func (h *CommentsHandler) Save() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		form := CreateCommentForm{
			Content: r.FormValue("content"),
		}
		if !form.Validate() {
			h.sessions.Put(r.Context(), "form", form)
			http.Redirect(w, r, r.Referer(), http.StatusFound)
			return
		}

		pidStr := chi.URLParam(r, "postID")
		pid, err := uuid.Parse(pidStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.store.SaveComment(&goreddit.Comment{
			ID:      uuid.New(),
			PostID:  pid,
			Content: form.Content,
			Votes:   0,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.sessions.Put(r.Context(), "flash", "Your comment has been submitted.")
		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}

func (h *CommentsHandler) Vote() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c, err := h.store.GetComment(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		dir := r.URL.Query().Get("dir")
		if dir == "up" {
			c.Votes++
		} else if dir == "down" {
			c.Votes--
		}

		if err := h.store.UpdateComment(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
