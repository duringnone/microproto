// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: mail/mail.proto

package mail

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 共用响应格式
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code"` // 错误码
	Msg  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`    // 错误信息
	Data []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`  // 数组 返回数据
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

// 发送邮件
type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail     string `protobuf:"bytes,3,opt,name=mail,proto3" json:"mail"`         // 邮箱
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`       // 邮件标题
	Content  string `protobuf:"bytes,5,opt,name=content,proto3" json:"content"`   // 邮件内容
	SerialNo string `protobuf:"bytes,8,opt,name=SerialNo,proto3" json:"SerialNo"` // 流水号
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{1}
}

func (x *SendMailRequest) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *SendMailRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendMailRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendMailRequest) GetSerialNo() string {
	if x != nil {
		return x.SerialNo
	}
	return ""
}

type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code"` // 错误码
	Msg  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`    // 错误信息
	Data []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`  // 数组 返回数据
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{2}
}

func (x *SendMailResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SendMailResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SendMailResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

// 获取邮件模板列表/详情
type GetMailConfigListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EId        int64  `protobuf:"varint,2,opt,name=eId,proto3" json:"eId"`              // 模板表ID(详情)
	EmailName  string `protobuf:"bytes,3,opt,name=emailName,proto3" json:"emailName"`   // 邮件模板名(模糊查询)
	EmailTitle string `protobuf:"bytes,4,opt,name=emailTitle,proto3" json:"emailTitle"` // 邮件模板标题(模糊)
	Page       int64  `protobuf:"varint,8,opt,name=page,proto3" json:"page"`            //当前页数
	PageSize   int64  `protobuf:"varint,9,opt,name=pageSize,proto3" json:"pageSize"`    // 每页显示数
	SerialNo   string `protobuf:"bytes,10,opt,name=SerialNo,proto3" json:"SerialNo"`    // 流水号 [用于微服务全链路日志追踪]
}

func (x *GetMailConfigListRequest) Reset() {
	*x = GetMailConfigListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMailConfigListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailConfigListRequest) ProtoMessage() {}

func (x *GetMailConfigListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailConfigListRequest.ProtoReflect.Descriptor instead.
func (*GetMailConfigListRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{3}
}

func (x *GetMailConfigListRequest) GetEId() int64 {
	if x != nil {
		return x.EId
	}
	return 0
}

func (x *GetMailConfigListRequest) GetEmailName() string {
	if x != nil {
		return x.EmailName
	}
	return ""
}

func (x *GetMailConfigListRequest) GetEmailTitle() string {
	if x != nil {
		return x.EmailTitle
	}
	return ""
}

func (x *GetMailConfigListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMailConfigListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetMailConfigListRequest) GetSerialNo() string {
	if x != nil {
		return x.SerialNo
	}
	return ""
}

type GetMailConfigListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64           `protobuf:"varint,1,opt,name=code,proto3" json:"code"` // 错误码
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`    // 错误信息
	Data *MailConfigList `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`  // 分页信息
}

func (x *GetMailConfigListResponse) Reset() {
	*x = GetMailConfigListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMailConfigListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailConfigListResponse) ProtoMessage() {}

func (x *GetMailConfigListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailConfigListResponse.ProtoReflect.Descriptor instead.
func (*GetMailConfigListResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{4}
}

func (x *GetMailConfigListResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetMailConfigListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetMailConfigListResponse) GetData() *MailConfigList {
	if x != nil {
		return x.Data
	}
	return nil
}

// 分页列表 {TotalCount:30,Page:1,PageSize:20,List:[map,map,map,...]}
type MailConfigList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64               `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount"` // 总数
	Page       int64               `protobuf:"varint,2,opt,name=page,proto3" json:"page"`             // 当前页数
	PageSize   int64               `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize"`     // 每页显示数
	List       []*MailConfigSingle `protobuf:"bytes,4,rep,name=list,proto3" json:"list"`
}

func (x *MailConfigList) Reset() {
	*x = MailConfigList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailConfigList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailConfigList) ProtoMessage() {}

func (x *MailConfigList) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailConfigList.ProtoReflect.Descriptor instead.
func (*MailConfigList) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{5}
}

func (x *MailConfigList) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *MailConfigList) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *MailConfigList) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MailConfigList) GetList() []*MailConfigSingle {
	if x != nil {
		return x.List
	}
	return nil
}

