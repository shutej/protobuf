// Code generated by protoc-gen-go. DO NOT EDIT.
// source: my_test/test.proto

/*
Package my_test is a generated protocol buffer package.

This package holds interesting messages.

It is generated from these files:
	my_test/test.proto

It has these top-level messages:
	Request
	Reply
	OtherBase
	ReplyExtensions
	OtherReplyExtensions
	OldReply
	Communique
*/
package my_test

import proto "github.com/shutej/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/shutej/protobuf/protoc-gen-go/testdata/multi"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HatType int32

const (
	// deliberately skipping 0
	HatType_FEDORA HatType = 1
	HatType_FEZ    HatType = 2
)

var HatType_name = map[int32]string{
	1: "FEDORA",
	2: "FEZ",
}
var HatType_value = map[string]int32{
	"FEDORA": 1,
	"FEZ":    2,
}

func (x HatType) Enum() *HatType {
	p := new(HatType)
	*p = x
	return p
}
func (x HatType) String() string {
	return proto.EnumName(HatType_name, int32(x))
}
func (x *HatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HatType_value, data, "HatType")
	if err != nil {
		return err
	}
	*x = HatType(value)
	return nil
}

// This enum represents days of the week.
type Days int32

const (
	Days_MONDAY  Days = 1
	Days_TUESDAY Days = 2
	Days_LUNDI   Days = 1
)

var Days_name = map[int32]string{
	1: "MONDAY",
	2: "TUESDAY",
	// Duplicate value: 1: "LUNDI",
}
var Days_value = map[string]int32{
	"MONDAY":  1,
	"TUESDAY": 2,
	"LUNDI":   1,
}

func (x Days) Enum() *Days {
	p := new(Days)
	*p = x
	return p
}
func (x Days) String() string {
	return proto.EnumName(Days_name, int32(x))
}
func (x *Days) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Days_value, data, "Days")
	if err != nil {
		return err
	}
	*x = Days(value)
	return nil
}

type Request_Color int32

const (
	Request_RED   Request_Color = 0
	Request_GREEN Request_Color = 1
	Request_BLUE  Request_Color = 2
)

var Request_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Request_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Request_Color) Enum() *Request_Color {
	p := new(Request_Color)
	*p = x
	return p
}
func (x Request_Color) String() string {
	return proto.EnumName(Request_Color_name, int32(x))
}
func (x *Request_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Color_value, data, "Request_Color")
	if err != nil {
		return err
	}
	*x = Request_Color(value)
	return nil
}

type Reply_Entry_Game int32

const (
	Reply_Entry_FOOTBALL Reply_Entry_Game = 1
	Reply_Entry_TENNIS   Reply_Entry_Game = 2
)

var Reply_Entry_Game_name = map[int32]string{
	1: "FOOTBALL",
	2: "TENNIS",
}
var Reply_Entry_Game_value = map[string]int32{
	"FOOTBALL": 1,
	"TENNIS":   2,
}

func (x Reply_Entry_Game) Enum() *Reply_Entry_Game {
	p := new(Reply_Entry_Game)
	*p = x
	return p
}
func (x Reply_Entry_Game) String() string {
	return proto.EnumName(Reply_Entry_Game_name, int32(x))
}
func (x *Reply_Entry_Game) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Entry_Game_value, data, "Reply_Entry_Game")
	if err != nil {
		return err
	}
	*x = Reply_Entry_Game(value)
	return nil
}

