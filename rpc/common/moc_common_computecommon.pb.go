// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_common_computecommon.proto

package common

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

type UserType int32

const (
	UserType_ROOT UserType = 0
	UserType_USER UserType = 1
)

var UserType_name = map[int32]string{
	0: "ROOT",
	1: "USER",
}

var UserType_value = map[string]int32{
	"ROOT": 0,
	"USER": 1,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}

func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{0}
}

type ProcessorType int32

const (
	ProcessorType_None    ProcessorType = 0
	ProcessorType_Intel   ProcessorType = 1
	ProcessorType_Intel64 ProcessorType = 2
	ProcessorType_AMD     ProcessorType = 3
	ProcessorType_AMD64   ProcessorType = 4
	ProcessorType_ARM     ProcessorType = 5
	ProcessorType_ARM64   ProcessorType = 6
)

var ProcessorType_name = map[int32]string{
	0: "None",
	1: "Intel",
	2: "Intel64",
	3: "AMD",
	4: "AMD64",
	5: "ARM",
	6: "ARM64",
}

var ProcessorType_value = map[string]int32{
	"None":    0,
	"Intel":   1,
	"Intel64": 2,
	"AMD":     3,
	"AMD64":   4,
	"ARM":     5,
	"ARM64":   6,
}

func (x ProcessorType) String() string {
	return proto.EnumName(ProcessorType_name, int32(x))
}

func (ProcessorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{1}
}

type OperatingSystemBootstrapEngine int32

const (
	OperatingSystemBootstrapEngine_CLOUD_INIT           OperatingSystemBootstrapEngine = 0
	OperatingSystemBootstrapEngine_WINDOWS_ANSWER_FILES OperatingSystemBootstrapEngine = 1
)

var OperatingSystemBootstrapEngine_name = map[int32]string{
	0: "CLOUD_INIT",
	1: "WINDOWS_ANSWER_FILES",
}

var OperatingSystemBootstrapEngine_value = map[string]int32{
	"CLOUD_INIT":           0,
	"WINDOWS_ANSWER_FILES": 1,
}

func (x OperatingSystemBootstrapEngine) String() string {
	return proto.EnumName(OperatingSystemBootstrapEngine_name, int32(x))
}

func (OperatingSystemBootstrapEngine) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{2}
}

type OperatingSystemType int32

const (
	OperatingSystemType_WINDOWS OperatingSystemType = 0
	OperatingSystemType_LINUX   OperatingSystemType = 1
)

var OperatingSystemType_name = map[int32]string{
	0: "WINDOWS",
	1: "LINUX",
}

var OperatingSystemType_value = map[string]int32{
	"WINDOWS": 0,
	"LINUX":   1,
}

func (x OperatingSystemType) String() string {
	return proto.EnumName(OperatingSystemType_name, int32(x))
}

func (OperatingSystemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{3}
}

type VirtualMachineSizeType int32

