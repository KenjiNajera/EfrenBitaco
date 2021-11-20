package datastore

import (
	"context"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetRoles(ctx context.Context) ([]*model.Role, error) {
	objects := []*model.Role{}
	if err := r.db.Raw("SELECT * FROM roles").Scan(&objects).Error; err != nil {
		return nil, err
	}

	for _, element := range objects {
		sqlModules := `SELECT modulerolid, moduleid FROM modulerols WHERE roleid = $1;`
		if err := r.db.Raw(sqlModules, element.Roleid).Scan(&element.Modules).Error; err != nil {
			return nil, err
		}

		for _, item := range element.Modules {
			sqlModulepermissions := `SELECT * FROM modpermimodrols WHERE modulerolid = $1`
			if err := r.db.Raw(sqlModulepermissions, item.Modulerolid).Scan(&item.Modulepermissions).Error; err != nil {
				return nil, err
			}
		}
	}

	return objects, nil
}

func (r *roleRepository) CreateRole(ctx context.Context, object *model.Role) (int, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	sql := "INSERT INTO roles( rolename, leader ) VALUES($1, $2) RETURNING roleid"
	if err := tx.Raw(sql, object.Rolename, object.Leader).Scan(&object).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, element := range object.Modules {
		sqlModuleRol := `INSERT INTO modulerols(roleid, moduleid) VALUES($1, $2) RETURNING modulerolid;`
		if err := tx.Raw(sqlModuleRol, object.Roleid, element.Moduleid).Scan(&element).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		for _, item := range element.Permissions {
			modulepermission := struct {
				Modulepermissionid int
			}{}
			sqlModulePermission := `
			with s AS (
				SELECT modulepermissionid FROM modulepermissions AS m  WHERE moduleid = $1 AND permissionid = $2
			),  i AS (
				INSERT INTO modulepermissions(moduleid, permissionid)
				SELECT $1, $2 
				WHERE NOT EXISTS (SELECT modulepermissionid FROM modulepermissions AS m  WHERE moduleid = $1 AND permissionid = $2)
				RETURNING modulepermissionid
			)
			SELECT modulepermissionid
			FROM i
			UNION ALL
			SELECT modulepermissionid
			FROM s`
			tx.Raw(sqlModulePermission, element.Moduleid, item).Scan(&modulepermission)

			element.Modulepermissions = append(element.Modulepermissions, modulepermission)
		}
		for _, item := range element.Modulepermissions {
			sqlModuleRolPermission := `INSERT INTO modpermimodrols(modulepermissionid, modulerolid) VALUES($1, $2);`
			if err := tx.Exec(sqlModuleRolPermission, item.Modulepermissionid, element.Modulerolid).Error; err != nil {
				tx.Rollback()
				return 0, err
			}
		}

	}

	return object.Roleid, tx.Commit().Error
}

func (r *roleRepository) UpdateRole(ctx context.Context, object *model.Role) (bool, error) {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	sqlDeleteALL := `DELETE FROM Modulerols WHERE roleid = $1`
	if err := tx.Exec(sqlDeleteALL, object.Roleid).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	sqlUpdate := `UPDATE roles SET rolename = $1, leader = $2 WHERE roleid = $3`
	if err := tx.Exec(sqlUpdate, object.Rolename, object.Leader, object.Roleid).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	for _, element := range object.Modules {
		sqlModuleRol := `INSERT INTO modulerols(roleid, moduleid) VALUES($1, $2) RETURNING modulerolid;`
		if err := tx.Raw(sqlModuleRol, object.Roleid, element.Moduleid).Scan(&element).Error; err != nil {
			tx.Rollback()
			return false, err
		}

		for _, item := range element.Permissions {
			modulepermission := struct {
				Modulepermissionid int
			}{}
			sqlModulePermission := `
			with s AS (
				SELECT modulepermissionid FROM modulepermissions AS m  WHERE moduleid = $1 AND permissionid = $2
			),  i AS (
				INSERT INTO modulepermissions(moduleid, permissionid)
				SELECT $1, $2 
				WHERE NOT EXISTS (SELECT modulepermissionid FROM modulepermissions AS m  WHERE moduleid = $1 AND permissionid = $2)
				RETURNING modulepermissionid
			)
			SELECT modulepermissionid
			FROM i
			UNION ALL
			SELECT modulepermissionid
			FROM s`
			tx.Raw(sqlModulePermission, element.Moduleid, item).Scan(&modulepermission)

			element.Modulepermissions = append(element.Modulepermissions, modulepermission)
		}

		for _, item := range element.Modulepermissions {
			sqlModuleRolPermission := `INSERT INTO modpermimodrols(modulepermissionid, modulerolid) VALUES($1, $2);`
			if err := tx.Exec(sqlModuleRolPermission, item.Modulepermissionid, element.Modulerolid).Error; err != nil {
				tx.Rollback()
				return false, err
			}
		}

	}

	return true, nil
}

func (r *roleRepository) DeleteRole(ctx context.Context, id int) error {
	sql := "DELETE FROM roles WHERE roleid = $1"
	err := r.db.Exec(sql, id).Error
	return err
}
