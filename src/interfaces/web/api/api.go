package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/paraggarg37/order_app/src/domain/models"
	"github.com/paraggarg37/order_app/src/interfaces/config"
	"github.com/paraggarg37/order_app/src/interfaces/database"
	"github.com/paraggarg37/order_app/src/usecases/definitions"
	"log"
	"net/http"
)

var router *httprouter.Router

type Interactor struct {
	OrderInteractor definitions.Order
}

type API struct {
	Cfg        *config.MainConfig
	Interactor *Interactor
	DB         database.DBRepository
}

func New(this *API) *API {
	return &API{Cfg: this.Cfg, Interactor: this.Interactor, DB: this.DB}
}

func init() {
	router = httprouter.New()
}

func (a *API) Run() {
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, error)

func ToJson(callback Handle) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		resp, err := callback(w, r, ps)
		w.Header().Set("Content-Type", "application/json")
		if err == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResp{
				Error: err.Error(),
			})
		}
	}
}
