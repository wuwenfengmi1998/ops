package models

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type File_info struct {
	ID     uint      `gorm:"primaryKey;autoIncrement"`
	Name   string    `gorm:"not null;size:256;index"` // 前端报告的文件名
	Path   string    `gorm:"not null;size:300"`       //
	Sha256 string    `gorm:"not null;size:256;index"` //
	Mime   string    `gorm:"size:64;index"`
	Type   string    `gorm:"size:64;index"`
	Const  uint      `gorm:"default:1;index"`
	Per    uint      `gorm:"default:1"`
	UserID uint      `gorm:"not null;index"`
	Date   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"` // 默认当前时间
}

type User struct {
	ID    uint      `gorm:"primaryKey;autoIncrement"`                // 自增主键
	Name  string    `gorm:"size:100;uniqueIndex"`                    // 唯一约束索引
	Email string    `gorm:"size:255;index"`                          // 字符串长度限制100 索引
	Pass  string    `gorm:"size:128"`                                // 建议存储哈希后的密码
	Date  time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"` // 默认当前时间
}

type User_info struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UserID     uint      `gorm:"not null;uniqueIndex"`
	FirstName  string    `gorm:"size:50;null"`
	Username   string    `gorm:"size:30;null"`
	Birthdate  time.Time `gorm:"type:datetime;null"`
	Gender     string    `gorm:"type:char(1);check:gender IN ('M', 'F', 'U');default:'U'"`
	AvatarPath string    `gorm:"size:255"`
	Region     string    `gorm:"size:50"`
	Language   string    `gorm:"size:10;default:'zh-CN'"`
	CreatedAt  time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;column:created_at"`
}

// var def_user_info = User_info{
// 	ID:0,
// 	UserID:0,
// }

type Cookie struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserID       uint      `gorm:"size:16;not null"`
	Name         string    `gorm:"size:255;not null;index"`
	Value        string    `gorm:"size:255;not null;index"`
	Domain       string    `gorm:"size:255;not null"`
	Path         string    `gorm:"size:255;not null;default:/"`
	ExpiresAt    time.Time `gorm:"type:datetime;index"`
	CreatedAt    time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:datetime;index;not null;default:CURRENT_TIMESTAMP"`
	SecureFlag   bool      `gorm:"not null;default:false"`
	HttpOnly     bool      `gorm:"not null;default:false"`
	SameSite     string    `gorm:"size:10;not null;default:'Lax'"`
	PartitionKey string    `gorm:"size:50"`
}

type Warehouse struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt    time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	Name         string    `gorm:"type:varchar(255);not null;index"`  // 仓库名称
	CreatorID    uint      `gorm:"not null;index"`                    // 创建用户ID（外键）
	Location     string    `gorm:"type:varchar(500);not null"`        // 仓库位置
	Capacity     int       `gorm:"not null;check:capacity >= 0"`      // 仓库总容量
	UsedCapacity int       `gorm:"not null;check:used_capacity >= 0"` // 已用容量
	ImagePath    string    `gorm:"type:varchar(1000)"`                // 仓库图片路径
	Info         string    `gorm:"type:text"`                         // 仓库详细信息

}

// 物品状态枚举
type ItemStatus string

const (
	ItemStatusAvailable   ItemStatus = "available"
	ItemStatusMaintenance ItemStatus = "maintenance"
	ItemStatusRetired     ItemStatus = "retired"
)

