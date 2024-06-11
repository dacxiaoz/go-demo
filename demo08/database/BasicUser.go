package database

type BasicUser struct {
	Id       uint   `gorm:"column:id;type:bigint;primary_key;auto_increment" json:"id" binding:"required"`
	Account  string `gorm:"column:account;type:varchar(11); not null" json:"account" binding:"required"`
	Password string `gorm:"column:password;type:varchar(16); not null" json:"password" binding:"required"`
	Name     string `gorm:"column:name;type:varchar(200); not null" json:"name" binding:"required"`
	State    string `gorm:"column:state;type:varchar(20); not null" json:"state" binding:"required"`
	Email    string `gorm:"column:email;type:varchar(40); not null" json:"email" binding:"required"`
	Address  string `gorm:"column:address;type:varchar(200); not null" json:"address" binding:"required"`
}

// type GeekProduct struct {
// 	gorm.Model        // 主键
// 	Title      string `gorm:"type:varchar(20); not null" json:"title" binding:"required"`
// 	Remark     string `gorm:"type:varchar(20); not null" json:"remark" binding:"required"`
// }
