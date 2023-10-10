package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009162159_objectives",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009162159_objectives_up,
		Down:      mig_20231009162159_objectives_down,
	})
}

func mig_20231009162159_objectives_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS objectives (
		id SERIAL NOT NULL,
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
		deleted_at timestamptz DEFAULT NULL,
		deleted_by VARCHAR(255) DEFAULT NULL,
		CONSTRAINT objectives_pkey PRIMARY KEY (id)
	)`).Error

	return err
}

func mig_20231009162159_objectives_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS objectives`).Error
	return err
}
