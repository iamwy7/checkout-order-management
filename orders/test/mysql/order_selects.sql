-- Get all products from an order
SELECT 
    op.op_order_id AS order_id,
    p.prod_id AS product_id,
    p.name AS product_name,
    p.price AS product_price,
    p.status AS product_status,
    p.created_at AS product_created_at,
    p.updated_at AS product_updated_at,
    op.op_ordered_prod_quant AS ordered_quantity
FROM 
    order_products op
INNER JOIN 
    products p
ON 
    op.op_prod_id = p.prod_id
WHERE 
    op.op_order_id = '8c8f5d8f-bcec-478c-8d56-5c8390d0f938';


-- Get all information about all products of an order, and the order itself
SELECT 
    o.order_id AS order_id,
    o.zone AS order_zone,
    o.state AS order_state,
    o.status AS order_status,
    o.created_at AS order_created_at,
    o.updated_at AS order_updated_at,
    p.prod_id AS product_id,
    p.name AS product_name,
    p.price AS product_price,
    p.status AS product_status,
    p.created_at AS product_created_at,
    p.updated_at AS product_updated_at,
    op.op_ordered_prod_quant AS ordered_quantity
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
WHERE 
    op.op_order_id = '8c8f5d8f-bcec-478c-8d56-5c8390d0f938';


-- Get all information about all products of an order, and the order itself, and the distribution center of each product
SELECT 
    o.order_id AS order_id,
    o.zone AS order_zone,
    o.state AS order_state,
    o.status AS order_status,
    o.created_at AS order_created_at,
    o.updated_at AS order_updated_at,
    p.prod_id AS product_id,
    p.name AS product_name,
    p.price AS product_price,
    op.op_ordered_prod_quant AS ordered_quantity,
    dc.dist_cen_id AS distribution_center_id,
    dc.name AS distribution_center_name,
    dc.zone AS distribution_center_zone,
    dc.state AS distribution_center_state,
    dc.status AS distribution_center_status
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
    op.op_order_id = '8c8f5d8f-bcec-478c-8d56-5c8390d0f938';