package domain

import (
	"service/stores/case1/datasource/cassandra"
	"service/stores/case1/internal/models"

	"github.com/gocql/gocql"
)

const (
	queryCqlInsertOrder = "INSERT INTO stores.store_login JSON ?"
)

func GetStoreLoginByEmail(db cassandra.DB, email string) (*models.StoreLogin, error) {
	var storeLogin models.StoreLogin

	query := db.Query(`SELECT admin_user_email, password FROM store_login WHERE admin_user_email = ?`, email)
	err := query.Scan(&storeLogin.EmployeeUserName, &storeLogin.Password)

	if err == gocql.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &storeLogin, nil
	}
}

func GetStoreLogin(session *gocql.Session, email string, resultChan chan<- *models.StoreLogin, errorChan chan<- error) {
	var storeLogin models.StoreLogin

	query := session.Query(`SELECT admin_user_email, password FROM store_login WHERE admin_user_email = ?`, email)
	err := query.Scan(&storeLogin.EmployeeUserName, &storeLogin.Password)

	if err == gocql.ErrNotFound {
		resultChan <- nil
	} else if err != nil {
		errorChan <- err
	} else {
		resultChan <- &storeLogin
	}
}

func GetAllStoreLogins(db cassandra.DB) ([]*models.StoreLogin, error) {
	var storeLogins []*models.StoreLogin

	query := db.Query(`SELECT admin_user_email, password FROM store_login`)
	iter := query.Iter()

	var adminUserEmail, password string
	for iter.Scan(&adminUserEmail, &password) {
		storeLogin := &models.StoreLogin{
			EmployeeUserName: adminUserEmail,
			Password:         password,
		}
		storeLogins = append(storeLogins, storeLogin)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return storeLogins, nil
}

func CreateStoreLogin(db cassandra.DB, storeLogin *models.StoreLogin) error {
	query := db.Query(`INSERT INTO store_login (admin_user_email, password) VALUES (?, ?)`, storeLogin.EmployeeUserName, storeLogin.Password)
	return query.Exec()
}

func UpdateStoreLogin(db cassandra.DB, storeLogin *models.StoreLogin) error {
	query := db.Query(`UPDATE store_login SET password = ? WHERE admin_user_email = ?`, storeLogin.Password, storeLogin.EmployeeUserName)
	return query.Exec()
}

func DeleteStoreLoginByEmail(db cassandra.DB, email string) error {
	query := db.Query(`DELETE FROM store_login WHERE admin_user_email = ?`, email)
	return query.Exec()
}

// func CreateCustomerFinalStoreOrderDao(db cassandra.DB, orderCompleted *models.CustomerOrderCompleted) error {

// 	co := models.CustomerOrderCompleted{
// 		ID: orderCompleted.ID,
// 		Inventory: models.Inventory{
// 			InventoryID: orderCompleted.Inventory.InventoryID,
// 			InventoryProduct: models.Product{
// 				ProductId:             orderCompleted.Inventory.InventoryProduct.ProductId,
// 				ProductTitle:          orderCompleted.Inventory.InventoryProduct.ProductTitle,
// 				ProductClassification: orderCompleted.Inventory.InventoryProduct.ProductClassification,
// 				ProductCategory:       orderCompleted.Inventory.InventoryProduct.ProductCategory,
// 				ProductBrand:          orderCompleted.Inventory.InventoryProduct.ProductBrand,
// 				ProductModel:          orderCompleted.Inventory.InventoryProduct.ProductModel,
// 				ProductMaterial:       orderCompleted.Inventory.InventoryProduct.ProductMaterial,
// 				ProductDescription:    orderCompleted.Inventory.InventoryProduct.ProductDescription,
// 				ProductImage:          orderCompleted.Inventory.InventoryProduct.ProductImage,
// 				ProductPrice:          orderCompleted.Inventory.InventoryProduct.ProductPrice,
// 				ProductQuantity:       orderCompleted.Inventory.InventoryProduct.ProductQuantity,
// 				ProductSerialNumber:   orderCompleted.Inventory.InventoryProduct.ProductSerialNumber,
// 				ProductCreatedAt:      orderCompleted.Inventory.InventoryProduct.ProductCreatedAt,
// 				ProductUpdatedAt:      orderCompleted.Inventory.InventoryProduct.ProductUpdatedAt,
// 				SupplierName:          orderCompleted.Inventory.InventoryProduct.SupplierName,
// 				Sku:                   orderCompleted.Inventory.InventoryProduct.Sku,
// 			},
// 			CreatedAt: orderCompleted.Inventory.CreatedAt,
// 		},
// 		StoreInfomation: models.StoreInfomation{
// 			ID:               orderCompleted.StoreInfomation.ID,
// 			StoreName:        orderCompleted.StoreInfomation.StoreName,
// 			StoreCode:        orderCompleted.StoreInfomation.StoreCode,
// 			StorePhoneNumber: orderCompleted.StoreInfomation.StorePhoneNumber,
// 			StoreEmail:       orderCompleted.StoreInfomation.StoreEmail,
// 			Permissions:      orderCompleted.StoreInfomation.Permissions,
// 			StoreAddress:     orderCompleted.StoreInfomation.StoreAddress,
// 		},
// 		Employee: models.EmployeeInternalInformation{
// 			ID:           orderCompleted.Employee.ID,
// 			PersonalInfo: orderCompleted.Employee.PersonalInfo,
// 			Area:         orderCompleted.Employee.Area,
// 			Designation:  orderCompleted.Employee.Designation,
// 			Manager:      orderCompleted.Employee.Manager,
// 			Permissions: models.Permissions{
// 				ID:                 orderCompleted.Employee.Permissions.ID,
// 				ReadCashier:        orderCompleted.Employee.Permissions.ReadCashier,
// 				ReadWriteCashier:   orderCompleted.Employee.Permissions.ReadWriteCashier,
// 				ReadAdminUser:      orderCompleted.Employee.Permissions.ReadAdminUser,
// 				ReadWriteAdminUser: orderCompleted.Employee.Permissions.ReadWriteAdminUser,
// 				ReadInventory:      orderCompleted.Employee.Permissions.ReadInventory,
// 				ReadWriteInventory: orderCompleted.Employee.Permissions.ReadWriteInventory,
// 			},
// 		},
// 		CreatedAt: orderCompleted.CreatedAt,
// 	}
// 	query := db.Query(queryCqlInsertOrder, co)
// 	return query.Exec()
// }
