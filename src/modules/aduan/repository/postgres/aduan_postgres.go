package postgres

import (
	"alqinsidev/jsa-mini-project/aduan/domain"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type AduanRepository struct {
	db *sql.DB
}

func NewAduanRepository(db *sql.DB) *AduanRepository {
	return &AduanRepository{db: db}
}

var queryDetail = `SELECT
adu.id,
adu.complain_id,
adu.title,
adu.description,
adu.evidence,
adu.complainant_id,
com."name" AS complainant_name,
com.phone AS complainant_phone,
com.social_media_id AS complainant_social_media_id,
com.social_media_name AS complainant_social_media_name,
com.social_media_link AS complainant_social_media_link,
adu.complainant_position_lat,
adu.complainant_position_lon,
adu.complained_city,
adu.complained_district,
adu.complained_sub_district,
adu.complained_address,
adu.complained_gmap_link,
adu.category_id,
cat."name" AS category_name,
adu.sub_category_id,
sc."name" AS sub_category_name,
adu.status,
sh.notes AS reason,
adu.created_at,
adu.updated_at
FROM aduan adu
INNER JOIN (
SELECT c.id, c.name, c.phone, c.social_media_id, sm.name AS social_media_name, c.social_media_link
FROM complainant c
INNER JOIN social_media sm ON c.social_media_id = sm.id
) com ON adu.complainant_id = com.id 
INNER JOIN category cat ON adu.category_id = cat.id 
INNER JOIN sub_category sc ON adu.sub_category_id = sc.id
LEFT JOIN (
SELECT
	sh.aduan_id,
	sh.status,
	sh.notes,
	ROW_NUMBER() OVER (PARTITION BY sh.aduan_id ORDER BY sh.created_at DESC) AS row_num
FROM status_history sh
) sh ON adu.id = sh.aduan_id AND sh.row_num = 1
WHERE 1=1
`
var queryTable = `SELECT a.id, a.complain_id, com.name AS complainant_name, cat.name AS category, a.status, a.created_at FROM aduan a INNER JOIN complainant com on a.complainant_id = com.id INNER JOIN category cat on a.category_id = cat.id WHERE 1=1`
var countQuery = `SELECT COUNT(*) FROM aduan a INNER JOIN complainant com on a.complainant_id = com.id WHERE 1=1`

func (ar *AduanRepository) Fetch(q *domain.RequestQuery) ([]domain.AduanTableResponse, int64, error) {

	startDate := q.StartDate
	endDate := q.EndDate
	sortBy := q.SortBy
	sortDir := q.SortDirection
	pageNumber := q.Page
	limit := q.Limit
	keyword := q.Keyword
	offset := (pageNumber - 1) * limit
	status := q.Filters["status"]

	var queryArgs []interface{}
	var totalCountQueryArgs []interface{}
	query := queryTable
	totalCountQuery := countQuery

	if status != nil {
		query += ` AND a.created_at BETWEEN $1 AND $2 AND (LOWER(com.name) LIKE $3 OR LOWER(a.complain_id) LIKE $4) AND a.status = $5 ORDER BY ` + sortBy + ` ` + sortDir + ` LIMIT $6 OFFSET $7`
		queryArgs = append(queryArgs, startDate, endDate, "%"+keyword+"%", "%"+keyword+"%", status, limit, offset)

		totalCountQuery += ` AND a.created_at BETWEEN $1 AND $2 AND (LOWER(com.name) LIKE $3 OR LOWER(a.complain_id) LIKE $4) AND a.status = $5`
		totalCountQueryArgs = append(totalCountQueryArgs, startDate, endDate, "%"+keyword+"%", "%"+keyword+"%", status)
	} else {
		query += ` AND a.created_at BETWEEN $1 AND $2 AND (LOWER(com.name) LIKE $3 OR LOWER(a.complain_id) LIKE $4) ORDER BY ` + sortBy + ` ` + sortDir + ` LIMIT $5 OFFSET $6`
		queryArgs = append(queryArgs, startDate, endDate, "%"+keyword+"%", "%"+keyword+"%", limit, offset)

		totalCountQuery += ` AND a.created_at BETWEEN $1 AND $2 AND (LOWER(com.name) LIKE $3 OR LOWER(a.complain_id) LIKE $4)`
		totalCountQueryArgs = append(totalCountQueryArgs, startDate, endDate, "%"+keyword+"%", "%"+keyword+"%")
	}

	var totalCount int64
	err := ar.db.QueryRow(totalCountQuery, totalCountQueryArgs...).Scan(&totalCount)
	if err != nil {
		log.Error().Err(err).Msg("error counting aduan")
		return nil, 0, err
	}

	rows, err := ar.db.Query(query, queryArgs...)
	if err != nil {
		log.Error().Err(err).Msg("fail to query aduan")
		return nil, 0, err
	}
	defer rows.Close()

	var result []domain.AduanTableResponse

	for rows.Next() {
		u := domain.AduanTableResponse{}
		err = rows.Scan(
			&u.ID,
			&u.ComplainID,
			&u.ComplainantName,
			&u.CategoryName,
			&u.Status,
			&u.CreatedAt,
		)
		if err != nil {
			log.Error().Err(err)
			return nil, 0, err
		}
		result = append(result, u)
	}

	return result, totalCount, nil
}

func (ar *AduanRepository) FindById(id uuid.UUID) (*domain.AduanDetail, error) {
	query := queryDetail + ` AND adu.id = $1`
	row := ar.db.QueryRow(query, id)
	aduan := &domain.AduanDetail{}

	err := row.Scan(
		&aduan.ID,
		&aduan.ComplainID,
		&aduan.Title,
		&aduan.Description,
		&aduan.Evidence,
		&aduan.ComplainantID,
		&aduan.ComplainantName,
		&aduan.ComplainantPhone,
		&aduan.ComplainantSocialMediaID,
		&aduan.ComplainantSocialMediaName,
		&aduan.ComplainantSocialMediaLink,
		&aduan.ComplainantPositionLat,
		&aduan.ComplainantPositionLon,
		&aduan.ComplainedCity,
		&aduan.ComplainedDistrict,
		&aduan.ComplainedSubDistrict,
		&aduan.ComplainedAddress,
		&aduan.ComplainedGMapLink,
		&aduan.CategoryID,
		&aduan.CategoryName,
		&aduan.SubCategoryID,
		&aduan.SubCategoryName,
		&aduan.Status,
		&aduan.Reason,
		&aduan.CreatedAt,
		&aduan.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrAduanNotFound
		}
		log.Error().Err(err).Msg("err")
		return nil, err
	}

	return aduan, nil
}

