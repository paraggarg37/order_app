package dao

import (
	"context"
	db "github.com/paraggarg37/order_app/src/db_modles"
	"github.com/paraggarg37/order_app/src/domain/models"
	"github.com/paraggarg37/order_app/src/interfaces/database"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type OrderRepo struct {
	database.DBRepository
}

func New(db database.DBRepository) *OrderRepo {
	return &OrderRepo{db}
}

func (o *OrderRepo) CreateOrder(ctx context.Context, req *models.OrderReq, distance int, exec boil.ContextExecutor) (*db.Order, error) {
	order := &db.Order{
		Distance:       distance,
		OriginLat:      req.Origin[0],
		OriginLNG:      req.Origin[1],
		DestinationLat: req.Destination[0],
		DestinationLNG: req.Destination[1],
		Status:         db.OrdersStatusUNASSIGNED,
	}

	err := order.Insert(ctx, exec, boil.Infer())

	return order, err
}

func (o *OrderRepo) UpdateStatus(ctx context.Context, req *db.Order, status string, exec boil.ContextExecutor) (*db.Order, error) {
	req.Status = status
	_, err := req.Update(ctx, exec, boil.Infer())
	return req, err
}

func (o *OrderRepo) GetOrderForUpdate(ctx context.Context, id int64, exec boil.ContextExecutor) (*db.Order, error) {
	return db.Orders(qm.Where("id = ?", id), qm.For("UPDATE")).One(ctx, exec)
}

func (o *OrderRepo) GetOrders(ctx context.Context, offset int, limit int, exec boil.ContextExecutor) ([]*db.Order, error) {

	return db.Orders(qm.OrderBy("id asc"), qm.Offset(offset), qm.Limit(limit)).All(ctx, exec)
}
