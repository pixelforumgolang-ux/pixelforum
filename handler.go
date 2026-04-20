package main

import (
	"html/template"
	"net/http"
	"strconv"
	"path/filepath"
	"time"

)

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

func homeHandler(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandler1(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/loginPage.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	
	tpl.Execute(w, nil)
}

func homeHandler2(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/aboutUsPage.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandler3(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/subject.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandler4(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles(filepath.Join(templateDir, "page/adminPage.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tpl.Execute(w, nil)
}