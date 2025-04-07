DROP DATABASE IF EXISTS orders_db;

CREATE DATABASE orders_db;

USE orders_db;

CREATE TABLE IF NOT EXISTS orders (
    order_id CHAR(36) NOT NULL PRIMARY KEY, -- UUID
    order_zone CHAR(2) NOT NULL,            
    order_state VARCHAR(2) NOT NULL,         
    order_status VARCHAR(10) NOT NULL,      -- order processing actual status. Can be PENDING, APPROVED and CANCELED
    order_prod_count INT NOT NULL,          -- number of products in the order
    order_created_at VARCHAR(20) NOT NULL,
    order_updated_at VARCHAR(20) NOT NULL
);


CREATE TABLE IF NOT EXISTS products (
    prod_id CHAR(36) NOT NULL PRIMARY KEY,  -- Occurrence product ID
    prod_prod_catalog_id CHAR(36) NOT NULL, -- Common product ID
    prod_name VARCHAR(100) NOT NULL,                
    prod_price FLOAT NOT NULL
);


CREATE TABLE IF NOT EXISTS distribution_centers (
    dist_cen_id CHAR(36) NOT NULL PRIMARY KEY,    -- UUID
    dist_cen_name CHAR(5) NOT NULL,             
    dist_cen_zone CHAR(2) NOT NULL,             
    dist_cen_state VARCHAR(2) NOT NULL,           
    dist_cen_status VARCHAR(10) NOT NULL          -- distribution center operational status. Can be ACTIVE and INACTIVE
);


CREATE TABLE IF NOT EXISTS order_products (
    op_order_id CHAR(36) NOT NULL,
    op_prod_id CHAR(36) NOT NULL,
    op_ordered_prod_quant INT NOT NULL,
    PRIMARY KEY (op_order_id, op_prod_id),
    FOREIGN KEY (op_order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
    FOREIGN KEY (op_prod_id) REFERENCES products(prod_id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS products_distribuition_centers (
    pdc_prod_id CHAR(36) NOT NULL,
    pdc_dist_cen_id CHAR(36) NOT NULL,
    PRIMARY KEY (pdc_prod_id, pdc_dist_cen_id),
    FOREIGN KEY (pdc_prod_id) REFERENCES products(prod_id) ON DELETE CASCADE,
    FOREIGN KEY (pdc_dist_cen_id) REFERENCES distribution_centers(dist_cen_id) ON DELETE CASCADE
);


INSERT INTO orders (order_id, order_zone, order_state, order_status, order_prod_count, order_created_at, order_updated_at) VALUES 
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', 'S1', 'SP', 'PENDING', 4 ,'2025-01-02 15:04:05', '2025-01-02 15:04:05')
  ;


INSERT INTO products (prod_id, prod_prod_catalog_id, prod_name, prod_price) VALUES
  ('0cb0a0d0-a815-41c6-9801-d4704c381cee','0f99276b-aa53-44f6-8bb6-5b4ededc9615', 'Perfume Masculino', 175.00),
  ('987313b0-c694-49b5-9fee-55e0e75234b4','c60ce040-e2e4-4828-b959-a500996816b8', 'Camisa Social', 200.00),
  ('25067e02-4de5-40fd-aafc-3f47bb2abe61','f51f1901-16f5-4af4-b080-b1a14ad2b4fc', 'Kit Festa facil de montar', 150.90),
  ('3ed31420-bd65-418d-81ac-84dd1f05fc78','8ebb99aa-cae8-48b8-8499-ec215b4f1edc', 'PS5 versao digital', 2499.99),
  ('8d01d7bb-fe91-43a1-a5da-be408f32d952','8ebb99aa-cae8-48b8-8499-ec215b4f1edc', 'PS5 versao digital Black Friday', 999.99)
  ;

INSERT INTO distribution_centers (dist_cen_id, dist_cen_name, dist_cen_zone, dist_cen_state, dist_cen_status) VALUES
  ('a5e3b8d4-6f25-4d2c-8e3a-1a2b3c4d5e6f', 'CD5', 'C1', 'SP', 'ACTIVE'),
  ('b2d3e4f5-6a7f-1c9d-8b0c-3d4e5b2a7f6c', 'CD12', 'N1', 'SP', 'ACTIVE'),
  ('e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c', 'CD15', 'N2', 'SP', 'ACTIVE')
  ;


INSERT INTO order_products (op_order_id, op_prod_id, op_ordered_prod_quant) VALUES
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', '0cb0a0d0-a815-41c6-9801-d4704c381cee', 1),
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', '987313b0-c694-49b5-9fee-55e0e75234b4', 2),
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', '25067e02-4de5-40fd-aafc-3f47bb2abe61', 3),
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', '8d01d7bb-fe91-43a1-a5da-be408f32d952', 1)
;

INSERT INTO products_distribuition_centers(pdc_prod_id, pdc_dist_cen_id) VALUES
  ('0f99276b-aa53-44f6-8bb6-5b4ededc9615', 'e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c'),
  ('c60ce040-e2e4-4828-b959-a500996816b8', 'b2d3e4f5-6a7f-1c9d-8b0c-3d4e5b2a7f6c'),
  ('f51f1901-16f5-4af4-b080-b1a14ad2b4fc', 'e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c'),
  ('8ebb99aa-cae8-48b8-8499-ec215b4f1edc', 'a5e3b8d4-6f25-4d2c-8e3a-1a2b3c4d5e6f')
;
