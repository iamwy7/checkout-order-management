package db_repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Database driver
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
	err := oa.CreateOrder(order)
	if err != nil {
		log.Println("error creating order:", err)
		return err
	}
	for _, product := range order.Products {
		err := oa.CreateProduct(product)
		if err != nil {
			log.Println("error creating product:", err)
			return err
		}
		err = oa.CreateDistributuionCenter(product.DistributionCenter)
		if err != nil {
			log.Println("error creating dc:", err)
			return err
		}
		err = oa.LinkProductToOrder(order.Id, product.Id, product.Quantity)

		if err != nil {
			log.Println("error linking product to order:", err)
			return err
		}
		err = oa.LinkDistributionCenterToProduct(product.Id, product.DistributionCenter.Id)
		if err != nil {
			log.Println("error linking dc to product:", err)
			return err
		}
	}
	return nil
}

func (oa *OrderAdapter) CreateOrder(order domain.Order) error {
	query := `
		INSERT INTO orders (order_id, order_zone, order_state, order_prod_count, order_status, order_created_at, order_updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := oa.db.Exec(query, order.Id, order.Zone, order.State, order.ProductsCount, order.Status, order.CreatedAt.Format("2000-01-01 00:00:00"), order.UpdatedAt.Format("2000-01-01 00:00:00"))
	return err
}

func (oa *OrderAdapter) CreateProduct(product domain.Product) error {
	query := `
		INSERT INTO products (prod_id, prod_prod_catalog_id, prod_name, prod_price)
		VALUES (?, ?, ?, ?)
	`
	_, err := oa.db.Exec(query, product.Id, product.CatalogProductId, product.Name, product.Price)
	return err
}

func (oa *OrderAdapter) CreateDistributuionCenter(dc domain.DistributionCenter) error {
	query := `
		INSERT INTO distribution_centers (dist_cen_id, dist_cen_name, dist_cen_zone, dist_cen_state, dist_cen_status)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE dist_cen_name = VALUES(dist_cen_name), dist_cen_zone = VALUES(dist_cen_zone), dist_cen_state = VALUES(dist_cen_state), dist_cen_status = VALUES(dist_cen_status);
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
		return nil, domain.ErrUnexpectedError
	}
	if len(*products) == 0 {
		return nil, domain.ErrOrderInvalidId
	}
	order, err := oa.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}
	order.Products = *products
	return order, nil
}

func (oa *OrderAdapter) GetProductsByOrderId(orderId string) (*[]domain.Product, error) {
	query := `
	SELECT 
		p.prod_id AS product_id,
		p.prod_prod_catalog_id AS product_catalog_id,
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
		err = rows.Scan(&product.Id, &product.CatalogProductId, &product.Name, &product.Price, &product.Quantity, &dc.Id, &dc.Name, &dc.Zone)
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
	if err != nil || !rows.Next() {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrOrderInvalidId
		}
		return nil, err
	}
	defer rows.Close()

	var order domain.Order
	var createdAt, updatedAt string // Temporary variables to hold the raw date values
	err = rows.Scan(&order.Id, &order.Zone, &order.State, &order.Status, &order.ProductsCount, &createdAt, &updatedAt)
	if err != nil {
		return nil, domain.ErrUnexpectedError
	}

	// Parse the raw date strings into time.Time
	order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		log.Printf("failed to parse order_created_at: %v", err)
		return nil, domain.ErrUnexpectedError
	}

	order.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
	if err != nil {
		log.Printf("failed to parse order_updated_at: %v", err)
		return nil, domain.ErrUnexpectedError
	}

	return &order, nil
}
