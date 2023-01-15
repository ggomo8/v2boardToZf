package tools

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
var db *gorm.DB

type V2ServerV2Ray struct {
	Id              int            `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	GroupId         string         `gorm:"column:group_id;type:varchar(255);NOT NULL" json:"group_id"`
	RouteId         sql.NullString `gorm:"column:route_id;type:varchar(255)" json:"route_id"`
	Name            string         `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	ParentId        sql.NullInt32  `gorm:"column:parent_id;type:int(11)" json:"parent_id"`
	Host            string         `gorm:"column:host;type:varchar(255);NOT NULL" json:"host"`
	Port            string         `gorm:"column:port;type:varchar(11);NOT NULL" json:"port"`
	ServerPort      int            `gorm:"column:server_port;type:int(11);NOT NULL" json:"server_port"`
	Tls             int            `gorm:"column:tls;type:tinyint(4);default:0;NOT NULL" json:"tls"`
	Tags            sql.NullString `gorm:"column:tags;type:varchar(255)" json:"tags"`
	Rate            string         `gorm:"column:rate;type:varchar(11);NOT NULL" json:"rate"`
	Network         string         `gorm:"column:network;type:text;NOT NULL" json:"network"`
	Rules           sql.NullString `gorm:"column:rules;type:text" json:"rules"`
	NetworkSettings sql.NullString `gorm:"column:networkSettings;type:text" json:"networkSettings"`
	TlsSettings     sql.NullString `gorm:"column:tlsSettings;type:text" json:"tlsSettings"`
	RuleSettings    sql.NullString `gorm:"column:ruleSettings;type:text" json:"ruleSettings"`
	DnsSettings     sql.NullString `gorm:"column:dnsSettings;type:text" json:"dnsSettings"`
	Show            int            `gorm:"column:show;type:tinyint(1);default:0;NOT NULL" json:"show"`
	Sort            sql.NullInt32  `gorm:"column:sort;type:int(11)" json:"sort"`
	CreatedAt       int            `gorm:"column:created_at;type:int(11);NOT NULL" json:"created_at"`
	UpdatedAt       int            `gorm:"column:updated_at;type:int(11);NOT NULL" json:"updated_at"`
}

func (m *V2ServerV2Ray) TableName() string {
	return "v2_server_v2ray"
}

type V2ServerShadowsocks struct {
	Id           int            `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	GroupId      string         `gorm:"column:group_id;type:varchar(255);NOT NULL" json:"group_id"`
	RouteId      sql.NullString `gorm:"column:route_id;type:varchar(255)" json:"route_id"`
	ParentId     sql.NullInt32  `gorm:"column:parent_id;type:int(11)" json:"parent_id"`
	Tags         sql.NullString `gorm:"column:tags;type:varchar(255)" json:"tags"`
	Name         string         `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Rate         string         `gorm:"column:rate;type:varchar(11);NOT NULL" json:"rate"`
	Host         string         `gorm:"column:host;type:varchar(255);NOT NULL" json:"host"`
	Port         string         `gorm:"column:port;type:varchar(11);NOT NULL" json:"port"`
	ServerPort   int            `gorm:"column:server_port;type:int(11);NOT NULL" json:"server_port"`
	Cipher       string         `gorm:"column:cipher;type:varchar(255);NOT NULL" json:"cipher"`
	Obfs         sql.NullString `gorm:"column:obfs;type:char(11)" json:"obfs"`
	ObfsSettings sql.NullString `gorm:"column:obfs_settings;type:varchar(255)" json:"obfs_settings"`
	Show         int            `gorm:"column:show;type:tinyint(4);default:0;NOT NULL" json:"show"`
	Sort         sql.NullInt32  `gorm:"column:sort;type:int(11)" json:"sort"`
	CreatedAt    int            `gorm:"column:created_at;type:int(11);NOT NULL" json:"created_at"`
	UpdatedAt    int            `gorm:"column:updated_at;type:int(11);NOT NULL" json:"updated_at"`
}

// ForWard 读取json文件的结构体
type ForWard struct {
	Forwards []Forwards `json:"forwards"`
}
type Forwards struct {
	ID                int           `json:"id"`
	UserID            interface{}   `json:"userId"`
	PortID            interface{}   `json:"portId"`
	ServerID          interface{}   `json:"serverId"`
	LocalPort         int           `json:"localPort"`
	ServerName        interface{}   `json:"serverName"`
	ServerHost        interface{}   `json:"serverHost"`
	ServerDisplayHost interface{}   `json:"serverDisplayHost"`
	InternetPort      int           `json:"internetPort"`
	Username          interface{}   `json:"username"`
	RemoteIP          string        `json:"remoteIp"`
	RemoteHost        string        `json:"remoteHost"`
	DataUsage         interface{}   `json:"dataUsage"`
	DataUsageInput    interface{}   `json:"dataUsageInput"`
	ForwardType       int           `json:"forwardType"`
	RemotePort        int           `json:"remotePort"`
	CreateTime        string        `json:"createTime"`
	UpdateTime        string        `json:"updateTime"`
	Deleted           bool          `json:"deleted"`
	Disabled          bool          `json:"disabled"`
	Iperf3            bool          `json:"iperf3"`
	State             int           `json:"state"`
	Remark            string        `json:"remark"`
	SendProxy         bool          `json:"sendProxy"`
	AcceptProxy       bool          `json:"acceptProxy"`
	SpeedLimit        int           `json:"speedLimit"`
	BalanceList       []interface{} `json:"balanceList"`
	IsBalance         bool          `json:"isBalance"`
	BalanceType       interface{}   `json:"balanceType"`
	Reason            string        `json:"reason"`
	HasDynamic        bool          `json:"hasDynamic"`
	IsServer          bool          `json:"isServer"`
	Secure            bool          `json:"secure"`
	UseServerCert     bool          `json:"useServerCert"`
	CustomHost        interface{}   `json:"customHost"`
	CustomSni         interface{}   `json:"customSni"`
	CustomPath        interface{}   `json:"customPath"`
	Crt               string        `json:"crt"`
	Key               string        `json:"key"`
	Ping              string        `json:"ping"`
	IsLine            bool          `json:"isLine"`
	StartLogs         interface{}   `json:"startLogs"`
}

type WsConfig struct {
	Path    string  `json:"path"`
	Headers Headers `json:"headers"`
}
type Headers struct {
	Host string `json:"Host"`
}

// 包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {
	MYSQL_IP := Cfg.Section("").Key("MYSQL_IP").MustString("127.0.0.1")
	MYSQL_PORT := Cfg.Section("").Key("MYSQL_PORT").MustString("3306")
	MYSQL_USER_NAME := Cfg.Section("").Key("MYSQL_USER_NAME").MustString("username..")
	MYSQL_PASSWORD := Cfg.Section("").Key("MYSQL_PASSWORD").MustString("SdhxpDsDiRb7Ztsy")
	MYSQL_DB_NAME := Cfg.Section("").Key("MYSQL_DB_NAME").MustString("DBNAME")

	dsn := MYSQL_USER_NAME + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_IP + ":" + MYSQL_PORT + ")/" +
		MYSQL_DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("数据库连接信息: ", dsn)
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

// GetDB 获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
// 不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return db
}
