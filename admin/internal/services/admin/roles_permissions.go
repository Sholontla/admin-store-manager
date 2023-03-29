package admin

import (
	"service/admin/case1/internal/models"
)

func RolesPermissionsAssignment(user models.AdminUsers) []string {
	permissionsMap := map[string]string{
		"CreateInventory": "create_inventory",
		"DeleteInventory": "delete_inventory",
		"ReadInventory":   "read_inventory",
		"UpdateInventory": "update_inventory",
		"CreateStores":    "create_stores",
		"DeleteStores":    "delete_stores",
		"ReadStores":      "read_stores",
		"UpdateStores":    "update_stores",
	}

	inventoryPer := []string{}

	if user.Permissions.CreateInventory {
		inventoryPer = append(inventoryPer, permissionsMap["CreateInventory"])
	}
	if user.Permissions.DeleteInventory {
		inventoryPer = append(inventoryPer, permissionsMap["DeleteInventory"])
	}
	if user.Permissions.ReadInventory {
		inventoryPer = append(inventoryPer, permissionsMap["ReadInventory"])
	}
	if user.Permissions.UpdateInventory {
		inventoryPer = append(inventoryPer, permissionsMap["UpdateInventory"])
	}
	if user.Permissions.CreateStore {
		inventoryPer = append(inventoryPer, permissionsMap["CreateStore"])
	}
	if user.Permissions.DeleteStore {
		inventoryPer = append(inventoryPer, permissionsMap["DeleteStore"])
	}
	if user.Permissions.ReadStore {
		inventoryPer = append(inventoryPer, permissionsMap["ReadStore"])
	}
	if user.Permissions.UpdateStore {
		inventoryPer = append(inventoryPer, permissionsMap["UpdateStore"])
	}

	return inventoryPer
}
