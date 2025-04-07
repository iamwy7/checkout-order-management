-- Get an order by id
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
    o.order_id = '8c8f5d8f-bcec-478c-8d56-5c8390d0f938';

-- Get all information about all products of an order and the distribution center of each product
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
    op.op_order_id = '8c8f5d8f-bcec-478c-8d56-5c8390d0f938';