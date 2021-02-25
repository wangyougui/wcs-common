package wcs

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// 库位实时信息表
type BizSiteInfo struct {
	gorm.Model
	SiteCode	string	`gorm:"column:site_code;type:varchar(60);unique" json:"siteCode"`      	//comment:位置编码
	SiteGrid	string	`gorm:"not null;type:varchar(60);column:site_grid;" json:"siteGrid"`	//comment:位置坐标
	SiteArea	string	`gorm:"not null;type:varchar(60);column:site_area;" json:"siteArea"`   	//comment:位置区域
	State		string	`gorm:"not null;type:varchar(60);column:state" json:"state"`			//comment:当前状态
	PalletCode	string	`gorm:"type:varchar(60);column:pallet_code" json:"palletCode"`			//comment:货架编码
	TaskCode	string	`gorm:"type:varchar(60);column:task_code" json:"taskCode"`				//comment:当前任务编码
	ExtData		json.RawMessage	`gorm:"type:json;column:ext_data" json:"extData"`				//comment:扩展信息
	AgvCode		string	`gorm:"type:varchar(60);column:agv_code" json:"agvCode"`				//comment:小车编码（有小车临停时有值）
	FbArea		string	`gorm:"type:varchar(60);column:fb_area" json:"fbArea"`					//comment:所在区域
	MtrMpCode	string	`gorm:"type:varchar(60);column:mtr_mp_code" json:"mtrMpCode"`			//comment:对应落料机台编码
	Floor		int		`gorm:"column:floor" json:"floor"`										//comment:楼层
	SiteType	string	`gorm:"type:varchar(150);column:site_type" json:"siteType"`				//comment:站点类型
	StationCode	string	`gorm:"type:varchar(60);column:station_code" json:"stationCode"`		//comment:工位号
	Position	int		`gorm:"column:position" json:"position"`								//comment:相对位置(相对提升机/电梯/输送线的位置) 1: 提升机左  2:提升机右 0：提升机内部
	WhId		string	`gorm:"type:varchar(60);column:wh_id" json:"whId"`						//comment:仓库ID(wms确定)
	StorageRoomType		string	`gorm:"type:varchar(250);column:storage_room_type" json:"storageRoomType"`	//comment:库房类型(wms确定)
	StorageRoomTag		string	`gorm:"type:varchar(250);column:storage_room_tag" json:"storageRoomTag"`	//comment:仓库ID(wms确定)
	SiteGroup			string	`gorm:"type:varchar(60);column:site_group" json:"siteGroup"`				//comment:货位分组
	GroupCount			int		`gorm:"column:group_count" json:"groupCount"`								//comment:货位组货位个数
	HoisterMpCode		string	`gorm:"type:varchar(160);column:hoister_mp_code" json:"hoisterMpCode"`		//comment:提升机编码
	ConveyorMpCode		string	`gorm:"type:varchar(160);column:conveyor_mp_code" json:"conveyorMpCode"` 	//comment:输送线编码(链条式)
	Picking				int		`gorm:"column:picking" json:"picking"`										//comment:是否可拣货(出库)
	Putaway				int		`gorm:"column:putaway" json:"putaway"`										//comment:是否可上架（入库）
	LeftSiteCode		string	`gorm:"type:varchar(160);column:left_site_code" json:"leftSiteCode"`		//comment:当前站点左方站点编号
	RightSiteCode		string	`gorm:"type:varchar(160);column:right_site_code" json:"rightSiteCode"`		//comment:当前站点右方站点编号
}

func (BizSiteInfo) TableName() string {
	return "biz_site_info"
}
