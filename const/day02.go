package main

import (
	"encoding/json"
	"fmt"
)

// 结构体定义
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

	fmt.Println(response.Result)

	for i, i2 := range response.Result {
		fmt.Println(i, i2.DT)
	}
}
