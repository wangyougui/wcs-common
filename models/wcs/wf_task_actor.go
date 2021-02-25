package wcs

// 任务参与者表
type WfTaskActor struct {
	TaskId	string	`gorm:"type:varchar(32);column:task_id；unique" json:"taskId"`		//comment:任务ID
	ActorId	string	`gorm:"type:varchar(100);column:actor_Id；unique" json:"actorId"`	//comment:参与者ID
}

func (wfTaskActor *WfTaskActor) TableName() string {
	return "wf_task_actor"
}
