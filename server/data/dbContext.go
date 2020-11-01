package data

import (
	"github.com/jinzhu/gorm"
)

type Context struct {
	GormSession *gorm.DB
}

func (c *Context) Close() {
	c.GormSession.Close()
}

func (c *Context) DbCollection(name string) *gorm.DB {
	return c.GormSession
}

func NewContext() *Context {
	session := GetSession()
	context := &Context{
		GormSession: session,
	}
	return context
}
