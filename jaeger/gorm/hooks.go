package gorm

import (
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

// https://blog.csdn.net/yes169yes123/article/details/108016990

const gormSpanKey = "__gorm_span"

/*
	hook methods
*/

func before(db *gorm.DB) {
	name := "gorm"
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, name)

	// store span
	db.InstanceSet(gormSpanKey, span)
}

func after(db *gorm.DB) {
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}

	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	// Finish
	defer span.Finish()

	if db.Error != nil {
		span.LogFields(tracerLog.Error(db.Error))
	}

	// sql --> 写法来源GORM V2的日志
	span.LogFields(tracerLog.String("sql",
		db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
}
