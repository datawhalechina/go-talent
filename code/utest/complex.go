/*
 * @Author: 光城
 * @Date: 2020-11-08 14:36:52
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-08 14:46:17
 * @Description:
 * @FilePath: /go-talent/code/unit_test/complex.go
 */
package utest

type Complex struct {
	Real float32
	Imag float32
}

func Add(a, b Complex) *Complex {
	return &Complex{
		Real: a.Real + b.Real,
		Imag: a.Imag + b.Imag,
	}
}