const (
	VirtualMachineSizeType_Default          VirtualMachineSizeType = 0
	VirtualMachineSizeType_Standard_A2_v2   VirtualMachineSizeType = 2
	VirtualMachineSizeType_Standard_A4_v2   VirtualMachineSizeType = 3
	VirtualMachineSizeType_Standard_D2s_v3  VirtualMachineSizeType = 4
	VirtualMachineSizeType_Standard_D4s_v3  VirtualMachineSizeType = 5
	VirtualMachineSizeType_Standard_D8s_v3  VirtualMachineSizeType = 6
	VirtualMachineSizeType_Standard_D16s_v3 VirtualMachineSizeType = 7
	VirtualMachineSizeType_Standard_D32s_v3 VirtualMachineSizeType = 8
	VirtualMachineSizeType_Standard_DS2_v2  VirtualMachineSizeType = 9
	VirtualMachineSizeType_Standard_DS3_v2  VirtualMachineSizeType = 10
	VirtualMachineSizeType_Standard_DS4_v2  VirtualMachineSizeType = 11
	VirtualMachineSizeType_Standard_DS5_v2  VirtualMachineSizeType = 12
	VirtualMachineSizeType_Standard_DS13_v2 VirtualMachineSizeType = 13
	VirtualMachineSizeType_Standard_K8S_v1  VirtualMachineSizeType = 14
	VirtualMachineSizeType_Standard_K8S2_v1 VirtualMachineSizeType = 15
	VirtualMachineSizeType_Standard_K8S3_v1 VirtualMachineSizeType = 16
	VirtualMachineSizeType_Standard_K8S4_v1 VirtualMachineSizeType = 17
	VirtualMachineSizeType_Standard_NK6     VirtualMachineSizeType = 18
	VirtualMachineSizeType_Standard_NK12    VirtualMachineSizeType = 19
	VirtualMachineSizeType_Standard_NV6     VirtualMachineSizeType = 20
	VirtualMachineSizeType_Standard_NV12    VirtualMachineSizeType = 21
	VirtualMachineSizeType_Standard_K8S5_v1 VirtualMachineSizeType = 22
	VirtualMachineSizeType_Custom           VirtualMachineSizeType = 98
	VirtualMachineSizeType_Unsupported      VirtualMachineSizeType = 99
)

var VirtualMachineSizeType_name = map[int32]string{
	0:  "Default",
	2:  "Standard_A2_v2",
	3:  "Standard_A4_v2",
	4:  "Standard_D2s_v3",
	5:  "Standard_D4s_v3",
	6:  "Standard_D8s_v3",
	7:  "Standard_D16s_v3",
	8:  "Standard_D32s_v3",
	9:  "Standard_DS2_v2",
	10: "Standard_DS3_v2",
	11: "Standard_DS4_v2",
	12: "Standard_DS5_v2",
	13: "Standard_DS13_v2",
	14: "Standard_K8S_v1",
	15: "Standard_K8S2_v1",
	16: "Standard_K8S3_v1",
	17: "Standard_K8S4_v1",
	18: "Standard_NK6",
	19: "Standard_NK12",
	20: "Standard_NV6",
	21: "Standard_NV12",
	22: "Standard_K8S5_v1",
	98: "Custom",
	99: "Unsupported",
}

var VirtualMachineSizeType_value = map[string]int32{
	"Default":          0,
	"Standard_A2_v2":   2,
	"Standard_A4_v2":   3,
	"Standard_D2s_v3":  4,
	"Standard_D4s_v3":  5,
	"Standard_D8s_v3":  6,
	"Standard_D16s_v3": 7,
	"Standard_D32s_v3": 8,
	"Standard_DS2_v2":  9,
	"Standard_DS3_v2":  10,
	"Standard_DS4_v2":  11,
	"Standard_DS5_v2":  12,
	"Standard_DS13_v2": 13,
	"Standard_K8S_v1":  14,
	"Standard_K8S2_v1": 15,
	"Standard_K8S3_v1": 16,
	"Standard_K8S4_v1": 17,
	"Standard_NK6":     18,
	"Standard_NK12":    19,
	"Standard_NV6":     20,
	"Standard_NV12":    21,
	"Standard_K8S5_v1": 22,
	"Custom":           98,
	"Unsupported":      99,
}

func (x VirtualMachineSizeType) String() string {
	return proto.EnumName(VirtualMachineSizeType_name, int32(x))
}

func (VirtualMachineSizeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{4}
}

type PowerState int32

const (
	PowerState_Unknown  PowerState = 0
	PowerState_Running  PowerState = 1
	PowerState_Off      PowerState = 2
	PowerState_Paused   PowerState = 3
	PowerState_Critical PowerState = 4
)

var PowerState_name = map[int32]string{
	0: "Unknown",
	1: "Running",
	2: "Off",
	3: "Paused",
	4: "Critical",
}

