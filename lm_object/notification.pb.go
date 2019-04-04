// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package lm_object

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Notification_NotificationType int32

const (
	Notification_ACTIVATION             Notification_NotificationType = 0
	Notification_PASSWORD_RESET         Notification_NotificationType = 1
	Notification_CANDIDATE_SELECTED     Notification_NotificationType = 2
	Notification_CANDIDATE_NOT_SELECTED Notification_NotificationType = 3
	Notification_CANDIDATE_DESELECTED   Notification_NotificationType = 4
	Notification_NEW_MISSION            Notification_NotificationType = 5
	Notification_NEW_CONTRACT_TO_SIGN   Notification_NotificationType = 6
	Notification_DEFINE_PASSWORD        Notification_NotificationType = 7
	Notification_NEW_CANDIDATE          Notification_NotificationType = 8
	Notification_MISSION_VALIDATED      Notification_NotificationType = 9
	Notification_MISSION_REFUSED        Notification_NotificationType = 10
	Notification_OVERTIME               Notification_NotificationType = 11
	Notification_MISSION_TO_VALIDATE    Notification_NotificationType = 12
	Notification_MISSION_DELETED        Notification_NotificationType = 13
	Notification_NEW_CONTRACT_TO_EDIT   Notification_NotificationType = 14
	Notification_NEW_INSTITUTION        Notification_NotificationType = 15
	Notification_CALL_REQUEST           Notification_NotificationType = 16
	Notification_INVITE                 Notification_NotificationType = 17
	Notification_DPAE_AVAILABLE         Notification_NotificationType = 18
	Notification_CANDIDATE_UNAPPLY      Notification_NotificationType = 19
	Notification_REMIND                 Notification_NotificationType = 20
	Notification_NEW_CONTRACT_SIGNED    Notification_NotificationType = 21
)

var Notification_NotificationType_name = map[int32]string{
	0:  "ACTIVATION",
	1:  "PASSWORD_RESET",
	2:  "CANDIDATE_SELECTED",
	3:  "CANDIDATE_NOT_SELECTED",
	4:  "CANDIDATE_DESELECTED",
	5:  "NEW_MISSION",
	6:  "NEW_CONTRACT_TO_SIGN",
	7:  "DEFINE_PASSWORD",
	8:  "NEW_CANDIDATE",
	9:  "MISSION_VALIDATED",
	10: "MISSION_REFUSED",
	11: "OVERTIME",
	12: "MISSION_TO_VALIDATE",
	13: "MISSION_DELETED",
	14: "NEW_CONTRACT_TO_EDIT",
	15: "NEW_INSTITUTION",
	16: "CALL_REQUEST",
	17: "INVITE",
	18: "DPAE_AVAILABLE",
	19: "CANDIDATE_UNAPPLY",
	20: "REMIND",
	21: "NEW_CONTRACT_SIGNED",
}

var Notification_NotificationType_value = map[string]int32{
	"ACTIVATION":             0,
	"PASSWORD_RESET":         1,
	"CANDIDATE_SELECTED":     2,
	"CANDIDATE_NOT_SELECTED": 3,
	"CANDIDATE_DESELECTED":   4,
	"NEW_MISSION":            5,
	"NEW_CONTRACT_TO_SIGN":   6,
	"DEFINE_PASSWORD":        7,
	"NEW_CANDIDATE":          8,
	"MISSION_VALIDATED":      9,
	"MISSION_REFUSED":        10,
	"OVERTIME":               11,
	"MISSION_TO_VALIDATE":    12,
	"MISSION_DELETED":        13,
	"NEW_CONTRACT_TO_EDIT":   14,
	"NEW_INSTITUTION":        15,
	"CALL_REQUEST":           16,
	"INVITE":                 17,
	"DPAE_AVAILABLE":         18,
	"CANDIDATE_UNAPPLY":      19,
	"REMIND":                 20,
	"NEW_CONTRACT_SIGNED":    21,
}

func (x Notification_NotificationType) String() string {
	return proto.EnumName(Notification_NotificationType_name, int32(x))
}

func (Notification_NotificationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0, 0}
}

