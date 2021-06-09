package model

const MigrationHistoryTableName = "migration_history"

type MigrationHistory struct {
	MigrationID string `gorm:"type:varchar(255); column:migration_id"`
}

func (m *MigrationHistory) TableName() string {
	return MigrationHistoryTableName
}
