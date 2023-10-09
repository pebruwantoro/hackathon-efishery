package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009163256_levels",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009163256_levels_up,
		Down:      mig_20231009163256_levels_down,
	})
}

func mig_20231009163256_levels_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS levels (
		id INT NOT NULL,
		level INT NOT NULL,
		max_exp int8 NOT NULL,
		CONSTRAINT levels_pkey PRIMARY KEY (id)
	)`).Error

	return err
}

func mig_20231009163256_levels_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS levels`).Error
	return err
}
