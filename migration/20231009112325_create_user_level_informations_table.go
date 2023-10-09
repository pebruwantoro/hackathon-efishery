package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009112325_create_user_level_informations_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009112325_create_user_level_informations_table_up,
		Down:      mig_20231009112325_create_user_level_informations_table_down,
	})
}

func mig_20231009112325_create_user_level_informations_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS user_level_informations (
		id bigserial NOT NULL,
		max_exp int8 NOT NULL,
		CONSTRAINT user_level_informations_pkey PRIMARY KEY (id)
	)`).Error

	return err
}

func mig_20231009112325_create_user_level_informations_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS user_level_informations`).Error
	return err
}
