package repository

import (
	"context"
	

	"github.com/Dmitry-dms/zkh-plus/models"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyListMongo struct {
	db *mongo.Database
}

func NewCompanyListMongo(db *mongo.Database) *CompanyListMongo {
	return &CompanyListMongo{db: db}
}

func (r *CompanyListMongo) GetAllCompanies() ([]models.Company, error) {
	var lists []models.Company
	filter := bson.M{}
	cur, err := r.db.Collection(companyCollection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()){
		var c models.Company
		err := cur.Decode(&c)
		if err != nil {
			return nil, err
		}
		lists = append(lists, c)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())

	return lists, err
}

func (r *CompanyListMongo) GetCompanyById(companyId int) (models.Company, error) {
	var company models.Company
	filter := bson.M{"_id" : companyId}
	r.db.Collection(companyCollection).FindOne(context.Background(), filter).Decode(&company)
	return company, nil
}

func (r *CompanyListMongo) CreateNotification(companyId int, notification models.Notification) error {
	// query := fmt.Sprintf("INSERT INTO %s (company_id, article, description, date_full) values ($1, $2, $3, $4)", notificationsTable)
	// row := r.db.QueryRow(query, companyId, notification.Article, notification.Description, notification.FullDate)
	// if row == nil {
	// 	return errors.New("failed to create notification")
	// }
	return nil
}