// Item 仓库物品表模型
type Item struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`

	// 基本信息
	Name        string `gorm:"type:varchar(255);not null;comment:物品名称"`
	Description string `gorm:"type:text;comment:物品描述"`

	// 库存信息
	WarehouseID uint   `gorm:"not null;index;comment:仓库ID"`
	Quantity    uint   `gorm:"not null;default:0;comment:库存数量"`
	Location    string `gorm:"type:varchar(100);comment:库内位置"`

	// 状态管理
	Status ItemStatus `gorm:"type:varchar(20);not null;default:'available';index;comment:物品状态"`

	// 扩展信息
	BatchNumber    string    `gorm:"type:varchar(50);comment:批次号"`
	ExpirationDate time.Time `gorm:"type:date;comment:失效日期"`

	// 关联关系
	//Warehouse      Warehouse `gorm:"foreignKey:WarehouseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type WarehouseItem struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	CreatedByID   uint      `gorm:"not null;index:idx_created_by;comment:创建用户ID"`
	WarehouseID   uint      `gorm:"not null;index:idx_warehouse;comment:仓库ID"`
	CreatedAt     time.Time `gorm:"autoCreateTime;index:idx_created_at;comment:创建时间"`
	Name          string    `gorm:"type:varchar(255);not null;index:idx_name;comment:物品名称"`
	SerialNumber  string    `gorm:"type:varchar(100);index;null;comment:序列号"`
	Description   string    `gorm:"type:text;null;comment:物品描述"`
	ShelfLocation string    `gorm:"type:varchar(50);null;index:idx_shelf;comment:货架位置"`
	Quantity      int       `gorm:"type:int unsigned;null;comment:物品数量"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime;index:idx_updated_at;comment:最后更新时间"`
	Status        string    `gorm:"type:varchar(10);default:'正常';index:idx_status;comment:物品状态（正常/损坏/维修/报废）"`
	Color         string    `gorm:"type:varchar(16);default:'bg-success';comment:状态颜色"`
	Destiny       string    `gorm:"type:varchar(100);index:idx_destiny;null;comment:物品归宿"`
	ItemValue     int       `gorm:"type:int;null;comment:物品价值（单位：分）"`
}

// 工单表
type Ticket struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`                           // 自增ID
	Title        string    `gorm:"type:varchar(255);index:idx_title,length:191" json:"title"`    // 带索引标题
	Info         string    `gorm:"type:text;not null" json:"info"`                               // 详细信息
	Type         string    `gorm:"type:varchar(16);not null;default:'normal'" json:"type"`       // 工单类型
	CreatedAt    time.Time `gorm:"autoCreateTime;index;not null;" json:"createdAt"`              // 创建日期
	UpdatedAt    time.Time `gorm:"autoCreateTime;index;not null;" json:"updatedAt"`              // 最后更新时间
	UserID       uint      `gorm:"index;not null" json:"userId"`                                 // 创建用户ID
	Status       string    `gorm:"type:varchar(16);index;not null;default:'open'" json:"status"` // 最后状态
	Color        string    `gorm:"type:varchar(32);not null;default:'#3498db'" json:"color"`     // 状态颜色
	ItemID       uint      `gorm:"null" json:"itemId"`                                           // 关联物件ID
	CommentCount uint      `gorm:"not null;default:0" json:"commentCount"`                       // 评论数量
}

func init() {

}

func Init_database() {

	fmt.Println("database_init")
	//var database_config map[string]interface{}=Configs["database"]
	Database_configs = Configs["database"].(map[string]interface{})

	if Database_configs["type"].(string) == "sqlite" {
		//sqlite init
		fmt.Println("sqlite")
		DB, err = gorm.Open(sqlite.Open(Database_configs["path"].(string)), &gorm.Config{})
	} else if Database_configs["type"].(string) == "mysql" {
		//mysql init
		fmt.Println("mysql")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Database_configs["user"].(string), Database_configs["pass"].(string), Database_configs["host"].(string), Database_configs["port"].(string), Database_configs["name"].(string))
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}

	// 检查 users 表是否存在
	// if !DB.Migrator().HasTable(&User{}) {
	// 	// 自动创建表结构
	// 	err := DB.AutoMigrate(&User{})
	// 	if err != nil {
	// 		panic("创建表失败: " + err.Error())
	// 	}
	// 	fmt.Println("users 表已创建")
	// } else {
	// 	fmt.Println("users 表已存在")
	// }

	// 自动创建表结构
	DB.AutoMigrate(&User{})

	DB.AutoMigrate(&User_info{})

	DB.AutoMigrate(&Cookie{})

	DB.AutoMigrate(&Warehouse{})

	DB.AutoMigrate(&Item{})

	DB.AutoMigrate(&WarehouseItem{})

	DB.AutoMigrate(&Ticket{})

	DB.AutoMigrate(&File_info{})

}
