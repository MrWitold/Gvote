package handlers

import (
	"net/http"

	"github.com/MrWitold/Gvote/models"
	"github.com/gorilla/mux"
)

// ShowAll handles GET requests and returns all current items
func (it *Items) ShowAll(rw http.ResponseWriter, r *http.Request) {
	it.l.Println("[DEBUG] get all records")

	items := models.GetList()

	err := models.ToJSON(items, rw)
	if err != nil {
		it.l.Println("[ERROR] serializing product", err)
	}
}

// ShowSingle handles GET requests and returns one item
func (it *Items) ShowSingle(rw http.ResponseWriter, r *http.Request) {
	it.l.Println("[DEBUG] get all records")

	vars := mux.Vars(r)
	id, _ := vars["id"]

	items := models.GetSingleItme(id)

	err := models.ToJSON(items, rw)
	if err != nil {
		it.l.Println("[ERROR] serializing product", err)
	}
}
