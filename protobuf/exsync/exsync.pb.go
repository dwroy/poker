// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/exsync/exsync.proto

/*
Package exsync is a generated protocol buffer package.

It is generated from these files:
	protobuf/exsync/exsync.proto

It has these top-level messages:
	ReqPing
	Pong
	Req
	Resp
	ReqMakeOrder
	RespMakeOrder
	ReqCancelOrder
	ReqTrades
	ReqOrders
	RespIndex
	RespOrders
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
	TradeAction_ActionUnkown TradeAction = 0
	TradeAction_OpenLong     TradeAction = 1
	TradeAction_OpenShort    TradeAction = 2
	TradeAction_CloseLong    TradeAction = 3
	TradeAction_CloseShort   TradeAction = 4
	TradeAction_Buy          TradeAction = 5
	TradeAction_Sell         TradeAction = 6
)

var TradeAction_name = map[int32]string{
	0: "ActionUnkown",
	1: "OpenLong",
	2: "OpenShort",
	3: "CloseLong",
	4: "CloseShort",
	5: "Buy",
	6: "Sell",
}
var TradeAction_value = map[string]int32{
	"ActionUnkown": 0,
	"OpenLong":     1,
	"OpenShort":    2,
	"CloseLong":    3,
	"CloseShort":   4,
	"Buy":          5,
	"Sell":         6,
}

func (x TradeAction) String() string {
	return proto.EnumName(TradeAction_name, int32(x))
}
func (TradeAction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PositionType int32

const (
	PositionType_PositionNone PositionType = 0
	PositionType_Long         PositionType = 1
	PositionType_Short        PositionType = 2
)

var PositionType_name = map[int32]string{
	0: "PositionNone",
	1: "Long",
	2: "Short",
}
var PositionType_value = map[string]int32{
	"PositionNone": 0,
	"Long":         1,
	"Short":        2,
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
	OrderStatus_StatusUnkown OrderStatus = 0
	OrderStatus_Created      OrderStatus = 1
	OrderStatus_Partial      OrderStatus = 2
	OrderStatus_Complete     OrderStatus = 3
	OrderStatus_Canceled     OrderStatus = 4
	OrderStatus_Canceling    OrderStatus = 5
)

var OrderStatus_name = map[int32]string{
	0: "StatusUnkown",
	1: "Created",
	2: "Partial",
	3: "Complete",
	4: "Canceled",
	5: "Canceling",
}
var OrderStatus_value = map[string]int32{
	"StatusUnkown": 0,
	"Created":      1,
	"Partial":      2,
	"Complete":     3,
	"Canceled":     4,
	"Canceling":    5,
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

type Resp struct {
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ReqMakeOrder struct {
	Exname  string      `protobuf:"bytes,1,opt,name=exname" json:"exname,omitempty"`
	TAction TradeAction `protobuf:"varint,2,opt,name=tAction,enum=exsync.TradeAction" json:"tAction,omitempty"`
	Amount  float64     `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	Price   float64     `protobuf:"fixed64,4,opt,name=price" json:"price,omitempty"`
}

func (m *ReqMakeOrder) Reset()                    { *m = ReqMakeOrder{} }
func (m *ReqMakeOrder) String() string            { return proto.CompactTextString(m) }
func (*ReqMakeOrder) ProtoMessage()               {}
func (*ReqMakeOrder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqMakeOrder) GetExname() string {
	if m != nil {
		return m.Exname
	}
	return ""
}

func (m *ReqMakeOrder) GetTAction() TradeAction {
	if m != nil {
		return m.TAction
	}
	return TradeAction_ActionUnkown
}

func (m *ReqMakeOrder) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqMakeOrder) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type RespMakeOrder struct {
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
}

func (m *RespMakeOrder) Reset()                    { *m = RespMakeOrder{} }
func (m *RespMakeOrder) String() string            { return proto.CompactTextString(m) }
func (*RespMakeOrder) ProtoMessage()               {}
func (*RespMakeOrder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RespMakeOrder) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type ReqCancelOrder struct {
	Exname string   `protobuf:"bytes,1,opt,name=exname" json:"exname,omitempty"`
	Ids    []string `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
}

func (m *ReqCancelOrder) Reset()                    { *m = ReqCancelOrder{} }
func (m *ReqCancelOrder) String() string            { return proto.CompactTextString(m) }
func (*ReqCancelOrder) ProtoMessage()               {}
func (*ReqCancelOrder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqCancelOrder) GetExname() string {
	if m != nil {
		return m.Exname
	}
	return ""
}

func (m *ReqCancelOrder) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ReqTrades struct {
	Exname string `protobuf:"bytes,1,opt,name=exname" json:"exname,omitempty"`
	Since  string `protobuf:"bytes,2,opt,name=since" json:"since,omitempty"`
	Num    int32  `protobuf:"varint,3,opt,name=num" json:"num,omitempty"`
}

func (m *ReqTrades) Reset()                    { *m = ReqTrades{} }
func (m *ReqTrades) String() string            { return proto.CompactTextString(m) }
func (*ReqTrades) ProtoMessage()               {}
func (*ReqTrades) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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

type ReqOrders struct {
	Exname string   `protobuf:"bytes,1,opt,name=exname" json:"exname,omitempty"`
	Ids    []string `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
}

func (m *ReqOrders) Reset()                    { *m = ReqOrders{} }
func (m *ReqOrders) String() string            { return proto.CompactTextString(m) }
func (*ReqOrders) ProtoMessage()               {}
func (*ReqOrders) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReqOrders) GetExname() string {
	if m != nil {
		return m.Exname
	}
	return ""
}

func (m *ReqOrders) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type RespIndex struct {
	Index float64 `protobuf:"fixed64,1,opt,name=index" json:"index,omitempty"`
}

func (m *RespIndex) Reset()                    { *m = RespIndex{} }
func (m *RespIndex) String() string            { return proto.CompactTextString(m) }
func (*RespIndex) ProtoMessage()               {}
func (*RespIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RespIndex) GetIndex() float64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type RespOrders struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *RespOrders) Reset()                    { *m = RespOrders{} }
func (m *RespOrders) String() string            { return proto.CompactTextString(m) }
func (*RespOrders) ProtoMessage()               {}
func (*RespOrders) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RespOrders) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type RespTrades struct {
	Trades []*Trade `protobuf:"bytes,1,rep,name=trades" json:"trades,omitempty"`
}

func (m *RespTrades) Reset()                    { *m = RespTrades{} }
func (m *RespTrades) String() string            { return proto.CompactTextString(m) }
func (*RespTrades) ProtoMessage()               {}
func (*RespTrades) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

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
	TAction    TradeAction `protobuf:"varint,2,opt,name=tAction,enum=exsync.TradeAction" json:"tAction,omitempty"`
	Amount     float64     `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	Price      float64     `protobuf:"fixed64,4,opt,name=price" json:"price,omitempty"`
	DealAmount float64     `protobuf:"fixed64,5,opt,name=dealAmount" json:"dealAmount,omitempty"`
	DealMoney  float64     `protobuf:"fixed64,6,opt,name=dealMoney" json:"dealMoney,omitempty"`
	AvgPrice   float64     `protobuf:"fixed64,7,opt,name=avgPrice" json:"avgPrice,omitempty"`
	Fee        float64     `protobuf:"fixed64,8,opt,name=fee" json:"fee,omitempty"`
	Status     OrderStatus `protobuf:"varint,9,opt,name=status,enum=exsync.OrderStatus" json:"status,omitempty"`
	CreateTime *Timestamp  `protobuf:"bytes,10,opt,name=createTime" json:"createTime,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetTAction() TradeAction {
	if m != nil {
		return m.TAction
	}
	return TradeAction_ActionUnkown
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

func (m *Order) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_StatusUnkown
}

func (m *Order) GetCreateTime() *Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type Trade struct {
	Id         string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TAction    TradeAction `protobuf:"varint,2,opt,name=tAction,enum=exsync.TradeAction" json:"tAction,omitempty"`
	Amount     float64     `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	Price      float64     `protobuf:"fixed64,4,opt,name=price" json:"price,omitempty"`
	Fee        float64     `protobuf:"fixed64,5,opt,name=fee" json:"fee,omitempty"`
	CreateTime *Timestamp  `protobuf:"bytes,6,opt,name=createTime" json:"createTime,omitempty"`
}

func (m *Trade) Reset()                    { *m = Trade{} }
func (m *Trade) String() string            { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()               {}
func (*Trade) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Trade) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trade) GetTAction() TradeAction {
	if m != nil {
		return m.TAction
	}
	return TradeAction_ActionUnkown
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
	Asks []*Trade `protobuf:"bytes,1,rep,name=asks" json:"asks,omitempty"`
	Bids []*Trade `protobuf:"bytes,2,rep,name=bids" json:"bids,omitempty"`
}

func (m *RespDepth) Reset()                    { *m = RespDepth{} }
func (m *RespDepth) String() string            { return proto.CompactTextString(m) }
func (*RespDepth) ProtoMessage()               {}
func (*RespDepth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *RespDepth) GetAsks() []*Trade {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *RespDepth) GetBids() []*Trade {
	if m != nil {
		return m.Bids
	}
	return nil
}

type Balance struct {
	Amount       float64      `protobuf:"fixed64,1,opt,name=amount" json:"amount,omitempty"`
	Deposit      float64      `protobuf:"fixed64,2,opt,name=deposit" json:"deposit,omitempty"`
	RealProfit   float64      `protobuf:"fixed64,3,opt,name=realProfit" json:"realProfit,omitempty"`
	UnrealProfit float64      `protobuf:"fixed64,4,opt,name=unrealProfit" json:"unrealProfit,omitempty"`
	RiskRate     float64      `protobuf:"fixed64,5,opt,name=riskRate" json:"riskRate,omitempty"`
	Currency     CurrencyUnit `protobuf:"varint,6,opt,name=currency,enum=exsync.CurrencyUnit" json:"currency,omitempty"`
}

func (m *Balance) Reset()                    { *m = Balance{} }
func (m *Balance) String() string            { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()               {}
func (*Balance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

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

func (m *Balance) GetRealProfit() float64 {
	if m != nil {
		return m.RealProfit
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

func (m *Balance) GetCurrency() CurrencyUnit {
	if m != nil {
		return m.Currency
	}
	return CurrencyUnit_CNY
}

type Position struct {
	Id              string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PType           PositionType `protobuf:"varint,2,opt,name=pType,enum=exsync.PositionType" json:"pType,omitempty"`
	Amount          float64      `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	AvailableAmount float64      `protobuf:"fixed64,4,opt,name=availableAmount" json:"availableAmount,omitempty"`
	AvgPrice        float64      `protobuf:"fixed64,5,opt,name=avgPrice" json:"avgPrice,omitempty"`
	Money           float64      `protobuf:"fixed64,6,opt,name=money" json:"money,omitempty"`
	Deposit         float64      `protobuf:"fixed64,7,opt,name=deposit" json:"deposit,omitempty"`
	Leverage        float64      `protobuf:"fixed64,8,opt,name=leverage" json:"leverage,omitempty"`
	ForceClosePrice float64      `protobuf:"fixed64,9,opt,name=forceClosePrice" json:"forceClosePrice,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *Position) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Position) GetPType() PositionType {
	if m != nil {
		return m.PType
	}
	return PositionType_PositionNone
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

func (m *Position) GetDeposit() float64 {
	if m != nil {
		return m.Deposit
	}
	return 0
}

func (m *Position) GetLeverage() float64 {
	if m != nil {
		return m.Leverage
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
func (*RespPosition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

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
func (*RespBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

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
	proto.RegisterType((*Resp)(nil), "exsync.Resp")
	proto.RegisterType((*ReqMakeOrder)(nil), "exsync.ReqMakeOrder")
	proto.RegisterType((*RespMakeOrder)(nil), "exsync.RespMakeOrder")
	proto.RegisterType((*ReqCancelOrder)(nil), "exsync.ReqCancelOrder")
	proto.RegisterType((*ReqTrades)(nil), "exsync.ReqTrades")
	proto.RegisterType((*ReqOrders)(nil), "exsync.ReqOrders")
	proto.RegisterType((*RespIndex)(nil), "exsync.RespIndex")
	proto.RegisterType((*RespOrders)(nil), "exsync.RespOrders")
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
	MakeOrder(ctx context.Context, in *ReqMakeOrder, opts ...grpc.CallOption) (*RespMakeOrder, error)
	CancelOrder(ctx context.Context, in *ReqCancelOrder, opts ...grpc.CallOption) (*Resp, error)
	GetOrders(ctx context.Context, in *ReqOrders, opts ...grpc.CallOption) (*RespOrders, error)
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

func (c *syncServiceClient) MakeOrder(ctx context.Context, in *ReqMakeOrder, opts ...grpc.CallOption) (*RespMakeOrder, error) {
	out := new(RespMakeOrder)
	err := grpc.Invoke(ctx, "/exsync.SyncService/MakeOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) CancelOrder(ctx context.Context, in *ReqCancelOrder, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/exsync.SyncService/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetOrders(ctx context.Context, in *ReqOrders, opts ...grpc.CallOption) (*RespOrders, error) {
	out := new(RespOrders)
	err := grpc.Invoke(ctx, "/exsync.SyncService/GetOrders", in, out, c.cc, opts...)
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
	MakeOrder(context.Context, *ReqMakeOrder) (*RespMakeOrder, error)
	CancelOrder(context.Context, *ReqCancelOrder) (*Resp, error)
	GetOrders(context.Context, *ReqOrders) (*RespOrders, error)
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

func _SyncService_MakeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMakeOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).MakeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/MakeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).MakeOrder(ctx, req.(*ReqMakeOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCancelOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).CancelOrder(ctx, req.(*ReqCancelOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exsync.SyncService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetOrders(ctx, req.(*ReqOrders))
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
			MethodName: "MakeOrder",
			Handler:    _SyncService_MakeOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _SyncService_CancelOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _SyncService_GetOrders_Handler,
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
	Metadata: "protobuf/exsync/exsync.proto",
}

func init() { proto.RegisterFile("protobuf/exsync/exsync.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1081 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcf, 0x72, 0xe3, 0xc4,
	0x13, 0xb6, 0x2c, 0xcb, 0xb6, 0xda, 0x4e, 0x56, 0x99, 0xf5, 0xef, 0x57, 0x2a, 0xd7, 0x42, 0x65,
	0x07, 0x16, 0x42, 0x80, 0x65, 0xf1, 0x6e, 0xaa, 0xa8, 0xe5, 0xb4, 0xeb, 0xad, 0x4a, 0x51, 0xec,
	0x1f, 0x33, 0x4e, 0x0e, 0x54, 0x71, 0x99, 0x48, 0x1d, 0xaf, 0x2a, 0xf2, 0x48, 0x91, 0xe4, 0x90,
	0x9c, 0x79, 0x0b, 0xde, 0x84, 0x1b, 0x2f, 0xc0, 0x99, 0xd7, 0xa1, 0xe6, 0x8f, 0xe4, 0xb1, 0x97,
	0x50, 0x70, 0x80, 0x93, 0xfb, 0xeb, 0xef, 0xeb, 0x99, 0xee, 0x9e, 0x1e, 0x6b, 0xe0, 0x5e, 0x5e,
	0x64, 0x55, 0x76, 0xb6, 0x3a, 0xff, 0x02, 0xaf, 0xcb, 0x1b, 0x11, 0x99, 0x9f, 0x87, 0xca, 0x4d,
	0xba, 0x1a, 0x51, 0x1f, 0x7a, 0x0c, 0x2f, 0x67, 0x89, 0x58, 0xd0, 0x2e, 0x74, 0x66, 0x99, 0x58,
	0xd0, 0xf7, 0xc0, 0x65, 0x78, 0x49, 0xfe, 0x0f, 0x5d, 0xbc, 0x16, 0x7c, 0x89, 0xa1, 0xb3, 0xef,
	0x1c, 0xf8, 0xcc, 0x20, 0x29, 0x63, 0x58, 0xe6, 0xf4, 0x27, 0x07, 0x86, 0x0c, 0x2f, 0x5f, 0xf1,
	0x0b, 0x7c, 0x53, 0xc4, 0x58, 0xdc, 0x16, 0x40, 0x3e, 0x87, 0x5e, 0xf5, 0x2c, 0xaa, 0x92, 0x4c,
	0x84, 0xed, 0x7d, 0xe7, 0x60, 0x77, 0x72, 0xf7, 0xa1, 0x49, 0xe5, 0xa4, 0xe0, 0x31, 0x6a, 0x8a,
	0xd5, 0x1a, 0xb9, 0x0c, 0x5f, 0x66, 0x2b, 0x51, 0x85, 0xee, 0xbe, 0x73, 0xe0, 0x30, 0x83, 0xc8,
	0x08, 0xbc, 0xbc, 0x48, 0x22, 0x0c, 0x3b, 0xca, 0xad, 0x01, 0x7d, 0x02, 0x3b, 0x32, 0x9b, 0x75,
	0x16, 0x1f, 0x80, 0x97, 0x49, 0x43, 0x25, 0x31, 0x98, 0xec, 0xd4, 0x7b, 0x29, 0x96, 0x69, 0x8e,
	0x3e, 0x85, 0x5d, 0x86, 0x97, 0x53, 0x2e, 0x22, 0x4c, 0xff, 0x3a, 0xf9, 0x00, 0xdc, 0x24, 0x2e,
	0xc3, 0xf6, 0xbe, 0x7b, 0xe0, 0x33, 0x69, 0xd2, 0x6f, 0xc1, 0x67, 0x78, 0xa9, 0x52, 0x2f, 0x6f,
	0x0d, 0x1b, 0x81, 0x57, 0x26, 0x22, 0x42, 0x55, 0xb1, 0xcf, 0x34, 0x90, 0x8b, 0x89, 0xd5, 0x52,
	0xd5, 0xe5, 0x31, 0x69, 0xd2, 0x23, 0xb5, 0x98, 0x4a, 0xa1, 0xfc, 0x07, 0x39, 0xdc, 0x97, 0x61,
	0x65, 0xfe, 0x8d, 0x88, 0xf1, 0x5a, 0xee, 0x95, 0x48, 0x43, 0x45, 0x39, 0x4c, 0x03, 0xfa, 0x18,
	0x40, 0x4a, 0xcc, 0xd2, 0x0f, 0xa0, 0xab, 0x2a, 0x2f, 0x43, 0x67, 0xdf, 0x7d, 0xb7, 0x2d, 0x86,
	0xac, 0x83, 0x4c, 0x71, 0x0f, 0xa0, 0x5b, 0x29, 0x6b, 0x3b, 0x48, 0xf1, 0xcc, 0x90, 0xf4, 0x6b,
	0xf0, 0x4f, 0x92, 0x25, 0x96, 0x15, 0x5f, 0xe6, 0x24, 0x84, 0x5e, 0x89, 0x51, 0x26, 0xe2, 0x52,
	0xa5, 0xe3, 0xb2, 0x1a, 0xca, 0x34, 0x05, 0x17, 0x59, 0xa9, 0x5a, 0xe2, 0x32, 0x0d, 0xe8, 0xaf,
	0x6d, 0xf0, 0xf4, 0x09, 0xec, 0x42, 0x3b, 0x89, 0x4d, 0xe5, 0xed, 0x24, 0xfe, 0x57, 0xc7, 0x86,
	0xbc, 0x0f, 0x10, 0x23, 0x4f, 0x9f, 0xe9, 0x08, 0x4f, 0x51, 0x96, 0x87, 0xdc, 0x03, 0x5f, 0xa2,
	0x57, 0x99, 0xc0, 0x9b, 0xb0, 0xab, 0xe8, 0xb5, 0x83, 0x8c, 0xa1, 0xcf, 0xaf, 0x16, 0x33, 0xb5,
	0x6c, 0x4f, 0x91, 0x0d, 0x96, 0x87, 0x75, 0x8e, 0x18, 0xf6, 0x95, 0x5b, 0x9a, 0xe4, 0x53, 0xe8,
	0x96, 0x15, 0xaf, 0x56, 0x65, 0xe8, 0x6f, 0xd6, 0xa1, 0xea, 0x9e, 0x2b, 0x8a, 0x19, 0x09, 0xf9,
	0x12, 0x20, 0x2a, 0x90, 0x57, 0x28, 0x5b, 0x1a, 0x82, 0x9a, 0xe1, 0xbd, 0xa6, 0xf0, 0xba, 0xcd,
	0xcc, 0x12, 0xd1, 0x5f, 0x1c, 0xf0, 0x54, 0x4b, 0xfe, 0xdb, 0x16, 0x9a, 0x42, 0xbd, 0x75, 0xa1,
	0x9b, 0xb9, 0x77, 0xff, 0x4e, 0xee, 0xdf, 0xe9, 0x41, 0x7e, 0x81, 0x79, 0xf5, 0x96, 0xdc, 0x87,
	0x0e, 0x2f, 0x2f, 0x6e, 0x99, 0x36, 0x45, 0x49, 0xc9, 0x59, 0x7d, 0x17, 0xde, 0x95, 0x48, 0x8a,
	0xfe, 0xe6, 0x40, 0xef, 0x39, 0x4f, 0xe5, 0xdd, 0xb6, 0x2a, 0x72, 0x36, 0x2a, 0x0a, 0xa1, 0x17,
	0x63, 0x9e, 0x95, 0x49, 0xa5, 0x1a, 0xe3, 0xb0, 0x1a, 0xca, 0xc1, 0x28, 0x90, 0xa7, 0xb3, 0x22,
	0x3b, 0x4f, 0xea, 0x3e, 0x58, 0x1e, 0x42, 0x61, 0xb8, 0x12, 0x96, 0x42, 0xb7, 0x64, 0xc3, 0x27,
	0xc7, 0xa3, 0x48, 0xca, 0x0b, 0xc6, 0xab, 0xba, 0x3d, 0x0d, 0x26, 0x8f, 0xa0, 0x1f, 0xad, 0x8a,
	0x02, 0x45, 0xa4, 0xe7, 0x6a, 0x77, 0x32, 0xaa, 0x8b, 0x98, 0x1a, 0xff, 0xa9, 0x48, 0x2a, 0xd6,
	0xa8, 0xe8, 0xcf, 0x6d, 0xe8, 0xcf, 0x64, 0x6e, 0xf2, 0x88, 0xb6, 0x4f, 0xf8, 0x10, 0xbc, 0xfc,
	0xe4, 0x26, 0x47, 0x73, 0xbe, 0xcd, 0x5a, 0x75, 0x80, 0xe4, 0x98, 0x96, 0xdc, 0x7a, 0xbc, 0x07,
	0x70, 0x87, 0x5f, 0xf1, 0x24, 0xe5, 0x67, 0x29, 0x9a, 0x0b, 0xa1, 0xab, 0xda, 0x76, 0x6f, 0xcc,
	0xbd, 0xb7, 0x35, 0xf7, 0x23, 0xf0, 0x96, 0xd6, 0x6d, 0xd1, 0xc0, 0x6e, 0x74, 0x6f, 0xb3, 0xd1,
	0x63, 0xe8, 0xa7, 0x78, 0x85, 0x05, 0x5f, 0xd4, 0x97, 0xa5, 0xc1, 0x32, 0xa3, 0xf3, 0xac, 0x88,
	0x70, 0x9a, 0x66, 0x25, 0xea, 0xed, 0x7c, 0x9d, 0xd1, 0x96, 0x9b, 0xfe, 0x20, 0xbf, 0x41, 0x65,
	0xde, 0xf4, 0xe7, 0x43, 0xe8, 0xa4, 0x99, 0x58, 0x98, 0x3f, 0xff, 0x60, 0xbb, 0x1d, 0x4c, 0xb1,
	0xe4, 0x23, 0xf0, 0xca, 0xb7, 0x59, 0xa1, 0x0f, 0xff, 0xcf, 0x64, 0x9a, 0xa6, 0x5f, 0xc1, 0x40,
	0xae, 0x5e, 0x4f, 0xd3, 0x27, 0xd0, 0x3b, 0xd3, 0xa6, 0x59, 0xff, 0x4e, 0x1d, 0x68, 0x14, 0xac,
	0xe6, 0x0f, 0x05, 0x0c, 0xac, 0x2b, 0x46, 0x02, 0x18, 0x6a, 0xeb, 0x54, 0x5c, 0x64, 0x3f, 0x8a,
	0xa0, 0x45, 0x86, 0xd0, 0x7f, 0x93, 0xa3, 0x78, 0x99, 0x89, 0x45, 0xe0, 0x90, 0x1d, 0xf0, 0x25,
	0x9a, 0xcb, 0x5d, 0x83, 0xb6, 0x84, 0xaa, 0x46, 0xc5, 0xba, 0x64, 0x17, 0x40, 0x41, 0x4d, 0x77,
	0x48, 0x0f, 0xdc, 0xe7, 0xab, 0x9b, 0xc0, 0x23, 0x7d, 0xe8, 0xcc, 0x31, 0x4d, 0x83, 0xee, 0xe1,
	0x11, 0x0c, 0xed, 0x23, 0x97, 0x1b, 0xd6, 0xf8, 0x75, 0x26, 0x30, 0x68, 0x49, 0xad, 0xd9, 0xcc,
	0x07, 0xcf, 0x6c, 0x74, 0x38, 0x81, 0xa1, 0x3d, 0x75, 0x72, 0xe5, 0xe9, 0xeb, 0xef, 0x83, 0x96,
	0x34, 0x4e, 0xe7, 0x2f, 0x02, 0x47, 0xed, 0x75, 0x32, 0x0d, 0xda, 0xd2, 0x78, 0x79, 0x32, 0x0d,
	0xdc, 0xc3, 0x18, 0x06, 0xd6, 0x1f, 0x97, 0xdc, 0x49, 0x5b, 0x4d, 0x69, 0x03, 0xe8, 0x4d, 0xd5,
	0x0d, 0x8f, 0x03, 0x47, 0x82, 0x19, 0x2f, 0xaa, 0x84, 0xa7, 0x41, 0x5b, 0x16, 0x3d, 0xcd, 0x96,
	0x79, 0x8a, 0x15, 0x06, 0xae, 0x42, 0xea, 0x0b, 0x8c, 0x71, 0xd0, 0x51, 0x35, 0x2b, 0x94, 0x88,
	0x45, 0xe0, 0x4d, 0x7e, 0x77, 0x61, 0x30, 0xbf, 0x11, 0xd1, 0x1c, 0x8b, 0x2b, 0x39, 0x5e, 0x1f,
	0x43, 0x47, 0x3e, 0x52, 0x48, 0xd3, 0x72, 0xf3, 0x6a, 0x19, 0x0f, 0xd7, 0x87, 0x27, 0x16, 0xb4,
	0x45, 0x9e, 0x82, 0xbf, 0x7e, 0x0c, 0x8c, 0x2c, 0x75, 0xe3, 0x1d, 0xff, 0x6f, 0xed, 0xb5, 0x5e,
	0x0e, 0xb4, 0x45, 0x8e, 0x60, 0xb0, 0xf1, 0x26, 0xb0, 0xa2, 0x2d, 0xff, 0x7a, 0x4b, 0xf5, 0x0e,
	0x6a, 0x91, 0x27, 0xe0, 0x1f, 0x63, 0x65, 0xbe, 0xb4, 0x7b, 0x56, 0x90, 0x76, 0x8d, 0x89, 0xad,
	0xd7, 0xbe, 0x26, 0xca, 0x7c, 0x6a, 0xed, 0x28, 0xed, 0xda, 0x8c, 0xd2, 0x3e, 0xda, 0x22, 0x9f,
	0x41, 0xff, 0x18, 0x2b, 0xfd, 0x7f, 0x39, 0xb0, 0x82, 0xc6, 0x7b, 0xb6, 0x5c, 0xf1, 0x8d, 0x5a,
	0x3f, 0x13, 0x6e, 0x57, 0x2b, 0x9e, 0xb6, 0xc8, 0x04, 0x06, 0xc7, 0x58, 0x35, 0x77, 0x69, 0x23,
	0x60, 0x64, 0x07, 0xd4, 0x12, 0xda, 0x22, 0x8f, 0x00, 0x8e, 0xb1, 0xaa, 0x6f, 0xc8, 0x46, 0xc8,
	0x5d, 0x3b, 0xc4, 0x28, 0x68, 0xeb, 0xac, 0xab, 0x1e, 0xa0, 0x8f, 0xff, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x54, 0xee, 0x4f, 0x3f, 0xa0, 0x0a, 0x00, 0x00,
}
