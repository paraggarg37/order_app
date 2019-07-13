package models

import (
	"errors"
	db "github.com/paraggarg37/order_app/src/db_modles"
	"strconv"
)

const (
	TAKE_ORDER_SUCCESS = "SUCCESS"
)

type OrderReq struct {
	Origin      []string `json:"origin"`
	Destination []string `json:"destination"`
}

type OrderResp struct {
	Id       int64  `json:"id"`
	Distance int    `json:"distance"`
	Status   string `json:"status"`
}

type TakeOrderResp struct {
	Status string `json:"status"`
}

type TakeOrderReq struct {
	Status string `json:"status"`
}

func (t *TakeOrderReq) Validate() error {
	if t.Status != db.OrdersStatusTAKEN {
		return errors.New("invalid status")
	}
	return nil
}

type ErrorResp struct {
	Error string `json:"error"`
}

func (o *OrderReq) Validate() error {
	if len(o.Origin) != 2 {
		return errors.New("invalid origin arguments")
	}
	if len(o.Destination) != 2 {
		return errors.New("invalid destination arguments")
	}

	if err := validateLatLng(o.Origin[0], o.Origin[1]); err != nil {
		return errors.New("invalid origin: " + err.Error())
	}

	if err := validateLatLng(o.Destination[0], o.Destination[1]); err != nil {
		return errors.New("invalid destination: " + err.Error())
	}
	return nil
}

func validateLatLng(lat string, lng string) error {

	if oLat, err := strconv.ParseFloat(lat, 64); err != nil {
		return errors.New("not a valid origin latitude")
	} else if oLang, err := strconv.ParseFloat(lng, 64); err != nil {
		return errors.New("not a valid origin longitude")
	} else {
		if oLat < -90 || oLat > 90 {
			return errors.New("latitude value must be between -90 and 90")
		}

		if oLang < -180 || oLang > 180 {
			return errors.New("longitude value must be between -180 and 180")
		}
	}

	return nil
}
