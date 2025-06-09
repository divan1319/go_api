package handlers

import (
	"encoding/json"
	"goapi/internal/models"
	"net/http"
	"sync"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextId   = 1
)

func init() {
	teachers[nextId] = models.Teacher{
		Id:        nextId,
		FirstName: "Mark",
		LastName:  "Smith",
		Class:     "10th",
		Subject:   "Math",
	}
	nextId++
	teachers[nextId] = models.Teacher{
		Id:        nextId,
		FirstName: "Kate",
		LastName:  "Smith",
		Class:     "10th",
		Subject:   "English",
	}
	nextId++
	teachers[nextId] = models.Teacher{
		Id:        nextId,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "10th",
		Subject:   "Science",
	}
	nextId++
}

func TeachersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//call get method fn
		getTeachers(w, r)
	case http.MethodPost:

		addTeacherHandler(w, r)
	case http.MethodPut:
		w.Write([]byte("put teacher method"))
	case http.MethodPatch:
		w.Write([]byte("patch teacher method"))
	case http.MethodDelete:
		w.Write([]byte("delete teacher method"))
	}

}

func getTeachers(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	teacherList := make([]models.Teacher, 0, len(teachers))
	for _, teacher := range teachers {
		if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
			teacherList = append(teacherList, teacher)
		}
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teachers),
		Data:   teacherList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []models.Teacher

	err := json.NewDecoder(r.Body).Decode(&newTeachers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addedTeacher := make([]models.Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		newTeacher.Id = nextId
		teachers[nextId] = newTeacher
		addedTeacher[i] = newTeacher
		nextId++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeacher),
		Data:   addedTeacher,
	}

	json.NewEncoder(w).Encode(response)
}
