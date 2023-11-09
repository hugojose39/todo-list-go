package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"todo-list/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()
	temp.ExecuteTemplate(w, "Index", tasks)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		description := r.FormValue("description")
		date := r.FormValue("date")
		done := r.FormValue("done")

		doneBool, err := strconv.ParseBool(done)

		if err != nil {
			log.Println("Erro na conversão da tarefa concluida", err)
		}

		models.CreateTask(description, date, doneBool)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("id")

	models.DeleteTask(idTask)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("id")
	task := models.EditTask(idTask)
	temp.ExecuteTemplate(w, "Edit", task)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		description := r.FormValue("description")
		date := r.FormValue("date")
		done := r.FormValue("done")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		doneBool, err := strconv.ParseBool(done)

		if err != nil {
			log.Println("Erro na conversão da tarefa concluida", err)
		}

		models.UpdateTask(idInt, description, date, doneBool)
	}

	http.Redirect(w, r, "/", 301)
}