// This is a message that might be sent somewhere.
type Request struct {
	Key []int64 `protobuf:"varint,1,rep,name=key" json:"key,omitempty"`
	//  optional imp.ImportedMessage imported_message = 2;
	Hue *Request_Color `protobuf:"varint,3,opt,name=hue,enum=my.test.Request_Color" json:"hue,omitempty"`
	Hat *HatType       `protobuf:"varint,4,opt,name=hat,enum=my.test.HatType,def=1" json:"hat,omitempty"`
	//  optional imp.ImportedMessage.Owner owner = 6;
	Deadline  *float32           `protobuf:"fixed32,7,opt,name=deadline,def=inf" json:"deadline,omitempty"`
	Somegroup *Request_SomeGroup `protobuf:"group,8,opt,name=SomeGroup,json=somegroup" json:"somegroup,omitempty"`
	// This is a map field. It will generate map[int32]string.
	NameMapping map[int32]string `protobuf:"bytes,14,rep,name=name_mapping,json=nameMapping" json:"name_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// This is a map field whose value type is a message.
	MsgMapping map[int64]*Reply `protobuf:"bytes,15,rep,name=msg_mapping,json=msgMapping" json:"msg_mapping,omitempty" protobuf_key:"zigzag64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Reset_     *int32           `protobuf:"varint,12,opt,name=reset" json:"reset,omitempty"`
	// This field should not conflict with any getters.
	GetKey_          *string `protobuf:"bytes,16,opt,name=get_key,json=getKey" json:"get_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

const Default_Request_Hat HatType = HatType_FEDORA

var Default_Request_Deadline float32 = float32(math.Inf(1))

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetHue() Request_Color {
	if m != nil && m.Hue != nil {
		return *m.Hue
	}
	return Request_RED
}

func (m *Request) GetHat() HatType {
	if m != nil && m.Hat != nil {
		return *m.Hat
	}
	return Default_Request_Hat
}

func (m *Request) GetDeadline() float32 {
	if m != nil && m.Deadline != nil {
		return *m.Deadline
	}
	return Default_Request_Deadline
}

func (m *Request) GetSomegroup() *Request_SomeGroup {
	if m != nil {
		return m.Somegroup
	}
	return nil
}

func (m *Request) GetNameMapping() map[int32]string {
	if m != nil {
		return m.NameMapping
	}
	return nil
}

func (m *Request) GetMsgMapping() map[int64]*Reply {
	if m != nil {
		return m.MsgMapping
	}
	return nil
}

func (m *Request) GetReset_() int32 {
	if m != nil && m.Reset_ != nil {
		return *m.Reset_
	}
	return 0
}

func (m *Request) GetGetKey_() string {
	if m != nil && m.GetKey_ != nil {
		return *m.GetKey_
	}
	return ""
}

type Request_SomeGroup struct {
	GroupField       *int32 `protobuf:"varint,9,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request_SomeGroup) Reset()         { *m = Request_SomeGroup{} }
func (m *Request_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*Request_SomeGroup) ProtoMessage()    {}

func (m *Request_SomeGroup) GetGroupField() int32 {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return 0
}

type Reply struct {
	Found                        []*Reply_Entry `protobuf:"bytes,1,rep,name=found" json:"found,omitempty"`
	CompactKeys                  []int32        `protobuf:"varint,2,rep,packed,name=compact_keys,json=compactKeys" json:"compact_keys,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}

var extRange_Reply = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Reply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Reply
}

func (m *Reply) GetFound() []*Reply_Entry {
	if m != nil {
		return m.Found
	}
	return nil
}

func (m *Reply) GetCompactKeys() []int32 {
	if m != nil {
		return m.CompactKeys
	}
	return nil
}

type Reply_Entry struct {
	KeyThatNeeds_1234Camel_CasIng *int64 `protobuf:"varint,1,req,name=key_that_needs_1234camel_CasIng,json=keyThatNeeds1234camelCasIng" json:"key_that_needs_1234camel_CasIng,omitempty"`
	Value                         *int64 `protobuf:"varint,2,opt,name=value,def=7" json:"value,omitempty"`
	XMyFieldName_2                *int64 `protobuf:"varint,3,opt,name=_my_field_name_2,json=MyFieldName2" json:"_my_field_name_2,omitempty"`
	XXX_unrecognized              []byte `json:"-"`
}

func (m *Reply_Entry) Reset()         { *m = Reply_Entry{} }
func (m *Reply_Entry) String() string { return proto.CompactTextString(m) }
func (*Reply_Entry) ProtoMessage()    {}

const Default_Reply_Entry_Value int64 = 7

func (m *Reply_Entry) GetKeyThatNeeds_1234Camel_CasIng() int64 {
	if m != nil && m.KeyThatNeeds_1234Camel_CasIng != nil {
		return *m.KeyThatNeeds_1234Camel_CasIng
	}
	return 0
}

func (m *Reply_Entry) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_Reply_Entry_Value
}

func (m *Reply_Entry) GetXMyFieldName_2() int64 {
	if m != nil && m.XMyFieldName_2 != nil {
		return *m.XMyFieldName_2
	}
	return 0
}

type OtherBase struct {
	Name                         *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *OtherBase) Reset()         { *m = OtherBase{} }
func (m *OtherBase) String() string { return proto.CompactTextString(m) }
func (*OtherBase) ProtoMessage()    {}

var extRange_OtherBase = []proto.ExtensionRange{
	{100, 536870911},
}

func (*OtherBase) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OtherBase
}

func (m *OtherBase) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ReplyExtensions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReplyExtensions) Reset()         { *m = ReplyExtensions{} }
func (m *ReplyExtensions) String() string { return proto.CompactTextString(m) }
func (*ReplyExtensions) ProtoMessage()    {}

var E_ReplyExtensions_Time = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*float64)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.time",
	Tag:           "fixed64,101,opt,name=time",
	Filename:      "my_test/test.proto",
}

var E_ReplyExtensions_Carrot = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         105,
	Name:          "my.test.ReplyExtensions.carrot",
	Tag:           "bytes,105,opt,name=carrot",
	Filename:      "my_test/test.proto",
}

var E_ReplyExtensions_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*OtherBase)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.donut",
	Tag:           "bytes,101,opt,name=donut",
	Filename:      "my_test/test.proto",
}

type OtherReplyExtensions struct {
	Key              *int32 `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OtherReplyExtensions) Reset()         { *m = OtherReplyExtensions{} }
func (m *OtherReplyExtensions) String() string { return proto.CompactTextString(m) }
func (*OtherReplyExtensions) ProtoMessage()    {}

func (m *OtherReplyExtensions) GetKey() int32 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

type OldReply struct {
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *OldReply) Reset()         { *m = OldReply{} }
func (m *OldReply) String() string { return proto.CompactTextString(m) }
func (*OldReply) ProtoMessage()    {}

func (m *OldReply) Marshal() ([]byte, error) {
	return proto.MarshalMessageSet(&m.XXX_InternalExtensions)
}
func (m *OldReply) Unmarshal(buf []byte) error {
	return proto.UnmarshalMessageSet(buf, &m.XXX_InternalExtensions)
}
func (m *OldReply) MarshalJSON() ([]byte, error) {
	return proto.MarshalMessageSetJSON(&m.XXX_InternalExtensions)
}
func (m *OldReply) UnmarshalJSON(buf []byte) error {
	return proto.UnmarshalMessageSetJSON(buf, &m.XXX_InternalExtensions)
}

// ensure OldReply satisfies proto.Marshaler and proto.Unmarshaler
var _ proto.Marshaler = (*OldReply)(nil)
var _ proto.Unmarshaler = (*OldReply)(nil)

var extRange_OldReply = []proto.ExtensionRange{
	{100, 2147483646},
}

func (*OldReply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OldReply
}

type Communique struct {
	MakeMeCry *bool `protobuf:"varint,1,opt,name=make_me_cry,json=makeMeCry" json:"make_me_cry,omitempty"`
	// This is a oneof, called "union".
	//
	// Types that are valid to be assigned to Union:
	//	*Communique_Number
	//	*Communique_Name
	//	*Communique_Data
	//	*Communique_TempC
	//	*Communique_Height
	//	*Communique_Today
	//	*Communique_Maybe
	//	*Communique_Delta_
	//	*Communique_Msg
	//	*Communique_Somegroup
	Union            isCommunique_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Communique) Reset()         { *m = Communique{} }
func (m *Communique) String() string { return proto.CompactTextString(m) }
func (*Communique) ProtoMessage()    {}

type isCommunique_Union interface {
	isCommunique_Union()
}

type Communique_Number struct {
	Number int32 `protobuf:"varint,5,opt,name=number,oneof"`
}
type Communique_Name struct {
	Name string `protobuf:"bytes,6,opt,name=name,oneof"`
}
type Communique_Data struct {
	Data []byte `protobuf:"bytes,7,opt,name=data,oneof"`
}
type Communique_TempC struct {
	TempC float64 `protobuf:"fixed64,8,opt,name=temp_c,json=tempC,oneof"`
}
type Communique_Height struct {
	Height float32 `protobuf:"fixed32,9,opt,name=height,oneof"`
}
type Communique_Today struct {
	Today Days `protobuf:"varint,10,opt,name=today,enum=my.test.Days,oneof"`
}
type Communique_Maybe struct {
	Maybe bool `protobuf:"varint,11,opt,name=maybe,oneof"`
}
type Communique_Delta_ struct {
	Delta int32 `protobuf:"zigzag32,12,opt,name=delta,oneof"`
}
type Communique_Msg struct {
	Msg *Reply `protobuf:"bytes,13,opt,name=msg,oneof"`
}
type Communique_Somegroup struct {
	Somegroup *Communique_SomeGroup `protobuf:"group,14,opt,name=SomeGroup,json=somegroup,oneof"`
}

func (*Communique_Number) isCommunique_Union()    {}
func (*Communique_Name) isCommunique_Union()      {}
func (*Communique_Data) isCommunique_Union()      {}
func (*Communique_TempC) isCommunique_Union()     {}
func (*Communique_Height) isCommunique_Union()    {}
func (*Communique_Today) isCommunique_Union()     {}
func (*Communique_Maybe) isCommunique_Union()     {}
func (*Communique_Delta_) isCommunique_Union()    {}
func (*Communique_Msg) isCommunique_Union()       {}
func (*Communique_Somegroup) isCommunique_Union() {}

func (m *Communique) GetUnion() isCommunique_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Communique) GetMakeMeCry() bool {
	if m != nil && m.MakeMeCry != nil {
		return *m.MakeMeCry
	}
	return false
}

