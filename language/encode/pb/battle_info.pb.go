// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: battle_info.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BattleType int32

const (
	BattleType_Match BattleType = 0
	BattleType_Rank  BattleType = 1
)

// Enum value maps for BattleType.
var (
	BattleType_name = map[int32]string{
		0: "Match",
		1: "Rank",
	}
	BattleType_value = map[string]int32{
		"Match": 0,
		"Rank":  1,
	}
)

func (x BattleType) Enum() *BattleType {
	p := new(BattleType)
	*p = x
	return p
}

func (x BattleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattleType) Descriptor() protoreflect.EnumDescriptor {
	return file_battle_info_proto_enumTypes[0].Descriptor()
}

func (BattleType) Type() protoreflect.EnumType {
	return &file_battle_info_proto_enumTypes[0]
}

func (x BattleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattleType.Descriptor instead.
func (BattleType) EnumDescriptor() ([]byte, []int) {
	return file_battle_info_proto_rawDescGZIP(), []int{0}
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID    int64  `protobuf:"varint,1,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	ZoneID    int32  `protobuf:"varint,2,opt,name=ZoneID,proto3" json:"ZoneID,omitempty"`
	KillNum   int32  `protobuf:"varint,3,opt,name=KillNum,proto3" json:"KillNum,omitempty"`
	DeadNum   int32  `protobuf:"varint,4,opt,name=DeadNum,proto3" json:"DeadNum,omitempty"`
	AssistNum int32  `protobuf:"varint,5,opt,name=AssistNum,proto3" json:"AssistNum,omitempty"`
	HeroLvl   int32  `protobuf:"varint,6,opt,name=HeroLvl,proto3" json:"HeroLvl,omitempty"`
	RobotType int32  `protobuf:"varint,7,opt,name=RobotType,proto3" json:"RobotType,omitempty"`
	Money     int32  `protobuf:"varint,8,opt,name=Money,proto3" json:"Money,omitempty"`
	HeroHurt  int32  `protobuf:"varint,9,opt,name=HeroHurt,proto3" json:"HeroHurt,omitempty"`
	IsMvp     bool   `protobuf:"varint,10,opt,name=IsMvp,proto3" json:"IsMvp,omitempty"`
	BeHurt    int32  `protobuf:"varint,11,opt,name=BeHurt,proto3" json:"BeHurt,omitempty"`
	Score     int32  `protobuf:"varint,12,opt,name=Score,proto3" json:"Score,omitempty"`
	RoadSort  int32  `protobuf:"varint,13,opt,name=RoadSort,proto3" json:"RoadSort,omitempty"`
	BlackNum  int32  `protobuf:"varint,14,opt,name=BlackNum,proto3" json:"BlackNum,omitempty"`
	RoomID    string `protobuf:"bytes,15,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
	GradeType int32  `protobuf:"varint,16,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	MultiKill int32  `protobuf:"varint,17,opt,name=MultiKill,proto3" json:"MultiKill,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_battle_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_battle_info_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInfo) GetRoleID() int64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *PlayerInfo) GetZoneID() int32 {
	if x != nil {
		return x.ZoneID
	}
	return 0
}

func (x *PlayerInfo) GetKillNum() int32 {
	if x != nil {
		return x.KillNum
	}
	return 0
}

func (x *PlayerInfo) GetDeadNum() int32 {
	if x != nil {
		return x.DeadNum
	}
	return 0
}

func (x *PlayerInfo) GetAssistNum() int32 {
	if x != nil {
		return x.AssistNum
	}
	return 0
}

func (x *PlayerInfo) GetHeroLvl() int32 {
	if x != nil {
		return x.HeroLvl
	}
	return 0
}

func (x *PlayerInfo) GetRobotType() int32 {
	if x != nil {
		return x.RobotType
	}
	return 0
}

func (x *PlayerInfo) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *PlayerInfo) GetHeroHurt() int32 {
	if x != nil {
		return x.HeroHurt
	}
	return 0
}

func (x *PlayerInfo) GetIsMvp() bool {
	if x != nil {
		return x.IsMvp
	}
	return false
}

func (x *PlayerInfo) GetBeHurt() int32 {
	if x != nil {
		return x.BeHurt
	}
	return 0
}

