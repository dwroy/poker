// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/exsync/exsync.proto

/*
Package exsync is a generated protocol buffer package.

It is generated from these files:
	proto/exsync/exsync.proto

It has these top-level messages:
	ReqPing
	Pong
	Req
	ReqTrades
	ReqOrder
	RespIndex
	RespOrder
	RespTrades
	Timestamp
	Order
	Trade
	RespDepth
	Balance
	Position
	RespPosition
	RespBalance
*/
package exsync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TradeAction int32

const (
	TradeAction_OpenLong   TradeAction = 0
	TradeAction_OpenShort  TradeAction = 1
	TradeAction_CloseLong  TradeAction = 2
	TradeAction_CloseShort TradeAction = 3
	TradeAction_Buy        TradeAction = 4
	TradeAction_Sell       TradeAction = 5
)

var TradeAction_name = map[int32]string{
	0: "OpenLong",
	1: "OpenShort",
	2: "CloseLong",
	3: "CloseShort",
	4: "Buy",
	5: "Sell",
}
var TradeAction_value = map[string]int32{
	"OpenLong":   0,
	"OpenShort":  1,
	"CloseLong":  2,
	"CloseShort": 3,
	"Buy":        4,
	"Sell":       5,
}

func (x TradeAction) String() string {
	return proto.EnumName(TradeAction_name, int32(x))
}
func (TradeAction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PositionType int32

const (
	PositionType_Long  PositionType = 0
	PositionType_Short PositionType = 1
)

var PositionType_name = map[int32]string{
	0: "Long",
	1: "Short",
}
var PositionType_value = map[string]int32{
	"Long":  0,
	"Short": 1,
}

func (x PositionType) String() string {
	return proto.EnumName(PositionType_name, int32(x))
}
func (PositionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CurrencyUnit int32

const (
	CurrencyUnit_CNY CurrencyUnit = 0
	CurrencyUnit_USD CurrencyUnit = 1
	CurrencyUnit_BTC CurrencyUnit = 2
	CurrencyUnit_LTC CurrencyUnit = 3
)

var CurrencyUnit_name = map[int32]string{
	0: "CNY",
	1: "USD",
	2: "BTC",
	3: "LTC",
}
var CurrencyUnit_value = map[string]int32{
	"CNY": 0,
	"USD": 1,
	"BTC": 2,
	"LTC": 3,
}

func (x CurrencyUnit) String() string {
	return proto.EnumName(CurrencyUnit_name, int32(x))
}
func (CurrencyUnit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type OrderStatus int32

const (
	OrderStatus_OrderStatusCreated   OrderStatus = 0
	OrderStatus_OrderStatusPartial   OrderStatus = 1
	OrderStatus_OrderStatusComplete  OrderStatus = 2
	OrderStatus_OrderStatusCanceled  OrderStatus = 3
	OrderStatus_OrderStatusCanceling OrderStatus = 4
)

var OrderStatus_name = map[int32]string{
	0: "OrderStatusCreated",
	1: "OrderStatusPartial",
	2: "OrderStatusComplete",
	3: "OrderStatusCanceled",
	4: "OrderStatusCanceling",
}
var OrderStatus_value = map[string]int32{
	"OrderStatusCreated":   0,
	"OrderStatusPartial":   1,
	"OrderStatusComplete":  2,
	"OrderStatusCanceled":  3,
	"OrderStatusCanceling": 4,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}
func (OrderStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ReqPing struct {
}

func (m *ReqPing) Reset()                    { *m = ReqPing{} }
func (m *ReqPing) String() string            { return proto.CompactTextString(m) }
func (*ReqPing) ProtoMessage()               {}
func (*ReqPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Pong struct {
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Req struct {
	Exname string `protobuf:"bytes,1,opt,name=exname" json:"exname,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Req) GetExname() string {
	if m != nil {
		return m.Exname
	}
	return ""
}

type ReqTrades struct {
	Exname string `protobuf:"bytes,1,opt,name=exname" json:"exname,omitempty"`
	Since  string `protobuf:"bytes,2,opt,name=since" json:"since,omitempty"`
	Num    int32  `protobuf:"varint,3,opt,name=num" json:"num,omitempty"`
}

func (m *ReqTrades) Reset()                    { *m = ReqTrades{} }
func (m *ReqTrades) String() string            { return proto.CompactTextString(m) }
func (*ReqTrades) ProtoMessage()               {}
func (*ReqTrades) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReqTrades) GetExname() string {
	if m != nil {
		return m.Exname
	}
	return ""
}

func (m *ReqTrades) GetSince() string {
	if m != nil {
		return m.Since
	}
	return ""
}

func (m *ReqTrades) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ReqOrder struct {
	Exname string `protobuf:"bytes,1,opt,name=exname" json:"exname,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *ReqOrder) Reset()                    { *m = ReqOrder{} }
func (m *ReqOrder) String() string            { return proto.CompactTextString(m) }
func (*ReqOrder) ProtoMessage()               {}
func (*ReqOrder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqOrder) GetExname() string {
	if m != nil {
		return m.Exname
	}
	return ""
}

func (m *ReqOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RespIndex struct {
	Index float64 `protobuf:"fixed64,1,opt,name=index" json:"index,omitempty"`
}

func (m *RespIndex) Reset()                    { *m = RespIndex{} }
func (m *RespIndex) String() string            { return proto.CompactTextString(m) }
func (*RespIndex) ProtoMessage()               {}
func (*RespIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RespIndex) GetIndex() float64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type RespOrder struct {
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
}

func (m *RespOrder) Reset()                    { *m = RespOrder{} }
func (m *RespOrder) String() string            { return proto.CompactTextString(m) }
func (*RespOrder) ProtoMessage()               {}
func (*RespOrder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RespOrder) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type RespTrades struct {
	Trades []*Trade `protobuf:"bytes,1,rep,name=trades" json:"trades,omitempty"`
}

func (m *RespTrades) Reset()                    { *m = RespTrades{} }
func (m *RespTrades) String() string            { return proto.CompactTextString(m) }
func (*RespTrades) ProtoMessage()               {}
func (*RespTrades) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RespTrades) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

type Timestamp struct {
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanos   int64 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int64 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type Order struct {
	Id         string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Taction    TradeAction `protobuf:"varint,2,opt,name=taction,enum=exsync.TradeAction" json:"taction,omitempty"`
	Amount     float64     `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	Price      float64     `protobuf:"fixed64,4,opt,name=price" json:"price,omitempty"`
	DealAmount float64     `protobuf:"fixed64,5,opt,name=dealAmount" json:"dealAmount,omitempty"`
	DealMoney  float64     `protobuf:"fixed64,6,opt,name=dealMoney" json:"dealMoney,omitempty"`
	AvgPrice   float64     `protobuf:"fixed64,7,opt,name=avgPrice" json:"avgPrice,omitempty"`
	Fee        float64     `protobuf:"fixed64,8,opt,name=fee" json:"fee,omitempty"`
	Status     int32       `protobuf:"varint,9,opt,name=status" json:"status,omitempty"`
	CreateTime *Timestamp  `protobuf:"bytes,10,opt,name=createTime" json:"createTime,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetTaction() TradeAction {
	if m != nil {
		return m.Taction
	}
	return TradeAction_OpenLong
}

func (m *Order) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetDealAmount() float64 {
	if m != nil {
		return m.DealAmount
	}
	return 0
}

func (m *Order) GetDealMoney() float64 {
	if m != nil {
		return m.DealMoney
	}
	return 0
}

func (m *Order) GetAvgPrice() float64 {
	if m != nil {
		return m.AvgPrice
	}
	return 0
}

func (m *Order) GetFee() float64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Order) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Order) GetCreateTime() *Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type Trade struct {
	Id         string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Taction    TradeAction `protobuf:"varint,2,opt,name=taction,enum=exsync.TradeAction" json:"taction,omitempty"`
	Amount     float64     `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	Price      float64     `protobuf:"fixed64,4,opt,name=price" json:"price,omitempty"`
	Fee        float64     `protobuf:"fixed64,5,opt,name=fee" json:"fee,omitempty"`
	CreateTime *Timestamp  `protobuf:"bytes,6,opt,name=createTime" json:"createTime,omitempty"`
}

func (m *Trade) Reset()                    { *m = Trade{} }
func (m *Trade) String() string            { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()               {}
func (*Trade) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Trade) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trade) GetTaction() TradeAction {
	if m != nil {
		return m.Taction
	}
	return TradeAction_OpenLong
}

func (m *Trade) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Trade) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Trade) GetFee() float64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Trade) GetCreateTime() *Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type RespDepth struct {
	Asks []*Order `protobuf:"bytes,1,rep,name=asks" json:"asks,omitempty"`
	Bids []*Order `protobuf:"bytes,2,rep,name=bids" json:"bids,omitempty"`
}

func (m *RespDepth) Reset()                    { *m = RespDepth{} }
func (m *RespDepth) String() string            { return proto.CompactTextString(m) }
func (*RespDepth) ProtoMessage()               {}
func (*RespDepth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RespDepth) GetAsks() []*Order {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *RespDepth) GetBids() []*Order {
	if m != nil {
		return m.Bids
	}
	return nil
}

type Balance struct {
	Amount       float64      `protobuf:"fixed64,1,opt,name=amount" json:"amount,omitempty"`
	Deposit      float64      `protobuf:"fixed64,2,opt,name=deposit" json:"deposit,omitempty"`
	RealProfil   float64      `protobuf:"fixed64,3,opt,name=realProfil" json:"realProfil,omitempty"`
	UnrealProfit float64      `protobuf:"fixed64,4,opt,name=UnrealProfit,json=unrealProfit" json:"UnrealProfit,omitempty"`
	RiskRate     float64      `protobuf:"fixed64,5,opt,name=RiskRate,json=riskRate" json:"RiskRate,omitempty"`
	Currentcy    CurrencyUnit `protobuf:"varint,6,opt,name=currentcy,enum=exsync.CurrencyUnit" json:"currentcy,omitempty"`
}

func (m *Balance) Reset()                    { *m = Balance{} }
func (m *Balance) String() string            { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()               {}
func (*Balance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Balance) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Balance) GetDeposit() float64 {
	if m != nil {
		return m.Deposit
	}
	return 0
}

func (m *Balance) GetRealProfil() float64 {
	if m != nil {
		return m.RealProfil
	}
	return 0
}

func (m *Balance) GetUnrealProfit() float64 {
	if m != nil {
		return m.UnrealProfit
	}
	return 0
}

func (m *Balance) GetRiskRate() float64 {
	if m != nil {
		return m.RiskRate
	}
	return 0
}

func (m *Balance) GetCurrentcy() CurrencyUnit {
	if m != nil {
		return m.Currentcy
	}
	return CurrencyUnit_CNY
}

type Position struct {
	Id              string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Ptype           PositionType `protobuf:"varint,2,opt,name=ptype,enum=exsync.PositionType" json:"ptype,omitempty"`
	Amount          float64      `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	AvailableAmount float64      `protobuf:"fixed64,4,opt,name=availableAmount" json:"availableAmount,omitempty"`
	AvgPrice        float64      `protobuf:"fixed64,5,opt,name=avgPrice" json:"avgPrice,omitempty"`
	Money           float64      `protobuf:"fixed64,6,opt,name=money" json:"money,omitempty"`
	Deposti         float64      `protobuf:"fixed64,7,opt,name=deposti" json:"deposti,omitempty"`
	Leverge         float64      `protobuf:"fixed64,8,opt,name=leverge" json:"leverge,omitempty"`
	ForceClosePrice float64      `protobuf:"fixed64,9,opt,name=forceClosePrice" json:"forceClosePrice,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Position) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Position) GetPtype() PositionType {
	if m != nil {
		return m.Ptype
	}
	return PositionType_Long
}

func (m *Position) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Position) GetAvailableAmount() float64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func (m *Position) GetAvgPrice() float64 {
	if m != nil {
		return m.AvgPrice
	}
	return 0
}

func (m *Position) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *Position) GetDeposti() float64 {
	if m != nil {
		return m.Deposti
	}
	return 0
}

func (m *Position) GetLeverge() float64 {
	if m != nil {
		return m.Leverge
	}
	return 0
}

func (m *Position) GetForceClosePrice() float64 {
	if m != nil {
		return m.ForceClosePrice
	}
	return 0
}

type RespPosition struct {
	Long  *Position `protobuf:"bytes,1,opt,name=long" json:"long,omitempty"`
	Short *Position `protobuf:"bytes,2,opt,name=short" json:"short,omitempty"`
}

func (m *RespPosition) Reset()                    { *m = RespPosition{} }
func (m *RespPosition) String() string            { return proto.CompactTextString(m) }
func (*RespPosition) ProtoMessage()               {}
func (*RespPosition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *RespPosition) GetLong() *Position {
	if m != nil {
		return m.Long
	}
	return nil
}

func (m *RespPosition) GetShort() *Position {
	if m != nil {
		return m.Short
	}
	return nil
}

type RespBalance struct {
	Balance *Balance `protobuf:"bytes,1,opt,name=balance" json:"balance,omitempty"`
}

func (m *RespBalance) Reset()                    { *m = RespBalance{} }
func (m *RespBalance) String() string            { return proto.CompactTextString(m) }
func (*RespBalance) ProtoMessage()               {}
func (*RespBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *RespBalance) GetBalance() *Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqPing)(nil), "exsync.ReqPing")
	proto.RegisterType((*Pong)(nil), "exsync.Pong")
	proto.RegisterType((*Req)(nil), "exsync.Req")
	proto.RegisterType((*ReqTrades)(nil), "exsync.ReqTrades")
	proto.RegisterType((*ReqOrder)(nil), "exsync.ReqOrder")
	proto.RegisterType((*RespIndex)(nil), "exsync.RespIndex")
	proto.RegisterType((*RespOrder)(nil), "exsync.RespOrder")
	proto.RegisterType((*RespTrades)(nil), "exsync.RespTrades")
	proto.RegisterType((*Timestamp)(nil), "exsync.Timestamp")
	proto.RegisterType((*Order)(nil), "exsync.Order")
	proto.RegisterType((*Trade)(nil), "exsync.Trade")
	proto.RegisterType((*RespDepth)(nil), "exsync.RespDepth")
	proto.RegisterType((*Balance)(nil), "exsync.Balance")
	proto.RegisterType((*Position)(nil), "exsync.Position")
	proto.RegisterType((*RespPosition)(nil), "exsync.RespPosition")
	proto.RegisterType((*RespBalance)(nil), "exsync.RespBalance")
	proto.RegisterEnum("exsync.TradeAction", TradeAction_name, TradeAction_value)
	proto.RegisterEnum("exsync.PositionType", PositionType_name, PositionType_value)
	proto.RegisterEnum("exsync.CurrencyUnit", CurrencyUnit_name, CurrencyUnit_value)
	proto.RegisterEnum("exsync.OrderStatus", OrderStatus_name, OrderStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SyncService service

type SyncServiceClient interface {
	Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*Pong, error)
	GetOrder(ctx context.Context, in *ReqOrder, opts ...grpc.CallOption) (*RespOrder, error)
	GetTrades(ctx context.Context, in *ReqTrades, opts ...grpc.CallOption) (*RespTrades, error)
	GetDepth(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespDepth, error)
	GetIndex(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespIndex, error)
	GetPosition(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespPosition, error)
	GetBalance(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespBalance, error)
}

type syncServiceClient struct {
	cc *grpc.ClientConn
}

func NewSyncServiceClient(cc *grpc.ClientConn) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/exsync.SyncService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetOrder(ctx context.Context, in *ReqOrder, opts ...grpc.CallOption) (*RespOrder, error) {
	out := new(RespOrder)
	err := grpc.Invoke(ctx, "/exsync.SyncService/GetOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetTrades(ctx context.Context, in *ReqTrades, opts ...grpc.CallOption) (*RespTrades, error) {
	out := new(RespTrades)
	err := grpc.Invoke(ctx, "/exsync.SyncService/GetTrades", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetDepth(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespDepth, error) {
	out := new(RespDepth)
	err := grpc.Invoke(ctx, "/exsync.SyncService/GetDepth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetIndex(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespIndex, error) {
	out := new(RespIndex)
	err := grpc.Invoke(ctx, "/exsync.SyncService/GetIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetPosition(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespPosition, error) {
	out := new(RespPosition)
	err := grpc.Invoke(ctx, "/exsync.SyncService/GetPosition", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetBalance(ctx context.Context, in *Req, opts ...grpc.CallOption) (*RespBalance, error) {
	out := new(RespBalance)
	err := grpc.Invoke(ctx, "/exsync.SyncService/GetBalance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SyncService service

type SyncServiceServer interface {
	Ping(context.Context, *ReqPing) (*Pong, error)
	GetOrder(context.Context, *ReqOrder) (*RespOrder, error)
	GetTrades(context.Context, *ReqTrades) (*RespTrades, error)
	GetDepth(context.Context, *Req) (*RespDepth, error)
	GetIndex(context.Context, *Req) (*RespIndex, error)
	GetPosition(context.Context, *Req) (*RespPosition, error)
	GetBalance(context.Context, *Req) (*RespBalance, error)
}

func RegisterSyncServiceServer(s *grpc.Server, srv SyncServiceServer) {
	s.RegisterService(&_SyncService_serviceDesc, srv)
}

func _SyncService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Ping(ctx, req.(*ReqPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetOrder(ctx, req.(*ReqOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqTrades)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/GetTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetTrades(ctx, req.(*ReqTrades))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetDepth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetDepth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/GetDepth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetDepth(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/GetIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetIndex(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/GetPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetPosition(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetBalance(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exsync.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SyncService_Ping_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _SyncService_GetOrder_Handler,
		},
		{
			MethodName: "GetTrades",
			Handler:    _SyncService_GetTrades_Handler,
		},
		{
			MethodName: "GetDepth",
			Handler:    _SyncService_GetDepth_Handler,
		},
		{
			MethodName: "GetIndex",
			Handler:    _SyncService_GetIndex_Handler,
		},
		{
			MethodName: "GetPosition",
			Handler:    _SyncService_GetPosition_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _SyncService_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/exsync/exsync.proto",
}

func init() { proto.RegisterFile("proto/exsync/exsync.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 982 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x8e, 0xed, 0x38, 0x89, 0x4f, 0xd2, 0xad, 0x3b, 0xbb, 0x2a, 0x66, 0x05, 0x68, 0x3b, 0xe5,
	0x67, 0x89, 0xa0, 0xb4, 0x29, 0x17, 0x48, 0x5c, 0xb5, 0xa9, 0xb4, 0x42, 0x14, 0x1a, 0x26, 0xd9,
	0x0b, 0xa4, 0xde, 0xcc, 0xda, 0x67, 0x53, 0xab, 0xce, 0xd8, 0x6b, 0x4f, 0x56, 0x9b, 0x47, 0xe0,
	0x0d, 0x78, 0x07, 0x9e, 0x80, 0xa7, 0xe0, 0x19, 0x78, 0x13, 0x34, 0x3f, 0x76, 0x9c, 0xb4, 0x2b,
	0x71, 0xd5, 0xab, 0xcc, 0x77, 0xbe, 0xef, 0x9c, 0x99, 0xf3, 0x33, 0x9e, 0xc0, 0xc7, 0x45, 0x99,
	0xcb, 0xfc, 0x3b, 0xbc, 0xa9, 0x36, 0x22, 0xb6, 0x3f, 0x8f, 0xb4, 0x8d, 0xf4, 0x0c, 0xa2, 0x01,
	0xf4, 0x19, 0x5e, 0xcd, 0x52, 0xb1, 0xa4, 0x3d, 0xe8, 0xce, 0x72, 0xb1, 0xa4, 0x9f, 0x82, 0xc7,
	0xf0, 0x8a, 0xdc, 0x87, 0x1e, 0xde, 0x08, 0xbe, 0xc2, 0xc8, 0x39, 0x71, 0x4e, 0x03, 0x66, 0x11,
	0xfd, 0x19, 0x02, 0x86, 0x57, 0x8b, 0x92, 0x27, 0x58, 0xdd, 0x26, 0x22, 0x47, 0xe0, 0x57, 0xa9,
	0x88, 0x31, 0x72, 0xb5, 0xd9, 0x00, 0x12, 0x82, 0x27, 0xd6, 0xab, 0xc8, 0x3b, 0x71, 0x4e, 0x7d,
	0xa6, 0x96, 0x74, 0x02, 0x03, 0x86, 0x57, 0xaf, 0xca, 0x04, 0xcb, 0x5b, 0x63, 0x1d, 0x80, 0x9b,
	0x26, 0x36, 0x90, 0x9b, 0x26, 0xf4, 0x81, 0x3a, 0x40, 0x55, 0xfc, 0x24, 0x12, 0xbc, 0x51, 0x1b,
	0xa5, 0x6a, 0xa1, 0x7d, 0x1c, 0x66, 0x00, 0x7d, 0x6c, 0x24, 0x26, 0xee, 0x43, 0xf0, 0x73, 0xb5,
	0xd0, 0x92, 0xe1, 0xe4, 0xce, 0x23, 0x5b, 0x08, 0xcd, 0x32, 0xc3, 0xd1, 0xa7, 0x00, 0xca, 0xc3,
	0xa6, 0xf5, 0x05, 0xf4, 0xa4, 0x5e, 0x45, 0xce, 0x89, 0xd7, 0xf6, 0xd1, 0x3c, 0xb3, 0x24, 0xfd,
	0x11, 0x82, 0x45, 0xba, 0xc2, 0x4a, 0xf2, 0x55, 0x41, 0x22, 0xe8, 0x57, 0x18, 0xe7, 0x22, 0xa9,
	0xf4, 0x46, 0x1e, 0xab, 0xa1, 0x3a, 0xa3, 0xe0, 0x22, 0xaf, 0x74, 0x0e, 0x1e, 0x33, 0x80, 0xfe,
	0xe5, 0x82, 0x6f, 0x0e, 0x68, 0x12, 0x74, 0xea, 0x04, 0xc9, 0xb7, 0xd0, 0x97, 0x3c, 0x96, 0x69,
	0x2e, 0xb4, 0xc7, 0xc1, 0xe4, 0x70, 0x67, 0xfb, 0x67, 0x9a, 0x62, 0xb5, 0x46, 0xd5, 0x8d, 0xaf,
	0xf2, 0xb5, 0x90, 0xba, 0xb0, 0x0e, 0xb3, 0x48, 0x6d, 0x5b, 0x94, 0x69, 0x8c, 0x51, 0xd7, 0x94,
	0x46, 0x03, 0xf2, 0x19, 0x40, 0x82, 0x3c, 0x7b, 0x66, 0x3c, 0x7c, 0x4d, 0xb5, 0x2c, 0xe4, 0x13,
	0x08, 0x14, 0xfa, 0x25, 0x17, 0xb8, 0x89, 0x7a, 0x9a, 0xde, 0x1a, 0xc8, 0x31, 0x0c, 0xf8, 0xf5,
	0x72, 0xa6, 0xc3, 0xf6, 0x35, 0xd9, 0x60, 0xd5, 0xdd, 0x4b, 0xc4, 0x68, 0xa0, 0xcd, 0x6a, 0xa9,
	0x4e, 0x56, 0x49, 0x2e, 0xd7, 0x55, 0x14, 0xe8, 0x96, 0x5b, 0x44, 0x9e, 0x00, 0xc4, 0x25, 0x72,
	0x89, 0xaa, 0x7a, 0x11, 0xe8, 0xb6, 0xdc, 0x6b, 0x72, 0xac, 0x2b, 0xca, 0x5a, 0x22, 0xfa, 0xb7,
	0x03, 0xbe, 0xce, 0xfe, 0xc3, 0x56, 0xcb, 0xe6, 0xe4, 0x6f, 0x73, 0xda, 0x3d, 0x7b, 0xef, 0xff,
	0x9c, 0xfd, 0x37, 0x33, 0x8d, 0x2f, 0xb0, 0x90, 0x6f, 0xc8, 0x03, 0xe8, 0xf2, 0xea, 0xed, 0x3b,
	0x83, 0x65, 0x86, 0x51, 0x53, 0x4a, 0x72, 0x91, 0x26, 0x6a, 0x5c, 0xde, 0x27, 0x51, 0x14, 0xfd,
	0xc7, 0x81, 0xfe, 0x73, 0x9e, 0x71, 0x75, 0xab, 0xb6, 0x19, 0x39, 0x3b, 0x19, 0x45, 0xd0, 0x4f,
	0xb0, 0xc8, 0xab, 0x54, 0xea, 0xc2, 0x38, 0xac, 0x86, 0x6a, 0x06, 0x4a, 0xe4, 0xd9, 0xac, 0xcc,
	0x2f, 0xd3, 0xcc, 0xd6, 0xa1, 0x65, 0x21, 0x14, 0x46, 0xe7, 0xa2, 0xc1, 0xd2, 0x96, 0x64, 0xb4,
	0x6e, 0xd9, 0xd4, 0x24, 0xb0, 0xb4, 0x7a, 0xcb, 0xb8, 0xac, 0xcb, 0x33, 0x28, 0x2d, 0x26, 0x13,
	0x08, 0xe2, 0x75, 0x59, 0xa2, 0x90, 0xb1, 0x99, 0xa1, 0x83, 0xc9, 0x51, 0x9d, 0xc5, 0x54, 0x13,
	0xf1, 0xe6, 0x5c, 0xa4, 0x92, 0x6d, 0x65, 0xf4, 0x4f, 0x17, 0x06, 0x33, 0x75, 0x3a, 0xd5, 0xa4,
	0xfd, 0x1e, 0x8f, 0xc1, 0x2f, 0xe4, 0xa6, 0x40, 0xdb, 0xe1, 0x26, 0x58, 0xed, 0xb0, 0xd8, 0x14,
	0xc8, 0x8c, 0xe4, 0xd6, 0x06, 0x9f, 0xc2, 0x5d, 0x7e, 0xcd, 0xd3, 0x8c, 0x5f, 0x64, 0x68, 0xa7,
	0xdf, 0xe4, 0xb5, 0x6f, 0xde, 0x19, 0x72, 0x7f, 0x6f, 0xc8, 0x8f, 0xc0, 0x5f, 0xb5, 0xae, 0x86,
	0x01, 0x4d, 0xa9, 0x65, 0x6a, 0x6f, 0x45, 0x0d, 0x15, 0x93, 0xe1, 0x35, 0x96, 0xcb, 0xfa, 0x62,
	0xd4, 0x50, 0x9d, 0xe7, 0x32, 0x2f, 0x63, 0x9c, 0x66, 0x79, 0x85, 0x66, 0xb3, 0xc0, 0x9c, 0x67,
	0xcf, 0x4c, 0x5f, 0xc3, 0x48, 0xcd, 0x4f, 0x53, 0x9d, 0xcf, 0xa1, 0x9b, 0xe5, 0x62, 0x69, 0xbf,
	0x67, 0xe1, 0x7e, 0x31, 0x98, 0x66, 0xc9, 0x97, 0xe0, 0x57, 0x6f, 0xf2, 0xd2, 0x34, 0xff, 0x7d,
	0x32, 0x43, 0xd3, 0x1f, 0x60, 0xa8, 0xa2, 0xd7, 0xd3, 0xf4, 0x35, 0xf4, 0x2f, 0xcc, 0xd2, 0xc6,
	0xbf, 0x5b, 0x3b, 0x5a, 0x05, 0xab, 0xf9, 0xf1, 0x6b, 0x18, 0xb6, 0xae, 0x18, 0x19, 0xc1, 0xe0,
	0x55, 0x81, 0xe2, 0x65, 0x2e, 0x96, 0x61, 0x87, 0xdc, 0x81, 0x40, 0xa1, 0xb9, 0xda, 0x23, 0x74,
	0x14, 0xd4, 0x19, 0x69, 0xd6, 0x25, 0x07, 0x00, 0x1a, 0x1a, 0xda, 0x23, 0x7d, 0xf0, 0x9e, 0xaf,
	0x37, 0x61, 0x97, 0x0c, 0xa0, 0x3b, 0xc7, 0x2c, 0x0b, 0xfd, 0xf1, 0x43, 0x18, 0xb5, 0xdb, 0xab,
	0x18, 0x1b, 0x3a, 0x00, 0xdf, 0x86, 0x1d, 0x4f, 0x60, 0xd4, 0x1e, 0x28, 0x15, 0x67, 0xfa, 0xeb,
	0xef, 0x61, 0x47, 0x2d, 0xce, 0xe7, 0x2f, 0x42, 0x47, 0x47, 0x5e, 0x4c, 0x43, 0x57, 0x2d, 0x5e,
	0x2e, 0xa6, 0xa1, 0x37, 0xfe, 0xc3, 0x81, 0xa1, 0xbe, 0x4b, 0x73, 0xf3, 0x35, 0xba, 0x0f, 0xa4,
	0x05, 0xa7, 0xfa, 0xde, 0x26, 0x61, 0x67, 0xcf, 0x3e, 0xe3, 0xa5, 0x4c, 0x79, 0x16, 0x3a, 0xe4,
	0x23, 0x38, 0x6c, 0xeb, 0xf3, 0x55, 0x91, 0xa1, 0xc4, 0xd0, 0xdd, 0x27, 0x54, 0x8d, 0x32, 0x4c,
	0x42, 0x8f, 0x44, 0x70, 0xf4, 0x0e, 0x91, 0x8a, 0x65, 0xd8, 0x9d, 0xfc, 0xeb, 0xc2, 0x70, 0xbe,
	0x11, 0xf1, 0x1c, 0xcb, 0x6b, 0x35, 0x5e, 0x5f, 0x41, 0x57, 0xbd, 0xc5, 0xa4, 0x29, 0xba, 0x7d,
	0x9c, 0x8f, 0x47, 0xdb, 0xf6, 0x89, 0x25, 0xed, 0x90, 0x27, 0x30, 0x38, 0x43, 0x69, 0xde, 0x8f,
	0xb0, 0x25, 0xd6, 0x96, 0xe3, 0x7b, 0x5b, 0x8b, 0x7d, 0x05, 0x69, 0x87, 0x7c, 0x0f, 0xc1, 0x19,
	0x4a, 0xfb, 0xc2, 0xb5, 0x14, 0xf6, 0x2d, 0x3f, 0x26, 0x6d, 0x27, 0x63, 0xa3, 0x1d, 0xf2, 0x8d,
	0xde, 0xc8, 0x7c, 0xbb, 0x86, 0x2d, 0xa7, 0xdd, 0x3d, 0x34, 0xdf, 0xa8, 0xcd, 0xd3, 0x7c, 0xbb,
	0x5a, 0xf3, 0xb4, 0x43, 0x26, 0x30, 0x3c, 0x43, 0xd9, 0xcc, 0xf5, 0x8e, 0xc3, 0x51, 0xdb, 0xa1,
	0x96, 0xd0, 0x0e, 0x79, 0x0c, 0x70, 0x86, 0xb2, 0x9e, 0xd6, 0x1d, 0x97, 0xc3, 0xb6, 0x8b, 0x55,
	0xd0, 0xce, 0x45, 0x4f, 0xff, 0xe3, 0x79, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xec,
	0x81, 0x4d, 0x0e, 0x09, 0x00, 0x00,
}