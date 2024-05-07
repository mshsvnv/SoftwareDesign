package mypostgres

import (
	"context"
	"src_new/internal/model"
	"src_new/internal/repository"
	"src_new/pkg/storage/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	*postgres.Postgres
}

func NewOrderRepository(db *postgres.Postgres) repository.IOrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(ctx context.Context, order *model.Order) error {

	query := r.Builder.
		Insert(orderTable).
		Columns(
			userIDField,
			deliveryDateField,
			addressField,
			recepientNameField,
			statusField,
			totalPriceField).
		Values(
			order.UserID,
			order.OrderInfo.DeliveryDate,
			order.OrderInfo.Address,
			order.OrderInfo.RecepientName,
			order.Status,
			order.TotalPrice).
		Suffix("returning id")

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(
		&order.ID,
	)

	if err != nil {
		return err
	}

	for _, line := range order.Lines {

		query := r.Builder.
			Insert(orderRacketTable).
			Columns(
				racketIDField,
				orderIDField,
				quantityField).
			Values(
				line.RacketID,
				order.ID,
				line.Quantity).
			Suffix("returning order_id")

		sql, args, err := query.ToSql()

		if err != nil {
			return err
		}

		row := r.Pool.QueryRow(ctx, sql, args...)

		err = row.Scan(
			&line.OrderID,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) Update(ctx context.Context, order *model.Order) error {

	query := r.Builder.
		Update(orderTable).
		Set(statusField, order.Status).
		Where(squirrel.And{squirrel.Eq{
			orderIDField: order.ID,
			userIDField:  order.UserID,
		}})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, orderID int) error {

	query := r.Builder.
		Delete(orderTable).
		Where(squirrel.Eq{orderIDField: orderID})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetMyOrders(ctx context.Context, userID int) ([]*model.Order, error) {

	query := r.Builder.
		Select("*").
		From(orderTable).
		Where(squirrel.Eq{userIDField: userID})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	var orders []*model.Order

	for rows.Next() {

		order, err := r.rowToModel(rows)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) Remove(ctx context.Context, orderID int) error {

	query := r.Builder.
		Delete(orderTable).
		Where(squirrel.Eq{idField: orderID})

	sql, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, orderID int) (*model.Order, error) {

	query := r.Builder.
		Select("*").
		From(orderTable).
		Where(squirrel.Eq{idField: orderID})

	sql, args, err := query.ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	order, err := r.rowToModel(row)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) rowToModel(row pgx.Row) (*model.Order, error) {

	var order model.Order
	var info model.OrderInfo

	err := row.Scan(
		&order.ID,
		&order.UserID,
		&info.DeliveryDate,
		&info.Address,
		&info.RecepientName,
		&order.Status,
		&order.TotalPrice,
	)

	order.OrderInfo = &info

	if err != nil {
		return nil, err
	}

	return &order, nil
}

// func (r *OrderRepository) rowToModelOrderRacket(row pgx.Row) (*model.OrderLine, error) {

// 	var orderLine model.OrderLine

// 	err := row.Scan(
// 		&orderLine.OrderID,
// 		&orderLine.RacketID,
// 		&orderLine.Price,
// 		&orderLine.Quantity,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &orderLine, nil
// }
