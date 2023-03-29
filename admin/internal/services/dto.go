package services

import "go.mongodb.org/mongo-driver/bson/primitive"

type AdminUsersResponse struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	AdminUserName     string             `bson:"admin_user_name" json:"admin_user_name"`
	AdminUserLastName string             `bson:"admin_user_last_name" json:"admin_user_last_name"`
	AdminUserEmail    string             `bson:"admin_user_email" json:"admin_user_email"`
	Password          string             `bson:"password" json:"password,omitempty"`
	CreatedAt         string             `bson:"created_at" json:"created_at"`
	UpdatedAt         string             `bson:"updated_at" json:"updated_at"`
	Permissions       string             `bson:"permissions" json:"permissions"`
}

type AdminUsersPermissionsResponse struct {
	ID                       primitive.ObjectID `bson:"_id" json:"id"`
	AdminUserPermissionWrite string             `bson:"admin_user_permission_write" json:"admin_user_permission_write"`
	AdminUserPermissionRead  string             `bson:"admin_user_permission_read" json:"admin_user_permission_read"`
	AdminUserPermissionAll   string             `bson:"admin_user_permission_all" json:"admin_user_permission_all"`
	SuperAdminUserPermission string             `bson:"super_admin_user_permission" json:"super_admin_user_permission"`
}
