package handlers

import (
	// "encoding/json"
	// "fmt"
	"github.com/7IBBE77S/web-app/internal/config"
	"github.com/7IBBE77S/web-app/internal/models"
	"github.com/7IBBE77S/web-app/internal/render"
	// "log"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}
// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
	
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w,r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}


