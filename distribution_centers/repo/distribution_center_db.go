package repo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Database driver
	"github.com/iamwy7/meli-challenge/distribution_centers/dtos"
	custom_errors "github.com/iamwy7/meli-challenge/distribution_centers/errors"
)

type DistributionCenterDB struct {
	db *sql.DB
}

func NewDistributionCenterDBFactory(db *sql.DB) (*DistributionCenterDB, error) {
	return &DistributionCenterDB{db: db}, nil
}

func (oa *DistributionCenterDB) GetDistributionCentersByProductId(prodId string) (*dtos.DistributionCenters, error) {
	query := `
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
		WHERE pdc.pdc_prod_id = ?
	`
	rows, err := oa.db.Query(query, prodId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dcs []dtos.DistributionCenter
	for rows.Next() {
		var dc dtos.DistributionCenter
		err = rows.Scan(&dc.Id, &dc.Name, &dc.Zone, &dc.State, &dc.Status, &dc.Quantity)
		if err != nil {
			return nil, custom_errors.ErrUnexpectedError
		}
		dcs = append(dcs, dc)
	}
	if len(dcs) == 0 {
		return nil, custom_errors.ErrProductInvalidId
	}

	dcList := dtos.NewDistributionCenters(prodId, dcs)
	return dcList, nil
}
