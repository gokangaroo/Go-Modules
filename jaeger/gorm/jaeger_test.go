package gorm

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Test_GormTracing(t *testing.T) {
	closer, err := initJaeger()
	if err != nil {
		t.Fatal(err)
	}
	// flush buffer
	defer closer.Close()

	dsn := "root:123456@tcp(localhost:3306)/repo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// use plugin: hooks
	_ = db.Use(&OpentracingPlugin{})

	// db schema
	_ = db.AutoMigrate(&Product{})

	// trace span start: global tracer
	span := opentracing.StartSpan("gormTracing unit test")
	defer span.Finish()

	// child span: ctx.WithValue
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	// ---> 下面就是GORM的范例
	// db exec set ctx
	session := db.WithContext(ctx)

	// Create
	session.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	session.First(&product, 1)                 // 根据整形主键查找
	session.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	session.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	session.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	session.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	session.Delete(&product, 1)
}
