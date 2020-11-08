/*
 * @Author: 光城
 * @Date: 2020-11-08 14:36:52
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-08 15:07:00
 * @Description:
 * @FilePath: /go-talent/code/utest/complex_test.go
 */
package utest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a Complex
		b Complex
	}
	tests := []struct {
		name string
		args args
		want *Complex
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				a: Complex{
					Real: 1.0,
					Imag: 2.0,
				},
				b: Complex{
					Real: 1.0,
					Imag: 1.0,
				},
			},
			want: &Complex{
				Real: 2.0,
				Imag: 3.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkComplex(t *testing.B) {

	for i := 0; i < t.N; i++ {
		fmt.Sprintf("hello")
	}
}
