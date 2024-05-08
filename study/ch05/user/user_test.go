package user

import (
	"fmt"
	"testing"
)

//单元测试中，可以直接写测试函数

// 固定写法
func TestAdd(t *testing.T) {

	re := add(1, 2)
	fmt.Println(re)
}

func BenchmarkAdd(bb *testing.B) {

	var a, b, c int
	a = 123
	b = 443
	c = 677

	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {

			fmt.Println(a, b, c)
		}
	}

}
