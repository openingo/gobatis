//This file was generated by xfali/gobatis-cmd at
//2020-02-20 15:55:57.3524245 +0800 CST m=+0.049892501

package test

import (
	"github.com/xfali/gobatis"
	"time"
)

type TestTable struct {
	//TableName gobatis.ModelName `test_table`
	Id         int       `xfield:"id"`
	Username   string    `xfield:"username"`
	Password   string    `xfield:"password"`
	Createtime time.Time `xfield:"createtime"`
}

func (m *TestTable) Select(sess *gobatis.Session) ([]TestTable, error) {
	return SelectTestTable(sess, *m)
}

func (m *TestTable) Count(sess *gobatis.Session) (int64, error) {
	return SelectTestTableCount(sess, *m)
}

func (m *TestTable) Insert(sess *gobatis.Session) (int64, int64, error) {
	return InsertTestTable(sess, *m)
}

func (m *TestTable) Update(sess *gobatis.Session) (int64, error) {
	return UpdateTestTable(sess, *m)
}

func (m *TestTable) Delete(sess *gobatis.Session) (int64, error) {
	return DeleteTestTable(sess, *m)
}