func (x *PlayerInfo) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *PlayerInfo) GetRoadSort() int32 {
	if x != nil {
		return x.RoadSort
	}
	return 0
}

func (x *PlayerInfo) GetBlackNum() int32 {
	if x != nil {
		return x.BlackNum
	}
	return 0
}

func (x *PlayerInfo) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *PlayerInfo) GetGradeType() int32 {
	if x != nil {
		return x.GradeType
	}
	return 0
}

func (x *PlayerInfo) GetMultiKill() int32 {
	if x != nil {
		return x.MultiKill
	}
	return 0
}

type BattleTeamInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalKill      int32         `protobuf:"varint,1,opt,name=TotalKill,proto3" json:"TotalKill,omitempty"`
	TotalMoney     int32         `protobuf:"varint,2,opt,name=TotalMoney,proto3" json:"TotalMoney,omitempty"`
	PlayerInfoList []*PlayerInfo `protobuf:"bytes,3,rep,name=PlayerInfoList,proto3" json:"PlayerInfoList,omitempty"`
}

func (x *BattleTeamInfo) Reset() {
	*x = BattleTeamInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleTeamInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleTeamInfo) ProtoMessage() {}

func (x *BattleTeamInfo) ProtoReflect() protoreflect.Message {
	mi := &file_battle_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleTeamInfo.ProtoReflect.Descriptor instead.
func (*BattleTeamInfo) Descriptor() ([]byte, []int) {
	return file_battle_info_proto_rawDescGZIP(), []int{1}
}

func (x *BattleTeamInfo) GetTotalKill() int32 {
	if x != nil {
		return x.TotalKill
	}
	return 0
}

func (x *BattleTeamInfo) GetTotalMoney() int32 {
	if x != nil {
		return x.TotalMoney
	}
	return 0
}

func (x *BattleTeamInfo) GetPlayerInfoList() []*PlayerInfo {
	if x != nil {
		return x.PlayerInfoList
	}
	return nil
}

type BattleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              int64           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BattleID        string          `protobuf:"bytes,2,opt,name=BattleID,proto3" json:"BattleID,omitempty"`
	BattleStartTime int64           `protobuf:"varint,3,opt,name=BattleStartTime,proto3" json:"BattleStartTime,omitempty"`
	BattleLastTime  int32           `protobuf:"varint,4,opt,name=BattleLastTime,proto3" json:"BattleLastTime,omitempty"`
	BattleType      BattleType      `protobuf:"varint,5,opt,name=battleType,proto3,enum=encode.BattleType" json:"battleType,omitempty"`
	IsBlack         bool            `protobuf:"varint,6,opt,name=isBlack,proto3" json:"isBlack,omitempty"`
	IsWarm          bool            `protobuf:"varint,7,opt,name=isWarm,proto3" json:"isWarm,omitempty"`
	WinCamp         int32           `protobuf:"varint,8,opt,name=winCamp,proto3" json:"winCamp,omitempty"`
	RedCampInfo     *BattleTeamInfo `protobuf:"bytes,9,opt,name=RedCampInfo,proto3" json:"RedCampInfo,omitempty"`
	BlueCampInfo    *BattleTeamInfo `protobuf:"bytes,10,opt,name=BlueCampInfo,proto3" json:"BlueCampInfo,omitempty"`
}

func (x *BattleInfo) Reset() {
	*x = BattleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_battle_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleInfo) ProtoMessage() {}

