package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{}) //确保在数据库有我们之前的Product结构，如果字段有缺失则添加，不会删除已有的字段

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	//db.First(&product, 1)                 // 根据整型主键查找，查找主键id为1的First记录，并存放到product结构体中
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录并且存放到product结构体

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段（两种方式）
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)//如果product的id不是1则没有任何记录会被删除
	db.Delete(&product) //直接删除product
}
