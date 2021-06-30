package handler

import (
	"Event-Management-System-Go-PSQL/storage"
	"html/template"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
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
	UnableToFindHtmlTemplate(tmp)
	usr, err := s.store.GetUser()
	UnableToGetData(err)
	tempData := user{
		User: usr,
	}
	err = tmp.Execute(w, tempData)
	ExcutionTemplateError(err)

}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Method : Create user called.")
	data := UserFormData{
		CSRFField: csrf.TemplateField(r),
	}
	s.loadUserTemplate(w, r, data)

}

func (s *Server) saveUser(w http.ResponseWriter, r *http.Request) {
	ParseFormData(r)
	// decode form data
	var form storage.User
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln("Decoding error")
	}
	// validation
	if err := form.Validate(); err != nil {
		vErrs := map[string]string{}
		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					vErrs[key] = value.Error()
				}
			}
		}

		data := UserFormData{
			CSRFField:  csrf.TemplateField(r),
			Form:       form,
			FormErrors: vErrs,
		}
		s.loadUserTemplate(w, r, data)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	/* hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 8)
	form.Password = string(hashedPassword) */

	// call database query
	_, err := s.store.CreateUser(form)
	UnableToInsertData(err)
	http.Redirect(w, r, "/event", http.StatusSeeOther)
}

func (s *Server) loadUserTemplate(w http.ResponseWriter, r *http.Request, form UserFormData) {
	tmpl := s.templates.Lookup("user-form.html")
	UnableToFindHtmlTemplate(tmpl)
	err := tmpl.Execute(w, form)
	ExcutionTemplateError(err)
}
