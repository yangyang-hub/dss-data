package model

//定时任务执行记录
type TaskInfo struct {
	Id        int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	TaskName  string `json:"task_name" gorm:"column:task_name"`   //任务名称
	Date      string `json:"date" gorm:"column:date"`             //日期
	SpendTime string `json:"spend_time" gorm:"column:spend_time"` //话费时间
}

func (taskInfo TaskInfo) TableName() string {
	return "task_info"
}
