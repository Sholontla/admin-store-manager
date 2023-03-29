package services

import (
	"context"
	"fmt"
	"log"
	"service/admin/case1/internal/db"
	"service/admin/case1/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterAdminUser(ctx context.Context, requestChan <-chan models.AdminUsers) error {

	select {
	case requests := <-requestChan:
		mongoCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		conn := db.MongoConn.Database("admin_db")
		coll := conn.Collection("admin_users")

		adminUser := models.AdminUsers{
			AdminUserName:     requests.AdminUserName,
			AdminUserLastName: requests.AdminUserLastName,
			AdminUserEmail:    requests.AdminUserEmail,
			Password:          requests.Password,
			CreatedAt:         time.Now().Local().String(),
			AdminUserRole:     requests.AdminUserRole,
			ID:                primitive.NewObjectID(),
			Permissions: models.Permissions{
				ID:                 primitive.NewObjectID(),
				UpdateStore:        requests.Permissions.UpdateStore,
				CreateStore:        requests.Permissions.CreateStore,
				DeleteStore:        requests.Permissions.DeleteStore,
				ReadStore:          requests.Permissions.ReadStore,
				UpdateInventory:    requests.Permissions.UpdateInventory,
				CreateInventory:    requests.Permissions.CreateInventory,
				DeleteInventory:    requests.Permissions.DeleteInventory,
				ReadInventory:      requests.Permissions.ReadInventory,
				ReadAdminUser:      requests.Permissions.ReadAdminUser,
				ReadWriteAdminUser: requests.Permissions.ReadWriteAdminUser,
				ReadConfig:         requests.Permissions.ReadConfig,
				ReadWriteConfig:    requests.Permissions.ReadWriteConfig,
			},
		}

		hash, err := models.SetPassword(requests.Password)
		if err != nil {
			log.Println(err)
		}
		adminUser.Password = string(hash)

		_, err = coll.InsertOne(mongoCtx, adminUser)
		if err != nil {
			return fmt.Errorf("failed to execute query: %v", err)
		}

		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func FindAdminUserByEmail(email string) {
	go func() {
		var adminUser models.AdminUsers
		condition := bson.M{
			"admin_user_email": email,
		}
		mongoCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		conn := db.MongoConn.Database("admin_db")
		coll := conn.Collection("admin_users")
		err := coll.FindOne(mongoCtx, condition).Decode(&adminUser)
		if err != nil {
			log.Println("No User found ...", err.Error())
		}
	}()
}
