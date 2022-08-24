package handler

import (
	"net/http"

	model "github.com/Precious025/go-web_app/models"
	"github.com/Precious025/go-web_app/pkg/config"
	"github.com/Precious025/go-web_app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplates(w, "home.page.tmpl", &model.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"

	m.App.Session.GetString(r.Context(), "remote_ip")
	render.RenderTemplates(w, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})
}
