package datastore

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/grupogit/RMNGR002001/Backend/domain/model"
	"github.com/grupogit/RMNGR002001/Backend/domain/repository"
	"github.com/jinzhu/gorm"
)

type binnacleRepository struct {
	db *gorm.DB
}

func NewBinnacleRepository(db *gorm.DB) repository.BinnacleRepository {
	return &binnacleRepository{db}
}

func (r *binnacleRepository) CreateBinnacleByID(ctx context.Context, CreateBinnacle *model.CreateBinnacle) error {
	sqlCheckBinnacle := `
	with s as(
		SELECT status from binnacles as b where userid = $1 ORDER BY binnacleid DESC LIMIT 1
	)
	select 
		CASE 
			WHEN NOT EXISTS(select status from s) THEN true 
			WHEN EXISTS(SELECT status from s) THEN (select status from s) else false END AS status`
	if err := r.db.Raw(sqlCheckBinnacle, CreateBinnacle.Userid).Scan(&CreateBinnacle).Error; err != nil {
		return err
	}

	if CreateBinnacle.Status == true {
		sql := `INSERT INTO binnacles(monthname, month, year, userid) VALUES(to_char(to_timestamp ($1::text, 'MM'), $1, $2, $3)`
		if err := r.db.Exec(sql, CreateBinnacle.Month, CreateBinnacle.Year, CreateBinnacle.Userid).Error; err != nil {
			return err
		}
	} else {
		return errors.New("No se puede crear porque ya tiene una bitacora activa.")
	}

	return nil
}
func (r *binnacleRepository) FinishBinnacleByID(ctx context.Context, id int) error {
	binnacle := &model.FisnisBinnacle{}
	sqlCheckBinnacle := `SELECT status, binnacleid FROM binnacles 
							WHERE userid = $1 ORDER BY binnacleid DESC LIMIT 1`
	if err := r.db.Raw(sqlCheckBinnacle, id).Scan(&binnacle).Error; err != nil {
		return err
	}

	if binnacle.Status == false {
		// calcula las horas del mes
		sqlHoursDays := `SELECT 
							extract(Hour FROM SUM(totalhoursday)) + extract(MINUTE FROM SUM(totalhoursday))/100 
							AS hourstotal 
						FROM days WHERE binnacleid = $1`
		if err := r.db.Raw(sqlHoursDays, binnacle.Binnacleid).Scan(&binnacle.Hourstotal).Error; err != nil {
			return err
		}
		// actulaiza la bitacora osea finaliza
		sql := `UPDATE binnacles SET status = true, totalhours = $1 WHERE userid = $2`
		if err := r.db.Exec(sql, binnacle.Hourstotal, id).Error; err != nil {
			return err
		}
	} else {
		return errors.New("La bitacora no se puede finalizar porque no tiene bitacoras activas")
	}

	return nil
}

