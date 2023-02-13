package encode

import (
	"encoding/json"
	"fmt"
	"github.com/lazymank54/golang-backroom/language/encode/pb"
	"google.golang.org/protobuf/proto"
	"time"
)

func PbEncodeSolution() {
	//battleInfo := DataForJson()
	battleInfoPb := dataForPb()
	time1 := time.Now()
	for i := 0; i < 1000000; i++ {
		//printDataForJson(battleInfo)
		printDataForPb(battleInfoPb)
	}
	dura := time.Now().Sub(time1).Milliseconds()
	fmt.Println(dura)
	//printDataForPb()
}
func printDataForJson(data *BattleInfo) {
	battleInfo := DataForJson()
	_, err := json.Marshal(battleInfo)
	if err != nil {
		fmt.Println("err ", err.Error())
	}
	//fmt.Println(len(data))
	//strData := string(data)
	//fmt.Println(strData)
}
func printDataForPb(data *pb.BattleInfo) {

	proto.Marshal(data)
	//fmt.Println(len(data))
	//strData := string(data)
	//resByte := []byte(strData)
	//
	//var target pb.BattleInfo
	//proto.Unmarshal(resByte, &target)
	//fmt.Println(target.BattleID)
	//fmt.Printf("target %+v", target)
}
func DataForJson() *BattleInfo {
	var redPlayerList []*PlayerInfo
	var bluePlayerList []*PlayerInfo
	for i := 0; i < 10; i++ {
		player := &PlayerInfo{
			RoleID:    1066877304,
			ZoneID:    13275,
			KillNum:   17,
			DeadNum:   4,
			AssistNum: 23,
			HeroLvl:   15,
			RobotType: 0,
			Money:     20220,
			HeroHurt:  167141,
			IsMvp:     false,
			BeHurt:    98288,
			Score:     1182,
			RoadSort:  5,
			BlackNum:  2,
			RoomID:    "331525241245",
			GradeType: 1,
			MultiKill: 4,
		}
		if i < 5 {
			redPlayerList = append(redPlayerList, player)
		} else {
			bluePlayerList = append(bluePlayerList, player)
		}

	}

	redBattleTeam := &BattleTeamInfo{
		TotalKill:      123,
		TotalMoney:     10220,
		PlayerInfoList: redPlayerList,
	}
	blueBattleTeam := &BattleTeamInfo{
		TotalKill:      123,
		TotalMoney:     10220,
		PlayerInfoList: redPlayerList,
	}
	battleInfo := &BattleInfo{
		ID:              123,
		BattleID:        "670008624904188330",
		BattleStartTime: time.Unix(1675408079, 0),
		BattleLastTime:  1807,
		BattleType:      1,
		IsBlack:         false,
		IsWarm:          false,
		WinCamp:         1,
		RedCampInfo:     redBattleTeam,
		BlueCampInfo:    blueBattleTeam,
	}
	return battleInfo
}

func dataForPb() *pb.BattleInfo {
	var redPlayerList []*pb.PlayerInfo
	var bluePlayerList []*pb.PlayerInfo
	for i := 0; i < 10; i++ {
		player := &pb.PlayerInfo{
			RoleID:    1066877304,
			ZoneID:    13275,
			KillNum:   17,
			DeadNum:   4,
			AssistNum: 23,
			HeroLvl:   15,
			RobotType: 0,
			Money:     20220,
			HeroHurt:  167141,
			IsMvp:     false,
			BeHurt:    98288,
			Score:     1182,
			RoadSort:  5,
			BlackNum:  2,
			RoomID:    "331525241245",
			GradeType: 1,
			MultiKill: 4,
		}
		if i < 5 {
			redPlayerList = append(redPlayerList, player)
		} else {
			bluePlayerList = append(bluePlayerList, player)
		}

	}

	redBattleTeam := &pb.BattleTeamInfo{
		TotalKill:      123,
		TotalMoney:     10220,
		PlayerInfoList: redPlayerList,
	}
	blueBattleTeam := &pb.BattleTeamInfo{
		TotalKill:      123,
		TotalMoney:     10220,
		PlayerInfoList: redPlayerList,
	}
	battleInfo := &pb.BattleInfo{
		ID:              123,
		BattleID:        "670008624904188330",
		BattleStartTime: 1675408079,
		BattleLastTime:  1807,
		BattleType:      1,
		IsBlack:         false,
		IsWarm:          false,
		WinCamp:         1,
		RedCampInfo:     redBattleTeam,
		BlueCampInfo:    blueBattleTeam,
	}
	return battleInfo
}
