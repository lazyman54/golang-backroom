package encode

import "time"

type BattleInfo struct {
	ID              int64
	BattleID        string
	BattleStartTime time.Time
	BattleLastTime  int32
	BattleType      int8
	IsBlack         bool
	IsWarm          bool
	WinCamp         int8
	RedCampInfo     *BattleTeamInfo
	BlueCampInfo    *BattleTeamInfo
}

type BattleTeamInfo struct {
	TotalKill      int32
	TotalMoney     int32
	PlayerInfoList []*PlayerInfo
}
type PlayerInfo struct {
	RoleID    int64
	ZoneID    int32
	KillNum   int16
	DeadNum   int16
	AssistNum int16
	HeroLvl   int8
	RobotType int8
	Money     int32
	HeroHurt  int32
	IsMvp     bool
	BeHurt    int32
	Score     int32
	RoadSort  int8
	BlackNum  int8
	RoomID    string
	GradeType int8
	MultiKill int8
}
