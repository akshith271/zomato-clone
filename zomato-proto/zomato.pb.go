// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zomato.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type NewUser struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUser) Reset()         { *m = NewUser{} }
func (m *NewUser) String() string { return proto.CompactTextString(m) }
func (*NewUser) ProtoMessage()    {}
func (*NewUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{1}
}

func (m *NewUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUser.Unmarshal(m, b)
}
func (m *NewUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUser.Marshal(b, m, deterministic)
}
func (m *NewUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUser.Merge(m, src)
}
func (m *NewUser) XXX_Size() int {
	return xxx_messageInfo_NewUser.Size(m)
}
func (m *NewUser) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUser.DiscardUnknown(m)
}

var xxx_messageInfo_NewUser proto.InternalMessageInfo

func (m *NewUser) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NewUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewUser) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *NewUser) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NewUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type VoidUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoidUserRequest) Reset()         { *m = VoidUserRequest{} }
func (m *VoidUserRequest) String() string { return proto.CompactTextString(m) }
func (*VoidUserRequest) ProtoMessage()    {}
func (*VoidUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{2}
}

func (m *VoidUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoidUserRequest.Unmarshal(m, b)
}
func (m *VoidUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoidUserRequest.Marshal(b, m, deterministic)
}
func (m *VoidUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoidUserRequest.Merge(m, src)
}
func (m *VoidUserRequest) XXX_Size() int {
	return xxx_messageInfo_VoidUserRequest.Size(m)
}
func (m *VoidUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoidUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoidUserRequest proto.InternalMessageInfo

type OrderItem struct {
	DishId               int64    `protobuf:"varint,1,opt,name=DishId,proto3" json:"DishId,omitempty"`
	Quantity             int64    `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{3}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetDishId() int64 {
	if m != nil {
		return m.DishId
	}
	return 0
}

func (m *OrderItem) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type UserOrder struct {
	OrderItems           []*OrderItem `protobuf:"bytes,1,rep,name=orderItems,proto3" json:"orderItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserOrder) Reset()         { *m = UserOrder{} }
func (m *UserOrder) String() string { return proto.CompactTextString(m) }
func (*UserOrder) ProtoMessage()    {}
func (*UserOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{4}
}

func (m *UserOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOrder.Unmarshal(m, b)
}
func (m *UserOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOrder.Marshal(b, m, deterministic)
}
func (m *UserOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrder.Merge(m, src)
}
func (m *UserOrder) XXX_Size() int {
	return xxx_messageInfo_UserOrder.Size(m)
}
func (m *UserOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrder.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrder proto.InternalMessageInfo

func (m *UserOrder) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

type NewRestaurant struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Timings              string   `protobuf:"bytes,3,opt,name=Timings,proto3" json:"Timings,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	IsActive             bool     `protobuf:"varint,6,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRestaurant) Reset()         { *m = NewRestaurant{} }
func (m *NewRestaurant) String() string { return proto.CompactTextString(m) }
func (*NewRestaurant) ProtoMessage()    {}
func (*NewRestaurant) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{5}
}

func (m *NewRestaurant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRestaurant.Unmarshal(m, b)
}
func (m *NewRestaurant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRestaurant.Marshal(b, m, deterministic)
}
func (m *NewRestaurant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRestaurant.Merge(m, src)
}
func (m *NewRestaurant) XXX_Size() int {
	return xxx_messageInfo_NewRestaurant.Size(m)
}
func (m *NewRestaurant) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRestaurant.DiscardUnknown(m)
}

var xxx_messageInfo_NewRestaurant proto.InternalMessageInfo

func (m *NewRestaurant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewRestaurant) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NewRestaurant) GetTimings() string {
	if m != nil {
		return m.Timings
	}
	return ""
}

func (m *NewRestaurant) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewRestaurant) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *NewRestaurant) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type Restaurant struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Timings              string   `protobuf:"bytes,3,opt,name=Timings,proto3" json:"Timings,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	IsActive             bool     `protobuf:"varint,6,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Restaurant) Reset()         { *m = Restaurant{} }
func (m *Restaurant) String() string { return proto.CompactTextString(m) }
func (*Restaurant) ProtoMessage()    {}
func (*Restaurant) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{6}
}

func (m *Restaurant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restaurant.Unmarshal(m, b)
}
func (m *Restaurant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restaurant.Marshal(b, m, deterministic)
}
func (m *Restaurant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restaurant.Merge(m, src)
}
func (m *Restaurant) XXX_Size() int {
	return xxx_messageInfo_Restaurant.Size(m)
}
func (m *Restaurant) XXX_DiscardUnknown() {
	xxx_messageInfo_Restaurant.DiscardUnknown(m)
}

