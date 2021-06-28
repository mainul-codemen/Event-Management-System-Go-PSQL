package handler

import (
	"Event-Management-System-Go-PSQL/storage"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type (
	user struct {
		User []storage.User
	}
)

type UserFormData struct {
	CSRFField  template.HTML
	Form       storage.User
	FormErrors map[string]string
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {

	tmp := s.templates.Lookup("user_list.html")

	if tmp == nil {
		log.Println("Unable to look feedback_list.html")
		return
	}
	usr, err := s.store.GetUser()

	fmt.Printf("%+v", usr)

	if err != nil {
		log.Println("Unable to get event type.  ", err)
	}

	tempData := user{
		User: usr,
	}

	err = tmp.Execute(w, tempData)
	if err != nil {
		log.Println("Error executing tempalte:", err)
		return
	}

}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Method : Create user called.")

	data := UserFormData{
		CSRFField: csrf.TemplateField(r),
	}

	s.loadUserTemplate(w, r, data)

}

func (s *Server) saveUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("Parsing error")
	}

	var form storage.User
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln("Decoding error")
	}
	fmt.Printf("%#v", form)
	if form.FirstName == "" {
		data := UserFormData{
			CSRFField: csrf.TemplateField(r),
			Form:      form,
			FormErrors: map[string]string{
				"FirstName": "First Name is required.",
			},
		}
		s.loadUserTemplate(w, r, data)
	}
	id, err := s.store.CreateUser(form)
	if err != nil {
		log.Fatalln("Unable to save data :", err)

	}
	fmt.Printf("%#v", id)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (s *Server) loadUserTemplate(w http.ResponseWriter, r *http.Request, form UserFormData) {
	tmpl := s.templates.Lookup("user-form.html")
	if tmpl == nil {
		log.Println("Unable to find form")
		return
	}
	if err := tmpl.Execute(w, form); err != nil {
		log.Println("Error executing tempalte : ", err)
		return
	}
}