func (x *BattleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_battle_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleInfo.ProtoReflect.Descriptor instead.
func (*BattleInfo) Descriptor() ([]byte, []int) {
	return file_battle_info_proto_rawDescGZIP(), []int{2}
}

func (x *BattleInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BattleInfo) GetBattleID() string {
	if x != nil {
		return x.BattleID
	}
	return ""
}

func (x *BattleInfo) GetBattleStartTime() int64 {
	if x != nil {
		return x.BattleStartTime
	}
	return 0
}

func (x *BattleInfo) GetBattleLastTime() int32 {
	if x != nil {
		return x.BattleLastTime
	}
	return 0
}

func (x *BattleInfo) GetBattleType() BattleType {
	if x != nil {
		return x.BattleType
	}
	return BattleType_Match
}

func (x *BattleInfo) GetIsBlack() bool {
	if x != nil {
		return x.IsBlack
	}
	return false
}

func (x *BattleInfo) GetIsWarm() bool {
	if x != nil {
		return x.IsWarm
	}
	return false
}

func (x *BattleInfo) GetWinCamp() int32 {
	if x != nil {
		return x.WinCamp
	}
	return 0
}

func (x *BattleInfo) GetRedCampInfo() *BattleTeamInfo {
	if x != nil {
		return x.RedCampInfo
	}
	return nil
}

func (x *BattleInfo) GetBlueCampInfo() *BattleTeamInfo {
	if x != nil {
		return x.BlueCampInfo
	}
	return nil
}

var File_battle_info_proto protoreflect.FileDescriptor

var file_battle_info_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xc8, 0x03, 0x0a, 0x0a,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x69,
	0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4b, 0x69, 0x6c,
	0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x44, 0x65, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x48, 0x65, 0x72, 0x6f, 0x4c, 0x76, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x48,
	0x65, 0x72, 0x6f, 0x4c, 0x76, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x6f, 0x62, 0x6f, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x65,
	0x72, 0x6f, 0x48, 0x75, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x48, 0x65,
	0x72, 0x6f, 0x48, 0x75, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x4d, 0x76, 0x70, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x4d, 0x76, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x42, 0x65, 0x48, 0x75, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x42, 0x65,
	0x48, 0x75, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f,
	0x61, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x52, 0x6f,
	0x61, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x4b, 0x69, 0x6c, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x4b, 0x69, 0x6c, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x4b, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x4b, 0x69, 0x6c, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x80, 0x03, 0x0a, 0x0a, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x28,
	0x0a, 0x0f, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x32, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x57, 0x61, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x57, 0x61, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x43, 0x61, 0x6d,
	0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x43, 0x61, 0x6d, 0x70,
	0x12, 0x38, 0x0a, 0x0b, 0x52, 0x65, 0x64, 0x43, 0x61, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x52,
	0x65, 0x64, 0x43, 0x61, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0c, 0x42, 0x6c,
	0x75, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x42, 0x6c, 0x75, 0x65, 0x43, 0x61,
	0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x21, 0x0a, 0x0a, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x10, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_battle_info_proto_rawDescOnce sync.Once
	file_battle_info_proto_rawDescData = file_battle_info_proto_rawDesc
)

func file_battle_info_proto_rawDescGZIP() []byte {
	file_battle_info_proto_rawDescOnce.Do(func() {
		file_battle_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_battle_info_proto_rawDescData)
	})
	return file_battle_info_proto_rawDescData
}

var file_battle_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_battle_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_battle_info_proto_goTypes = []interface{}{
	(BattleType)(0),        // 0: encode.BattleType
	(*PlayerInfo)(nil),     // 1: encode.PlayerInfo
	(*BattleTeamInfo)(nil), // 2: encode.BattleTeamInfo
	(*BattleInfo)(nil),     // 3: encode.BattleInfo
}
var file_battle_info_proto_depIdxs = []int32{
	1, // 0: encode.BattleTeamInfo.PlayerInfoList:type_name -> encode.PlayerInfo
	0, // 1: encode.BattleInfo.battleType:type_name -> encode.BattleType
	2, // 2: encode.BattleInfo.RedCampInfo:type_name -> encode.BattleTeamInfo
	2, // 3: encode.BattleInfo.BlueCampInfo:type_name -> encode.BattleTeamInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_battle_info_proto_init() }
func file_battle_info_proto_init() {
	if File_battle_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_battle_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_battle_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleTeamInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_battle_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_battle_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_battle_info_proto_goTypes,
		DependencyIndexes: file_battle_info_proto_depIdxs,
		EnumInfos:         file_battle_info_proto_enumTypes,
		MessageInfos:      file_battle_info_proto_msgTypes,
	}.Build()
	File_battle_info_proto = out.File
	file_battle_info_proto_rawDesc = nil
	file_battle_info_proto_goTypes = nil
	file_battle_info_proto_depIdxs = nil
}
