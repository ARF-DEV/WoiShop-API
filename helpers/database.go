package helpers

import (
	"azura-lab-intern/study-case-1/seeder"
	"database/sql"
)

func MigrateDB(DB *sql.DB) {
	seeder.MigrateCart(DB)
	seeder.MigrateCategory(DB)
	seeder.MigrateDeliveryMethod(DB)
	seeder.MigrateOrder(DB)
	seeder.MigrateProduct(DB)
	seeder.MigrateStore(DB)
	seeder.MigrateUser(DB)
}
