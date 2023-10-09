package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009111118_create_objective_users_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009111118_create_objective_users_table_up,
		Down:      mig_20231009111118_create_objective_users_table_down,
	})
}

func mig_20231009111118_create_objective_users_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS objective_users (
		objective_uuid VARCHAR(50) NOT NULL,
		user_uuid VARCHAR(50) NOT NULL,
		CONSTRAINT objective_users_pkey PRIMARY KEY (objective_uuid, user_uuid)
	)`).Error

	return err
}

func mig_20231009111118_create_objective_users_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS objective_users`).Error
	return err
}
