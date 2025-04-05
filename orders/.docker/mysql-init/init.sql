DROP DATABASE IF EXISTS orders_db;

CREATE DATABASE orders_db;

USE orders_db;

CREATE TABLE IF NOT EXISTS orders (
    order_id CHAR(36) NOT NULL PRIMARY KEY, -- UUID
    zone CHAR(2) NOT NULL,            
    state VARCHAR(2) NOT NULL,         
    status VARCHAR(10) NOT NULL,        -- order current status, can be updated if needed PENDING, APPROVED, CANCELED   \  idea: dedicated table
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- creation date of the order
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- date of last update
);



CREATE TABLE IF NOT EXISTS products (
    prod_id CHAR(36) NOT NULL PRIMARY KEY,      -- UUID
    name VARCHAR(100) NOT NULL,                
    price FLOAT NOT NULL,                      
    status VARCHAR(10) NOT NULL,            -- product status, can be ACTIVE, INACTIVE and DELETED
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- creation date of the product
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- date of last update
);


CREATE TABLE IF NOT EXISTS distribution_centers (
    dist_cen_id CHAR(36) NOT NULL PRIMARY KEY,   -- UUID
    name CHAR(5) NOT NULL,             
    zone CHAR(2) NOT NULL,             
    state VARCHAR(2) NOT NULL,           
    status VARCHAR(10) NOT NULL,            -- distribuition_center status, can be ACTIVE, INACTIVE
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- creation date of a distribution_center
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- date of last update
);


CREATE TABLE IF NOT EXISTS order_products (
    op_order_id CHAR(36) NOT NULL,       -- Chave estrangeira para a tabela `orders`
    op_prod_id CHAR(36) NOT NULL,     -- Chave estrangeira para a tabela `products`
    op_ordered_prod_quant INT NOT NULL, -- Quantidade do produto daquele pedido
    PRIMARY KEY (op_order_id, op_prod_id), -- Chave primária composta
    FOREIGN KEY (op_order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
    FOREIGN KEY (op_prod_id) REFERENCES products(prod_id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS products_distribuition_centers (
    pdc_prod_id CHAR(36) NOT NULL,                     -- Chave estrangeira para a tabela `orders`
    pdc_dist_cen_id CHAR(36) NOT NULL,       -- Chave estrangeira para a tabela `products`
    PRIMARY KEY (pdc_prod_id, pdc_dist_cen_id), -- Chave primária composta
    FOREIGN KEY (pdc_prod_id) REFERENCES products(prod_id) ON DELETE CASCADE,
    FOREIGN KEY (pdc_dist_cen_id) REFERENCES distribution_centers(dist_cen_id) ON DELETE CASCADE
);


INSERT INTO orders (order_id, zone, state, status) VALUES ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', 'S1', 'SP', 'PENDING');


INSERT INTO products (prod_id, name, price, status) VALUES
  ('0f99276b-aa53-44f6-8bb6-5b4ededc9615', 'Perfume Masculino', 175.00, 'ACTIVE'),
  ('c60ce040-e2e4-4828-b959-a500996816b8', 'Camisa Social', 200.00, 'ACTIVE'),
  ('f51f1901-16f5-4af4-b080-b1a14ad2b4fc', 'Kit Festa facil de montar', 150.90, 'ACTIVE'),
  ('8ebb99aa-cae8-48b8-8499-ec215b4f1edc', 'PS5 versao digital', 2499.99, 'ACTIVE')
  ;

INSERT INTO distribution_centers (dist_cen_id, name, zone, state, status) VALUES
    ('a5e3b8d4-6f25-4d2c-8e3a-1a2b3c4d5e6f', 'CD5', 'C1', 'SP', 'ACTIVE'),
    ('b2d3e4f5-6a7f-1c9d-8b0c-3d4e5b2a7f6c', 'CD12', 'N1', 'SP', 'ACTIVE'),
    ('e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c', 'CD15', 'N2', 'SP', 'ACTIVE')
    ;


INSERT INTO order_products (op_order_id, op_prod_id, op_ordered_prod_quant) VALUES
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', '0f99276b-aa53-44f6-8bb6-5b4ededc9615', 1),
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', 'c60ce040-e2e4-4828-b959-a500996816b8', 2),
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', 'f51f1901-16f5-4af4-b080-b1a14ad2b4fc', 3),
  ('8c8f5d8f-bcec-478c-8d56-5c8390d0f938', '8ebb99aa-cae8-48b8-8499-ec215b4f1edc', 1)
;

INSERT INTO products_distribuition_centers(pdc_prod_id, pdc_dist_cen_id) VALUES
  ('0f99276b-aa53-44f6-8bb6-5b4ededc9615', 'e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c'),
  ('c60ce040-e2e4-4828-b959-a500996816b8', 'b2d3e4f5-6a7f-1c9d-8b0c-3d4e5b2a7f6c'),
  ('f51f1901-16f5-4af4-b080-b1a14ad2b4fc', 'e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c'),
  ('8ebb99aa-cae8-48b8-8499-ec215b4f1edc', 'a5e3b8d4-6f25-4d2c-8e3a-1a2b3c4d5e6f')
;