type Notification struct {
	NotificationType Notification_NotificationType `protobuf:"varint,1,opt,name=notification_type,json=notificationType,proto3,enum=lm.object.Notification_NotificationType" json:"notification_type,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*Notification_Mission
	//	*Notification_SimpleContent
	//	*Notification_NewInstitution
	//	*Notification_CallRequest
	//	*Notification_Invitation
	//	*Notification_Missions
	Content              isNotification_Content `protobuf_oneof:"content"`
	Recipients           []*People              `protobuf:"bytes,2,rep,name=recipients,proto3" json:"recipients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetNotificationType() Notification_NotificationType {
	if m != nil {
		return m.NotificationType
	}
	return Notification_ACTIVATION
}

type isNotification_Content interface {
	isNotification_Content()
}

type Notification_Mission struct {
	Mission *Mission `protobuf:"bytes,3,opt,name=mission,proto3,oneof"`
}

type Notification_SimpleContent struct {
	SimpleContent *SimpleContent `protobuf:"bytes,4,opt,name=simple_content,json=simpleContent,proto3,oneof"`
}

type Notification_NewInstitution struct {
	NewInstitution *NewInstitution `protobuf:"bytes,5,opt,name=new_institution,json=newInstitution,proto3,oneof"`
}

type Notification_CallRequest struct {
	CallRequest *CallRequest `protobuf:"bytes,6,opt,name=call_request,json=callRequest,proto3,oneof"`
}

type Notification_Invitation struct {
	Invitation *Invitation `protobuf:"bytes,7,opt,name=invitation,proto3,oneof"`
}

type Notification_Missions struct {
	Missions *Missions `protobuf:"bytes,8,opt,name=missions,proto3,oneof"`
}

func (*Notification_Mission) isNotification_Content() {}

func (*Notification_SimpleContent) isNotification_Content() {}

func (*Notification_NewInstitution) isNotification_Content() {}

func (*Notification_CallRequest) isNotification_Content() {}

func (*Notification_Invitation) isNotification_Content() {}

func (*Notification_Missions) isNotification_Content() {}

func (m *Notification) GetContent() isNotification_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Notification) GetMission() *Mission {
	if x, ok := m.GetContent().(*Notification_Mission); ok {
		return x.Mission
	}
	return nil
}

func (m *Notification) GetSimpleContent() *SimpleContent {
	if x, ok := m.GetContent().(*Notification_SimpleContent); ok {
		return x.SimpleContent
	}
	return nil
}

func (m *Notification) GetNewInstitution() *NewInstitution {
	if x, ok := m.GetContent().(*Notification_NewInstitution); ok {
		return x.NewInstitution
	}
	return nil
}

func (m *Notification) GetCallRequest() *CallRequest {
	if x, ok := m.GetContent().(*Notification_CallRequest); ok {
		return x.CallRequest
	}
	return nil
}

func (m *Notification) GetInvitation() *Invitation {
	if x, ok := m.GetContent().(*Notification_Invitation); ok {
		return x.Invitation
	}
	return nil
}

func (m *Notification) GetMissions() *Missions {
	if x, ok := m.GetContent().(*Notification_Missions); ok {
		return x.Missions
	}
	return nil
}

func (m *Notification) GetRecipients() []*People {
	if m != nil {
		return m.Recipients
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Notification) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Notification_Mission)(nil),
		(*Notification_SimpleContent)(nil),
		(*Notification_NewInstitution)(nil),
		(*Notification_CallRequest)(nil),
		(*Notification_Invitation)(nil),
		(*Notification_Missions)(nil),
	}
}

type Missions struct {
	Missions             []*Mission `protobuf:"bytes,1,rep,name=missions,proto3" json:"missions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Missions) Reset()         { *m = Missions{} }
func (m *Missions) String() string { return proto.CompactTextString(m) }
func (*Missions) ProtoMessage()    {}
func (*Missions) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *Missions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Missions.Unmarshal(m, b)
}
func (m *Missions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Missions.Marshal(b, m, deterministic)
}
func (m *Missions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Missions.Merge(m, src)
}
func (m *Missions) XXX_Size() int {
	return xxx_messageInfo_Missions.Size(m)
}
func (m *Missions) XXX_DiscardUnknown() {
	xxx_messageInfo_Missions.DiscardUnknown(m)
}

var xxx_messageInfo_Missions proto.InternalMessageInfo

func (m *Missions) GetMissions() []*Mission {
	if m != nil {
		return m.Missions
	}
	return nil
}

type SimpleContent struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Link                 string   `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleContent) Reset()         { *m = SimpleContent{} }
func (m *SimpleContent) String() string { return proto.CompactTextString(m) }
func (*SimpleContent) ProtoMessage()    {}
func (*SimpleContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{2}
}

