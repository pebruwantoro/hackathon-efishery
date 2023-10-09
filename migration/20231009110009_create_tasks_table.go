package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009110009_create_tasks_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009110009_create_tasks_table_up,
		Down:      mig_20231009110009_create_tasks_table_down,
	})
}

func mig_20231009110009_create_tasks_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		uuid VARCHAR(50) NOT NULL,
		objective_uuid VARCHAR(50) NOT NULL,
		subtask_uuid VARCHAR(50) NULL,
		name VARCHAR(255) NOT NULL,
		description TEXT NULL,
		level VARCHAR(50) NULL,
		point int8 NULL,
		assignee VARCHAR(50) NULL,
		status VARCHAR(50) NULL,
		start_date timestamptz NOT NULL,
		end_date timestamptz NOT NULL,
		due_date timestamptz NOT NULL,
		created_at timestamptz NOT NULL,
		created_by VARCHAR(255) NOT NULL,
		updated_at timestamptz NOT NULL,
		updated_by VARCHAR(255) NOT NULL,
		deleted_at timestamptz NOT NULL,
		deleted_by VARCHAR(255) NULL,
		CONSTRAINT tasks_pkey PRIMARY KEY (uuid)
	)`).Error

	return err
}

func mig_20231009110009_create_tasks_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS tasks`).Error
	return err
}
