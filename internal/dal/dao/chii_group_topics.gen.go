// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameGroupTopic = "chii_group_topics"

// GroupTopic mapped from table <chii_group_topics>
type GroupTopic struct {
	ID          uint32 `gorm:"column:grp_tpc_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	GroupID     uint32 `gorm:"column:grp_tpc_gid;type:mediumint(8) unsigned;not null"`
	UID         uint32 `gorm:"column:grp_tpc_uid;type:mediumint(8) unsigned;not null"`
	Title       string `gorm:"column:grp_tpc_title;type:varchar(80);not null"`
	CreatedTime uint32 `gorm:"column:grp_tpc_dateline;type:int(10) unsigned;not null"`
	UpdatedTime uint32 `gorm:"column:grp_tpc_lastpost;type:int(10) unsigned;not null"`
	Replies     uint32 `gorm:"column:grp_tpc_replies;type:mediumint(8) unsigned;not null"`
	State       uint8  `gorm:"column:grp_tpc_state;type:tinyint(1) unsigned;not null"`
	Display     uint8  `gorm:"column:grp_tpc_display;type:tinyint(1) unsigned;not null;default:1"`
}

// TableName GroupTopic's table name
func (*GroupTopic) TableName() string {
	return TableNameGroupTopic
}