var xxx_messageInfo_Restaurant proto.InternalMessageInfo

func (m *Restaurant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Restaurant) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Restaurant) GetTimings() string {
	if m != nil {
		return m.Timings
	}
	return ""
}

func (m *Restaurant) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Restaurant) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Restaurant) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type NewDish struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	Price                int64    `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
	RestaurantId         int64    `protobuf:"varint,5,opt,name=RestaurantId,proto3" json:"RestaurantId,omitempty"`
	IsAvailable          bool     `protobuf:"varint,6,opt,name=IsAvailable,proto3" json:"IsAvailable,omitempty"`
	IsVeg                bool     `protobuf:"varint,7,opt,name=IsVeg,proto3" json:"IsVeg,omitempty"`
	Cuisine              string   `protobuf:"bytes,8,opt,name=Cuisine,proto3" json:"Cuisine,omitempty"`
	Category             string   `protobuf:"bytes,9,opt,name=Category,proto3" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewDish) Reset()         { *m = NewDish{} }
func (m *NewDish) String() string { return proto.CompactTextString(m) }
func (*NewDish) ProtoMessage()    {}
func (*NewDish) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{7}
}

func (m *NewDish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewDish.Unmarshal(m, b)
}
func (m *NewDish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewDish.Marshal(b, m, deterministic)
}
func (m *NewDish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewDish.Merge(m, src)
}
func (m *NewDish) XXX_Size() int {
	return xxx_messageInfo_NewDish.Size(m)
}
func (m *NewDish) XXX_DiscardUnknown() {
	xxx_messageInfo_NewDish.DiscardUnknown(m)
}

var xxx_messageInfo_NewDish proto.InternalMessageInfo

func (m *NewDish) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewDish) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NewDish) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *NewDish) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *NewDish) GetRestaurantId() int64 {
	if m != nil {
		return m.RestaurantId
	}
	return 0
}

func (m *NewDish) GetIsAvailable() bool {
	if m != nil {
		return m.IsAvailable
	}
	return false
}

func (m *NewDish) GetIsVeg() bool {
	if m != nil {
		return m.IsVeg
	}
	return false
}

func (m *NewDish) GetCuisine() string {
	if m != nil {
		return m.Cuisine
	}
	return ""
}

func (m *NewDish) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type Dish struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	Price                int64    `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
	RestaurantId         int64    `protobuf:"varint,5,opt,name=RestaurantId,proto3" json:"RestaurantId,omitempty"`
	IsAvailable          bool     `protobuf:"varint,6,opt,name=IsAvailable,proto3" json:"IsAvailable,omitempty"`
	IsVeg                bool     `protobuf:"varint,7,opt,name=IsVeg,proto3" json:"IsVeg,omitempty"`
	Cuisine              string   `protobuf:"bytes,8,opt,name=Cuisine,proto3" json:"Cuisine,omitempty"`
	Category             string   `protobuf:"bytes,9,opt,name=Category,proto3" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dish) Reset()         { *m = Dish{} }
func (m *Dish) String() string { return proto.CompactTextString(m) }
func (*Dish) ProtoMessage()    {}
func (*Dish) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{8}
}

func (m *Dish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dish.Unmarshal(m, b)
}
func (m *Dish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dish.Marshal(b, m, deterministic)
}
func (m *Dish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dish.Merge(m, src)
}
func (m *Dish) XXX_Size() int {
	return xxx_messageInfo_Dish.Size(m)
}
func (m *Dish) XXX_DiscardUnknown() {
	xxx_messageInfo_Dish.DiscardUnknown(m)
}

var xxx_messageInfo_Dish proto.InternalMessageInfo

func (m *Dish) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dish) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Dish) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Dish) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Dish) GetRestaurantId() int64 {
	if m != nil {
		return m.RestaurantId
	}
	return 0
}

func (m *Dish) GetIsAvailable() bool {
	if m != nil {
		return m.IsAvailable
	}
	return false
}

func (m *Dish) GetIsVeg() bool {
	if m != nil {
		return m.IsVeg
	}
	return false
}

func (m *Dish) GetCuisine() string {
	if m != nil {
		return m.Cuisine
	}
	return ""
}

func (m *Dish) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type VoidDishResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoidDishResponse) Reset()         { *m = VoidDishResponse{} }
func (m *VoidDishResponse) String() string { return proto.CompactTextString(m) }
func (*VoidDishResponse) ProtoMessage()    {}
func (*VoidDishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf5f31cd82a50922, []int{9}
}

