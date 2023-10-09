package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009102314_create_users_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009102314_create_users_table_up,
		Down:      mig_20231009102314_create_users_table_down,
	})
}

func mig_20231009102314_create_users_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS users (
		uuid VARCHAR(50) NOT NULL,
		name VARCHAR(50) NOT NULL,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		access_role VARCHAR(50) NOT NULL,
		password TEXT NOT NULL,
		created_at timestamptz NOT NULL,
		created_by VARCHAR(255) NOT NULL,
		updated_at timestamptz NOT NULL,
		updated_by VARCHAR(255) NOT NULL,
		deleted_at timestamptz NULL,
		deleted_by VARCHAR(255) NULL,
		CONSTRAINT users_pkey PRIMARY KEY (uuid)
	)`).Error

	return err
}

func mig_20231009102314_create_users_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS users`).Error
	return err
}
