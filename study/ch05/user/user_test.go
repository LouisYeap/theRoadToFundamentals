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