func (m *VoidDishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoidDishResponse.Unmarshal(m, b)
}
func (m *VoidDishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoidDishResponse.Marshal(b, m, deterministic)
}
func (m *VoidDishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoidDishResponse.Merge(m, src)
}
func (m *VoidDishResponse) XXX_Size() int {
	return xxx_messageInfo_VoidDishResponse.Size(m)
}
func (m *VoidDishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VoidDishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VoidDishResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "zomatoDB.User")
	proto.RegisterType((*NewUser)(nil), "zomatoDB.NewUser")
	proto.RegisterType((*VoidUserRequest)(nil), "zomatoDB.VoidUserRequest")
	proto.RegisterType((*OrderItem)(nil), "zomatoDB.OrderItem")
	proto.RegisterType((*UserOrder)(nil), "zomatoDB.UserOrder")
	proto.RegisterType((*NewRestaurant)(nil), "zomatoDB.NewRestaurant")
	proto.RegisterType((*Restaurant)(nil), "zomatoDB.Restaurant")
	proto.RegisterType((*NewDish)(nil), "zomatoDB.NewDish")
	proto.RegisterType((*Dish)(nil), "zomatoDB.Dish")
	proto.RegisterType((*VoidDishResponse)(nil), "zomatoDB.VoidDishResponse")
}

func init() { proto.RegisterFile("zomato.proto", fileDescriptor_cf5f31cd82a50922) }

