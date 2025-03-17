package routes

import (
	"github.com/kitshy/treasure-backend/api/common"
	"github.com/kitshy/treasure-backend/api/models"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
)

func (h Routes) GetDepositTokensListHandler(w http.ResponseWriter, r *http.Request) {
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")
	order := r.URL.Query().Get("order")

	page, _ := strconv.Atoi(pageQuery)
	size, _ := strconv.Atoi(pageSizeQuery)
	request := models.QueryListParams{
		Page:     page,
		PageSize: size,
		Order:    order,
	}
	resp, err := h.svc.GetDepositTokensList(&request)
	if err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	err = common.JsonResponse(w, resp, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}
