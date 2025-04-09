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
WHERE pdc.pdc_prod_id = '1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p';

-- Others to try
-- 1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p - Has 6 DCs, but the CD4 must won
-- 2b3c4d5e-6f7g-8h9i-0j1k-2l3m4n5o6p7q - Has 4 DCs, but the CD13 must won
-- 3c4d5e6f-7g8h-9i0j-1k2l-3m4n5o6p7q8r - Has 2 DCs, but the CD16 must won
-- 4d5e6f7g-8h9i-0j1k-2l3m-4n5o6p7q8r9s - Has 1 DC, and the CD2 must won
;