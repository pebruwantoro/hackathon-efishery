package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009111547_create_user_points_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009111547_create_user_points_table_up,
		Down:      mig_20231009111547_create_user_points_table_down,
	})
}

func mig_20231009111547_create_user_points_table_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS user_points (
		uuid VARCHAR(50) NOT NULL,
		user_uuid VARCHAR(50) NOT NULL,
		level int8 NOT NULL,
		total_hp int8 NOT NULL,
		total_exp int8 NOT NULL,
		CONSTRAINT user_points_pkey PRIMARY KEY (uuid)
	)`).Error

	return err
}

func mig_20231009111547_create_user_points_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS user_points`).Error
	return err
}
