package repository

import (
	"context"

	"github.com/Dmitry-dms/zkh-plus/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type UserRequestsMongo struct {
	db *mongo.Database
}

func NewUserRequestsMongo(db *mongo.Database) *UserRequestsMongo {
	return &UserRequestsMongo{db: db}
}

// func (r *UserRequestsMongo) UpdateUserCompany(userId, companyId int) error {
// 	query := fmt.Sprintf("UPDATE %s SET company_id=$1 WHERE user_id = $2", usersTable)
// 	row := r.db.QueryRow(query, companyId, userId)
// 	if row == nil {
// 		return errors.New("failed to update company_id")
// 	}
// 	return nil
// }
func (r *UserRequestsMongo) GetUserInfo(userId primitive.ObjectID) (models.User, models.Company, error) {
	var user models.User
	var company models.Company
	filter := bson.M{"_id":userId}
	err := r.db.Collection(usersCollection).FindOne(context.TODO(), filter).Decode(&user)
	filter2 := bson.M{"_id":user.CompanyId}
	err = r.db.Collection(companyCollection).FindOne(context.TODO(), filter2).Decode(&company)
	return user, company, err
}
// func (r *UserRequestsMongo) CreateUserAddress(userId int, address models.UserAddress) (int, error) {
// 	var addressId int
// 	var checkAddress models.UserAddress
// 	checkQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND city=$2 AND street=$3 AND home_number=$4 AND flat=$5 limit 1", usersAddressTable)
// 	err := r.db.Get(&checkAddress, checkQuery, userId, address.City, address.Street, address.HomeNumber, address.Flat)
// 	if err != nil {
// 		query := fmt.Sprintf("INSERT INTO %s (user_id, city, street, home_number, flat) values ($1, $2, $3, $4, $5) RETURNING address_id", usersAddressTable)
// 		row := r.db.QueryRow(query, userId, address.City, address.Street, address.HomeNumber, address.Flat)
// 		if err := row.Scan(&addressId); err != nil {
// 			return 0, err
// 		}
// 		return addressId, nil
// 	}
// 	return 0, nil
// }

// func (r *UserRequestsMongo) GetAllUserAddress(userId int) ([]models.UserAddress, error) {
// 	var lists []models.UserAddress
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", usersAddressTable)
// 	err := r.db.Select(&lists, query, userId)
// 	return lists, err
// }



// func (r *UserRequestsMongo) InputVolumes(userId int, volume models.DataVolume) (models.VolumeResponse, error) {

// 	var arg1, arg2, arg3, arg4, arg5 string
// 	resp := new(models.VolumeResponse)
// 	var money float32

// 	lastValues, _ := r.GetUsersLastVolume(userId)
// 	if lastValues != nil {
// 		testVolume := lastValues[0]
// 		if volume.Electricity != nil {
// 			tEl, _ := strconv.ParseFloat(*testVolume.Electricity, 32)
// 			d := float32(tEl)
// 			arg1 = fmt.Sprintf("%s", *volume.Electricity)
// 			value, _ := strconv.ParseFloat(arg1, 32)
// 			pay := float32(value)
// 			resp.Electricity = electricity * (pay - d)
// 			money += electricity * (pay - d)
// 		} else {
// 			arg1 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.Gas != nil {
// 			tGas, _ := strconv.ParseFloat(*testVolume.Gas, 32)
// 			d := float32(tGas)
// 			arg2 = fmt.Sprintf("%s", *volume.Gas)
// 			value, _ := strconv.ParseFloat(arg2, 32)
// 			pay := float32(value)
// 			resp.Gas = gas * (pay - d)
// 			money += gas * (pay - d)
// 		} else {
// 			arg2 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.HotWater != nil {
// 			tHot, _ := strconv.ParseFloat(*testVolume.HotWater, 32)
// 			d := float32(tHot)
// 			arg3 = fmt.Sprintf("%s", *volume.HotWater)
// 			value, _ := strconv.ParseFloat(arg3, 32)
// 			pay := float32(value)
// 			resp.HotWater = hotW * (pay - d)
// 			money += hotW * (pay - d)
// 		} else {
// 			arg3 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.ColdWater != nil {
// 			tCold, _ := strconv.ParseFloat(*testVolume.ColdWater, 32)
// 			d := float32(tCold)
// 			arg4 = fmt.Sprintf("%s", *volume.ColdWater)
// 			value, _ := strconv.ParseFloat(arg4, 32)
// 			pay := float32(value)
// 			resp.ColdWater = coldW * (pay - d)
// 			money += coldW * (pay - d)
// 		} else {
// 			arg4 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.Warming != nil {
// 			tWarm, _ := strconv.ParseFloat(*testVolume.Warming, 32)
// 			d := float32(tWarm)
// 			arg5 = fmt.Sprintf("%s", *volume.Warming)
// 			value, _ := strconv.ParseFloat(arg5, 32)
// 			pay := float32(value)
// 			resp.Warming = warming * (pay - d)
// 			money += warming * (pay - d)
// 		} else {
// 			arg5 = fmt.Sprintf("%s", nullValue)
// 		}
// 		resp.Summ = money

