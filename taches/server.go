package todoliste

import "net/http"

func Server() {
	tasks := &[]Task{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r, tasks)
	})
	http.HandleFunc("/add-task", func(w http.ResponseWriter, r *http.Request) {
		AddTask(w, r, tasks)
	})
	http.HandleFunc("/delete-task", func(w http.ResponseWriter, r *http.Request) {
		DeleteTask(w, r, tasks)
	})
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