var PowerState_value = map[string]int32{
	"Unknown":  0,
	"Running":  1,
	"Off":      2,
	"Paused":   3,
	"Critical": 4,
}

func (x PowerState) String() string {
	return proto.EnumName(PowerState_name, int32(x))
}

func (PowerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{5}
}

type VirtualMachineOperation int32

const (
	VirtualMachineOperation_START VirtualMachineOperation = 0
	VirtualMachineOperation_STOP  VirtualMachineOperation = 1
	VirtualMachineOperation_RESET VirtualMachineOperation = 2
)

var VirtualMachineOperation_name = map[int32]string{
	0: "START",
	1: "STOP",
	2: "RESET",
}

var VirtualMachineOperation_value = map[string]int32{
	"START": 0,
	"STOP":  1,
	"RESET": 2,
}

func (x VirtualMachineOperation) String() string {
	return proto.EnumName(VirtualMachineOperation_name, int32(x))
}

func (VirtualMachineOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{6}
}

type VirtualMachineCustomSize struct {
	CpuCount             int32    `protobuf:"varint,1,opt,name=cpuCount,proto3" json:"cpuCount,omitempty"`
	MemoryMB             int32    `protobuf:"varint,2,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineCustomSize) Reset()         { *m = VirtualMachineCustomSize{} }
func (m *VirtualMachineCustomSize) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineCustomSize) ProtoMessage()    {}
func (*VirtualMachineCustomSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{0}
}

func (m *VirtualMachineCustomSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineCustomSize.Unmarshal(m, b)
}
func (m *VirtualMachineCustomSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineCustomSize.Marshal(b, m, deterministic)
}
func (m *VirtualMachineCustomSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineCustomSize.Merge(m, src)
}
func (m *VirtualMachineCustomSize) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineCustomSize.Size(m)
}
func (m *VirtualMachineCustomSize) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineCustomSize.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineCustomSize proto.InternalMessageInfo

func (m *VirtualMachineCustomSize) GetCpuCount() int32 {
	if m != nil {
		return m.CpuCount
	}
	return 0
}

func (m *VirtualMachineCustomSize) GetMemoryMB() int32 {
	if m != nil {
		return m.MemoryMB
	}
	return 0
}

func init() {
	proto.RegisterEnum("moc.UserType", UserType_name, UserType_value)
	proto.RegisterEnum("moc.ProcessorType", ProcessorType_name, ProcessorType_value)
	proto.RegisterEnum("moc.OperatingSystemBootstrapEngine", OperatingSystemBootstrapEngine_name, OperatingSystemBootstrapEngine_value)
	proto.RegisterEnum("moc.OperatingSystemType", OperatingSystemType_name, OperatingSystemType_value)
	proto.RegisterEnum("moc.VirtualMachineSizeType", VirtualMachineSizeType_name, VirtualMachineSizeType_value)
	proto.RegisterEnum("moc.PowerState", PowerState_name, PowerState_value)
	proto.RegisterEnum("moc.VirtualMachineOperation", VirtualMachineOperation_name, VirtualMachineOperation_value)
	proto.RegisterType((*VirtualMachineCustomSize)(nil), "moc.VirtualMachineCustomSize")
}

func init() { proto.RegisterFile("moc_common_computecommon.proto", fileDescriptor_7d1a061f6c82445b) }

var fileDescriptor_7d1a061f6c82445b = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xcb, 0x6f, 0xdb, 0x38,
	0x10, 0xc6, 0x2d, 0xbf, 0x33, 0xce, 0x63, 0xc2, 0x64, 0xb3, 0xc6, 0x1e, 0x82, 0xc5, 0x2e, 0x0a,
	0x14, 0x02, 0x1a, 0xc3, 0x8f, 0x18, 0xe9, 0xd1, 0xb1, 0x5d, 0x40, 0x4d, 0x6c, 0x19, 0xa2, 0x9d,
	0xb4, 0xbd, 0x08, 0x8a, 0x4c, 0x3b, 0x42, 0x2d, 0x52, 0xa0, 0xa8, 0x04, 0xe9, 0x1f, 0xde, 0x73,
	0x41, 0x39, 0x71, 0x2b, 0x9d, 0xc4, 0xf9, 0xf1, 0xe3, 0xc7, 0x0f, 0xd4, 0x0c, 0x9c, 0x87, 0xc2,
	0x77, 0x7d, 0x11, 0x86, 0x82, 0xeb, 0x4f, 0x94, 0x28, 0xb6, 0xad, 0x2e, 0x22, 0x29, 0x94, 0x20,
	0xa5, 0x50, 0xf8, 0xff, 0x39, 0xd0, 0xbc, 0x0b, 0xa4, 0x4a, 0xbc, 0xcd, 0xc4, 0xf3, 0x1f, 0x03,
	0xce, 0x86, 0x49, 0xac, 0x44, 0x48, 0x83, 0x1f, 0x8c, 0xfc, 0x03, 0x75, 0x3f, 0x4a, 0x86, 0x22,
	0xe1, 0xaa, 0x69, 0xfc, 0x6b, 0xbc, 0xaf, 0x38, 0xbb, 0x5a, 0xef, 0x85, 0x2c, 0x14, 0xf2, 0x65,
	0x72, 0xdd, 0x2c, 0x6e, 0xf7, 0xde, 0x6a, 0xf3, 0x1c, 0xea, 0x8b, 0x98, 0xc9, 0xf9, 0x4b, 0xc4,
	0x48, 0x1d, 0xca, 0x8e, 0x6d, 0xcf, 0xb1, 0xa0, 0x57, 0x0b, 0x3a, 0x76, 0xd0, 0x30, 0xbf, 0xc2,
	0xc1, 0x4c, 0x0a, 0x9f, 0xc5, 0xb1, 0xd8, 0x89, 0xa6, 0x82, 0x33, 0x2c, 0x90, 0x3d, 0xa8, 0x58,
	0x5c, 0xb1, 0x0d, 0x1a, 0xa4, 0x01, 0xb5, 0x74, 0xd9, 0xef, 0x61, 0x91, 0xd4, 0xa0, 0x34, 0x98,
	0x8c, 0xb0, 0xa4, 0x05, 0x83, 0xc9, 0xa8, 0xdf, 0xc3, 0x72, 0xca, 0x9c, 0x09, 0x56, 0x52, 0xe6,
	0x4c, 0xfa, 0x3d, 0xac, 0x9a, 0x9f, 0xe1, 0xdc, 0x8e, 0x98, 0xf4, 0x54, 0xc0, 0xd7, 0xf4, 0x25,
	0x56, 0x2c, 0xbc, 0x16, 0x42, 0xc5, 0x4a, 0x7a, 0xd1, 0x98, 0xaf, 0x03, 0xce, 0xc8, 0x21, 0xc0,
	0xf0, 0xd6, 0x5e, 0x8c, 0x5c, 0x6b, 0x6a, 0xe9, 0x58, 0x4d, 0x38, 0xbd, 0xb7, 0xa6, 0x23, 0xfb,
	0x9e, 0xba, 0x83, 0x29, 0xbd, 0x1f, 0x3b, 0xee, 0x27, 0xeb, 0x76, 0x4c, 0xd1, 0x30, 0x3f, 0xc0,
	0x49, 0xce, 0x2b, 0x0d, 0xdb, 0x80, 0xda, 0xeb, 0x81, 0x6d, 0xde, 0x5b, 0x6b, 0xba, 0xf8, 0x82,
	0x86, 0xf9, 0xb3, 0x04, 0x67, 0xd9, 0xa7, 0xd4, 0x8f, 0xf8, 0x76, 0x64, 0xc4, 0x56, 0x5e, 0xb2,
	0x51, 0x58, 0x20, 0x04, 0x0e, 0xa9, 0xf2, 0xf8, 0xd2, 0x93, 0x4b, 0x77, 0xd0, 0x71, 0x9f, 0x3a,
	0x58, 0xcc, 0xb2, 0x9e, 0x66, 0x25, 0x72, 0x02, 0x47, 0x3b, 0x36, 0xea, 0xc4, 0xee, 0x53, 0x17,
	0xcb, 0x59, 0xd8, 0x4b, 0x61, 0x25, 0x0b, 0xaf, 0x52, 0x58, 0x25, 0xa7, 0x80, 0xbf, 0x61, 0xbb,
	0x9f, 0xd2, 0x5a, 0x96, 0x76, 0xb7, 0xae, 0xf5, 0xac, 0x01, 0x4d, 0x33, 0xed, 0xe5, 0x60, 0x57,
	0x43, 0xc8, 0xc1, 0x34, 0x69, 0x23, 0x07, 0x2f, 0x35, 0xdc, 0xcf, 0xde, 0x44, 0xdb, 0xe9, 0xf9,
	0x83, 0x8c, 0xf4, 0xe6, 0x8a, 0xba, 0x4f, 0x6d, 0x3c, 0xcc, 0x48, 0x6f, 0xae, 0xf4, 0xfd, 0x6d,
	0x3c, 0xca, 0xd3, 0xae, 0xa6, 0x98, 0xa7, 0x3d, 0x4d, 0x8f, 0x09, 0xc2, 0xfe, 0x8e, 0x4e, 0x6f,
	0xfa, 0x48, 0xc8, 0x31, 0x1c, 0xfc, 0x41, 0xda, 0x1d, 0x3c, 0xc9, 0x8a, 0xee, 0xfa, 0x78, 0x9a,
	0x15, 0xdd, 0xb5, 0x3b, 0xf8, 0x57, 0xde, 0xff, 0x52, 0xfb, 0x9f, 0x11, 0x80, 0xea, 0x76, 0x2e,
	0xf0, 0x81, 0x1c, 0x41, 0x63, 0xc1, 0xe3, 0x24, 0x8a, 0x84, 0x54, 0x6c, 0x89, 0xbe, 0x69, 0x01,
	0xcc, 0xc4, 0x33, 0x93, 0x54, 0x79, 0x2a, 0xfd, 0xd7, 0x0b, 0xfe, 0x9d, 0x8b, 0x67, 0x8e, 0x05,
	0x5d, 0x38, 0x09, 0xe7, 0x01, 0x5f, 0xa3, 0xa1, 0xfb, 0xd5, 0x5e, 0xad, 0xb0, 0xa8, 0xdd, 0x66,
	0x5e, 0x12, 0xb3, 0x25, 0x96, 0xc8, 0x3e, 0xd4, 0x87, 0x32, 0x50, 0x81, 0xef, 0x6d, 0xb0, 0x6c,
	0x7e, 0x84, 0xbf, 0xb3, 0x2d, 0xf4, 0xda, 0x80, 0x82, 0xeb, 0x4e, 0xa3, 0xf3, 0x81, 0xf3, 0x3a,
	0x49, 0x74, 0x6e, 0xcf, 0xd0, 0xd0, 0xd0, 0x19, 0xd3, 0xf1, 0x1c, 0x8b, 0xd7, 0xef, 0xbe, 0xfd,
	0xbf, 0x0e, 0xd4, 0x63, 0xf2, 0x70, 0xe1, 0x8b, 0xb0, 0x15, 0x06, 0xbe, 0x14, 0xb1, 0x58, 0xa9,
	0x56, 0x28, 0xfc, 0x96, 0x8c, 0xfc, 0xd6, 0x76, 0xf4, 0x1f, 0xaa, 0xe9, 0xec, 0x77, 0x7f, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0x66, 0xd4, 0x3c, 0x1d, 0x04, 0x00, 0x00,
}