func (r *binnacleRepository) GetBinnacleByIdMovil(ctx context.Context, id int) (*model.BinnacleViewUser, error) {
	binnacleview := &model.BinnacleViewUser{}
	sqlBinnacleData := `
		SELECT b.binnacleid, b.monthname, b.year, b.totalhours, b.status FROM binnacles AS b
			INNER JOIN users AS u ON b.userid = u.userid
			WHERE b.userid = $1
		`
	if err := r.db.Raw(sqlBinnacleData, id).Scan(&binnacleview).Error; err != nil {
		return nil, err
	}

	sqlHoursDays := `SELECT 
						extract(Hour FROM SUM(totalhoursday)) + extract(MINUTE FROM SUM(totalhoursday))/100 
						AS totalhours 
					FROM days WHERE binnacleid = $1`
	if err := r.db.Raw(sqlHoursDays, binnacleview.Binnacleid).Scan(&binnacleview).Error; err != nil {
		return nil, err
	}

	sqlProjectnames := `SELECT distinct p.projectname  FROM binnacles AS b
	INNER JOIN days AS d ON b.binnacleid = d.binnacleid
	INNER JOIN dayprojects AS dp ON d.dayid = dp.dayid
	INNER JOIN projects AS p ON dp.projectid = p.projectid
	WHERE b.binnacleid = $1`
	if err := r.db.Raw(sqlProjectnames, binnacleview.Binnacleid).Scan(&binnacleview.Projectnames).Error; err != nil {
		return nil, err
	}

	var daydata []*model.Day
	sqlDayData := `
		SELECT d.dayid,  TO_CHAR(d.daydate, 'TMDay') AS dayname, d.daydate, d.createat, d.totalhoursday, d.dailytime FROM binnacles AS b
			INNER JOIN days AS d ON b.binnacleid = d.binnacleid
			WHERE b.binnacleid = $1
	`
	if err := r.db.Raw(sqlDayData, binnacleview.Binnacleid).Scan(&daydata).Error; err != nil {
		return nil, err
	}

	for _, element := range daydata {
		daybinnacleview := &model.DayBinnacleViewUser{
			Dayid:         element.Dayid,
			Dayname:       element.Dayname,
			Daydate:       element.Daydate,
			Totalhoursday: element.Totalhoursday,
		}

		sqlCheckTime := `SELECT checktime FROM checktimes WHERE dayid = $1`
		if err := r.db.Raw(sqlCheckTime, element.Dayid).Scan(&daybinnacleview).Error; err != nil {
			return nil, err
		}

		sqlMealTime := `SELECT mealtime FROM mealtimes WHERE dayid = $1`
		r.db.Raw(sqlMealTime, element.Dayid).Scan(&daybinnacleview)

		sqlDepartureTime := `SELECT departuretime FROM departuretimes WHERE dayid = $1`
		r.db.Raw(sqlDepartureTime, element.Dayid).Scan(&daybinnacleview)

		sqlOverTime := `SELECT overtime FROM overtimes WHERE dayid = $1`
		r.db.Raw(sqlOverTime, element.Dayid).Scan(&daybinnacleview)

		sqlActivity := `SELECT summary, invoice FROM activitys WHERE dayid = $1`
		r.db.Raw(sqlActivity, element.Dayid).Scan(&daybinnacleview)

		binnacleview.Days = append(binnacleview.Days, daybinnacleview)
	}
	sort.SliceStable(binnacleview.Days, func(i, j int) bool {
		return binnacleview.Days[i].Daydate.Before(binnacleview.Days[j].Daydate)
	})

	return binnacleview, nil
}

func (r *binnacleRepository) GetBinnacleById(ctx context.Context, id int) (*model.BinnacleView, error) {
	binnacleview := &model.BinnacleView{}
	sqlBinnacleData := `
		SELECT b.binnacleid, b.monthname, b.year, b.totalhours, b.status, b.userid, r.firstname, r.lastname,r.photouser FROM binnacles AS b
			INNER JOIN users AS u ON b.userid = u.userid
			INNER JOIN resourcedatas AS r ON u.resourcedataid = r.resourcedataid
			WHERE b.userid = $1
		`
	if err := r.db.Raw(sqlBinnacleData, id).Scan(&binnacleview.BinnacleData).Error; err != nil {
		return nil, err
	}

	sqlHoursDays := `SELECT 
						extract(Hour FROM SUM(totalhoursday)) + extract(MINUTE FROM SUM(totalhoursday))/100 
						AS Totalhours 
					FROM days WHERE binnacleid = $1`
	if err := r.db.Raw(sqlHoursDays, binnacleview.BinnacleData.Binnacleid).Scan(&binnacleview.BinnacleData).Error; err != nil {
		return nil, err
	}

	sqlProjectnames := `SELECT distinct p.projectname  FROM binnacles AS b
	INNER JOIN days AS d ON b.binnacleid = d.binnacleid
	INNER JOIN dayprojects AS dp ON d.dayid = dp.dayid
	INNER JOIN projects AS p ON dp.projectid = p.projectid
	WHERE b.binnacleid = $1`
	if err := r.db.Raw(sqlProjectnames, binnacleview.BinnacleData.Binnacleid).Scan(&binnacleview.Projectnames).Error; err != nil {
		return nil, err
	}

	var daydata []*model.Day
	sqlDayData := `
		SELECT d.dayid,  TO_CHAR(d.daydate, 'TMDay') AS dayname, d.daydate, d.createat, d.totalhoursday, d.dailytime FROM binnacles AS b
			INNER JOIN days AS d ON b.binnacleid = d.binnacleid
			WHERE b.binnacleid = $1
	`
	if err := r.db.Raw(sqlDayData, binnacleview.BinnacleData.Binnacleid).Scan(&daydata).Error; err != nil {
		return nil, err
	}

	for _, element := range daydata {
		daybinnacleview := &model.DayBinnacleView{
			Daydata: element,
		}

		sqlCheckTime := `SELECT * FROM checktimes WHERE dayid = $1`
		if err := r.db.Raw(sqlCheckTime, element.Dayid).Scan(&daybinnacleview.Checktime).Error; err != nil {
			return nil, err
		}

		sqlMealTime := `SELECT * FROM mealtimes WHERE dayid = $1`
		r.db.Raw(sqlMealTime, element.Dayid).Scan(&daybinnacleview.Mealtime)

		sqlDepartureTime := `SELECT * FROM departuretimes WHERE dayid = $1`
		r.db.Raw(sqlDepartureTime, element.Dayid).Scan(&daybinnacleview.Departuretime)

		sqlOverTime := `SELECT * FROM overtimes WHERE dayid = $1`
		r.db.Raw(sqlOverTime, element.Dayid).Scan(&daybinnacleview.Overtime)

		sqlActivity := `SELECT * FROM activitys WHERE dayid = $1`
		r.db.Raw(sqlActivity, element.Dayid).Scan(&daybinnacleview.Activity)

		binnacleview.Days = append(binnacleview.Days, daybinnacleview)
	}
	sort.SliceStable(binnacleview.Days, func(i, j int) bool {
		return binnacleview.Days[i].Daydata.Daydate.Before(binnacleview.Days[j].Daydata.Daydate)
	})

	return binnacleview, nil
}

