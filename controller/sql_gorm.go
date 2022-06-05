package controller

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GLOBAL_DB *gorm.DB

func SqlLianjie() {
	//连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:Yfl20020301@tcp(rm-bp1bux16i07sc08d50o.mysql.rds.aliyuncs.com:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: false,
	})

	GLOBAL_DB = db

	if err != nil {
		fmt.Println(err)
	}
}