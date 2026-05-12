package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	

)

var theUser User
var login bool = false

func redirectHandler(w http.ResponseWriter, r *http.Request) { // fini
	if r.Method == http.MethodPost {
		page := r.FormValue("page")
		if page != "" {
			http.Redirect(w, r, page, http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func homeHandlerIndex(w http.ResponseWriter, r *http.Request) { //fini

	var listePost []Post
	readPostAll(&listePost)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	tpl.Execute(w, listePost)
}

func homeHandlerLogin(w http.ResponseWriter, r *http.Request) { // fini
	
	if r.Method == http.MethodPost {
		
		switch r.FormValue("action") {
			case "connect":
				login(r.FormValue("loginMail"),r.FormValue("loginMdp"), &theUser)
			case "create":
				createActalUser(r.FormValue("userName"),r.FormValue("firstName"),r.FormValue("lastName"),r.FormValue("userMail"),r.FormValue("userMdp"),"client")

		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/login.html"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandlerUser(w http.ResponseWriter, r *http.Request) { //fini
	
	if r.Method == http.MethodPost {
		
		switch r.FormValue("action") {
			case "disconnect":
				disconnect(&theUser,&login)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			case "update":
				UpdateActualUser(theUser.id,r.FormValue("userName"), r.FormValue("firstName"),r.FormValue("lastName"),r.FormValue("userMail"),r.FormValue("userMdp"), "client", login)
				http.Redirect(w, r, "/account", http.StatusSeeOther)
				return
			case "delete":
				deleteActualUser(&theUser,&login)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
		} 
	}
	
	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/account.html"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	if (login == true) {
		tpl.Execute(w, nil)
	} else {
		tpl.Execute(w, theUser)
	}
	
}

func homeHandlerAbout(w http.ResponseWriter, r *http.Request) { //fini

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/aboutUsPage.html"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandlerError(w http.ResponseWriter, r *http.Request) { //fini
	
	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/Error"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandlerMakePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if(login == true){
			switch r.FormValue("action") {
				case "PostSomething":
					createPost(theUser.id,r.FormValue("subject "),r.FormValue("message"),r.FormValue("image"),Date.now())

			}

			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}
	}
	
	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/makeSubject.html"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandlerPost(w http.ResponseWriter, r *http.Request) {
	var subject PostPage
	id = r.URL.Query().Get("id")
	
	readPost(id, &thePost)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/subject.html"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	tpl.Execute(w, subject)
}

func homeHandlerAdmin(w http.ResponseWriter, r *http.Request) {
	
	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/adminPage.html"))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	tpl.Execute(w, nil)
}

