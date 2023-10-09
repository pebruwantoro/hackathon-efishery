package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009105703_create_objectives_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009105703_create_objectives_table_up,
		Down:      mig_20231009105703_create_objectives_table_down,
	})
}

func mig_20231009105703_create_objectives_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS objectives (
		uuid VARCHAR(50) NOT NULL,
		name VARCHAR(255) NOT NULL,
		description TEXT NULL,
		weight int8 NOT NULL,
		start_date timestamptz NOT NULL,
		end_date timestamptz NOT NULL,
		due_date timestamptz NOT NULL,
		created_at timestamptz NOT NULL,
		created_by VARCHAR(255) NOT NULL,
		updated_at timestamptz NOT NULL,
		updated_by VARCHAR(255) NOT NULL,
		deleted_at timestamptz NOT NULL,
		deleted_by VARCHAR(255) NULL,
		CONSTRAINT objectives_pkey PRIMARY KEY (uuid)
	)`).Error

	return err
}

func mig_20231009105703_create_objectives_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS objectives`).Error
	return err
}
