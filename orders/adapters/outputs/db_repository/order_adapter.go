package db_repository

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iamwy7/meli-challenge/orders/application/domain"
	out_ports "github.com/iamwy7/meli-challenge/orders/application/ports/outputs"
)

type OrderAdapter struct {
	db *sql.DB
}

func NewMySqlOrderAdapterFactory(db *sql.DB) (out_ports.OrderRepository, error) {
	return &OrderAdapter{db: db}, nil
}

func (oa *OrderAdapter) CreateAggregatedOrder(order domain.Order) error {
	oa.CreateOrder(order)
	for _, product := range order.Products {
		oa.CreateProduct(product)
		oa.CreateDistributuionCenter(product.DistributionCenter)
		oa.LinkProductToOrder(order.Id, product.Id, product.Quantity)
		oa.LinkDistributionCenterToProduct(product.Id, product.DistributionCenter.Id)
	}
	return nil
}

func (oa *OrderAdapter) CreateOrder(order domain.Order) error {
	query := `
		INSERT INTO orders (order_id, order_zone, order_state, order_status, order_created_at, order_updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := oa.db.Exec(query, order.Id, order.Zone, order.State, order.Status, order.CreatedAt, order.UpdatedAt)
	return err
}

func (oa *OrderAdapter) CreateProduct(product domain.Product) error {
	query := `
		INSERT INTO products (prod_id, prod_name, prod_price)
		VALUES (?, ?, ?, ?)
	`
	_, err := oa.db.Exec(query, product.Id, product.Name, product.Price)
	return err
}

func (oa *OrderAdapter) CreateDistributuionCenter(dc domain.DistributionCenter) error {
	query := `
		INSERT INTO distribution_centers (dist_cen_id, dist_cen_name, dist_cen_zone, dist_cen_state, dist_cen_status)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := oa.db.Exec(query, dc.Id, dc.Name, dc.Zone, dc.State, dc.Status)
	return err
}

func (oa *OrderAdapter) LinkProductToOrder(orderId string, prodId string, quantityOrdered int) error {
	query := `
		INSERT INTO order_products (op_order_id, op_prod_id, op_ordered_prod_quant)
		VALUES (?, ?, ?)
	`
	_, err := oa.db.Exec(query, orderId, prodId, quantityOrdered)
	return err
}

func (oa *OrderAdapter) LinkDistributionCenterToProduct(prodId string, dcId string) error {
	query := `
		INSERT INTO products_distribuition_centers(pdc_prod_id, pdc_dist_cen_id)
		VALUES (?, ?)
	`
	_, err := oa.db.Exec(query, prodId, dcId)
	return err
}

func (oa *OrderAdapter) GetAggregatedOrderById(orderId string) (*domain.Order, error) {
	products, err := oa.GetProductsByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	order, err := oa.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}
	order.Products = *products
	order.ProductsCount = len(*products)
	return order, nil
}

func (oa *OrderAdapter) GetProductsByOrderId(orderId string) (*[]domain.Product, error) {
	query := `
	SELECT 
		p.prod_id AS product_id,
		p.prod_name AS product_name,
		p.prod_price AS product_price,
		op.op_ordered_prod_quant AS ordered_prod_quant,
		dc.dist_cen_id AS distribution_center_id,
		dc.dist_cen_name AS distribution_center_name,
		dc.dist_cen_zone AS distribution_center_zone
	FROM 
		order_products op
	INNER JOIN 
		products p
	ON 
		op.op_prod_id = p.prod_id
	INNER JOIN 
		orders o
	ON 
		op.op_order_id = o.order_id
	INNER JOIN 
		products_distribuition_centers pdc
	ON 
		p.prod_id = pdc.pdc_prod_id
	INNER JOIN 
		distribution_centers dc
	ON 
		pdc.pdc_dist_cen_id = dc.dist_cen_id
	WHERE 
		op.op_order_id = ?
	`
	rows, err := oa.db.Query(query, orderId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrOrderInvalidId
		}
		return nil, err
	}
	defer rows.Close()
	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		var dc domain.DistributionCenter
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &dc.Id, &dc.Name, &dc.Zone)
		if err != nil {
			return nil, err
		}
		product.DistributionCenter = dc
		products = append(products, product)
	}
	return &products, nil
}

func (oa *OrderAdapter) GetOrderById(orderId string) (*domain.Order, error) {
	query := `
	SELECT 
		o.order_id AS order_id,
		o.order_zone AS order_zone,
		o.order_state AS order_state,
		o.order_status AS order_status,
		o.order_prod_count AS order_product_count,
		o.order_created_at AS order_created_at,
		o.order_updated_at AS order_updated_at
	FROM 
		orders o
	WHERE 
		o.order_id = ?
`
	rows, err := oa.db.Query(query, orderId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrOrderInvalidId
		}
		return nil, err
	}
	defer rows.Close()

	var order domain.Order
	err = rows.Scan(&order.Id, &order.Zone, &order.State, &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
