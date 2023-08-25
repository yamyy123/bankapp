package models

import (
	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Customer struct {
// 	ID          primitive.ObjectID `json:"id" bson:"id"`
// 	Customer_ID int64              `json:"customer_id" bson:"customer_id"`
// 	Balance     int64              `json:"balance" bson:"balance"`
// 	Bank_ID     int64              `json:"bank_id" bson:"bank_id"`
// 	Password    string             `json:"password" bson:"password"`
// 	Name        string             `json:"name" bson:"name"`
// 	Account_ID  int64              `json:"account_id" bson:"account_id"`
// 	Transaction []CustTransaction  `json:"transaction" bson:"transaction"`
// }

// type CustTransaction struct {
// 	Transaction_id     int64     `json:"transaction_id" bson:"transaction_id"`
// 	Transaction_amount int64     `json:"transaction_amount" bson:"transaction_amount"`
// 	Date               time.Time `json:"date" bson:"date"`
// 	Transaction_type   string    `json:"transaction_type" bson:"transaction_type"`
// }


type Customer struct{

	BankID int  `json:"bankid" bson:"bankid"`
	Customer_Name string `json:"customer_name" bson:"customer_name"`
	DOB time.Time `json:"dob" bson:"dob"`
	Password []byte `json:"pasword" bson:"password"`
	Customer_ID int  `json:"customer_id" bson:"customer_id"`
	Balance int  `json:"balance bson:"balance"`
}