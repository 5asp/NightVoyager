// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: accountsvc.proto

package accountsvc

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

type GetByAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetByAccountReq) Reset() {
	*x = GetByAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByAccountReq) ProtoMessage() {}

func (x *GetByAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByAccountReq.ProtoReflect.Descriptor instead.
func (*GetByAccountReq) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{0}
}

func (x *GetByAccountReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type GetByAccountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Account `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetByAccountResp) Reset() {
	*x = GetByAccountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByAccountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByAccountResp) ProtoMessage() {}

func (x *GetByAccountResp) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByAccountResp.ProtoReflect.Descriptor instead.
func (*GetByAccountResp) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{1}
}

func (x *GetByAccountResp) GetData() *Account {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdReq) Reset() {
	*x = GetByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdReq) ProtoMessage() {}

func (x *GetByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdReq.ProtoReflect.Descriptor instead.
func (*GetByIdReq) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{2}
}

func (x *GetByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Account `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetByIdResp) Reset() {
	*x = GetByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdResp) ProtoMessage() {}

func (x *GetByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdResp.ProtoReflect.Descriptor instead.
func (*GetByIdResp) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIdResp) GetData() *Account {
	if x != nil {
		return x.Data
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Account   string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Status    int64  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{4}
}

func (x *Account) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Account) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Account) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Account) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Account `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{5}
}

func (x *CreateReq) GetData() *Account {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResp) Reset() {
	*x = CreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResp) ProtoMessage() {}

func (x *CreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResp.ProtoReflect.Descriptor instead.
func (*CreateResp) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{6}
}

func (x *CreateResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Account `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateReq) GetData() *Account {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteResp) Reset() {
	*x = DeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResp) ProtoMessage() {}

func (x *DeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResp.ProtoReflect.Descriptor instead.
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteResp) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AccountLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId int64  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	IpAddr    string `protobuf:"bytes,4,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	Device    string `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *AccountLog) Reset() {
	*x = AccountLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountLog) ProtoMessage() {}

func (x *AccountLog) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountLog.ProtoReflect.Descriptor instead.
func (*AccountLog) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{11}
}

func (x *AccountLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountLog) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *AccountLog) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AccountLog) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *AccountLog) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type InsertLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *AccountLog `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InsertLogReq) Reset() {
	*x = InsertLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertLogReq) ProtoMessage() {}

func (x *InsertLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertLogReq.ProtoReflect.Descriptor instead.
func (*InsertLogReq) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{12}
}

func (x *InsertLogReq) GetData() *AccountLog {
	if x != nil {
		return x.Data
	}
	return nil
}

type InsertLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InsertLogResp) Reset() {
	*x = InsertLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertLogResp) ProtoMessage() {}

func (x *InsertLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertLogResp.ProtoReflect.Descriptor instead.
func (*InsertLogResp) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{13}
}

func (x *InsertLogResp) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AccountLogListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccountLogListReq) Reset() {
	*x = AccountLogListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountLogListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountLogListReq) ProtoMessage() {}

func (x *AccountLogListReq) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountLogListReq.ProtoReflect.Descriptor instead.
func (*AccountLogListReq) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{14}
}

func (x *AccountLogListReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AccountLogListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AccountLog `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *AccountLogListResp) Reset() {
	*x = AccountLogListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accountsvc_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountLogListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountLogListResp) ProtoMessage() {}

func (x *AccountLogListResp) ProtoReflect() protoreflect.Message {
	mi := &file_accountsvc_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountLogListResp.ProtoReflect.Descriptor instead.
func (*AccountLogListResp) Descriptor() ([]byte, []int) {
	return file_accountsvc_proto_rawDescGZIP(), []int{15}
}

func (x *AccountLogListResp) GetList() []*AccountLog {
	if x != nil {
		return x.List
	}
	return nil
}

var File_accountsvc_proto protoreflect.FileDescriptor

var file_accountsvc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x22, 0x2b,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa5,
	0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1c, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x76, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x1c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x8b, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x3a, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x0d, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xbe, 0x02, 0x0a, 0x0a,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x76, 0x63, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76,
	0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76,
	0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32, 0xa2, 0x01, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x76, 0x63, 0x12, 0x40,
	0x0a, 0x09, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x76, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x4f, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x76,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accountsvc_proto_rawDescOnce sync.Once
	file_accountsvc_proto_rawDescData = file_accountsvc_proto_rawDesc
)

func file_accountsvc_proto_rawDescGZIP() []byte {
	file_accountsvc_proto_rawDescOnce.Do(func() {
		file_accountsvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_accountsvc_proto_rawDescData)
	})
	return file_accountsvc_proto_rawDescData
}