func (ar *AduanRepository) FetchSummary() (*domain.AduanSummaryResponse, error) {
	summaryQuery := `SELECT COUNT(*) FILTER (WHERE status = 1) AS total_waiting, COUNT(*) FILTER (WHERE status = 2) AS total_verified, COUNT(*) FILTER (WHERE status = 90) AS total_failed, COUNT(*) AS total_complain FROM aduan`
	row := ar.db.QueryRow(summaryQuery)
	summary := &domain.AduanSummaryResponse{}

	err := row.Scan(
		&summary.TotalWaiting,
		&summary.TotalVerified,
		&summary.TotalFailed,
		&summary.Total,
	)
	if err != nil {
		log.Error().Err(err).Msg("err")
		return nil, err
	}
	return summary, nil
}

func (ar *AduanRepository) UpdateStatus(update *domain.UpdateStatusPayload) (interface{}, error) {
	isExist := ar.IsAduanExist(update.ID)

	if !isExist {
		return nil, domain.ErrAduanNotFound
	}
	tx, err := ar.db.Begin()
	if err != nil {
		log.Error().Err(err).Msg("Error start trx in update status")
		return nil, err
	}

	queryStoreStatus := "INSERT INTO status_history (aduan_id, status, notes) VALUES ($1,$2,$3)"
	_, err = tx.Exec(queryStoreStatus, update.ID, update.Status, update.Reason)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Err update status insert history tx")
	}

	queryUpdateAduanStatus := "UPDATE aduan SET status = $1 WHERE id = $2"
	_, err = tx.Exec(queryUpdateAduanStatus, update.Status, update.ID)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("Err update status aduan tx")
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("Error commit update status")
	}
	return update, nil
}

func (ar *AduanRepository) IsAduanExist(id uuid.UUID) bool {
	simpleSelectQuery := "SELECT id FROM aduan WHERE id = $1 LIMIT 1"
	row := ar.db.QueryRow(simpleSelectQuery, id)

	var existingID uuid.UUID
	if err := row.Scan(&existingID); err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Error().Err(err).Msg("error querying to check aduan isexist")
			return false
		}
	}
	return true
}
