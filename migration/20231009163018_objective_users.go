package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009163018_objective_users",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009163018_objective_users_up,
		Down:      mig_20231009163018_objective_users_down,
	})
}

func mig_20231009163018_objective_users_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS objective_users (
		id INT NOT NULL,
		objective_id INT NOT NULL,
		user_id INT NOT NULL,
		CONSTRAINT objective_users_pkey PRIMARY KEY (id),
		CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
		CONSTRAINT fk_objective_id FOREIGN KEY(objective_id) REFERENCES objectives(id)
	)`).Error

	return err
}

func mig_20231009163018_objective_users_down(tx *gorm.DB) error {
	return nil
}
