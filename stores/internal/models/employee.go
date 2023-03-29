package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type StoreLogin struct {
	EmployeeUserName string `bson:"employee_user_email" json:"employee_user_email"`
	Password         string `bson:"password" json:"password,omitempty"`
}

type EmployeeInternalInformation struct {
	ID           primitive.ObjectID          `bson:"_id" json:"id"`
	PersonalInfo EmployeePersonalInformation `bson:"personal_info" json:"personal_info"`
	Area         string                      `bson:"area" json:"area"`
	Designation  string                      `bson:"designation" json:"designation"`
	Manager      string                      `bson:"manager" json:"manager"`
	Permissions  Permissions                 `bson:"permissions" json:"permissions"`
}

type EmployeePersonalInformation struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	FirstName     string             `bson:"first_name" json:"first_name"`
	LastName      string             `bson:"last_name" json:"last_name"`
	EmployeeEmail string             `bson:"employee_email" json:"employee_email"`
}
