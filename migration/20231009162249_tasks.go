package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009162249_tasks",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009162249_tasks_up,
		Down:      mig_20231009162249_tasks_down,
	})
}

func mig_20231009162249_tasks_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL NOT NULL,
		objective_id INT,
		subtask_id INT,
		name VARCHAR(255) NOT NULL,
		description TEXT NULL,
		point INT NULL,
		user_id INT NULL,
		status VARCHAR(50) NULL,
		start_date timestamptz DEFAULT NULL,
		end_date timestamptz DEFAULT NULL,
		due_date timestamptz DEFAULT NULL,
		created_at timestamptz NOT NULL,
		created_by VARCHAR(255) NOT NULL,
		updated_at timestamptz NOT NULL,
		updated_by VARCHAR(255) NOT NULL,
		deleted_at timestamptz DEFAULT NULL,
		deleted_by VARCHAR(255) DEFAULT NULL,
		CONSTRAINT tasks_pkey PRIMARY KEY (id),
		CONSTRAINT fk_objective_id FOREIGN KEY(objective_id) REFERENCES objectives(id),
		CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
	)`).Error

	return err
}

func mig_20231009162249_tasks_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS tasks`).Error
	return err
}
