package definitions

import (
	"context"
	"github.com/paraggarg37/order_app/src/domain/models"
)

type Order interface {
	CreateOrder(ctx context.Context, req *models.OrderReq) (*models.OrderResp, error)
	TakeOrder(ctx context.Context, req *models.TakeOrderReq, id int64) (*models.TakeOrderResp, error)
	ListOrders(ctx context.Context, page int, limit int) ([]*models.OrderResp, error)
}
