package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	err := render.DecodeJSON(r.Body, &request)
	println("printando")
	println(request.Name)

	email := r.Context().Value("email").(string)
	request.CreatedBy = email

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err
}
