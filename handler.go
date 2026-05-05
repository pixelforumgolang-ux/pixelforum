package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	

)

var oneUser User
var login bool = false

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		page := r.FormValue("page")
		if page != "" {
			http.Redirect(w, r, page, http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func homeHandlerIndex(w http.ResponseWriter, r *http.Request) {

	var listePost []Post
	readPostAll(listePost)

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, listePost)
}

func homeHandlerLogin(w http.ResponseWriter, r *http.Request) {
	
	var verif bool = verifLoginPage(r.FormValue("mail"),r.FormValue("mdp"))

	if (verif){
		result := db.Where("UserMdp = ?", r.FormValue("mdp")).First(&theUser)
		login = true
	} else {
		fmt.Println("information érroné")
	}

	if (verif){
		createUser(r.FormValue("userName"),r.FormValue("firstName"),r.FormValue("lastName"),r.FormValue("userMail"),r.FormValue("userMdp"))
	} else {
		fmt.Println("information érroné")
	}
	

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/login.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandlerUser(w http.ResponseWriter, r *http.Request) {

	if (login){
		switch r.FormValue("action") {
			case "disconnect":
				disconnect(theUser.id)
				login = false
			case "update":
				updateUser(theUser.id,r.FormValue("userName"), r.FormValue("firstName"),r.FormValue("lastName"),r.FormValue("userMail"),r.FormValue("userMdp"), "client")
			case "delete":
				deleteUser(theUser.id)
		}
	}
	
	
	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/account.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if (theUser == nil) {
		tpl.Execute(w, nil)
	} else {
		tpl.Execute(w, theUser)
	}
	
}

func homeHandlerAbout(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/aboutUsPage.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandlerPost(w http.ResponseWriter, r *http.Request) {
	var subject PostPage
	var id uint = 0
	MakeStructPostPage(id, subject)
	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/subject.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, subject)
}

func homeHandlerAdmin(w http.ResponseWriter, r *http.Request) {
	var alltab AdminPage
	MakeStructAdminPage(alltab)
	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/adminPage.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tpl.Execute(w, alltab)
}

