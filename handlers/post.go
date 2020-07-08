package handlers

import (
	"net/http"

	"github.com/MrWitold/Gvote/models"
)

// CreateVote handles POST requests and returns one item
func (it *Items) CreateVote(rw http.ResponseWriter, r *http.Request) {
	it.l.Println("[DEBUG] get all records")

	items := models.GetList()

	err := models.ToJSON(items, rw)
	if err != nil {
		it.l.Println("[ERROR] serializing product", err)
	}
}
