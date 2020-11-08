/*
 * @Author: 光城
 * @Date: 2020-11-08 15:19:03
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-08 16:15:32
 * @Description:
 * @FilePath: /go-talent/code/utest/db/db_test.go
 */
package db

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"

	"bou.ke/monkey"
)

func TestGetValue(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq(1)).Return("我是1的value", nil)

	if v, err := GetValue(m, 1); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestGetValue1(t *testing.T) {
	var h *Handler
	monkey.PatchInstanceMethod(reflect.TypeOf(h), "Get", func(handler *Handler, key int) (string, error) {
		return "我是1的value", nil
	})
	if v, err := GetValue(h, 1); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
