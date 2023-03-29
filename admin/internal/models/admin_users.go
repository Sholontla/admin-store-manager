package models

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AdminUserLogin struct {
	AdminUserEmail string `bson:"admin_user_email" json:"admin_user_email"`
	Password       string `bson:"password" json:"password,omitempty"`
}

type AdminUsers struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	AdminUserName     string             `bson:"admin_user_name" json:"admin_user_name"`
	AdminUserLastName string             `bson:"admin_user_last_name" json:"admin_user_last_name"`
	AdminUserEmail    string             `bson:"admin_user_email" json:"admin_user_email"`
	Password          string             `bson:"password" json:"password,omitempty"`
	CreatedAt         string             `bson:"created_at" json:"created_at"`
	UpdatedAt         string             `bson:"updated_at" json:"updated_at"`
	Permissions       Permissions        `bson:"permissions" json:"permissions"`
	AdminUserRole     bool               `bson:"role" json:"role"`
}

type AdminUserRole struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	RoleAdminUser     bool               `bson:"admin_user" json:"admin_user"`
	RoleInventoryRead bool               `bson:"role_inventory_read" json:"role_inventory_read"`
	RoleSuperAdmin    bool               `bson:"role" role_super_admin:"role_super_admin"`
}

type Permissions struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id"`
	AdminUserId        string             `bson:"admin_user_id" json:"admin_user_id"`
	UpdateStore        bool               `bson:"update_store" json:"update_store"`
	CreateStore        bool               `bson:"create_store" json:"create_store"`
	DeleteStore        bool               `bson:"delete_store" json:"delete_store"`
	ReadStore          bool               `bson:"read_store" json:"read_store"`
	UpdateInventory    bool               `bson:"update_inventory" json:"update_inventory"`
	CreateInventory    bool               `bson:"create_inventory" json:"create_inventory"`
	DeleteInventory    bool               `bson:"delete_inventory" json:"delete_inventory"`
	ReadInventory      bool               `bson:"read_inventory" json:"read_inventory"`
	ReadAdminUser      bool               `bson:"read_admin_user" json:"read_admin_user"`
	ReadWriteAdminUser bool               `bson:"read_write_admin_user" json:"read_write_admin_user"`
	ReadConfig         bool               `bson:"read_config" json:"read_config"`
	ReadWriteConfig    bool               `bson:"read_erite_config" json:"read_write_config"`
}

func (u AdminUsers) UsersValidations(ctx *fiber.Ctx) error {

	if len(u.AdminUserEmail) == 0 {
		return ctx.JSON(fiber.Map{"user_email": "User Email is required ..."})
	}
	if len(u.Password) < 8 {
		return ctx.JSON(fiber.Map{"password": "Password must be at least lenght of 6 elements ..."})
	}
	return nil
}

func SetPassword(password string) (string, error) {
	cost := 8
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Panic("Error ")
	}
	return string(hashedPassword), nil
}

func ComparePassword(s string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
}
