package model

import "time"

type TrJobsMdl struct {
	JobID       uint64    `gorm:"column:JobID;primaryKey;autoIncrement:true"`
	JobName     string    `gorm:"column:JobName;size:100;not null;uniqueIndex:idx_job_name"`
	APIUrl      string    `gorm:"column:APIUrl;size:100;not null"`
	IsExecuted  bool      `gorm:"column:IsExecuted;not null"`
	ExecutedAt  time.Time `gorm:"column:ExecutedAt;type:datetime;not null"`
	CreatedTime time.Time `gorm:"column:CreatedTime;type:datetime;not null"`
}

func (TrJobsMdl) TableName() string {
	return "TrJobs"
}