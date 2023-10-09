package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009112600_create_levels_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009112600_create_levels_table_up,
		Down:      mig_20231009112600_create_levels_table_down,
	})
}

func mig_20231009112600_create_levels_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS levels (
		id bigserial NOT NULL,
		name VARCHAR(50) NOT NULL,
		point int8 NOT NULL,
		CONSTRAINT levels_pkey PRIMARY KEY (id)
	)`).Error

	return err
}

func mig_20231009112600_create_levels_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS levels`).Error
	return err
}
