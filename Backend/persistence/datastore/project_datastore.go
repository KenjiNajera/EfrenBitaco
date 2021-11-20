package datastore

import (
	"context"
	"errors"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) repository.ProjectRepository {
	return &projectRepository{db}
}

func (r *projectRepository) GetProjects(ctx context.Context) ([]*model.ProjectCard, error) {
	projectsCard := []*model.ProjectCard{}
	sql := `SELECT p.*, c.photoclient  FROM projects as p INNER JOIN clients as c ON p.clientid = c.clientid`
	if err := r.db.Raw(sql).Scan(&projectsCard).Error; err != nil {
		return nil, err
	}

	sqlProjectHardskills := `SELECT h.* FROM projects AS p 
						INNER JOIN technologiesprojects AS t ON p.projectid = t.projectid
						INNER JOIN hardskills AS h ON t.hardskillid = h.hardskillid
						Where p.projectid = $1`

	sqlProjectResources := `SELECT u.userid, r.firstname, r.photouser FROM projects AS p 
						INNER JOIN members AS m ON p.projectid = m.projectid
						INNER JOIN users AS u ON m.userid = u.userid
						INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid  
						WHERE p.projectid = $1`

	for _, element := range projectsCard {
		if err := r.db.Raw(sqlProjectHardskills, element.Projectid).Scan(&element.Hardskills).Error; err != nil {
			return nil, err
		}
		if err := r.db.Raw(sqlProjectResources, element.Projectid).Scan(&element.Resources).Error; err != nil {
			return nil, err
		}
	}

	return projectsCard, nil
}

func (r *projectRepository) GetProjectByID(ctx context.Context, id int) (*model.ProjectCardDetail, error) {
	projectcard := &model.ProjectCardDetail{}
	sql := `SELECT * FROM projects WHERE projectid = $1`
	if err := r.db.Raw(sql, id).Scan(&projectcard).Error; err != nil {
		return nil, err
	}

	sqlResource := `SELECT u.userid, r.firstname, r.lastname, r.photouser, rol.rolename, u.status FROM projects AS p 
						INNER JOIN members AS m ON p.projectid = m.projectid
						INNER JOIN users AS u ON m.userid = u.userid
						INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid  
						INNER JOIN roles AS rol ON u.roleid = rol.roleid
						WHERE p.projectid = $1`

	sqlHardskill := `SELECT h.* FROM projects AS p 
						INNER JOIN technologiesprojects AS t ON p.projectid = t.projectid
						INNER JOIN hardskills AS h ON t.hardskillid = h.hardskillid
						Where p.projectid = $1`

	if err := r.db.Raw(sqlHardskill, id).Scan(&projectcard.Hardskills).Error; err != nil {
		return nil, err
	}
	if err := r.db.Raw(sqlResource, id).Scan(&projectcard.Resources).Error; err != nil {
		return nil, err
	}

	return projectcard, nil
}

func (r *projectRepository) CreateProject(ctx context.Context, object *model.Project) (bool, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Create(&object).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	sqlharsekill := `INSERT INTO technologiesprojects(projectid, hardskillid) values($1, $2)`
	for _, element := range object.Hardskills {
		if err := tx.Exec(sqlharsekill, object.Projectid, element).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit().Error
}

func (r *projectRepository) GetResources(ctx context.Context) ([]*model.UserProjectDetail, error) {

	resources := []*model.UserProjectDetail{}
	sql := `SELECT u.userid, r.firstname, r.lastname, r.photouser, rol.rolename FROM users AS u 
				INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid  
				INNER JOIN roles AS rol ON u.roleid = rol.roleid
				WHERE u.status = true and rol.leader = false;`

	if err := r.db.Raw(sql).Scan(&resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (r *projectRepository) GetLeaders(ctx context.Context) ([]*model.UserProjectDetail, error) {

	resources := []*model.UserProjectDetail{}
	sql := `SELECT u.userid, r.firstname, r.lastname, r.photouser, rol.rolename FROM users AS u 
	INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid  
	INNER JOIN roles AS rol ON u.roleid = rol.roleid
	WHERE u.status = true and rol.leader = true;`
	if err := r.db.Raw(sql).Scan(&resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (r *projectRepository) AllocateResource(ctx context.Context, member *model.Member) (bool, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	sqlCheck := `SELECT memberid FROM members WHERE userid = $1`
	tx.Raw(sqlCheck, member.Userid).Scan(&member)

	if member.Memberid == 0 {
		sql := `INSERT INTO members(userid, projectid) VALUES($1, $2)`
		if err := tx.Exec(sql, member.Userid, member.Projectid).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	} else {
		sqlupdate := `UPDATE members SET projectid = $1 WHERE memberid = $2;`
		if err := tx.Exec(sqlupdate, member.Projectid, member.Memberid).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit().Error
}

func (r *projectRepository) AllocateLeader(ctx context.Context, member *model.Member) (bool, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	sqlCheck := `SELECT r.leader FROM users AS u 
					INNER JOIN roles AS r ON u.roleid = r.roleid
					WHERE u.userid = $1;`
	if err := tx.Raw(sqlCheck, member.Userid).Scan(&member).Error; err != nil {
		return false, err
	}

	if member.Leader != false {
		sql := `INSERT INTO members(userid, projectid) VALUES($1, $2)`
		if err := tx.Exec(sql, member.Userid, member.Projectid).Error; err != nil {
			return false, err
		}
		sqlAssign := `UPDATE projects SET userid = $1 WHERE projectid = $2;`
		if err := tx.Exec(sqlAssign, member.Userid, member.Projectid).Error; err != nil {
			return false, err
		}
	} else {
		return false, errors.New("The user is not a leader")
	}

	return true, tx.Commit().Error
}

func (r *projectRepository) DeleteResource(ctx context.Context, id int) (bool, error) {

	sql := `DELETE FROM members WHERE userid = $1`
	if err := r.db.Exec(sql, id).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *projectRepository) FinishProject(ctx context.Context, id int) error {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	sqlfinishproject := `UPDATE projects SET status = true, finishat = current_date WHERE projectid = $1`
	if err := tx.Exec(sqlfinishproject, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlDeletemembers := `DELETE FROM members WHERE projectid = $1`
	if err := tx.Exec(sqlDeletemembers, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *projectRepository) UpdateProject(ctx context.Context, object *model.Project) (bool, error) {
	sql := ``
	if err := r.db.Exec(sql).Error; err != nil {
		return false, err
	}

	return true, nil
}
