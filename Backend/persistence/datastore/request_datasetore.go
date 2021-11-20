package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type requestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) repository.RequestRepository {
	return &requestRepository{db}
}

func (r *requestRepository) GetRequests(ctx context.Context) ([]*model.GetRequest, error) {
	requests := []*model.GetRequest{}
	sql := `SELECT r.requestid, r.createat, re.firstname, re.lastname, re.photouser, t.typerequestname, r.description FROM requests AS r 
				INNER JOIN users AS u ON r.userid = u.userid
				INNER JOIN resourcedatas AS re ON u.resourcedataid = re.resourcedataid
				INNER JOIN typerequests AS t ON r.typerequestid = t.typerequestid
				WHERE r.status IS null`
	if err := r.db.Raw(sql).Scan(&requests).Error; err != nil {
		return nil, err
	}
	for _, element := range requests {
		element.Format()
	}
	return requests, nil
}

func (r *requestRepository) CreateRequest(ctx context.Context, request *model.Request) (bool, error) {

	sql := `INSERT INTO requests(description, userid, typerequestid) values($1, $2, $3)`
	if err := r.db.Exec(sql, request.Description, request.Userid, request.Typerequestid).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *requestRepository) ResponseRequest(ctx context.Context, request *model.Request) (bool, error) {
	sql := `UPDATE requests SET status = $1, details = $2 WHERE requestid = $3`
	if err := r.db.Exec(sql, request.Status, request.Details, request.Requestid).Error; err != nil {
		return false, err
	}
	// Ejecutar funciones de firebase para notificaciones de app al aprobar o rechasar una peticion

	return true, nil
}

func (r *requestRepository) GetRequestsUser(ctx context.Context, id int) ([]*model.GetRequestUser, error) {
	requests := []*model.GetRequestUser{}
	sql := `SELECT r.requestid, t.typerequestname, r.description, r.status, r.details FROM requests AS r 
				INNER JOIN users AS u ON r.userid = u.userid
				INNER JOIN typerequests AS t ON r.typerequestid = t.typerequestid
				WHERE r.userid = $1 ORDER BY r.requestid DESC LIMIT 10`
	if err := r.db.Raw(sql, id).Scan(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}
