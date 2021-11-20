package datastore

import (
	"context"
	"fmt"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) repository.ClientRepository {
	return &clientRepository{db}
}

func (r *clientRepository) GetClients(ctx context.Context) ([]*model.Clients, error) {
	clients := []*model.Clients{}
	sql := `SELECT * FROM clients WHERE status = true;`

	if err := r.db.Raw(sql).Scan(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil

}

func (r *clientRepository) GetClientByID(ctx context.Context, id int) (*model.ClientProyect, error) {
	clientProyects := &model.ClientProyect{}
	sql := `SELECT * FROM clients WHERE clientid = $1`
	sqlprojects := `SELECT * FROM projects WHERE clientid = $1`

	if err := r.db.Raw(sql, id).Scan(&clientProyects.Cliente).Error; err != nil {
		return nil, err
	}
	if err := r.db.Raw(sqlprojects, id).Scan(&clientProyects.Projects).Error; err != nil {
		return nil, err
	}

	sqlHardskill := `SELECT h.* FROM projects AS p 
						INNER JOIN technologiesprojects AS t ON p.projectid = t.projectid
						INNER JOIN hardskills AS h ON t.hardskillid = h.hardskillid
						Where p.projectid = $1`

	sqlResource := `SELECT u.userid, r.resourcedataid, r.firstname, r.photouser FROM projects AS p 
						INNER JOIN members AS m ON p.projectid = m.projectid
						INNER JOIN users AS u ON m.userid = u.userid
						INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid  
						WHERE p.projectid = $1`

	for _, element := range clientProyects.Projects {
		if err := r.db.Raw(sqlHardskill, element.Projectid).Scan(&element.Hardskills).Error; err != nil {
			return nil, err
		}
		if err := r.db.Raw(sqlResource, element.Projectid).Scan(&element.Resources).Error; err != nil {
			return nil, err
		}
	}

	return clientProyects, nil
}

func (r *clientRepository) CreateClient(ctx context.Context, client *model.Client) (int, error) {
	if err := r.db.Create(&client).Error; err != nil {
		return 0, err
	}

	return client.Clientid, nil
}

func (r *clientRepository) UpdateClient(ctx context.Context, client *model.Client) (bool, error) {
	fmt.Println(client)
	sql := `UPDATE clients SET clientname = $1, socialreason = $2, photoclient = $3, 
				rfc = $4, address = $5, countryid = $6, city = $7, cp = $8, cellphone = $9 WHERE clientid = $10`
	if err := r.db.Exec(sql, client.Clientname, client.Socialreason, client.Photoclient,
		client.Rfc, client.Address, client.Countryid, client.City,
		client.Cp, client.Cellphone, client.Clientid).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *clientRepository) DeleteClient(ctx context.Context, id int) error {
	tx := r.db.Begin()

	sqlupdateClients := `UPDATE clients SET status = false WHERE clientid = $1;`
	if err := tx.Exec(sqlupdateClients, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	sqlDeletemembers := `DELETE FROM members WHERE projectid in (SELECT projectid FROM projects WHERE clientid = $1);`
	if err := tx.Exec(sqlDeletemembers, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	sqlUpdateprojects := `UPDATE projects SET status = true WHERE clientid = $1;`
	if err := tx.Exec(sqlUpdateprojects, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *clientRepository) Reactivate(ctx context.Context, id int) (bool, error) {
	sql := `UPDATE clients SET status = true WHERE clientid = $1;`
	if err := r.db.Exec(sql, id).Error; err != nil {
		return false, nil
	}
	return true, nil
}
