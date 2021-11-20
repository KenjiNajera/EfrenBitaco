package datastore

import (
	"context"
	"errors"
	"strconv"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/grupogit/RMNGR002001/Backend/helper"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (bool, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Create(&user.Resourcedata).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	user.Resourcedataid = user.Resourcedata.Resourcedataid

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	for _, element := range user.Resourcedata.Hardskills {
		element.Resourcedataid = user.Resourcedataid
		if err := tx.Create(&element).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}
	for _, element := range user.Resourcedata.Softskills {
		element.Resourcedataid = user.Resourcedataid
		if err := tx.Create(&element).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}
	for _, element := range user.Resourcedata.Experiences {
		element.Resourcedataid = user.Resourcedataid
		if err := tx.Create(&element).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}
	for _, element := range user.Resourcedata.Languages {
		element.Resourcedataid = user.Resourcedataid
		if err := tx.Create(&element).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}
	for _, element := range user.Resourcedata.Educations {
		element.Resourcedataid = user.Resourcedataid
		if err := tx.Create(&element).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}
	for _, element := range user.Resourcedata.Certificates {
		element.Resourcedataid = user.Resourcedataid
		if err := tx.Create(&element).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	// send email to create password personal.
	useridencript := helper.Encrypt(strconv.Itoa(user.Userid))

	templateData := struct {
		Name  string
		Email string
		URL   string
	}{
		Name:  user.Resourcedata.Firstname,
		Email: user.Resourcedata.Personalemail,
		URL:   "http://localhost:3000/auth/verify+account/" + useridencript,
	}
	subject := "Welcome to GrupoGIT"
	request := helper.NewRequest([]string{user.Resourcedata.Personalemail}, subject)
	if _, err := request.Send("./helper/templates/templateVerifyAccount.html", templateData); err != nil {
		tx.Rollback()
		return false, err
	}
	return true, tx.Commit().Error
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) (bool, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	sqlUpdateResourceData := `UPDATE resourcedatas 
		SET firstname = $1, lastname = $2, photouser = $3, cellphone = $4, 
			personalemail = $5, datebirth = $6, address = $7, country = $8, 
			city = $9, cp= $10 WHERE resourcedataid = $11;`
	if err := tx.Exec(sqlUpdateResourceData,
		user.Resourcedata.Firstname, user.Resourcedata.Lastname, user.Resourcedata.Photouser,
		user.Resourcedata.Cellphone, user.Resourcedata.Personalemail, user.Resourcedata.Datebirth,
		user.Resourcedata.Address, user.Resourcedata.Country, user.Resourcedata.City,
		user.Resourcedata.Cp, user.Resourcedataid).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	sqlUpdateUsers := `UPDATE users SET email = $1, roleid = $2 WHERE resourcedataid = $3;`
	if err := tx.Exec(sqlUpdateUsers, user.Email, user.Roleid, user.Resourcedataid).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	for _, element := range user.Resourcedata.Hardskills {
		sqlHardskills := `INSERT INTO hardsresources (domain, hardskillid, resourcedataid)
							VALUES($1, $2, $3) 
								ON CONFLICT (hardskillid, resourcedataid) 
							DO 
								UPDATE SET domain = $1;`
		if err := tx.Exec(sqlHardskills, element.Domain, element.Hardskillid, user.Resourcedataid).Error; err != nil {
			return false, err
		}
	}

	for _, element := range user.Resourcedata.Softskills {
		sqlSoftskills := `INSERT INTO softsresources (softskillid, resourcedataid)
							VALUES($1, $2) 
								ON CONFLICT (softskillid, resourcedataid) 
							DO 
								NOTHING;`
		if err := tx.Exec(sqlSoftskills, element.Softskillid, user.Resourcedataid).Error; err != nil {
			return false, err
		}
	}

	for _, element := range user.Resourcedata.Languages {
		sqlLanguages := `INSERT INTO languageresources (domain, languageid, resourcedataid)
							VALUES($1, $2, $3) 
								ON CONFLICT (languageid, resourcedataid) 
							DO 
								UPDATE SET domain = $1;`
		if err := tx.Exec(sqlLanguages, element.Domain, element.Languageid, user.Resourcedataid).Error; err != nil {
			return false, err
		}
	}

	for _, element := range user.Resourcedata.Educations {
		sqlEducations := `INSERT INTO educations (academicname, academictitle, year, resourcedataid)
							VALUES($1, $2, $3, $4) 
								ON CONFLICT (academicname, academictitle, resourcedataid) 
							DO 
								UPDATE SET academicname = $1, academictitle = $2, year = $3`
		if err := tx.Exec(sqlEducations, element.Academicname, element.Academictitle, element.Year, user.Resourcedataid).Error; err != nil {
			return false, err
		}
	}

	for _, element := range user.Resourcedata.Experiences {
		sqlEducations := `INSERT INTO experiences (projecttitle, description, yearelaboration, resourcedataid)
							VALUES($1, $2, $3, $4) 
								ON CONFLICT (projecttitle,resourcedataid) 
							DO 
								UPDATE SET 	projecttitle = $1, 
											description = $2, 
											yearelaboration = $3 `
		if err := tx.Exec(sqlEducations, element.Projecttitle, element.Description, element.Yearelaboration, user.Resourcedataid).Error; err != nil {
			return false, err
		}
	}

	for _, element := range user.Resourcedata.Certificates {
		sqlEducations := `INSERT INTO certificates (idcertificate, certificatename, expedition, expiration, resourcedataid)
							VALUES($1, $2, $3, $4, $5) 
								ON CONFLICT (idcertificate) 
							DO 
								UPDATE SET 	idcertificate = $1, 
											certificatename = $2, 
											expedition = $3, 
											expiration = $4`
		if err := tx.Exec(sqlEducations, element.Idcertificate, element.Certificatename,
			element.Expedition, element.Expiration, user.Resourcedataid).Error; err != nil {
			return false, err
		}
	}

	return true, tx.Commit().Error
}

func (r *userRepository) ActivateUser(ctx context.Context, id int) (bool, error) {
	user := &model.User{}
	sql := `SELECT * FROM users WHERE userid = $1`
	if err := r.db.Raw(sql, id).Scan(&user).Error; err != nil {
		return false, err
	}
	if user.Status != false {
		return false, errors.New("This user is already activated")
	}
	sqlActivate := `UPDATE users SET status = true WHERE userid = $1`
	if err := r.db.Exec(sqlActivate, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *userRepository) DesactivateUser(ctx context.Context, id int) (bool, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	user := &model.User{}
	sql := `SELECT * FROM users WHERE userid = $1`
	if err := tx.Raw(sql, id).Scan(&user).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if user.Status != true {
		return false, errors.New("This user is already deactivated")
	}

	sqlDesactivate := `UPDATE users SET status = false WHERE userid = $1`
	if err := tx.Exec(sqlDesactivate, id).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	// eliminarlo de members
	sqlDeleteofmember := `DELETE FROM members WHERE userid = $1`
	if err := tx.Exec(sqlDeleteofmember, id).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error
}

func (r *userRepository) GetUsers(ctx context.Context) ([]*model.GetUserDatas, error) {
	getusers := []*model.GetUserDatas{}
	sql := `SELECT u.userid, u.email, u.roleid, u.resourcedataid, rol.rolename, u.status FROM users AS u
				INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid
				INNER JOIN roles AS rol ON u.roleid = rol.roleid
				WHERE u.status = true;`
	if err := r.db.Raw(sql).Scan(&getusers).Error; err != nil {
		return nil, err
	}

	for _, item := range getusers {
		sqlResource := `SELECT * FROM resourcedatas WHERE resourcedataid = $1`
		if err := r.db.Raw(sqlResource, item.Resourcedataid).Scan(&item.Resourcedata).Error; err != nil {
			return nil, err
		}
		sqlEducation := `SELECT * FROM educations WHERE resourcedataid = $1`
		if err := r.db.Raw(sqlEducation, item.Resourcedataid).Scan(&item.Resourcedata.Educations).Error; err != nil {
			return nil, err
		}

		sqlExperience := `SELECT * FROM experiences  WHERE resourcedataid = $1`
		if err := r.db.Raw(sqlExperience, item.Resourcedataid).Scan(&item.Resourcedata.Experiences).Error; err != nil {
			return nil, err
		}

		sqlHardskills := `SELECT hr.domain, hr.hardsresourceid, hr.resourcedataid,
							h.hardskillid, h.hardskillname, h.image FROM hardskills AS h
							INNER JOIN hardsresources AS hr ON h.hardskillid = hr.hardskillid 
							WHERE hr.resourcedataid = $1`
		if err := r.db.Raw(sqlHardskills, item.Resourcedataid).Scan(&item.Resourcedata.Hardskills).Error; err != nil {
			return nil, err
		}

		sqlSoftskills := `SELECT sr.*, s.softskillname FROM softsresources AS sr
							INNER JOIN softskills AS s ON sr.softskillid = s.softskillid
							WHERE sr.resourcedataid = $1;`
		if err := r.db.Raw(sqlSoftskills, item.Resourcedataid).Scan(&item.Resourcedata.Softskills).Error; err != nil {
			return nil, err
		}

		sqlLanguages := `SELECT lr.*, l.languagename FROM languageresources AS lr
							INNER JOIN languages AS l ON lr.languageid = l.languageid
							WHERE lr.resourcedataid = $1;`
		if err := r.db.Raw(sqlLanguages, item.Resourcedataid).Scan(&item.Resourcedata.Languages).Error; err != nil {
			return nil, err
		}

		sqlCertificates := `SELECT * FROM certificates WHERE resourcedataid = $1`
		if err := r.db.Raw(sqlCertificates, item.Resourcedataid).Scan(&item.Resourcedata.Certificates).Error; err != nil {
			return nil, err
		}
	}

	return getusers, nil
}

func (r *userRepository) GetMyData(ctx context.Context, id int) (*model.MyData, error) {
	// sql := ``
	return nil, nil
}

func (r *userRepository) GetMyDataMovil(ctx context.Context, id int) (*model.MyData, error) {
	data := &model.MyData{}

	sqlQuery := `SELECT rolename, resourcedataid FROM users AS u 
		INNER JOIN roles AS r ON u.roleid = r.roleid
			WHERE userid = $1;`
	if err := r.db.Raw(sqlQuery, id).Scan(&data).Error; err != nil {
		return nil, err
	}

	sqlQuery = `SELECT * FROM resourcedatas WHERE resourcedataid = $1;`
	if err := r.db.Raw(sqlQuery, data.Resourcedataid).Scan(&data).Error; err != nil {
		return nil, err
	}

	sqlQuery = `SELECT * FROM experiences WHERE resourcedataid = $1;`
	if err := r.db.Raw(sqlQuery, data.Resourcedataid).Scan(&data.Experiences).Error; err != nil {
		return nil, err
	}

	sqlQuery = `SELECT * FROM educations WHERE resourcedataid = $1;`
	if err := r.db.Raw(sqlQuery, data.Resourcedataid).Scan(&data.Educations).Error; err != nil {
		return nil, err
	}

	sqlQuery = `SELECT * FROM certificates WHERE resourcedataid = $1;`
	if err := r.db.Raw(sqlQuery, data.Resourcedataid).Scan(&data.Certificates).Error; err != nil {
		return nil, err
	}

	sqlQuery = `SELECT hr.hardsresourceid, hr.domain, hs.hardskillname, hs.image FROM hardsresources AS hr
		INNER JOIN hardskills AS hs ON hr.hardskillid = hs.hardskillid 
			WHERE resourcedataid = $1;`
	if err := r.db.Raw(sqlQuery, data.Resourcedataid).Scan(&data.Hardskills).Error; err != nil {
		return nil, err
	}

	sqlQuery = `SELECT lr.languageresourceid, lr.domain, ls.languagename FROM languageresources AS lr
		INNER JOIN languages AS ls ON lr.languageid = ls.languageid 
			WHERE resourcedataid = $1;`
	if err := r.db.Raw(sqlQuery, data.Resourcedataid).Scan(&data.Languages).Error; err != nil {
		return nil, err
	}

	sqlQuery = `SELECT sr.softsresourceid, ss.softskillname FROM softsresources AS sr
		INNER JOIN softskills AS ss ON sr.softskillid = ss.softskillid 
			WHERE resourcedataid = $1;`
	if err := r.db.Raw(sqlQuery, data.Resourcedataid).Scan(&data.Softskills).Error; err != nil {
		return nil, err
	}

	return data, nil
}
func (r *userRepository) UpdateMyData(ctx context.Context, user *model.MyData) (bool, error) {
	return true, nil
}
