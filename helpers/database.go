package helpers

import (
	"azura-lab-intern/study-case-1/seeder"
	"database/sql"
)

func MigrateDB(DB *sql.DB) {
	seeder.MigrateUser(DB)
	seeder.MigrateDeliveryMethod(DB)
	seeder.MigrateCart(DB)
	seeder.MigrateCategory(DB)
	seeder.MigrateStore(DB)
	seeder.MigrateProduct(DB)
	seeder.MigrateOrder(DB)
}
