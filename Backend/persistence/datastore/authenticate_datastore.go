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

type authenticateRepository struct {
	db *gorm.DB
}

func NewAuthenticateRepository(db *gorm.DB) repository.AuthenticateRepository {
	return &authenticateRepository{db}
}

func (r *authenticateRepository) Authenticate(ctx context.Context, login *model.Login) (*model.Authenticate, error) {
	auth := &model.Authenticate{}
	user := &model.User{}
	if err := r.db.Raw("SELECT * FROM users WHERE email = $1", login.Email).Scan(&user).Error; err != nil {
		return nil, err
	}

	if state := helper.CheckPasswordHash(login.Password, user.Password); state != true {
		return nil, errors.New("the passwords do not match")
	}

	sql := `SELECT 
				userid, 
				resourcedataid, 
				u.roleid, 
				r.rolename, 
				status 
					FROM users AS u 
						INNER JOIN roles AS r 
							ON u.roleid = r.roleid 
								WHERE userid = $1`
	if err := r.db.Raw(sql, user.Userid).Scan(&auth).Error; err != nil {
		return nil, err
	}

	sqlupdate := `UPDATE users SET staterecovery = false WHERE userid = $1`
	if err := r.db.Exec(sqlupdate, user.Userid).Error; err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *authenticateRepository) AuthenticateMovil(ctx context.Context, obj *model.LoginMovil) (*model.AuthenticateMovil, error) {
	auth := &model.AuthenticateMovil{}
	user := &model.User{}
	if err := r.db.Raw("SELECT * FROM users WHERE email = $1", obj.Email).Scan(&user).Error; err != nil {
		return nil, err
	}

	if state := helper.CheckPasswordHash(obj.Password, user.Password); state != true {
		return nil, errors.New("the passwords do not match")
	}

	sql := `SELECT 
				userid, 
				resourcedataid, 
				u.roleid, 
				status 
					FROM users AS u 
						INNER JOIN roles AS r 
							ON u.roleid = r.roleid 
								WHERE userid = $1`
	if err := r.db.Raw(sql, user.Userid).Scan(&auth).Error; err != nil {
		return nil, err
	}

	if auth.Status != true {
		return nil, errors.New("Esta cuenta no esta activada.")
	}

	if obj.TokenMovil != "" {
		sql := `UPDATE users SET mobileid = $1 WHERE userid = $2`
		if err := r.db.Exec(sql, obj.TokenMovil, user.Userid).Error; err != nil {
			return nil, err
		}
	}

	return auth, nil
}

func (r *authenticateRepository) RecoveryPassword(ctx context.Context, email *model.RecoveryPassword) (bool, error) {

	sql := `SELECT u.userid, u.email, r.firstname, r.lastname  FROM users AS u
				INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid 
				WHERE u.email = $1`
	sqlupdate := `UPDATE users SET staterecovery = true WHERE userid = $1`
	if err := r.db.Raw(sql, email.Email).Scan(&email).Error; err != nil {
		return false, errors.New("user not found")
	}

	hash := helper.Encrypt(strconv.Itoa(email.Userid))

	templateData := struct {
		Name string
		URL  string
	}{
		Name: email.Firstname,
		URL:  "http://127.0.0.1:3000/auth/changepassword/" + hash,
	}
	subject := "Recovery Password"
	request := helper.NewRequest([]string{email.Email}, subject)

	if _, err := request.Send("./helper/templates/templateRecoveryPassword.html", templateData); err != nil {
		return false, err
	}

	if err := r.db.Exec(sqlupdate, email.Userid).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *authenticateRepository) ChangePassword(ctx context.Context, obj *model.ChangePassword) (bool, error) {
	u := &model.User{}
	id, _ := strconv.Atoi(helper.Decrypt(obj.Hashuserid))
	u.Userid = id
	sql := `SELECT * FROM users WHERE userid = $1`
	if err := r.db.Raw(sql, u.Userid).Scan(&u).Error; err != nil {
		return false, err
	}

	if u.Staterecovery != true {
		return false, errors.New("No request to change password")
	}

	sqlupdate := `UPDATE users SET password = $1, staterecovery = false WHERE userid = $2`
	newpass, _ := helper.HashPassword(obj.Newpassword)
	if err := r.db.Exec(sqlupdate, newpass, u.Userid).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *authenticateRepository) VerifyAccount(ctx context.Context, obj *model.ChangePassword) (bool, error) {
	u := &model.User{}
	id, _ := strconv.Atoi(helper.Decrypt(obj.Hashuserid))
	u.Userid = id
	sql := `SELECT * FROM users WHERE userid = $1`
	if err := r.db.Raw(sql, u.Userid).Scan(&u).Error; err != nil {
		return false, err
	}

	if u.Status == true {
		return false, errors.New("Esta cuenta ya esta verificada")
	}

	sqlupdate := `UPDATE users SET password = $1, status = true WHERE userid = $2`
	newpass, _ := helper.HashPassword(obj.Newpassword)
	if err := r.db.Exec(sqlupdate, newpass, u.Userid).Error; err != nil {
		return false, err
	}

	return true, nil
}