func createDay(userid int, daytime time.Time, tx *gorm.DB) (int, error) {

	binnacle := &model.CreateDay{}

	sqlBinnacleActive := `SELECT binnacleid, month, year FROM binnacles 
							WHERE userid = $1 and status = false ORDER BY binnacleid DESC LIMIT 1`
	if err := tx.Raw(sqlBinnacleActive, userid).Scan(&binnacle).Error; err != nil {
		return 0, err
	}

	year, month, _ := daytime.Date()

	if month == time.Month(binnacle.Month) && year == binnacle.Year {
		sqlCreateDay := `INSERT INTO days(daydate, binnacleid) VALUES($1, $2) RETURNING dayid;`
		if err := tx.Raw(sqlCreateDay, daytime, binnacle.Binnacleid).Scan(&binnacle).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		sqlCheckProjectsUser := `SELECT projectid FROM members WHERE userid = $1`
		if err := tx.Raw(sqlCheckProjectsUser, userid).Scan(&binnacle.Projectids).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		for _, id := range binnacle.Projectids {
			sqlCreateDayProjects := `INSERT INTO dayprojects(dayid, projectid) VALUES($1, $2);`
			if err := tx.Exec(sqlCreateDayProjects, binnacle.Dayid, id).Error; err != nil {
				tx.Rollback()
				return 0, err
			}
		}

		return binnacle.Dayid, nil
	}

	return 0, errors.New("La fecha ingresada no pertenece al mes de la bitacora")
}

