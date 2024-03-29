package util

import (
	"encoding/json"
	"fmt"
)

type A1 struct {
	Item  string `json:"item"`
	Extra A2     `json:"extra"`
}
type A2 struct {
	Scene   string `json:"scene"`
	Payload A3     `json:"payload"`
}
type A3 struct {
	Anchor string `json:"anchor"`
}

func JsonSolution() {
	//var str = "{\n    \"data\": {\n        \"report_type\": 2,\n        \"data_report\": {\n            \"current_range_level\": {\n                \"name\": \"Mythic I\",\n                \"rank_level\": 7,\n                \"sub_rank_level\": 30,\n                \"rank_stars\": 44,\n                \"level_max_stars\": 10,\n                \"total_stars\": 156,\n                \"rank_id\": 180\n            },\n            \"history_range_level\": {\n                \"name\": \"Mythic II\",\n                \"rank_level\": 7,\n                \"sub_rank_level\": 29,\n                \"rank_stars\": 30,\n                \"level_max_stars\": 10,\n                \"total_stars\": 142,\n                \"rank_id\": 166\n            },\n            \"rank_points\": [\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic II\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 29,\n                        \"rank_stars\": 34,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 146,\n                        \"rank_id\": 170\n                    },\n                    \"date_ts\": \"1662336000\"\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic II\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 29,\n                        \"rank_stars\": 36,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 148,\n                        \"rank_id\": 172\n                    },\n                    \"date_ts\": \"1662422400\"\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic II\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 29,\n                        \"rank_stars\": 35,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 147,\n                        \"rank_id\": 171\n                    },\n                    \"date_ts\": \"1662508800\"\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic I\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 30,\n                        \"rank_stars\": 40,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 152,\n                        \"rank_id\": 176\n                    },\n                    \"date_ts\": \"1662595200\"\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic I\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 30,\n                        \"rank_stars\": 42,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 154,\n                        \"rank_id\": 178\n                    },\n                    \"date_ts\": \"1662681600\",\n                    \"event\": {\n                        \"event_id\": \"rank-8\",\n                        \"word\": \"\",\n                        \"description\": \"Memainkan 5 pertandingan pada Hari Ke-5 dan memperoleh 2 bintang. Teruskan!\",\n                        \"emoji_url\": \"https://p16-g-va.ibyteimg.com/tos-maliva-i-npqod914p0/3ed6e160ff7646fd857abbb9b154ae68~tplv-npqod914p0-image.image\",\n                        \"bubble_url\": \"\"\n                    }\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic I\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 30,\n                        \"rank_stars\": 43,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 155,\n                        \"rank_id\": 179\n                    },\n                    \"date_ts\": \"1662768000\"\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic I\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 30,\n                        \"rank_stars\": 44,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 156,\n                        \"rank_id\": 180\n                    },\n                    \"date_ts\": \"1662854400\"\n                }\n            ],\n            \"battle_total\": 44,\n            \"rank_battle_total\": 40,\n            \"win_rate\": \"70%\",\n            \"mvp_times\": 15,\n            \"diff_stars\": 14,\n            \"y_line\": [\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythical Glory 600\",\n                        \"rank_level\": 8,\n                        \"sub_rank_level\": 31,\n                        \"rank_stars\": 0,\n                        \"level_max_stars\": 99999999,\n                        \"total_stars\": 162,\n                        \"rank_id\": 186\n                    }\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic II\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 29,\n                        \"rank_stars\": 0,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 142,\n                        \"rank_id\": 166\n                    }\n                },\n                {\n                    \"rank_info\": {\n                        \"name\": \"Mythic III\",\n                        \"rank_level\": 7,\n                        \"sub_rank_level\": 28,\n                        \"rank_stars\": 0,\n                        \"level_max_stars\": 10,\n                        \"total_stars\": 132,\n                        \"rank_id\": 156\n                    }\n                }\n            ]\n        },\n        \"start_ts\": \"1662336000\",\n        \"end_ts\": \"1662940799\"\n    },\n    \"msg\": \"success\",\n    \"status_code\": 0\n}"

	var str = "my bad"
	//var rspBody = make(map[string]interface{})
	//err := json.Unmarshal([]byte(str), &rspBody)
	//fmt.Println(err)
	//fmt.Println(rspBody)
	payload, _ := json.Marshal(str)

	fmt.Println(string(payload))

}

func jsonUnmarshal() {
	data := new(A1)
	//jsonStr := "{\"item\":\"730\",\"extra\":{\"scene\":\"player\"}}"
	jsonStr := "{\"item\":\"730\",\"extra\":{\"scene\":\"player\",\"payload\":{\"anchor\":\"aigc\"}}}"

	err := json.Unmarshal([]byte(jsonStr), data)
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	fmt.Printf("data:%v", data)
}
