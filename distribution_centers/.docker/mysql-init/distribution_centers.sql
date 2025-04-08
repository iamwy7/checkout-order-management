SELECT
	d.dist_cen_id AS distribution_center_id,
	d.dist_cen_name AS distribution_center_name,
    d.dist_cen_zone AS distribution_center_zone,
    d.dist_cen_state AS distribution_center_state,
    d.dist_cen_status AS distribution_center_status,
    pdc.pdc_dist_cen_id_prod_count AS distribution_center_product_quantity
FROM products_distribuition_centers pdc
INNER JOIN
	distribution_centers d
ON
	pdc.pdc_dist_cen_id = d.dist_cen_id
WHERE pdc.pdc_prod_id = '2l3m4n5o-6p7q-8r9s-0t1u-2v3w4x5y6z7a';




SELECT
*
FROM products p
WHERE p.prod_id = '2l3m4n5o-6p7q-8r9s-0t1u-2v3w4x5y6z7a';