// 		query := fmt.Sprintf("INSERT INTO %s (user_id,el_volume,gas_volume,hot_w_volume,cold_w_volume,warming_volume,date_full,date_year,date_month,date_day) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", volumeTable)

// 		_, err := r.db.Query(query, userId, arg1, arg2, arg3, arg4, arg5, volume.FullDate, volume.Year, volume.Month, volume.Day)
// 		return *resp, err
// 	} else {

// 		if volume.Electricity != nil {
// 			arg1 = fmt.Sprintf("%s", *volume.Electricity)
// 			value, _ := strconv.ParseFloat(arg1, 32)
// 			pay := float32(value)
// 			resp.Electricity = electricity * pay
// 			money += electricity * pay
// 		} else {
// 			arg1 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.Gas != nil {
// 			arg2 = fmt.Sprintf("%s", *volume.Gas)
// 			value, _ := strconv.ParseFloat(arg2, 32)
// 			pay := float32(value)
// 			resp.Gas = gas * pay
// 			money += gas * pay
// 		} else {
// 			arg2 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.HotWater != nil {
// 			arg3 = fmt.Sprintf("%s", *volume.HotWater)
// 			value, _ := strconv.ParseFloat(arg3, 32)
// 			pay := float32(value)
// 			resp.HotWater = hotW * pay
// 			money += hotW * pay
// 		} else {
// 			arg3 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.ColdWater != nil {
// 			arg4 = fmt.Sprintf("%s", *volume.ColdWater)
// 			value, _ := strconv.ParseFloat(arg4, 32)
// 			pay := float32(value)
// 			resp.ColdWater = coldW * pay
// 			money += coldW * pay
// 		} else {
// 			arg4 = fmt.Sprintf("%s", nullValue)
// 		}
// 		if volume.Warming != nil {
// 			arg5 = fmt.Sprintf("%s", *volume.Warming)
// 			value, _ := strconv.ParseFloat(arg5, 32)
// 			pay := float32(value)
// 			resp.Warming = warming * pay
// 			money += warming * pay
// 		} else {
// 			arg5 = fmt.Sprintf("%s", nullValue)
// 		}
// 		resp.Summ = money

// 		query := fmt.Sprintf("INSERT INTO %s (user_id,el_volume,gas_volume,hot_w_volume,cold_w_volume,warming_volume,date_full,date_year,date_month,date_day) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", volumeTable)

// 		_, err := r.db.Query(query, userId, arg1, arg2, arg3, arg4, arg5, volume.FullDate, volume.Year, volume.Month, volume.Day)
// 		return *resp, err
// 	}
// }

// func (r *UserRequestsMongo) GetUsersValuesByYearAndMonth(userId, year, month int) ([]models.DataVolume, error) {
// 	var lists []models.DataVolume
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND date_year=$2 AND date_month=$3", volumeTable)
// 	err := r.db.Select(&lists, query, userId, year, month)
// 	return lists, err
// }
// func (r *UserRequestsMongo) GetAllUserValues(userId int) ([]models.DataVolume, error) {
// 	var lists []models.DataVolume
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", volumeTable)
// 	err := r.db.Select(&lists, query, userId)
// 	return lists, err
// }
// func (r *UserRequestsMongo) GetUsersLastVolume(userId int) ([]models.DataVolume, error) {
// 	var value []models.DataVolume
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 ORDER BY id DESC LIMIT 1", volumeTable)
// 	err := r.db.Select(&value, query, userId)
// 	return value, err
// }

// func (r *UserRequestsMongo) GetNotifications(companyId int) ([]models.Notification, error) {
// 	var lists []models.Notification
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE company_id=$1", notificationsTable)
// 	err := r.db.Select(&lists, query, companyId)
// 	return lists, err
// }