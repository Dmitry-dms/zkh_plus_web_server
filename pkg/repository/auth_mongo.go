package repository

import (
	"context"
	
	"fmt"

	"github.com/Dmitry-dms/zkh-plus/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)


type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user models.User, companyId primitive.ObjectID) (interface{}, error) {
	fullName := user.Surname + " " + user.Name + " " + user.Patronymic
	user.FullName = fullName
	user.CompanyId = companyId
	res, err := r.db.Collection(usersCollection).InsertOne(context.TODO(), user)
	if err != nil {
		return 0, err
	}
	return res.InsertedID, nil
}

func (r *AuthMongo) GetUser(email, password string) (models.User, error) {
	var user models.User
	filter := bson.M{"email" :email,"password":password}
	err := r.db.Collection(usersCollection).FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}
func (r *AuthMongo) CreateCompany(company models.Company) (interface{}, error) {
	res, err := r.db.Collection(companyCollection).InsertOne(context.TODO(), company)
	if err != nil {
		fmt.Println("here error")
		return 0, err
	}
	fmt.Println(res.InsertedID)
	return res.InsertedID, nil
}

func (r *AuthMongo) GetCompany(email, password string) (models.Company, error) {
	var user models.Company
	filter := bson.M{"email" : email,"password": password}
	r.db.Collection(companyCollection).FindOne(context.TODO(), filter).Decode(&user)
	return user, nil
}
