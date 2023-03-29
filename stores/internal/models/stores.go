package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type StoreInfomation struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	StoreName        string             `bson:"store_name" json:"store_name"`
	StoreCode        string             `bson:"store_code" json:"store_code"`
	StorePhoneNumber string             `bson:"store_phone" json:"store_phone"`
	StoreEmail       string             `bson:"store_email" json:"store_email"`
	Permissions      Permissions        `bson:"permissions" json:"permissions"`
	StoreAddress     StoreAddress       `bson:"store_address" json:"store_address"`
}

type StoreAddress struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	StoreStreet   string             `bson:"store_street" json:"store_street"`
	StoreNumber   string             `bson:"store_number" json:"store_number"`
	StoreZipcode  string             `bson:"store_zip_code" json:"store_zip_code"`
	StoreAreaCode string             `bson:"store_area_code" json:"store_area_code"`
}

type StoreServiceArea struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	RoleAdminUser     bool               `bson:"admin_user" json:"admin_user"`
	RoleInventoryRead bool               `bson:"role_inventory_read" json:"role_inventory_read"`
	RoleSuperAdmin    bool               `bson:"role" role_super_admin:"role_super_admin"`
}

// func (u AdminUsers) UsersValidations(ctx *fiber.Ctx) error {

// 	if len(u.AdminUserEmail) == 0 {
// 		return ctx.JSON(fiber.Map{"user_email": "User Email is required ..."})
// 	}
// 	if len(u.Password) < 8 {
// 		return ctx.JSON(fiber.Map{"password": "Password must be at least lenght of 6 elements ..."})
// 	}
// 	return nil
// }

// func SetPassword(password string) (string, error) {
// 	cost := 8
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
// 	if err != nil {
// 		log.Panic("Error ")
// 	}
// 	return string(hashedPassword), nil
// }

// func ComparePassword(s string, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
// }
