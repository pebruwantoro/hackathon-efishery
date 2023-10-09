package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009162023_roles",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009162023_roles_up,
		Down:      mig_20231009162023_roles_down,
	})
}

func mig_20231009162023_roles_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS roles (
		id INT,
		name VARCHAR(50) NOT NULL,
		CONSTRAINT roles_pkey PRIMARY KEY (id)
	)`).Error

	return err
}

func mig_20231009162023_roles_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS roles`).Error
	return err
}
