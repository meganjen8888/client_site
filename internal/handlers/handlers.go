package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/meganjen8888/acupuncture/internal/config"
	"github.com/meganjen8888/acupuncture/internal/models"
	"github.com/meganjen8888/acupuncture/internal/render"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Services Reservation renders the make a reservation page and displays form
func (m *Repository) MakeAppointment(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-appointment.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	date := r.Form.Get("start")
	w.Write([]byte(fmt.Sprintf("start date is %s", date)))
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Services(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "services.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Head(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "head.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Hand(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "hands.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Legs(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "legs.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Other(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "other.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Torso(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "torso.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Feet(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "feet.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "searchavailability.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//renders search availability page
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/")
	w.Write(out)
}