func (m *SimpleContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleContent.Unmarshal(m, b)
}
func (m *SimpleContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleContent.Marshal(b, m, deterministic)
}
func (m *SimpleContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleContent.Merge(m, src)
}
func (m *SimpleContent) XXX_Size() int {
	return xxx_messageInfo_SimpleContent.Size(m)
}
func (m *SimpleContent) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleContent.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleContent proto.InternalMessageInfo

func (m *SimpleContent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SimpleContent) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type NewInstitution struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	InstitutionName      string   `protobuf:"bytes,4,opt,name=institution_name,json=institutionName,proto3" json:"institution_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewInstitution) Reset()         { *m = NewInstitution{} }
func (m *NewInstitution) String() string { return proto.CompactTextString(m) }
func (*NewInstitution) ProtoMessage()    {}
func (*NewInstitution) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{3}
}

func (m *NewInstitution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewInstitution.Unmarshal(m, b)
}
func (m *NewInstitution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewInstitution.Marshal(b, m, deterministic)
}
func (m *NewInstitution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewInstitution.Merge(m, src)
}
func (m *NewInstitution) XXX_Size() int {
	return xxx_messageInfo_NewInstitution.Size(m)
}
func (m *NewInstitution) XXX_DiscardUnknown() {
	xxx_messageInfo_NewInstitution.DiscardUnknown(m)
}

var xxx_messageInfo_NewInstitution proto.InternalMessageInfo

func (m *NewInstitution) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *NewInstitution) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *NewInstitution) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewInstitution) GetInstitutionName() string {
	if m != nil {
		return m.InstitutionName
	}
	return ""
}

type CallRequest struct {
	LastName             string               `protobuf:"bytes,1,opt,name=lastName,proto3" json:"lastName,omitempty"`
	FirstName            string               `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	Institution          string               `protobuf:"bytes,3,opt,name=institution,proto3" json:"institution,omitempty"`
	Poste                string               `protobuf:"bytes,4,opt,name=poste,proto3" json:"poste,omitempty"`
	Email                string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string               `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	ContactDateTime      *timestamp.Timestamp `protobuf:"bytes,7,opt,name=contactDateTime,proto3" json:"contactDateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{4}
}

func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (m *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(m, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CallRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CallRequest) GetInstitution() string {
	if m != nil {
		return m.Institution
	}
	return ""
}

func (m *CallRequest) GetPoste() string {
	if m != nil {
		return m.Poste
	}
	return ""
}

func (m *CallRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CallRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CallRequest) GetContactDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.ContactDateTime
	}
	return nil
}

