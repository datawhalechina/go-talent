/*
 * @Author: 光城
 * @Date: 2020-11-08 15:19:03
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-08 16:11:07
 * @Description:
 * @FilePath: /go-talent/code/utest/db/db.go
 */
package db

import (
	"errors"
)

type DB interface {
	Get(key int) (string, error)
}
type Handler struct {
}

func (handler *Handler) Get(key int) (string, error) {
	// TODO
	return "", nil
}

func GetValue(db DB, key int) (string, error) {
	value, err := db.Get(key)
	if err != nil {
		return "", errors.New("fail")
	}
	return value, nil
}
