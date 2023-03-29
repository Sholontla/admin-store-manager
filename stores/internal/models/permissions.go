package models

import "github.com/gocql/gocql"

type Permissions struct {
	ID                 gocql.UUID `cql:"_id" json:"id"`
	ReadCashier        bool       `cql:"read_cashier" json:"read_cashier"`
	ReadWriteCashier   bool       `cql:"read_write_cashier" json:"read_write_cashier"`
	ReadAdminUser      bool       `cql:"read_admin_user" json:"read_admin_user"`
	ReadWriteAdminUser bool       `cql:"read_write_admin_user" json:"read_write_admin_user"`
	ReadInventory      bool       `cql:"read_inventory" json:"read_inventory"`
	ReadWriteInventory bool       `cql:"read_write_inventory" json:"read_write_inventory"`
}