var fileDescriptor_cf5f31cd82a50922 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x95, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xc7, 0xb3, 0xdd, 0xb4, 0x4d, 0xa6, 0x2f, 0x4f, 0xeb, 0x56, 0x0f, 0xab, 0x08, 0xa1, 0xc8,
	0xa7, 0x5c, 0x9a, 0x8a, 0x56, 0xe2, 0xc0, 0x05, 0xd2, 0x04, 0xa1, 0xbd, 0x84, 0xb2, 0xa2, 0x3d,
	0xf4, 0xe6, 0x64, 0x47, 0xa9, 0x45, 0x76, 0xbd, 0xd8, 0x4e, 0xa3, 0xf2, 0x35, 0xf8, 0x06, 0x48,
	0x7c, 0x49, 0x24, 0x24, 0x64, 0x7b, 0xdf, 0x12, 0xc2, 0x15, 0x21, 0x71, 0x69, 0xf7, 0x37, 0x33,
	0xf6, 0xfc, 0xe7, 0xbf, 0x6b, 0x07, 0xf6, 0x3f, 0x8b, 0x84, 0x69, 0xd1, 0xcf, 0xa4, 0xd0, 0x82,
	0xb4, 0x1c, 0x8d, 0xae, 0x68, 0x06, 0xcd, 0x1b, 0x85, 0x92, 0x1c, 0xc2, 0x16, 0x8f, 0x03, 0xaf,
	0xeb, 0xf5, 0xfc, 0x68, 0x8b, 0xc7, 0x84, 0x40, 0x33, 0x65, 0x09, 0x06, 0x5b, 0x5d, 0xaf, 0xd7,
	0x8e, 0xec, 0x33, 0x39, 0x85, 0xed, 0xec, 0x5e, 0xa4, 0x18, 0x6c, 0xdb, 0xa0, 0x03, 0x12, 0xc0,
	0x2e, 0x8b, 0x63, 0x89, 0x4a, 0x05, 0xbe, 0x8d, 0x17, 0x68, 0xea, 0x31, 0x61, 0x7c, 0x1e, 0x34,
	0x5d, 0xbd, 0x05, 0xaa, 0x60, 0x77, 0x8c, 0xcb, 0x3f, 0xdc, 0xf4, 0x18, 0xfe, 0xbb, 0x15, 0x3c,
	0x36, 0x5d, 0x23, 0xfc, 0xb4, 0x40, 0xa5, 0xe9, 0x2b, 0x68, 0xbf, 0x93, 0x31, 0xca, 0x50, 0x63,
	0x42, 0xfe, 0x87, 0x9d, 0x11, 0x57, 0xf7, 0x61, 0xa1, 0x26, 0x27, 0xd2, 0x81, 0xd6, 0xfb, 0x05,
	0x4b, 0x35, 0xd7, 0x8f, 0x56, 0x95, 0x1f, 0x95, 0x4c, 0x5f, 0x43, 0xdb, 0xec, 0x67, 0x37, 0x21,
	0x97, 0x00, 0xa2, 0xd8, 0x4d, 0x05, 0x5e, 0xd7, 0xef, 0xed, 0x5d, 0x9c, 0xf4, 0x0b, 0x9b, 0xfb,
	0x65, 0xa7, 0xa8, 0x56, 0x46, 0xbf, 0x79, 0x70, 0x30, 0xc6, 0x65, 0x84, 0x4a, 0xb3, 0x85, 0x64,
	0xa9, 0x36, 0x0e, 0x8c, 0x8d, 0x03, 0x9e, 0x73, 0xc0, 0x3c, 0x9b, 0x59, 0x07, 0xf9, 0xac, 0xce,
	0x98, 0x02, 0x4d, 0xe6, 0x03, 0x4f, 0x78, 0x3a, 0x2b, 0x5d, 0xc8, 0xd1, 0xb8, 0xf0, 0xa6, 0xee,
	0x82, 0x05, 0x33, 0xcd, 0x35, 0x53, 0x6a, 0x29, 0x64, 0x9c, 0xdb, 0x59, 0xb2, 0xc9, 0x85, 0x6a,
	0x30, 0xd5, 0xfc, 0x01, 0x83, 0x9d, 0xae, 0xd7, 0x6b, 0x45, 0x25, 0xd3, 0xaf, 0x1e, 0xc0, 0x5f,
	0x2f, 0xf2, 0x87, 0x67, 0x3f, 0x2c, 0xf3, 0xe2, 0x36, 0x2a, 0xec, 0xc2, 0xde, 0x08, 0xd5, 0x54,
	0xf2, 0x4c, 0x73, 0x91, 0xe6, 0x2a, 0xeb, 0x21, 0xa3, 0x27, 0x4c, 0xd8, 0x0c, 0x73, 0x9d, 0x0e,
	0x4c, 0xf4, 0x5a, 0xf2, 0x29, 0x5a, 0x95, 0x7e, 0xe4, 0x80, 0x50, 0xd8, 0xaf, 0x1c, 0x09, 0x9d,
	0x52, 0x3f, 0x5a, 0x89, 0x99, 0x8e, 0xa1, 0x1a, 0x3c, 0x30, 0x3e, 0x67, 0x93, 0x79, 0x21, 0xb8,
	0x1e, 0xb2, 0x1d, 0xd5, 0x2d, 0xce, 0x82, 0x5d, 0x9b, 0x73, 0x60, 0x1c, 0x1b, 0x2e, 0xb8, 0xe2,
	0x29, 0x06, 0x2d, 0xe7, 0x58, 0x8e, 0x66, 0xfe, 0x21, 0xd3, 0x38, 0x13, 0xf2, 0x31, 0x68, 0x3b,
	0x6f, 0x0a, 0xa6, 0xdf, 0x3d, 0x68, 0xfe, 0xb3, 0xc3, 0x13, 0x38, 0x32, 0xe7, 0xdb, 0xcc, 0x1f,
	0xa1, 0xca, 0x44, 0xaa, 0xf0, 0xe2, 0x8b, 0x0f, 0xe4, 0xce, 0x1d, 0x40, 0xa6, 0xd9, 0x84, 0x29,
	0x1c, 0xca, 0x45, 0x4c, 0x9e, 0x03, 0x0c, 0x25, 0x32, 0x8d, 0xf6, 0x0a, 0x3a, 0xae, 0xce, 0x68,
	0x7e, 0x2b, 0x75, 0x0e, 0xab, 0x90, 0x61, 0xda, 0x20, 0x7d, 0x80, 0x9b, 0x2c, 0x2e, 0x96, 0xac,
	0xe5, 0x37, 0xd4, 0xbf, 0x80, 0x83, 0xb7, 0xa8, 0xcb, 0xcb, 0x41, 0xfd, 0xb2, 0xe4, 0x64, 0x95,
	0x6d, 0x15, 0x6d, 0x90, 0x21, 0x1c, 0x39, 0x69, 0xb5, 0xc3, 0xf6, 0x64, 0x45, 0x60, 0x95, 0xe8,
	0x9c, 0x56, 0x89, 0x2a, 0x4a, 0x1b, 0xd5, 0x7c, 0xf6, 0x63, 0x58, 0x9d, 0xcf, 0x84, 0xea, 0x7a,
	0x0d, 0xd7, 0xe7, 0xb3, 0x4b, 0xd6, 0xf2, 0x1b, 0xea, 0x5f, 0x02, 0x8c, 0x70, 0x8e, 0xbf, 0xa9,
	0xef, 0x54, 0xbc, 0xfe, 0x4e, 0x68, 0xe3, 0xea, 0xd9, 0xdd, 0xd3, 0x44, 0x4c, 0x3f, 0x9e, 0xcd,
	0x64, 0x36, 0x3d, 0x77, 0x85, 0x67, 0xf6, 0x47, 0xe9, 0xdc, 0xfe, 0x9d, 0xec, 0xd8, 0x7f, 0x97,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x05, 0xda, 0xcc, 0x85, 0xb1, 0x06, 0x00, 0x00,
}