type Invitation struct {
	Institution           string   `protobuf:"bytes,1,opt,name=institution,proto3" json:"institution,omitempty"`
	RegistrationBacklink  string   `protobuf:"bytes,2,opt,name=registrationBacklink,proto3" json:"registrationBacklink,omitempty"`
	DocumentationBacklink string   `protobuf:"bytes,3,opt,name=documentationBacklink,proto3" json:"documentationBacklink,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Invitation) Reset()         { *m = Invitation{} }
func (m *Invitation) String() string { return proto.CompactTextString(m) }
func (*Invitation) ProtoMessage()    {}
func (*Invitation) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{5}
}

func (m *Invitation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invitation.Unmarshal(m, b)
}
func (m *Invitation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invitation.Marshal(b, m, deterministic)
}
func (m *Invitation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invitation.Merge(m, src)
}
func (m *Invitation) XXX_Size() int {
	return xxx_messageInfo_Invitation.Size(m)
}
func (m *Invitation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invitation.DiscardUnknown(m)
}

var xxx_messageInfo_Invitation proto.InternalMessageInfo

func (m *Invitation) GetInstitution() string {
	if m != nil {
		return m.Institution
	}
	return ""
}

func (m *Invitation) GetRegistrationBacklink() string {
	if m != nil {
		return m.RegistrationBacklink
	}
	return ""
}

func (m *Invitation) GetDocumentationBacklink() string {
	if m != nil {
		return m.DocumentationBacklink
	}
	return ""
}

func init() {
	proto.RegisterEnum("lm.object.Notification_NotificationType", Notification_NotificationType_name, Notification_NotificationType_value)
	proto.RegisterType((*Notification)(nil), "lm.object.Notification")
	proto.RegisterType((*Missions)(nil), "lm.object.Missions")
	proto.RegisterType((*SimpleContent)(nil), "lm.object.SimpleContent")
	proto.RegisterType((*NewInstitution)(nil), "lm.object.NewInstitution")
	proto.RegisterType((*CallRequest)(nil), "lm.object.CallRequest")
	proto.RegisterType((*Invitation)(nil), "lm.object.Invitation")
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07) }

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 856 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcd, 0x72, 0xe2, 0x46,
	0x10, 0x06, 0xff, 0x01, 0x0d, 0x86, 0xa1, 0xfd, 0x13, 0x85, 0x24, 0x15, 0x17, 0x27, 0xe7, 0xc2,
	0xd6, 0x3a, 0xa9, 0x4a, 0x55, 0x52, 0x39, 0xc8, 0x68, 0x36, 0xa8, 0x0a, 0x0b, 0x32, 0x1a, 0xd8,
	0xca, 0x49, 0x25, 0x93, 0xb1, 0xa3, 0xac, 0x7e, 0x08, 0x1a, 0x67, 0x6b, 0x1f, 0x22, 0xf7, 0x3c,
	0x5c, 0x8e, 0x79, 0x81, 0xbc, 0xc1, 0xd6, 0x8c, 0x10, 0x1a, 0xbc, 0x3e, 0x41, 0xf7, 0xf7, 0x7d,
	0x3d, 0xdf, 0xb4, 0xa6, 0x1b, 0x30, 0xcd, 0x64, 0xf4, 0x10, 0xad, 0x42, 0x19, 0x65, 0xe9, 0x68,
	0xbd, 0xc9, 0x64, 0x86, 0xad, 0x38, 0x19, 0x65, 0xf7, 0x7f, 0x88, 0x95, 0x1c, 0x74, 0x8a, 0xdf,
	0x02, 0x18, 0x7c, 0xfd, 0x98, 0x65, 0x8f, 0xb1, 0x78, 0xa5, 0xa3, 0xfb, 0xa7, 0x87, 0x57, 0x32,
	0x4a, 0x44, 0x2e, 0xc3, 0x64, 0x5d, 0x10, 0x86, 0xff, 0x35, 0xa0, 0xe3, 0x19, 0x05, 0x71, 0x01,
	0x7d, 0xf3, 0x80, 0x40, 0x7e, 0x58, 0x0b, 0xab, 0x7e, 0x55, 0xbf, 0xee, 0xde, 0x5c, 0x8f, 0x76,
	0xc7, 0x8c, 0x4c, 0xcd, 0x5e, 0xc0, 0x3f, 0xac, 0x05, 0x23, 0xe9, 0xb3, 0x0c, 0x8e, 0xa0, 0x91,
	0x44, 0x79, 0x1e, 0x65, 0xa9, 0x75, 0x78, 0x55, 0xbf, 0x6e, 0xdf, 0xa0, 0x51, 0xec, 0xae, 0x40,
	0x26, 0x35, 0x56, 0x92, 0xd0, 0x86, 0x6e, 0x1e, 0x25, 0xeb, 0x58, 0x04, 0xab, 0x2c, 0x95, 0x22,
	0x95, 0xd6, 0x91, 0x96, 0x59, 0x86, 0xcc, 0xd7, 0x84, 0x71, 0x81, 0x4f, 0x6a, 0xec, 0x34, 0x37,
	0x13, 0xe8, 0x40, 0x2f, 0x15, 0xef, 0x83, 0x28, 0xcd, 0x65, 0x24, 0x9f, 0x94, 0x13, 0xeb, 0x58,
	0xd7, 0xf8, 0xdc, 0xbc, 0x87, 0x78, 0xef, 0x56, 0x84, 0x49, 0x8d, 0x75, 0xd3, 0xbd, 0x0c, 0xfe,
	0x08, 0x9d, 0x55, 0x18, 0xc7, 0xc1, 0x46, 0xfc, 0xf9, 0x24, 0x72, 0x69, 0x9d, 0xe8, 0x12, 0x97,
	0x46, 0x89, 0x71, 0x18, 0xc7, 0xac, 0x40, 0x27, 0x35, 0xd6, 0x5e, 0x55, 0x21, 0x7e, 0x0f, 0x10,
	0xa5, 0x7f, 0x45, 0x52, 0xf7, 0xc1, 0x6a, 0x68, 0xe9, 0x85, 0x21, 0x75, 0x77, 0xe0, 0xa4, 0xc6,
	0x0c, 0x2a, 0xbe, 0x86, 0xe6, 0xb6, 0x13, 0xb9, 0xd5, 0xd4, 0xb2, 0xb3, 0x4f, 0xfb, 0x95, 0x4f,
	0x6a, 0x6c, 0x47, 0xc3, 0xd7, 0x00, 0x1b, 0xb1, 0x8a, 0xd6, 0x91, 0x48, 0x65, 0x6e, 0x1d, 0x5c,
	0x1d, 0x5e, 0xb7, 0x6f, 0xfa, 0x86, 0x68, 0x2e, 0xb2, 0x75, 0x2c, 0x98, 0x41, 0x1a, 0xfe, 0x7b,
	0x08, 0xe4, 0xf9, 0xb7, 0xc3, 0x2e, 0x80, 0x3d, 0xe6, 0xee, 0xd2, 0xe6, 0xee, 0xcc, 0x23, 0x35,
	0x44, 0xe8, 0xce, 0x6d, 0xdf, 0x7f, 0x3b, 0x63, 0x4e, 0xc0, 0xa8, 0x4f, 0x39, 0xa9, 0xe3, 0x25,
	0xe0, 0xd8, 0xf6, 0x1c, 0xd7, 0xb1, 0x39, 0x0d, 0x7c, 0x3a, 0xa5, 0x63, 0x4e, 0x1d, 0x72, 0x80,
	0x03, 0xb8, 0xac, 0xf2, 0xde, 0x8c, 0x57, 0xd8, 0x21, 0x5a, 0x70, 0x5e, 0x61, 0x0e, 0xdd, 0x21,
	0x47, 0xd8, 0x83, 0xb6, 0x47, 0xdf, 0x06, 0x77, 0xae, 0xef, 0xab, 0x23, 0x8f, 0x15, 0x55, 0x25,
	0xc6, 0x33, 0x8f, 0x33, 0x7b, 0xcc, 0x03, 0x3e, 0x0b, 0x7c, 0xf7, 0x67, 0x8f, 0x9c, 0xe0, 0x19,
	0xf4, 0x1c, 0xfa, 0xc6, 0xf5, 0x68, 0x50, 0x7a, 0x22, 0x0d, 0xec, 0xc3, 0xa9, 0xa6, 0x97, 0xd5,
	0x49, 0x13, 0x2f, 0xa0, 0xbf, 0x2d, 0x17, 0x2c, 0xed, 0xa9, 0xce, 0x3a, 0xa4, 0xa5, 0xe4, 0x65,
	0x9a, 0xd1, 0x37, 0x0b, 0x9f, 0x3a, 0x04, 0xb0, 0x03, 0xcd, 0xd9, 0x92, 0x32, 0xee, 0xde, 0x51,
	0xd2, 0xc6, 0xcf, 0xe0, 0xac, 0xa4, 0xf0, 0xd9, 0x4e, 0x4c, 0x3a, 0xa6, 0xd6, 0xa1, 0x53, 0xaa,
	0x0a, 0x9e, 0xbe, 0xe4, 0x94, 0x3a, 0x2e, 0x27, 0x5d, 0x45, 0x57, 0x88, 0xeb, 0xf9, 0xdc, 0xe5,
	0x0b, 0xdd, 0xcb, 0x1e, 0x12, 0xe8, 0x8c, 0xed, 0xe9, 0x34, 0x60, 0xf4, 0x97, 0x05, 0xf5, 0x39,
	0x21, 0x08, 0x70, 0xe2, 0x7a, 0x4b, 0x97, 0x53, 0xd2, 0x57, 0x9d, 0x76, 0xe6, 0x36, 0x0d, 0xec,
	0xa5, 0xed, 0x4e, 0xed, 0xdb, 0x29, 0x25, 0xa8, 0x2e, 0x52, 0x75, 0x6d, 0xe1, 0xd9, 0xf3, 0xf9,
	0xf4, 0x57, 0x72, 0xa6, 0x64, 0x8c, 0xde, 0xb9, 0x9e, 0x43, 0xce, 0x95, 0xe3, 0x3d, 0x0f, 0xaa,
	0x55, 0xd4, 0x21, 0x17, 0xb7, 0x2d, 0x68, 0x6c, 0x87, 0x67, 0xf8, 0x03, 0x34, 0xcb, 0x47, 0x83,
	0x23, 0xe3, 0x6d, 0xd5, 0xf5, 0x33, 0x79, 0x61, 0x16, 0xab, 0x87, 0x35, 0xfc, 0x09, 0x4e, 0xf7,
	0x26, 0x0d, 0x2d, 0x68, 0x24, 0x22, 0xcf, 0xc3, 0xc7, 0x62, 0x31, 0xb4, 0x58, 0x19, 0x22, 0xc2,
	0x51, 0x1c, 0xa5, 0xef, 0xac, 0x03, 0x9d, 0xd6, 0xff, 0x87, 0x7f, 0xd7, 0xa1, 0xbb, 0x3f, 0x65,
	0xf8, 0x15, 0xc0, 0x43, 0xb4, 0xc9, 0x65, 0x90, 0x86, 0x49, 0x59, 0xa3, 0xa5, 0x33, 0x5e, 0x98,
	0x08, 0xfc, 0x02, 0x5a, 0x71, 0x58, 0xa2, 0x45, 0xa9, 0xa6, 0x4a, 0x68, 0xf0, 0x1c, 0x8e, 0x45,
	0x12, 0x46, 0xb1, 0x5e, 0x23, 0x2d, 0x56, 0x04, 0xf8, 0x0d, 0x10, 0x63, 0xce, 0x0b, 0xe5, 0x91,
	0x26, 0xf4, 0x8c, 0xbc, 0x2a, 0x30, 0xfc, 0xbf, 0x0e, 0x6d, 0x63, 0x64, 0x71, 0x00, 0xbb, 0xe2,
	0x5b, 0x2b, 0xd5, 0x61, 0x5f, 0x42, 0x65, 0x6b, 0xeb, 0xc4, 0xf0, 0x79, 0x05, 0x6d, 0x73, 0xb9,
	0x14, 0x86, 0xcc, 0x94, 0x32, 0xbb, 0xce, 0x72, 0x59, 0x7a, 0x29, 0x82, 0xea, 0x0a, 0xc7, 0xe6,
	0x15, 0x14, 0xf7, 0xf7, 0x2c, 0x15, 0x7a, 0xc3, 0x28, 0xae, 0x0a, 0xd4, 0x12, 0x53, 0xdf, 0x30,
	0x5c, 0x49, 0x27, 0x94, 0x82, 0x47, 0x89, 0xd8, 0xae, 0x91, 0xc1, 0xa8, 0x58, 0xed, 0xa3, 0x72,
	0xb5, 0x8f, 0x78, 0xb9, 0xda, 0xd9, 0x73, 0xc9, 0xf0, 0x9f, 0x3a, 0x40, 0xb5, 0x6b, 0x9e, 0x1b,
	0xaf, 0x7f, 0x6a, 0xfc, 0x06, 0xce, 0x37, 0xe2, 0x31, 0xca, 0xe5, 0x46, 0x2b, 0x6e, 0xc3, 0xd5,
	0x3b, 0xe3, 0xc3, 0xbe, 0x88, 0xe1, 0x77, 0x70, 0xf1, 0x5b, 0xb6, 0x7a, 0x4a, 0x44, 0x2a, 0xf7,
	0x45, 0x45, 0x63, 0x5e, 0x06, 0xef, 0x4f, 0xb4, 0xff, 0x6f, 0x3f, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xc8, 0x74, 0x6f, 0xd8, 0xd7, 0x06, 0x00, 0x00,
}