func (m *Communique) GetNumber() int32 {
	if x, ok := m.GetUnion().(*Communique_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Communique) GetName() string {
	if x, ok := m.GetUnion().(*Communique_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Communique) GetData() []byte {
	if x, ok := m.GetUnion().(*Communique_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Communique) GetTempC() float64 {
	if x, ok := m.GetUnion().(*Communique_TempC); ok {
		return x.TempC
	}
	return 0
}

func (m *Communique) GetHeight() float32 {
	if x, ok := m.GetUnion().(*Communique_Height); ok {
		return x.Height
	}
	return 0
}

func (m *Communique) GetToday() Days {
	if x, ok := m.GetUnion().(*Communique_Today); ok {
		return x.Today
	}
	return Days_MONDAY
}

func (m *Communique) GetMaybe() bool {
	if x, ok := m.GetUnion().(*Communique_Maybe); ok {
		return x.Maybe
	}
	return false
}

func (m *Communique) GetDelta() int32 {
	if x, ok := m.GetUnion().(*Communique_Delta_); ok {
		return x.Delta
	}
	return 0
}

func (m *Communique) GetMsg() *Reply {
	if x, ok := m.GetUnion().(*Communique_Msg); ok {
		return x.Msg
	}
	return nil
}

func (m *Communique) GetSomegroup() *Communique_SomeGroup {
	if x, ok := m.GetUnion().(*Communique_Somegroup); ok {
		return x.Somegroup
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Communique) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Communique_OneofMarshaler, _Communique_OneofUnmarshaler, _Communique_OneofSizer, []interface{}{
		(*Communique_Number)(nil),
		(*Communique_Name)(nil),
		(*Communique_Data)(nil),
		(*Communique_TempC)(nil),
		(*Communique_Height)(nil),
		(*Communique_Today)(nil),
		(*Communique_Maybe)(nil),
		(*Communique_Delta_)(nil),
		(*Communique_Msg)(nil),
		(*Communique_Somegroup)(nil),
	}
}

func _Communique_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Communique)
	// union
	switch x := m.Union.(type) {
	case *Communique_Number:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Number))
	case *Communique_Name:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *Communique_Data:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Data)
	case *Communique_TempC:
		b.EncodeVarint(8<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.TempC))
	case *Communique_Height:
		b.EncodeVarint(9<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.Height)))
	case *Communique_Today:
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Today))
	case *Communique_Maybe:
		t := uint64(0)
		if x.Maybe {
			t = 1
		}
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Communique_Delta_:
		b.EncodeVarint(12<<3 | proto.WireVarint)
		b.EncodeZigzag32(uint64(x.Delta))
	case *Communique_Msg:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Msg); err != nil {
			return err
		}
	case *Communique_Somegroup:
		b.EncodeVarint(14<<3 | proto.WireStartGroup)
		if err := b.Marshal(x.Somegroup); err != nil {
			return err
		}
		b.EncodeVarint(14<<3 | proto.WireEndGroup)
	case nil:
	default:
		return fmt.Errorf("Communique.Union has unexpected type %T", x)
	}
	return nil
}

