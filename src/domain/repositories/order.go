package repositories

import (
	"context"
	db "github.com/paraggarg37/order_app/src/db_modles"
	"github.com/paraggarg37/order_app/src/domain/models"
	"github.com/paraggarg37/order_app/src/interfaces/database"
	"github.com/volatiletech/sqlboiler/boil"
)

type OrderRepository interface {
	database.DBRepository
	CreateOrder(ctx context.Context, req *models.OrderReq, distance int, exec boil.ContextExecutor) (*db.Order, error)
	GetOrderForUpdate(ctx context.Context, id int64, exec boil.ContextExecutor) (*db.Order, error)
	UpdateStatus(ctx context.Context, req *db.Order, status string, exec boil.ContextExecutor) (*db.Order, error)
	GetOrders(ctx context.Context, offset int, limit int, exec boil.ContextExecutor) ([]*db.Order, error)
}
