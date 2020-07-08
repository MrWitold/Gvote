package handlers

import (
	"net/http"

	"github.com/MrWitold/Gvote/models"
)

// UpdateVote handles PUT requests and returns one item
func (it *Items) UpdateVote(rw http.ResponseWriter, r *http.Request) {
	it.l.Println("[DEBUG] get all records")

	items := models.GetList()

	err := models.ToJSON(items, rw)
	if err != nil {
		it.l.Println("[ERROR] serializing product", err)
	}
}
