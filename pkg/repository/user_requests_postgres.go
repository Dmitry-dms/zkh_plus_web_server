package repository

import (
	"errors"
	"fmt"
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type UserRequestsPostgres struct {
	db *sqlx.DB
}

func NewUserRequestsPostgres(db *sqlx.DB) *UserRequestsPostgres {
	return &UserRequestsPostgres{db: db}
}

func (r *UserRequestsPostgres) UpdateUserCompany(userId, companyId int) error {
	query := fmt.Sprintf("UPDATE %s SET company_id=$1 WHERE user_id = $2", usersTable)
	row := r.db.QueryRow(query, companyId, userId)
	if row == nil {
		return errors.New("failed to update company_id")
	}
	return nil
}
func (r *UserRequestsPostgres) GetUserInfo(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT name,surname,patronymic,company_id FROM %s WHERE user_id=$1", usersTable)
	err := r.db.Get(&user, query, userId)
	return user, err
}
func (r *UserRequestsPostgres) CreateUserAddress(userId int, address models.UserAddress) (int, error) {
	var addressId int
	var checkAddress models.UserAddress
	checkQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND city=$2 AND street=$3 AND home_number=$4 AND flat=$5 limit 1", usersAddressTable)
	err := r.db.Get(&checkAddress, checkQuery, userId, address.City, address.Street, address.HomeNumber, address.Flat)
	if err != nil {
		query := fmt.Sprintf("INSERT INTO %s (user_id, city, street, home_number, flat) values ($1, $2, $3, $4, $5) RETURNING address_id", usersAddressTable)
		row := r.db.QueryRow(query, userId, address.City, address.Street, address.HomeNumber, address.Flat)
		if err := row.Scan(&addressId); err != nil {
			return 0, err
		}
		return addressId, nil
	}
	return 0, nil
}

func (r *UserRequestsPostgres) GetAllUserAddress(userId int) ([]models.UserAddress, error) {
	var lists []models.UserAddress
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", usersAddressTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

const nullValue = "nullValue"
const (
	coldW       = 35.78
	hotW        = 132.2
	warming     = 1968.67
	electricity = 3.41
	gas         = 5.03
)

func (r *UserRequestsPostgres) InputVolumes(userId int, volume models.DataVolume) (models.VolumeResponse, error) {

	var arg1, arg2, arg3, arg4, arg5 string
	resp := new(models.VolumeResponse)
	var money float32

	if volume.Electricity != nil {
		arg1 = fmt.Sprintf("%s", *volume.Electricity)
		value, _ := strconv.ParseFloat(arg1, 32)
		pay := float32(value)
		resp.Electricity = electricity * pay
		money += electricity * pay
	} else {
		arg1 = fmt.Sprintf("%s", nullValue)
	}
	if volume.Gas != nil {
		arg2 = fmt.Sprintf("%s", *volume.Gas)
		value, _ := strconv.ParseFloat(arg2, 32)
		pay := float32(value)
		resp.Gas = gas * pay
		money += gas * pay
	} else {
		arg2 = fmt.Sprintf("%s", nullValue)
	}
	if volume.HotWater != nil {
		arg3 = fmt.Sprintf("%s", *volume.HotWater)
		value, _ := strconv.ParseFloat(arg3, 32)
		pay := float32(value)
		resp.HotWater = hotW * pay
		money += hotW * pay
	} else {
		arg3 = fmt.Sprintf("%s", nullValue)
	}
	if volume.ColdWater != nil {
		arg4 = fmt.Sprintf("%s", *volume.ColdWater)
		value, _ := strconv.ParseFloat(arg4, 32)
		pay := float32(value)
		resp.ColdWater = coldW * pay
		money += coldW * pay
	} else {
		arg4 = fmt.Sprintf("%s", nullValue)
	}
	if volume.Warming != nil {
		arg5 = fmt.Sprintf("%s", *volume.Warming)
		value, _ := strconv.ParseFloat(arg5, 32)
		pay := float32(value)
		resp.Warming = warming * pay
		money += warming * pay
	} else {
		arg5 = fmt.Sprintf("%s", nullValue)
	}
	resp.Summ = money

	query := fmt.Sprintf("INSERT INTO %s (user_id,el_volume,gas_volume,hot_w_volume,cold_w_volume,warming_volume,date_full,date_year,date_month,date_day) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", volumeTable)

	_, err := r.db.Query(query, userId, arg1, arg2, arg3, arg4, arg5, volume.FullDate, volume.Year, volume.Month, volume.Day)
	return *resp, err
}

func (r *UserRequestsPostgres) GetUsersValuesByYearAndMonth(userId, year, month int) ([]models.DataVolume, error) {
	var lists []models.DataVolume
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND date_year=$2 AND date_month=$3", volumeTable)
	err := r.db.Select(&lists, query, userId, year, month)
	return lists, err
}
func (r *UserRequestsPostgres) GetAllUserValues(userId int) ([]models.DataVolume, error) {
	var lists []models.DataVolume
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", volumeTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *UserRequestsPostgres) GetNotifications(companyId int) ([]models.Notification, error) {
	var lists []models.Notification
	query := fmt.Sprintf("SELECT * FROM %s WHERE company_id=$1", notificationsTable)
	err := r.db.Select(&lists, query, companyId)
	return lists, err
}
