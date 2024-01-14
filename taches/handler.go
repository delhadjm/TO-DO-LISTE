package todoliste

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AddTask(w http.ResponseWriter, r *http.Request, tasks *[]Task) {
	task := Task{
		Id:          GetLastId(*tasks) + 1,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		State:       r.FormValue("state"),
		Deadline:    r.FormValue("date"),
	}
	*tasks = append(*tasks, task)
	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request, tasks *[]Task) {
	idToCompare, _ := strconv.Atoi(r.FormValue("id"))
	for i, task := range *tasks {
		if task.Id == idToCompare {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func HomeHandler(w http.ResponseWriter, r *http.Request, tasks *[]Task) {
	tmpl, err := template.ParseFiles("./pages/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, tasks)
}
