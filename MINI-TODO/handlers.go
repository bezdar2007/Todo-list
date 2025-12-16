package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

var tmp1 = template.Must(template.ParseFiles("templates/index.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Tолько GET", http.StatusMethodNotAllowed)
		return
	}
	data := struct{ Tasks []Task }{GetAllTasks()}
	tmp1.Execute(w, data)
}

func addFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Tолько POST", http.StatusMethodNotAllowed)
		return
	}
	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Пустое название", http.StatusBadRequest)
		return
	}
	AddTask(title)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetAllTasks())
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	if input.Title == "" {
		http.Error(w, "Пустое название", http.StatusBadRequest)
		return
	}
	created := AddTask(input.Title)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 || !DeleteTask(id) {
		http.Error(w, "Не найдено", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
