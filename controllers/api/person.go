package api

import (
	"encoding/json"
	"go_crud/db"
	"go_crud/db/model"
	"go_crud/utility"
	"net/http"

	"github.com/gorilla/mux"
)

// home page
func List(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	var student []model.Student
	db.Database.Find(&student)
	utility.SendResponse(w, r, status, student)
}

func Create(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)
	// add data int db
	db.Database.Create(&student)
	utility.SendResponse(w, r, status, student)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	var student model.Student

	db.Database.First(&student, mux.Vars(r)["id"])

	json.NewDecoder(r.Body).Decode(&student)

	db.Database.Save(&student)

	utility.SendResponse(w, r, status, student)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	var student model.Student
	db.Database.Delete(&student, mux.Vars(r)["id"])
	utility.SendResponse(w, r, status, "Successfully deleted!")
}