// 单条
type MailConfigSingle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EId          string `protobuf:"bytes,1,opt,name=e_id,json=eId,proto3" json:"e_id"`                            // 邮件模板表ID
	EAddDt       string `protobuf:"bytes,4,opt,name=e_add_dt,json=eAddDt,proto3" json:"e_add_dt"`                 // 添加时间
	EModifyDt    string `protobuf:"bytes,5,opt,name=e_modify_dt,json=eModifyDt,proto3" json:"e_modify_dt"`        // 更新时间
	EmailContent string `protobuf:"bytes,7,opt,name=email_content,json=emailContent,proto3" json:"email_content"` // 邮件内容
	EmailTitle   string `protobuf:"bytes,8,opt,name=email_title,json=emailTitle,proto3" json:"email_title"`       // 邮件标题
	EmailName    string `protobuf:"bytes,11,opt,name=email_name,json=emailName,proto3" json:"email_name"`         // 邮件模板名(唯一)
}

func (x *MailConfigSingle) Reset() {
	*x = MailConfigSingle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailConfigSingle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailConfigSingle) ProtoMessage() {}

func (x *MailConfigSingle) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailConfigSingle.ProtoReflect.Descriptor instead.
func (*MailConfigSingle) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{6}
}

func (x *MailConfigSingle) GetEId() string {
	if x != nil {
		return x.EId
	}
	return ""
}

func (x *MailConfigSingle) GetEAddDt() string {
	if x != nil {
		return x.EAddDt
	}
	return ""
}

func (x *MailConfigSingle) GetEModifyDt() string {
	if x != nil {
		return x.EModifyDt
	}
	return ""
}

func (x *MailConfigSingle) GetEmailContent() string {
	if x != nil {
		return x.EmailContent
	}
	return ""
}

func (x *MailConfigSingle) GetEmailTitle() string {
	if x != nil {
		return x.EmailTitle
	}
	return ""
}

func (x *MailConfigSingle) GetEmailName() string {
	if x != nil {
		return x.EmailName
	}
	return ""
}

// 添加邮件模板
type AddMailConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailName    string `protobuf:"bytes,2,opt,name=emailName,proto3" json:"emailName"`       // 邮件模板名
	EmailTitle   string `protobuf:"bytes,3,opt,name=emailTitle,proto3" json:"emailTitle"`     // 模板邮件标题
	EmailContent string `protobuf:"bytes,4,opt,name=emailContent,proto3" json:"emailContent"` // 模板邮件内容
	SerialNo     string `protobuf:"bytes,8,opt,name=SerialNo,proto3" json:"SerialNo"`         // 流水号 [用于微服务全链路日志追踪]
}

func (x *AddMailConfigRequest) Reset() {
	*x = AddMailConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMailConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMailConfigRequest) ProtoMessage() {}

func (x *AddMailConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMailConfigRequest.ProtoReflect.Descriptor instead.
func (*AddMailConfigRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{7}
}

func (x *AddMailConfigRequest) GetEmailName() string {
	if x != nil {
		return x.EmailName
	}
	return ""
}

func (x *AddMailConfigRequest) GetEmailTitle() string {
	if x != nil {
		return x.EmailTitle
	}
	return ""
}

func (x *AddMailConfigRequest) GetEmailContent() string {
	if x != nil {
		return x.EmailContent
	}
	return ""
}

func (x *AddMailConfigRequest) GetSerialNo() string {
	if x != nil {
		return x.SerialNo
	}
	return ""
}

type AddMailConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code"` // 错误码
	Msg  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`    // 错误信息
	Data []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`  // 数组 返回数据
}

func (x *AddMailConfigResponse) Reset() {
	*x = AddMailConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMailConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMailConfigResponse) ProtoMessage() {}

func (x *AddMailConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMailConfigResponse.ProtoReflect.Descriptor instead.
func (*AddMailConfigResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{8}
}

func (x *AddMailConfigResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddMailConfigResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AddMailConfigResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

// 更新模板内容
type UpdateMailConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EId          int64  `protobuf:"varint,2,opt,name=eId,proto3" json:"eId"`                  // 邮件模板表ID
	EmailName    string `protobuf:"bytes,3,opt,name=emailName,proto3" json:"emailName"`       // 邮件模板名
	EmailTitle   string `protobuf:"bytes,4,opt,name=emailTitle,proto3" json:"emailTitle"`     // 模板邮件标题
	EmailContent string `protobuf:"bytes,5,opt,name=emailContent,proto3" json:"emailContent"` // 模板邮件内容
	SerialNo     string `protobuf:"bytes,8,opt,name=SerialNo,proto3" json:"SerialNo"`         // 流水号 [用于微服务全链路日志追踪]
}

func (x *UpdateMailConfigRequest) Reset() {
	*x = UpdateMailConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMailConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMailConfigRequest) ProtoMessage() {}

func (x *UpdateMailConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMailConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateMailConfigRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateMailConfigRequest) GetEId() int64 {
	if x != nil {
		return x.EId
	}
	return 0
}

func (x *UpdateMailConfigRequest) GetEmailName() string {
	if x != nil {
		return x.EmailName
	}
	return ""
}

func (x *UpdateMailConfigRequest) GetEmailTitle() string {
	if x != nil {
		return x.EmailTitle
	}
	return ""
}

func (x *UpdateMailConfigRequest) GetEmailContent() string {
	if x != nil {
		return x.EmailContent
	}
	return ""
}

func (x *UpdateMailConfigRequest) GetSerialNo() string {
	if x != nil {
		return x.SerialNo
	}
	return ""
}

type UpdateMailConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code"` // 错误码
	Msg  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`    // 错误信息
	Data []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`  // 数组 返回数据
}