func (r *binnacleRepository) CheckIn(ctx context.Context, checkin *model.Checktime) error {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	dayid, err := createDay(checkin.Userid, checkin.Daydate, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	sqlCreateCheckIn := `INSERT INTO checktimes(checktime, dayid) VALUES($1, $2);`
	if err := tx.Exec(sqlCreateCheckIn, checkin.Checktime.Round(time.Hour), dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *binnacleRepository) CheckOut(ctx context.Context, departuretime *model.Departuretime) error {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	departuretimedata := &model.CreateDeparureTime{}

	sqlBinnacleActive := `SELECT binnacleid FROM binnacles 
							WHERE userid = $1 and status = false ORDER BY binnacleid DESC LIMIT 1`
	if err := tx.Raw(sqlBinnacleActive, departuretime.Userid).Scan(&departuretimedata).Error; err != nil {
		return err
	}

	sqlCheckExistDay := `SELECT dayid FROM days WHERE daydate = $1 AND binnacleid = $2;`
	if err := tx.Raw(sqlCheckExistDay, departuretime.Daydate, departuretimedata.Binnacleid).Scan(&departuretimedata).Error; err != nil {
		return err
	}

	checkout := departuretime.Departuretime.Round(time.Hour)
	overtime := checkout.Add(-18 * time.Hour).Round(time.Hour)
	// hour, _, _ := overtime.Clock()

	if departuretime.Overtime == true {
		// checkout = checkout.Add((time.Duration(hour) * time.Hour) * -1).Round(time.Hour)
		sqlInsertOvertime := `INSERT INTO overtimes(overtime, dayid) VALUES($1, $2)`
		if err := tx.Exec(sqlInsertOvertime, overtime, departuretimedata.Dayid).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if departuretime.Summary != "" {
		sqlInsertActivity := `INSERT INTO activitys(summary, invoice, dayid) VALUES($1, $2, $3)`
		if err := tx.Exec(sqlInsertActivity, departuretime.Summary, departuretime.Invoice, departuretimedata.Dayid).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	sqlInsertDeparture := `INSERT INTO departuretimes(departuretime, dayid) VALUES($1, $2)`
	if err := tx.Exec(sqlInsertDeparture, checkout, departuretimedata.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlCheckinDay := `SELECT extract(hour from checktime) * -1 AS checktime FROM checktimes WHERE dayid = $1`
	if err := tx.Raw(sqlCheckinDay, departuretimedata.Dayid).Scan(&departuretimedata).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlMealDay := `SELECT extract(hour from mealtime) - 1 + extract(hour from mealtime) * -1 AS mealtime FROM mealtimes WHERE dayid = $1`
	tx.Raw(sqlMealDay, departuretimedata.Dayid).Scan(&departuretimedata)

	dailytime := checkout.Add(time.Duration(departuretimedata.Checktime) * time.Hour).Round(time.Hour)
	totalhour := dailytime.Add(time.Duration(departuretimedata.Mealtime) * time.Hour).Round(time.Hour)
	sqlUpdateDay := `UPDATE days SET totalhoursday = $1, dailytime = $2 WHERE dayid = $3;`
	if err := tx.Exec(sqlUpdateDay, totalhour, dailytime, departuretimedata.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *binnacleRepository) MealDay(ctx context.Context, mealtime *model.Mealtime) error {
	mealtimedata := &model.CreateMealTime{}

	sqlBinnacleActive := `SELECT binnacleid FROM binnacles 
							WHERE userid = $1 and status = false ORDER BY binnacleid DESC LIMIT 1`
	if err := r.db.Raw(sqlBinnacleActive, mealtime.Userid).Scan(&mealtimedata).Error; err != nil {
		return err
	}

	sqlCheckExistDay := `SELECT dayid FROM days 
								WHERE daydate = $1 AND binnacleid = $2;`
	if err := r.db.Raw(sqlCheckExistDay, mealtime.Daydate, mealtimedata.Binnacleid).Scan(&mealtimedata).Error; err != nil {
		return errors.New("El dia aun no tiene un check in de entrada")
	}

	sqlCheckExistCheckout := `SELECT departuretimeid FROM departuretimes WHERE dayid =$1`
	r.db.Raw(sqlCheckExistCheckout, mealtimedata.Dayid).Scan(&mealtimedata)

	if mealtimedata.DeparturetimeId == 0 {
		return errors.New("El día ya esta finalizado")
	}

	sqlCreateMealtime := `INSERT INTO mealtimes(mealtime, dayid) VALUES($1, $2)`
	if err := r.db.Exec(sqlCreateMealtime, mealtime.Mealtime.Round(time.Hour), mealtimedata.Dayid).Error; err != nil {
		return err
	}

	return nil
}

func (r *binnacleRepository) UpdateDay(ctx context.Context, data *model.UpdateDay) error {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	sqlUpdateDay := `UPDATE days SET daydate = $1, totalhoursday = $2, dailytime = $3 WHERE dayid = $4;`
	if err := tx.Exec(sqlUpdateDay, data.Daydate, data.Totalhoursday, data.Dailytime, data.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlUpdateChecktime := `UPDATE checktimes SET checktime = $1 WHERE dayid = $2;`
	if err := tx.Exec(sqlUpdateChecktime, data.Checktime, data.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlUpdateMealtime := `UPDATE mealtimes SET mealtime = $1 WHERE dayid = $2;`
	if err := tx.Exec(sqlUpdateMealtime, data.Mealtime, data.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlUpdateDeparturetime := `UPDATE departuretimes SET departuretime = $1 WHERE dayid = $2;`
	if err := tx.Exec(sqlUpdateDeparturetime, data.Departuretime, data.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlUpdateOvertime := `UPDATE overtimes SET overtime = $1 WHERE dayid = $2;`
	if err := tx.Exec(sqlUpdateOvertime, data.Overtime, data.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	sqlUpdateActivity := `UPDATE activitys SET summary = $1, Invoice = $2 WHERE dayid = $3;`
	if err := tx.Exec(sqlUpdateActivity, data.Summary, data.Invoice, data.Dayid).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