func _Communique_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Communique)
	switch tag {
	case 5: // union.number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Number{int32(x)}
		return true, err
	case 6: // union.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &Communique_Name{x}
		return true, err
	case 7: // union.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Union = &Communique_Data{x}
		return true, err
	case 8: // union.temp_c
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Union = &Communique_TempC{math.Float64frombits(x)}
		return true, err
	case 9: // union.height
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Union = &Communique_Height{math.Float32frombits(uint32(x))}
		return true, err
	case 10: // union.today
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Today{Days(x)}
		return true, err
	case 11: // union.maybe
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Communique_Maybe{x != 0}
		return true, err
	case 12: // union.delta
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag32()
		m.Union = &Communique_Delta_{int32(x)}
		return true, err
	case 13: // union.msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Reply)
		err := b.DecodeMessage(msg)
		m.Union = &Communique_Msg{msg}
		return true, err
	case 14: // union.somegroup
		if wire != proto.WireStartGroup {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Communique_SomeGroup)
		err := b.DecodeGroup(msg)
		m.Union = &Communique_Somegroup{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Communique_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Communique)
	// union
	switch x := m.Union.(type) {
	case *Communique_Number:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Number))
	case *Communique_Name:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *Communique_Data:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Data)))
		n += len(x.Data)
	case *Communique_TempC:
		n += proto.SizeVarint(8<<3 | proto.WireFixed64)
		n += 8
	case *Communique_Height:
		n += proto.SizeVarint(9<<3 | proto.WireFixed32)
		n += 4
	case *Communique_Today:
		n += proto.SizeVarint(10<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Today))
	case *Communique_Maybe:
		n += proto.SizeVarint(11<<3 | proto.WireVarint)
		n += 1
	case *Communique_Delta_:
		n += proto.SizeVarint(12<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64((uint32(x.Delta) << 1) ^ uint32((int32(x.Delta) >> 31))))
	case *Communique_Msg:
		s := proto.Size(x.Msg)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Communique_Somegroup:
		n += proto.SizeVarint(14<<3 | proto.WireStartGroup)
		n += proto.Size(x.Somegroup)
		n += proto.SizeVarint(14<<3 | proto.WireEndGroup)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Communique_SomeGroup struct {
	Member           *string `protobuf:"bytes,15,opt,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Communique_SomeGroup) Reset()         { *m = Communique_SomeGroup{} }
func (m *Communique_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*Communique_SomeGroup) ProtoMessage()    {}

func (m *Communique_SomeGroup) GetMember() string {
	if m != nil && m.Member != nil {
		return *m.Member
	}
	return ""
}

type Communique_Delta struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Communique_Delta) Reset()         { *m = Communique_Delta{} }
func (m *Communique_Delta) String() string { return proto.CompactTextString(m) }
func (*Communique_Delta) ProtoMessage()    {}

var E_Tag = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*string)(nil),
	Field:         103,
	Name:          "my.test.tag",
	Tag:           "bytes,103,opt,name=tag",
	Filename:      "my_test/test.proto",
}

var E_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*OtherReplyExtensions)(nil),
	Field:         106,
	Name:          "my.test.donut",
	Tag:           "bytes,106,opt,name=donut",
	Filename:      "my_test/test.proto",
}

func init() {
	proto.RegisterType((*Request)(nil), "my.test.Request")
	proto.RegisterType((*Request_SomeGroup)(nil), "my.test.Request.SomeGroup")
	proto.RegisterType((*Reply)(nil), "my.test.Reply")
	proto.RegisterType((*Reply_Entry)(nil), "my.test.Reply.Entry")
	proto.RegisterType((*OtherBase)(nil), "my.test.OtherBase")
	proto.RegisterType((*ReplyExtensions)(nil), "my.test.ReplyExtensions")
	proto.RegisterType((*OtherReplyExtensions)(nil), "my.test.OtherReplyExtensions")
	proto.RegisterType((*OldReply)(nil), "my.test.OldReply")
	proto.RegisterType((*Communique)(nil), "my.test.Communique")
	proto.RegisterType((*Communique_SomeGroup)(nil), "my.test.Communique.SomeGroup")
	proto.RegisterType((*Communique_Delta)(nil), "my.test.Communique.Delta")
	proto.RegisterEnum("my.test.HatType", HatType_name, HatType_value)
	proto.RegisterEnum("my.test.Days", Days_name, Days_value)
	proto.RegisterEnum("my.test.Request_Color", Request_Color_name, Request_Color_value)
	proto.RegisterEnum("my.test.Reply_Entry_Game", Reply_Entry_Game_name, Reply_Entry_Game_value)
	proto.RegisterExtension(E_ReplyExtensions_Time)
	proto.RegisterExtension(E_ReplyExtensions_Carrot)
	proto.RegisterExtension(E_ReplyExtensions_Donut)
	proto.RegisterExtension(E_Tag)
	proto.RegisterExtension(E_Donut)
}