func (x *UpdateMailConfigResponse) Reset() {
	*x = UpdateMailConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMailConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMailConfigResponse) ProtoMessage() {}

func (x *UpdateMailConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMailConfigResponse.ProtoReflect.Descriptor instead.
func (*UpdateMailConfigResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateMailConfigResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateMailConfigResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpdateMailConfigResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_mail_mail_proto protoreflect.FileDescriptor

var file_mail_mail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x44, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x71, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x22, 0x4c, 0x0a, 0x10, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb6, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x6f, 0x22, 0x66, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x87, 0x01, 0x0a, 0x0e, 0x4d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x11, 0x0a, 0x04, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x08, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x5f, 0x64, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x41, 0x64, 0x64, 0x44, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x5f, 0x64, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x44, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x41,
	0x64, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x6f, 0x22, 0x51, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xa9, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x6f,
	0x22, 0x54, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x94, 0x02, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x31, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x15, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mail_mail_proto_rawDescOnce sync.Once
	file_mail_mail_proto_rawDescData = file_mail_mail_proto_rawDesc
)

func file_mail_mail_proto_rawDescGZIP() []byte {
	file_mail_mail_proto_rawDescOnce.Do(func() {
		file_mail_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mail_mail_proto_rawDescData)
	})
	return file_mail_mail_proto_rawDescData
}

var file_mail_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_mail_mail_proto_goTypes = []interface{}{
	(*Response)(nil),                  // 0: Response
	(*SendMailRequest)(nil),           // 1: SendMailRequest
	(*SendMailResponse)(nil),          // 2: SendMailResponse
	(*GetMailConfigListRequest)(nil),  // 3: GetMailConfigListRequest
	(*GetMailConfigListResponse)(nil), // 4: GetMailConfigListResponse
	(*MailConfigList)(nil),            // 5: MailConfigList
	(*MailConfigSingle)(nil),          // 6: MailConfigSingle
	(*AddMailConfigRequest)(nil),      // 7: AddMailConfigRequest
	(*AddMailConfigResponse)(nil),     // 8: AddMailConfigResponse
	(*UpdateMailConfigRequest)(nil),   // 9: UpdateMailConfigRequest
	(*UpdateMailConfigResponse)(nil),  // 10: UpdateMailConfigResponse
}
var file_mail_mail_proto_depIdxs = []int32{
	5,  // 0: GetMailConfigListResponse.data:type_name -> MailConfigList
	6,  // 1: MailConfigList.list:type_name -> MailConfigSingle
	1,  // 2: Mail.SendMail:input_type -> SendMailRequest
	3,  // 3: Mail.GetMailConfigList:input_type -> GetMailConfigListRequest
	7,  // 4: Mail.AddMailConfig:input_type -> AddMailConfigRequest
	9,  // 5: Mail.UpdateMailConfig:input_type -> UpdateMailConfigRequest
	2,  // 6: Mail.SendMail:output_type -> SendMailResponse
	4,  // 7: Mail.GetMailConfigList:output_type -> GetMailConfigListResponse
	8,  // 8: Mail.AddMailConfig:output_type -> AddMailConfigResponse
	10, // 9: Mail.UpdateMailConfig:output_type -> UpdateMailConfigResponse
	6,  // [6:10] is the sub-list for method output_type
	2,  // [2:6] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_mail_mail_proto_init() }
func file_mail_mail_proto_init() {
	if File_mail_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mail_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_mail_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailRequest); i {
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
		file_mail_mail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailResponse); i {
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
		file_mail_mail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMailConfigListRequest); i {
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
		file_mail_mail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMailConfigListResponse); i {
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
		file_mail_mail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailConfigList); i {
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
		file_mail_mail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailConfigSingle); i {
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
		file_mail_mail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMailConfigRequest); i {
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
		file_mail_mail_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMailConfigResponse); i {
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
		file_mail_mail_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMailConfigRequest); i {
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
		file_mail_mail_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMailConfigResponse); i {
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
			RawDescriptor: file_mail_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mail_mail_proto_goTypes,
		DependencyIndexes: file_mail_mail_proto_depIdxs,
		MessageInfos:      file_mail_mail_proto_msgTypes,
	}.Build()
	File_mail_mail_proto = out.File
	file_mail_mail_proto_rawDesc = nil
	file_mail_mail_proto_goTypes = nil
	file_mail_mail_proto_depIdxs = nil
}
