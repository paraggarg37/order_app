package api

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/paraggarg37/order_app/src/domain/models"
	"net/http"
	"strconv"
)

func (a *API) RegisterOrder() {
	router.POST("/orders/:id", ToJson(a.TakeOrder))
	router.POST("/orders", ToJson(a.CreateOrder))
	router.GET("/orders", ToJson(a.ListOrders))
}
func (a *API) CreateOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	ctx := r.Context()
	req := &models.OrderReq{}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, errors.New("invalid request data")
	}

	return a.Interactor.OrderInteractor.CreateOrder(ctx, req)
}

func (a *API) TakeOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {

	ctx := r.Context()
	req := &models.TakeOrderReq{}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, errors.New("invalid request data")
	}

	idStr := ps.ByName("id")

	orderID, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return nil, errors.New("invalid order id")
	}

	return a.Interactor.OrderInteractor.TakeOrder(ctx, req, orderID)
}

func (a *API) ListOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	ctx := r.Context()
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errors.New("invalid page")
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		return nil, errors.New("invalid limit")
	}

	return a.Interactor.OrderInteractor.ListOrders(ctx, page, limit)
}
