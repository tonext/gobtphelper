package gobtphelper

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetZoneCode() string {
	zoneCode := GetConfig("zone")
	zoneCodeEnv := GetArgValue("-zone=")
	if zoneCodeEnv != "" {
		zoneCode = zoneCodeEnv
	}
	if zoneCode == "" {
		log.Println("没有配置 -zone= 变量")
	}
	return zoneCode
}

func GetDB(database string) *gorm.DB {
	host := GetSectionConfig("mysql", "host")
	user := GetSectionConfig("mysql", "user")
	password := GetSectionConfig("mysql", "password")
	port := GetSectionConfig("mysql", "port")

	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)

	// 先连接到 MySQL 服务器，但是不指定数据库名
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, // 打印 SQL
	})
	if err != nil {
		log.Fatalf("数据库连接失败1: %v", err)
		return nil
	}

	// 获取底层的 *sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("数据库连接失败2: %v", err)
		return nil
	}
	defer sqlDB.Close()

	// 检查数据库是否已经存在
	var schemaName string
	err = sqlDB.QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", database).Scan(&schemaName)
	if err == sql.ErrNoRows {
		schemaName = ""
	} else if err != nil {
		log.Fatalf("查询数据库是否存在时出错: %v", err)
	}

	// 如果数据库不存在，则创建它，并指定字符集和排序规则
	if schemaName == "" {
		log.Printf("Database %s does not exist, creating now...", database)
		createDBSQL := fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci", database)
		if _, err := sqlDB.Exec(createDBSQL); err != nil {
			log.Fatalf("创建数据库时出错: %v", err)
		}
	}

	// 重新连接数据库，这次指定数据库名
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	db, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{
		QueryFields: true, // 打印 SQL
	})
	if err != nil {
		log.Fatalf("数据库连接失败4: %v", err)
		return nil
	}

	// 获取底层的 *sql.DB 连接
	sqlDB1, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层 *sql.DB 连接时出错: %v", err)
	}

	// 设置连接池参数
	s1 := GetSectionConfig("mysql", "max_open_conns")
	s2 := GetSectionConfig("mysql", "max_idle_conns")
	s3 := GetSectionConfig("mysql", "conn_max_life_time")
	if s1 == "" {
		s1 = "20"
	}
	if s2 == "" {
		s2 = "10"
	}
	if s3 == "" {
		s3 = "5"
	}
	n1, _ := strconv.Atoi(s1)
	n2, _ := strconv.Atoi(s2)
	n3, _ := strconv.Atoi(s3)
	sqlDB1.SetMaxOpenConns(n1)                                 // 设置最大打开的连接数
	sqlDB1.SetMaxIdleConns(n2)                                 // 设置最大空闲连接数
	sqlDB1.SetConnMaxLifetime(time.Duration(n3) * time.Minute) // 设置连接的最大生命周期

	log.Printf("数据库连接成功! database = %s", database)

	return db
}
