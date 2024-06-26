package models

import (
	"context"
	"errors"
	"time"

	"github.com/SohamGhugare/hackbangalore-24/database"
	"github.com/SohamGhugare/hackbangalore-24/utils"
)

// model for the investor signup request
type InvestorSignup struct {
	Name     string              `bson:"name" json:"name"`
	Email    string              `bson:"email" json:"email"`
	Password string              `bson:"password" json:"password"`
	Tags     map[string][]string `bson:"tags" json:"tags"`
}

// model for the investor login request
type InvestorLogin struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

// model for the investor
type Investor struct {
	Name      string              `bson:"name" json:"name"`
	Email     string              `bson:"email" json:"email"`
	Password  string              `bson:"password" json:"password"`
	Tags      map[string][]string `bson:"tags" json:"tags"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
}

// new method for InvestorSignup
func (Investor) New(investorSignup InvestorSignup) (Investor, error) {
	hash, _ := utils.HashPassword(investorSignup.Password)

	investor := Investor{
		Name:      investorSignup.Name,
		Email:     investorSignup.Email,
		Password:  hash,
		Tags:      investorSignup.Tags,
		CreatedAt: time.Now(),
	}

	coll := database.DatabaseClient.Database("hackbangalore").Collection("investors")
	_, err := coll.InsertOne(context.TODO(), investor)

	return investor, err
}

// DoesExist checks if the investor already exists
func (Investor) DoesExist(email string) bool {
	coll := database.DatabaseClient.Database("hackbangalore").Collection("investors")
	filter := map[string]string{"email": email}
	res := coll.FindOne(context.TODO(), filter)

	return res.Err() == nil
}

// Get method for InvestorLogin
func (Investor) Get(email string) (Investor, error) {
	var investor Investor

	coll := database.DatabaseClient.Database("hackbangalore").Collection("investors")
	filter := map[string]string{"email": email}
	res := coll.FindOne(context.TODO(), filter)

	if res.Err() != nil {
		return Investor{}, errors.New("Invalid email")
	}

	err := res.Decode(&investor)

	return investor, err
}