var file_accountsvc_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_accountsvc_proto_goTypes = []interface{}{
	(*GetByAccountReq)(nil),    // 0: accountsvc.GetByAccountReq
	(*GetByAccountResp)(nil),   // 1: accountsvc.GetByAccountResp
	(*GetByIdReq)(nil),         // 2: accountsvc.GetByIdReq
	(*GetByIdResp)(nil),        // 3: accountsvc.GetByIdResp
	(*Account)(nil),            // 4: accountsvc.Account
	(*CreateReq)(nil),          // 5: accountsvc.CreateReq
	(*CreateResp)(nil),         // 6: accountsvc.CreateResp
	(*UpdateReq)(nil),          // 7: accountsvc.UpdateReq
	(*UpdateResp)(nil),         // 8: accountsvc.UpdateResp
	(*DeleteReq)(nil),          // 9: accountsvc.DeleteReq
	(*DeleteResp)(nil),         // 10: accountsvc.DeleteResp
	(*AccountLog)(nil),         // 11: accountsvc.AccountLog
	(*InsertLogReq)(nil),       // 12: accountsvc.InsertLogReq
	(*InsertLogResp)(nil),      // 13: accountsvc.InsertLogResp
	(*AccountLogListReq)(nil),  // 14: accountsvc.AccountLogListReq
	(*AccountLogListResp)(nil), // 15: accountsvc.AccountLogListResp
}
var file_accountsvc_proto_depIdxs = []int32{
	4,  // 0: accountsvc.GetByAccountResp.data:type_name -> accountsvc.Account
	4,  // 1: accountsvc.GetByIdResp.data:type_name -> accountsvc.Account
	4,  // 2: accountsvc.CreateReq.data:type_name -> accountsvc.Account
	4,  // 3: accountsvc.UpdateReq.data:type_name -> accountsvc.Account
	11, // 4: accountsvc.InsertLogReq.data:type_name -> accountsvc.AccountLog
	11, // 5: accountsvc.AccountLogListResp.List:type_name -> accountsvc.AccountLog
	0,  // 6: accountsvc.AccountSvc.GetByAccount:input_type -> accountsvc.GetByAccountReq
	2,  // 7: accountsvc.AccountSvc.GetById:input_type -> accountsvc.GetByIdReq
	5,  // 8: accountsvc.AccountSvc.Create:input_type -> accountsvc.CreateReq
	7,  // 9: accountsvc.AccountSvc.Update:input_type -> accountsvc.UpdateReq
	9,  // 10: accountsvc.AccountSvc.Delete:input_type -> accountsvc.DeleteReq
	12, // 11: accountsvc.AccountLogSvc.InsertLog:input_type -> accountsvc.InsertLogReq
	14, // 12: accountsvc.AccountLogSvc.AccountLogList:input_type -> accountsvc.AccountLogListReq
	1,  // 13: accountsvc.AccountSvc.GetByAccount:output_type -> accountsvc.GetByAccountResp
	3,  // 14: accountsvc.AccountSvc.GetById:output_type -> accountsvc.GetByIdResp
	6,  // 15: accountsvc.AccountSvc.Create:output_type -> accountsvc.CreateResp
	8,  // 16: accountsvc.AccountSvc.Update:output_type -> accountsvc.UpdateResp
	10, // 17: accountsvc.AccountSvc.Delete:output_type -> accountsvc.DeleteResp
	13, // 18: accountsvc.AccountLogSvc.InsertLog:output_type -> accountsvc.InsertLogResp
	15, // 19: accountsvc.AccountLogSvc.AccountLogList:output_type -> accountsvc.AccountLogListResp
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_accountsvc_proto_init() }
func file_accountsvc_proto_init() {
	if File_accountsvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accountsvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByAccountReq); i {
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
		file_accountsvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByAccountResp); i {
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
		file_accountsvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdReq); i {
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
		file_accountsvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdResp); i {
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
		file_accountsvc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_accountsvc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_accountsvc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResp); i {
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
		file_accountsvc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_accountsvc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResp); i {
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
		file_accountsvc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_accountsvc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResp); i {
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
		file_accountsvc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountLog); i {
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
		file_accountsvc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertLogReq); i {
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
		file_accountsvc_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertLogResp); i {
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
		file_accountsvc_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountLogListReq); i {
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
		file_accountsvc_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountLogListResp); i {
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
			RawDescriptor: file_accountsvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_accountsvc_proto_goTypes,
		DependencyIndexes: file_accountsvc_proto_depIdxs,
		MessageInfos:      file_accountsvc_proto_msgTypes,
	}.Build()
	File_accountsvc_proto = out.File
	file_accountsvc_proto_rawDesc = nil
	file_accountsvc_proto_goTypes = nil
	file_accountsvc_proto_depIdxs = nil
}
