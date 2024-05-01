package main

import (
	"encoding/json"
	"fmt"
)

// Result 结构体定义
type Result struct {
	ID                int     `json:"id"`
	Station           int     `json:"station"`
	PlanGenerated     float64 `json:"plan_generated"`
	DT                int     `json:"dt"`
	Owner             int     `json:"owner"`
	PlanYearAllEnergy float64 `json:"plan_year_all_energy"`
	PlanYear          int     `json:"plan_year"`
}

type Response struct {
	Code   string   `json:"code"`
	Result []Result `json:"result"`
}

func main() {
	// 给定的 JSON 数据
	jsonData := `{
        "code": "00000",
        "result": [
            {
                "id": 17,
                "station": 54,
                "plan_generated": 4.706,
                "dt": 202212,
                "owner": 2,
                "plan_year_all_energy": 60.0,
                "plan_year": 2025
            },
            {
                "id": 16,
                "station": 54,
                "plan_generated": 4.706,
                "dt": 202211,
                "owner": 2,
                "plan_year_all_energy": 60.0,
                "plan_year": 2025
            },
            {
                "id": 15,
                "station": 54,
                "plan_generated": 4.706,
                "dt": 202210,
                "owner": 2,
                "plan_year_all_energy": 60.0,
                "plan_year": 2025
            }
        ]
    }`

	// 解析 JSON 数据
	var response Response
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		fmt.Println("解析 JSON 数据时出错:", err)
		return
	}

	//fmt.Println(response.Result)

	//在 Go 语言的 fmt 包中，Printf 函数用于格式化并输出字符串到标准输出（通常是控制台）。格式化字符串中的 %+v 是一个动词（verb），用于指示如何格式化并打印其后的参数。
	//
	//对于 %+v，它的特殊之处在于：
	//
	//%v：这是 fmt 包中用于打印值的通用动词。对于结构体（如你的 Result 类型），它会打印出结构体的字段值，但不会打印字段名。
	//+：这是一个修饰符，当与 %v 结合使用时（即 %+v），它会告诉 fmt 包在打印结构体时，不仅打印字段值，还要打印字段名。
	//因此，当你使用 fmt.Printf("%+v\n", i2) 来打印 Result 类型的变量 i2 时，你会看到如下输出（基于你提供的 Result 结构体定义和 JSON 数据）：
	//for _, i2 := range response.Result {
	//	fmt.Printf("%+v\n", i2)
	//}
	//fmt.Print("%T", response)
	jsonStr, _ := json.Marshal(response)
	println(string(jsonStr))

	//	json的转换都是byte数组进行转换
}
