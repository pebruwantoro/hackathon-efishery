package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009112049_create_roles_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009112049_create_roles_table_up,
		Down:      mig_20231009112049_create_roles_table_down,
	})
}

func mig_20231009112049_create_roles_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS roles (
		id bigserial NOT NULL,
		name VARCHAR(50) NOT NULL,
		CONSTRAINT roles_pkey PRIMARY KEY (id)
	)`).Error

	return err
}

func mig_20231009112049_create_roles_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS roles`).Error
	return err
}
