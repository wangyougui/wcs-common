package wcs

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

// OFS任务表
type BizOfsTask struct {
	gorm.Model
	TaskCode		string	`gorm:"type:varchar(60);column:task_code；unique" json:"taskCode"`		//comment:任务编码
	RcsSys			string	`gorm:"column:rcs_sys;type:varchar(60);" json:"rcsSys"`      			//comment:对应RCS系统
	FinishedTime	time.Time	`gorm:"column:finished_time;" json:"finishedTime"`				//comment:任务完成时间
	TaskType	string	`gorm:"not null;type:varchar(60);column:task_type;" json:"taskType"`   	//comment:任务类型
	TaskStatus	string	`gorm:"not null;type:varchar(60);column:task_status" json:"taskStatus"`	//comment:当前任务状态
	CurrStep	string	`gorm:"type:varchar(40);column:curr_step" json:"currStep"`				//comment:任务当前节点
	ExtData		json.RawMessage	`gorm:"type:json;column:ext_data" json:"extData"`				//comment:扩展信息
	InitialCode	string	`gorm:"type:varchar(60);column:initial_code" json:"initialCode"`		//comment:任务起点位置编码
	InitialGrid	string	`gorm:"type:varchar(60);column:initial_grid" json:"initialGrid"`		//comment:任务起点位置坐标
	TargetCode	string	`gorm:"type:varchar(60);column:target_code" json:"targetCode"`			//comment:任务终点位置编码
	TargetGrid	string	`gorm:"type:varchar(60);column:target_grid" json:"targetGrid"`			//comment:任务终点位置坐标
	Remark		string	`gorm:"type:longtext;column:remark" json:"remark"`						//comment:备注信息
	AgvCode		string	`gorm:"type:varchar(60);column:agv_code" json:"agvCode"`				//comment:小车号
	ErrorTime	time.Time	`gorm:"column:error_time;" json:"errorTime"`						//comment:任务出错时间
	CancelTime	time.Time	`gorm:"column:cancel_time;" json:"cancelTime"`						//comment:任务取消时间
	SnakerOrderId		string	`gorm:"type:varchar(100);column:snaker_order_id" json:"snakerOrderId"`			//comment:流程实例ID
	SnakerProcessName	string	`gorm:"type:varchar(250);column:snaker_process_name" json:"snakerProcessName"`	//comment:流程名称
	ProcessDisplayName	string	`gorm:"type:varchar(250);column:process_display_name" json:"processDisplayName"`//comment:流程显示名称
	IniData		json.RawMessage	`gorm:"type:json;column:ini_data" json:"iniData"`								//comment:初始化数据
	ParentTaskCode		string	`gorm:"type:varchar(250);column:parent_task_code" json:"parentTaskCode"`		//comment:上级任务编号
	ParentNodeActorId	string	`gorm:"type:varchar(250);column:parent_node_actor_id" json:"parentNodeActorId"`	//comment:上级任务节点ID
	Priority			int		`gorm:"column:picking" json:"priority"`										//comment:优先级
	TaskGroup			string	`gorm:"type:varchar(160);column:task_group" json:"taskGroup"`		//comment:任务分组
	CbTaskId			string	`gorm:"type:varchar(160);column:cb_task_id" json:"cbTaskId"`		//comment:最新回调节点任务ID
	CbData				string	`gorm:"type:varchar(160);column:cb_data" json:"cbData"`				//comment:最新回调数据
	ErrorTaskId			string	`gorm:"type:varchar(250);column:error_task_id" json:"errorTaskId"`	//comment:出错节点任务ID
	CbAddr				string	`gorm:"type:varchar(160);column:cb_addr" json:"cbAddr"`				//comment:回调地址
	InitFloor			int		`gorm:"column:init_floor" json:"initFloor"`							//comment:任务起点楼层
	TargetFloor			int		`gorm:"column:target_floor" json:"targetFloor"`						//comment:任务终点楼层
}

func (BizOfsTask) TableName() string {
	return "biz_ofs_task"
}

// 向任务表的扩展字体中添加数据。
func (bizOfsTask *BizOfsTask) AddExtData(key string, value interface{}) {
	var extDataMap map[string]interface{}
	if err := json.Unmarshal(bizOfsTask.ExtData, &extDataMap); err != nil {
		extDataMap = make(map[string]interface{})
	}
	extDataMap[key] = value
	if extData, err := json.Marshal(extDataMap); err != nil {
		bizOfsTask.ExtData = extData
	}
}

// 向任务表的扩展字体中添加Map数据。
func (bizOfsTask *BizOfsTask) AddExtDataMap(value map[string]interface{}) {
	var extDataMap map[string]interface{}
	if err := json.Unmarshal(bizOfsTask.ExtData, &extDataMap); err != nil {
		extDataMap = make(map[string]interface{})
	}
	for key, val := range value {
		extDataMap[key] = val
	}
	if extData, err := json.Marshal(extDataMap); err != nil {
		bizOfsTask.ExtData = extData
	}
}

// 将任务表中的扩展字体转换为Map
func (bizOfsTask *BizOfsTask) GetExtDataMap() map[string]interface{} {
	var extDataMap map[string]interface{}
	if err := json.Unmarshal(bizOfsTask.ExtData, &extDataMap); err != nil {
		return extDataMap
	}else {
		extDataMap = make(map[string]interface{})
		return extDataMap
	}
}
