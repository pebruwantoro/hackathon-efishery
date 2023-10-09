package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009163346_user_levels",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009163346_user_levels_up,
		Down:      mig_20231009163346_user_levels_down,
	})
}

func mig_20231009163346_user_levels_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS user_levels (
		id SERIAL NOT NULL,
		user_id INT NOT NULL,
		level_id INT NOT NULL,
		health_point int8 NOT NULL,
		experience_point int8 NOT NULL,
		CONSTRAINT user_levels_pkey PRIMARY KEY (id),
		CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
		CONSTRAINT fk_level_id FOREIGN KEY(level_id) REFERENCES levels(id)
	)`).Error

	return err
}

func mig_20231009163346_user_levels_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS user_levels`).Error
	return err
}
