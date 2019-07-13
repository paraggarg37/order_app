package order

import (
	"context"
	"errors"
	db "github.com/paraggarg37/order_app/src/db_modles"
	"github.com/paraggarg37/order_app/src/domain/models"
	"github.com/paraggarg37/order_app/src/domain/repositories"
	"github.com/paraggarg37/order_app/src/interfaces/distancematrix"
)

type ordersInteractor struct {
	DistanceService distancematrix.DistanceService
	OrderRepo       repositories.OrderRepository
}

func Init(d distancematrix.DistanceService, repo repositories.OrderRepository) *ordersInteractor {
	return &ordersInteractor{DistanceService: d, OrderRepo: repo}
}

func (o *ordersInteractor) CreateOrder(ctx context.Context, req *models.OrderReq) (*models.OrderResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	distance, err := o.DistanceService.CalculateDistance(ctx, req.Origin, req.Destination)

	if err != nil {
		return nil, err
	}

	order, err := o.OrderRepo.CreateOrder(ctx, req, distance, o.OrderRepo.GetMaster())

	if err != nil {
		return nil, err
	}

	return &models.OrderResp{
		Distance: order.Distance,
		Status:   order.Status,
		Id:       order.ID,
	}, err

}

func (o *ordersInteractor) TakeOrder(ctx context.Context, req *models.TakeOrderReq, id int64) (*models.TakeOrderResp, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	dbInstance := o.OrderRepo

	tx, err := dbInstance.GetTransaction(ctx)

	if err != nil {
		return nil, err
	}

	defer dbInstance.RollbackTransaction(tx)

	order, err := o.OrderRepo.GetOrderForUpdate(ctx, id, tx)

	if err != nil {
		return nil, err
	}

	if order.Status != db.OrdersStatusUNASSIGNED {
		return nil, errors.New("order is already " + order.Status)
	}

	_, err = o.OrderRepo.UpdateStatus(ctx, order, req.Status, tx)

	if err != nil {
		return nil, err
	}

	err = dbInstance.CommitTransaction(tx)

	if err != nil {
		return nil, err
	}

	return &models.TakeOrderResp{
		Status: models.TAKE_ORDER_SUCCESS,
	}, err

}

func (o *ordersInteractor) ListOrders(ctx context.Context, page int, limit int) ([]*models.OrderResp, error) {

	offset := (page - 1) * limit
	data, err := o.OrderRepo.GetOrders(ctx, offset, limit, o.OrderRepo.GetSlave())

	if err != nil {
		return nil, err
	}

	resp := make([]*models.OrderResp, 0)

	for _, i := range data {
		resp = append(resp, &models.OrderResp{
			Status:   i.Status,
			Distance: i.Distance,
			Id:       i.ID,
		})
	}

	return resp, err
}
