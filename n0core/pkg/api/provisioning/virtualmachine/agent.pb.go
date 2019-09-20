// Code generated by protoc-gen-go. DO NOT EDIT.
// source: n0core/pkg/api/provisioning/virtualmachine/agent.proto

package virtualmachine

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VirtualMachineState int32

const (
	VirtualMachineState_FAILED   VirtualMachineState = 0
	VirtualMachineState_UNKNOWN  VirtualMachineState = 1
	VirtualMachineState_SHUTDOWN VirtualMachineState = 2
	VirtualMachineState_RUNNING  VirtualMachineState = 3
	VirtualMachineState_PAUSED   VirtualMachineState = 4
)

var VirtualMachineState_name = map[int32]string{
	0: "FAILED",
	1: "UNKNOWN",
	2: "SHUTDOWN",
	3: "RUNNING",
	4: "PAUSED",
}

var VirtualMachineState_value = map[string]int32{
	"FAILED":   0,
	"UNKNOWN":  1,
	"SHUTDOWN": 2,
	"RUNNING":  3,
	"PAUSED":   4,
}

func (x VirtualMachineState) String() string {
	return proto.EnumName(VirtualMachineState_name, int32(x))
}

func (VirtualMachineState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{0}
}

type BlockDev struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	BootIndex            uint32   `protobuf:"varint,3,opt,name=boot_index,json=bootIndex,proto3" json:"boot_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockDev) Reset()         { *m = BlockDev{} }
func (m *BlockDev) String() string { return proto.CompactTextString(m) }
func (*BlockDev) ProtoMessage()    {}
func (*BlockDev) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{0}
}

func (m *BlockDev) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockDev.Unmarshal(m, b)
}
func (m *BlockDev) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockDev.Marshal(b, m, deterministic)
}
func (m *BlockDev) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDev.Merge(m, src)
}
func (m *BlockDev) XXX_Size() int {
	return xxx_messageInfo_BlockDev.Size(m)
}
func (m *BlockDev) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDev.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDev proto.InternalMessageInfo

func (m *BlockDev) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlockDev) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *BlockDev) GetBootIndex() uint32 {
	if m != nil {
		return m.BootIndex
	}
	return 0
}

type NetDev struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NetworkName          string   `protobuf:"bytes,2,opt,name=network_name,json=networkName,proto3" json:"network_name,omitempty"`
	HardwareAddress      string   `protobuf:"bytes,3,opt,name=hardware_address,json=hardwareAddress,proto3" json:"hardware_address,omitempty"`
	Ipv4AddressCidr      string   `protobuf:"bytes,4,opt,name=ipv4_address_cidr,json=ipv4AddressCidr,proto3" json:"ipv4_address_cidr,omitempty"`
	Ipv4Gateway          string   `protobuf:"bytes,5,opt,name=ipv4_gateway,json=ipv4Gateway,proto3" json:"ipv4_gateway,omitempty"`
	Nameservers          []string `protobuf:"bytes,6,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetDev) Reset()         { *m = NetDev{} }
func (m *NetDev) String() string { return proto.CompactTextString(m) }
func (*NetDev) ProtoMessage()    {}
func (*NetDev) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{1}
}

func (m *NetDev) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetDev.Unmarshal(m, b)
}
func (m *NetDev) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetDev.Marshal(b, m, deterministic)
}
func (m *NetDev) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetDev.Merge(m, src)
}
func (m *NetDev) XXX_Size() int {
	return xxx_messageInfo_NetDev.Size(m)
}
func (m *NetDev) XXX_DiscardUnknown() {
	xxx_messageInfo_NetDev.DiscardUnknown(m)
}

var xxx_messageInfo_NetDev proto.InternalMessageInfo

func (m *NetDev) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetDev) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

func (m *NetDev) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

func (m *NetDev) GetIpv4AddressCidr() string {
	if m != nil {
		return m.Ipv4AddressCidr
	}
	return ""
}

func (m *NetDev) GetIpv4Gateway() string {
	if m != nil {
		return m.Ipv4Gateway
	}
	return ""
}

func (m *NetDev) GetNameservers() []string {
	if m != nil {
		return m.Nameservers
	}
	return nil
}

type BootVirtualMachineRequest struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid                 string      `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Vcpus                uint32      `protobuf:"varint,3,opt,name=vcpus,proto3" json:"vcpus,omitempty"`
	MemoryBytes          uint64      `protobuf:"varint,4,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
	Blockdevs            []*BlockDev `protobuf:"bytes,5,rep,name=blockdevs,proto3" json:"blockdevs,omitempty"`
	Netdevs              []*NetDev   `protobuf:"bytes,6,rep,name=netdevs,proto3" json:"netdevs,omitempty"`
	LoginUsername        string      `protobuf:"bytes,7,opt,name=login_username,json=loginUsername,proto3" json:"login_username,omitempty"`
	SshAuthorizedKeys    []string    `protobuf:"bytes,8,rep,name=ssh_authorized_keys,json=sshAuthorizedKeys,proto3" json:"ssh_authorized_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BootVirtualMachineRequest) Reset()         { *m = BootVirtualMachineRequest{} }
func (m *BootVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*BootVirtualMachineRequest) ProtoMessage()    {}
func (*BootVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{2}
}

func (m *BootVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootVirtualMachineRequest.Unmarshal(m, b)
}
func (m *BootVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (m *BootVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootVirtualMachineRequest.Merge(m, src)
}
func (m *BootVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_BootVirtualMachineRequest.Size(m)
}
func (m *BootVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BootVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BootVirtualMachineRequest proto.InternalMessageInfo

func (m *BootVirtualMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BootVirtualMachineRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BootVirtualMachineRequest) GetVcpus() uint32 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *BootVirtualMachineRequest) GetMemoryBytes() uint64 {
	if m != nil {
		return m.MemoryBytes
	}
	return 0
}

func (m *BootVirtualMachineRequest) GetBlockdevs() []*BlockDev {
	if m != nil {
		return m.Blockdevs
	}
	return nil
}

func (m *BootVirtualMachineRequest) GetNetdevs() []*NetDev {
	if m != nil {
		return m.Netdevs
	}
	return nil
}

func (m *BootVirtualMachineRequest) GetLoginUsername() string {
	if m != nil {
		return m.LoginUsername
	}
	return ""
}

func (m *BootVirtualMachineRequest) GetSshAuthorizedKeys() []string {
	if m != nil {
		return m.SshAuthorizedKeys
	}
	return nil
}

type BootVirtualMachineResponse struct {
	State                VirtualMachineState `protobuf:"varint,1,opt,name=state,proto3,enum=n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineState" json:"state,omitempty"`
	WebsocketPort        uint32              `protobuf:"varint,2,opt,name=websocket_port,json=websocketPort,proto3" json:"websocket_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BootVirtualMachineResponse) Reset()         { *m = BootVirtualMachineResponse{} }
func (m *BootVirtualMachineResponse) String() string { return proto.CompactTextString(m) }
func (*BootVirtualMachineResponse) ProtoMessage()    {}
func (*BootVirtualMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{3}
}

func (m *BootVirtualMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootVirtualMachineResponse.Unmarshal(m, b)
}
func (m *BootVirtualMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootVirtualMachineResponse.Marshal(b, m, deterministic)
}
func (m *BootVirtualMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootVirtualMachineResponse.Merge(m, src)
}
func (m *BootVirtualMachineResponse) XXX_Size() int {
	return xxx_messageInfo_BootVirtualMachineResponse.Size(m)
}
func (m *BootVirtualMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BootVirtualMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BootVirtualMachineResponse proto.InternalMessageInfo

func (m *BootVirtualMachineResponse) GetState() VirtualMachineState {
	if m != nil {
		return m.State
	}
	return VirtualMachineState_FAILED
}

func (m *BootVirtualMachineResponse) GetWebsocketPort() uint32 {
	if m != nil {
		return m.WebsocketPort
	}
	return 0
}

type RebootVirtualMachineRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hard                 bool     `protobuf:"varint,2,opt,name=hard,proto3" json:"hard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RebootVirtualMachineRequest) Reset()         { *m = RebootVirtualMachineRequest{} }
func (m *RebootVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*RebootVirtualMachineRequest) ProtoMessage()    {}
func (*RebootVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{4}
}

func (m *RebootVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RebootVirtualMachineRequest.Unmarshal(m, b)
}
func (m *RebootVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RebootVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (m *RebootVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RebootVirtualMachineRequest.Merge(m, src)
}
func (m *RebootVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_RebootVirtualMachineRequest.Size(m)
}
func (m *RebootVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RebootVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RebootVirtualMachineRequest proto.InternalMessageInfo

func (m *RebootVirtualMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RebootVirtualMachineRequest) GetHard() bool {
	if m != nil {
		return m.Hard
	}
	return false
}

type RebootVirtualMachineResponse struct {
	State                VirtualMachineState `protobuf:"varint,1,opt,name=state,proto3,enum=n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RebootVirtualMachineResponse) Reset()         { *m = RebootVirtualMachineResponse{} }
func (m *RebootVirtualMachineResponse) String() string { return proto.CompactTextString(m) }
func (*RebootVirtualMachineResponse) ProtoMessage()    {}
func (*RebootVirtualMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{5}
}

func (m *RebootVirtualMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RebootVirtualMachineResponse.Unmarshal(m, b)
}
func (m *RebootVirtualMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RebootVirtualMachineResponse.Marshal(b, m, deterministic)
}
func (m *RebootVirtualMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RebootVirtualMachineResponse.Merge(m, src)
}
func (m *RebootVirtualMachineResponse) XXX_Size() int {
	return xxx_messageInfo_RebootVirtualMachineResponse.Size(m)
}
func (m *RebootVirtualMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RebootVirtualMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RebootVirtualMachineResponse proto.InternalMessageInfo

func (m *RebootVirtualMachineResponse) GetState() VirtualMachineState {
	if m != nil {
		return m.State
	}
	return VirtualMachineState_FAILED
}

type ShutdownVirtualMachineRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hard                 bool     `protobuf:"varint,2,opt,name=hard,proto3" json:"hard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownVirtualMachineRequest) Reset()         { *m = ShutdownVirtualMachineRequest{} }
func (m *ShutdownVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownVirtualMachineRequest) ProtoMessage()    {}
func (*ShutdownVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{6}
}

func (m *ShutdownVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownVirtualMachineRequest.Unmarshal(m, b)
}
func (m *ShutdownVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (m *ShutdownVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownVirtualMachineRequest.Merge(m, src)
}
func (m *ShutdownVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_ShutdownVirtualMachineRequest.Size(m)
}
func (m *ShutdownVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownVirtualMachineRequest proto.InternalMessageInfo

func (m *ShutdownVirtualMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShutdownVirtualMachineRequest) GetHard() bool {
	if m != nil {
		return m.Hard
	}
	return false
}

type ShutdownVirtualMachineResponse struct {
	State                VirtualMachineState `protobuf:"varint,1,opt,name=state,proto3,enum=n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ShutdownVirtualMachineResponse) Reset()         { *m = ShutdownVirtualMachineResponse{} }
func (m *ShutdownVirtualMachineResponse) String() string { return proto.CompactTextString(m) }
func (*ShutdownVirtualMachineResponse) ProtoMessage()    {}
func (*ShutdownVirtualMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{7}
}

func (m *ShutdownVirtualMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownVirtualMachineResponse.Unmarshal(m, b)
}
func (m *ShutdownVirtualMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownVirtualMachineResponse.Marshal(b, m, deterministic)
}
func (m *ShutdownVirtualMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownVirtualMachineResponse.Merge(m, src)
}
func (m *ShutdownVirtualMachineResponse) XXX_Size() int {
	return xxx_messageInfo_ShutdownVirtualMachineResponse.Size(m)
}
func (m *ShutdownVirtualMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownVirtualMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownVirtualMachineResponse proto.InternalMessageInfo

func (m *ShutdownVirtualMachineResponse) GetState() VirtualMachineState {
	if m != nil {
		return m.State
	}
	return VirtualMachineState_FAILED
}

type DeleteVirtualMachineRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// TODO: netdev の情報を QMP から取るまでは、とりあえず渡してもらう
	Netdevs              []*NetDev `protobuf:"bytes,8,rep,name=netdevs,proto3" json:"netdevs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteVirtualMachineRequest) Reset()         { *m = DeleteVirtualMachineRequest{} }
func (m *DeleteVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteVirtualMachineRequest) ProtoMessage()    {}
func (*DeleteVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0185a3add0d295d8, []int{8}
}

func (m *DeleteVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteVirtualMachineRequest.Unmarshal(m, b)
}
func (m *DeleteVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (m *DeleteVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteVirtualMachineRequest.Merge(m, src)
}
func (m *DeleteVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteVirtualMachineRequest.Size(m)
}
func (m *DeleteVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteVirtualMachineRequest proto.InternalMessageInfo

func (m *DeleteVirtualMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteVirtualMachineRequest) GetNetdevs() []*NetDev {
	if m != nil {
		return m.Netdevs
	}
	return nil
}

func init() {
	proto.RegisterEnum("n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineState", VirtualMachineState_name, VirtualMachineState_value)
	proto.RegisterType((*BlockDev)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.BlockDev")
	proto.RegisterType((*NetDev)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.NetDev")
	proto.RegisterType((*BootVirtualMachineRequest)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.BootVirtualMachineRequest")
	proto.RegisterType((*BootVirtualMachineResponse)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.BootVirtualMachineResponse")
	proto.RegisterType((*RebootVirtualMachineRequest)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.RebootVirtualMachineRequest")
	proto.RegisterType((*RebootVirtualMachineResponse)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.RebootVirtualMachineResponse")
	proto.RegisterType((*ShutdownVirtualMachineRequest)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.ShutdownVirtualMachineRequest")
	proto.RegisterType((*ShutdownVirtualMachineResponse)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.ShutdownVirtualMachineResponse")
	proto.RegisterType((*DeleteVirtualMachineRequest)(nil), "n0stack.internal.n0core.provisioning.virtual_machine.DeleteVirtualMachineRequest")
}

func init() {
	proto.RegisterFile("n0core/pkg/api/provisioning/virtualmachine/agent.proto", fileDescriptor_0185a3add0d295d8)
}

var fileDescriptor_0185a3add0d295d8 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xdf, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0xd7, 0xcd, 0x9f, 0xcd, 0x9e, 0x34, 0x25, 0x9d, 0xae, 0x2a, 0x93, 0xa5, 0x28, 0x58,
	0x42, 0x5a, 0x7a, 0x61, 0x57, 0x4b, 0xc5, 0x0d, 0x08, 0x29, 0x21, 0x61, 0x89, 0x16, 0xbc, 0x8b,
	0xd3, 0x14, 0x09, 0x21, 0x59, 0x8e, 0x7d, 0x48, 0xac, 0x24, 0x1e, 0x33, 0x33, 0x76, 0x08, 0x37,
	0x08, 0x71, 0xd5, 0xf7, 0xe0, 0x25, 0xb8, 0xe2, 0x8a, 0xa7, 0xe0, 0x05, 0x78, 0x0c, 0x34, 0x33,
	0x4e, 0xda, 0x15, 0xd9, 0x55, 0x09, 0x51, 0xef, 0x66, 0xbe, 0x99, 0xfc, 0xce, 0x7c, 0x73, 0xce,
	0x89, 0x07, 0x3e, 0x4a, 0x9e, 0x84, 0x94, 0xa1, 0x93, 0xce, 0x26, 0x4e, 0x90, 0xc6, 0x4e, 0xca,
	0x68, 0x1e, 0xf3, 0x98, 0x26, 0x71, 0x32, 0x71, 0xf2, 0x98, 0x89, 0x2c, 0x98, 0x2f, 0x82, 0x70,
	0x1a, 0x27, 0xe8, 0x04, 0x13, 0x4c, 0x84, 0x9d, 0x32, 0x2a, 0x28, 0x79, 0x9a, 0x3c, 0xe1, 0x22,
	0x08, 0x67, 0x76, 0x9c, 0x08, 0x64, 0x49, 0x30, 0xb7, 0x35, 0xc8, 0x7e, 0x15, 0x60, 0x17, 0x00,
	0xbf, 0x20, 0xb4, 0x4e, 0x26, 0x94, 0x4e, 0xe6, 0xe8, 0x28, 0xc6, 0x38, 0xfb, 0xde, 0xc1, 0x45,
	0x2a, 0x56, 0x1a, 0x69, 0x5d, 0x42, 0xad, 0x3b, 0xa7, 0xe1, 0xac, 0x87, 0x39, 0x21, 0x50, 0x4e,
	0x82, 0x05, 0x9a, 0x46, 0xdb, 0x38, 0x3d, 0xf2, 0xd4, 0x98, 0x34, 0xa1, 0x94, 0xb1, 0xb9, 0x79,
	0x47, 0x49, 0x72, 0x48, 0x1e, 0x01, 0x8c, 0x29, 0x15, 0x7e, 0x9c, 0x44, 0xf8, 0xa3, 0x59, 0x6a,
	0x1b, 0xa7, 0x0d, 0xef, 0x48, 0x2a, 0x03, 0x29, 0x58, 0x7f, 0x19, 0x50, 0x75, 0x51, 0xdc, 0xc4,
	0x7b, 0x0f, 0xee, 0x26, 0x28, 0x96, 0x94, 0xcd, 0x7c, 0xb5, 0xa6, 0xc1, 0xf5, 0x42, 0x73, 0xe5,
	0x96, 0x0f, 0xa0, 0x39, 0x0d, 0x58, 0xb4, 0x0c, 0x18, 0xfa, 0x41, 0x14, 0x31, 0xe4, 0x5c, 0x85,
	0x39, 0xf2, 0xde, 0x5a, 0xeb, 0x1d, 0x2d, 0x93, 0xc7, 0x70, 0x3f, 0x4e, 0xf3, 0xa7, 0xeb, 0x6d,
	0x7e, 0x18, 0x47, 0xcc, 0x2c, 0xeb, 0xbd, 0x72, 0xa1, 0xd8, 0xf7, 0x59, 0x1c, 0x31, 0x19, 0x59,
	0xed, 0x9d, 0x04, 0x02, 0x97, 0xc1, 0xca, 0xac, 0xe8, 0xc8, 0x52, 0x3b, 0xd7, 0x12, 0x69, 0x43,
	0x5d, 0x1e, 0x8a, 0x23, 0xcb, 0x91, 0x71, 0xb3, 0xda, 0x2e, 0xa9, 0xb3, 0xbd, 0x94, 0xac, 0x17,
	0x25, 0x78, 0xbb, 0x4b, 0xa9, 0x78, 0xae, 0xef, 0xf8, 0x2b, 0x7d, 0xc5, 0x1e, 0xfe, 0x90, 0x21,
	0x17, 0x5b, 0x0d, 0x13, 0x28, 0x67, 0x59, 0x1c, 0x15, 0x46, 0xd5, 0x98, 0x1c, 0x43, 0x25, 0x0f,
	0xd3, 0x8c, 0x17, 0xb7, 0xa7, 0x27, 0xf2, 0x80, 0x0b, 0x5c, 0x50, 0xb6, 0xf2, 0xc7, 0x2b, 0x81,
	0x5c, 0xf9, 0x28, 0x7b, 0x75, 0xad, 0x75, 0xa5, 0x44, 0xbe, 0x83, 0xa3, 0xb1, 0xcc, 0x56, 0x84,
	0x39, 0x37, 0x2b, 0xed, 0xd2, 0x69, 0xfd, 0xec, 0x53, 0x7b, 0x97, 0xa2, 0xb0, 0xd7, 0x49, 0xf7,
	0x5e, 0x02, 0xc9, 0x73, 0x38, 0x4c, 0x50, 0x28, 0x76, 0x55, 0xb1, 0x3f, 0xd9, 0x8d, 0xad, 0xd3,
	0xef, 0xad, 0x61, 0xe4, 0x7d, 0xb8, 0x37, 0xa7, 0x93, 0x38, 0xf1, 0x33, 0x2e, 0x21, 0x0b, 0x34,
	0x0f, 0xd5, 0x65, 0x34, 0x94, 0x3a, 0x2a, 0x44, 0x62, 0xc3, 0x03, 0xce, 0xa7, 0x7e, 0x90, 0x89,
	0x29, 0x65, 0xf1, 0x4f, 0x18, 0xf9, 0x33, 0x5c, 0x71, 0xb3, 0xa6, 0xb2, 0x70, 0x9f, 0xf3, 0x69,
	0x67, 0xb3, 0x72, 0x81, 0x2b, 0x6e, 0xfd, 0x66, 0x40, 0x6b, 0x5b, 0x2e, 0x78, 0x4a, 0x13, 0x8e,
	0xc4, 0x87, 0x0a, 0x17, 0x81, 0xd0, 0xd9, 0xb8, 0x77, 0x36, 0xd8, 0xcd, 0xcb, 0x75, 0xf8, 0x50,
	0x02, 0x3d, 0xcd, 0x95, 0xb6, 0x96, 0x38, 0xe6, 0x34, 0x9c, 0xa1, 0xf0, 0x53, 0xca, 0x84, 0xca,
	0x71, 0xc3, 0x6b, 0x6c, 0xd4, 0x2b, 0xca, 0x84, 0xd5, 0x87, 0x13, 0x0f, 0xc7, 0xff, 0xb5, 0x66,
	0x64, 0xa5, 0x2b, 0x5e, 0xcd, 0x53, 0x63, 0xeb, 0x67, 0x78, 0x67, 0x3b, 0xe6, 0x0d, 0xd9, 0xb5,
	0xce, 0xe1, 0xd1, 0x70, 0x9a, 0x89, 0x88, 0x2e, 0x93, 0xff, 0xe7, 0xe4, 0x17, 0x03, 0xde, 0xbd,
	0x89, 0xf4, 0xa6, 0xcc, 0xbc, 0x30, 0xe0, 0xa4, 0x87, 0x73, 0x14, 0xf8, 0xfa, 0x5e, 0x5e, 0x69,
	0x8f, 0xda, 0x1e, 0xdb, 0xe3, 0xf1, 0x08, 0x1e, 0x6c, 0x39, 0x29, 0x01, 0xa8, 0x7e, 0xde, 0x19,
	0x7c, 0xd9, 0xef, 0x35, 0x0f, 0x48, 0x1d, 0x0e, 0x47, 0xee, 0x85, 0x7b, 0xf9, 0x8d, 0xdb, 0x34,
	0xc8, 0x5d, 0xa8, 0x0d, 0xbf, 0x18, 0x3d, 0xeb, 0xc9, 0xd9, 0x1d, 0xb9, 0xe4, 0x8d, 0x5c, 0x77,
	0xe0, 0x9e, 0x37, 0x4b, 0xf2, 0x37, 0x57, 0x9d, 0xd1, 0xb0, 0xdf, 0x6b, 0x96, 0xcf, 0xfe, 0xae,
	0x40, 0xeb, 0x3a, 0xb7, 0x23, 0x3f, 0x25, 0x43, 0x64, 0x79, 0x1c, 0x22, 0xf9, 0xdd, 0x00, 0xf2,
	0xef, 0xee, 0x21, 0x97, 0x3b, 0xfe, 0x9d, 0xdc, 0x54, 0xdf, 0xad, 0xab, 0xfd, 0x01, 0x75, 0x71,
	0x58, 0x07, 0xe4, 0x0f, 0x03, 0x8e, 0xb7, 0x35, 0x03, 0xf9, 0x7a, 0xb7, 0x60, 0xb7, 0xf4, 0x67,
	0xcb, 0xdb, 0x27, 0x72, 0xe3, 0xe0, 0x4f, 0x03, 0x1e, 0x6e, 0xef, 0x01, 0x32, 0xdc, 0x2d, 0xe0,
	0xad, 0xbd, 0xd9, 0x7a, 0xb6, 0x5f, 0xe8, 0xc6, 0xc7, 0xaf, 0x06, 0x1c, 0x6f, 0xeb, 0xa3, 0x5d,
	0x33, 0x71, 0x4b, 0x4f, 0xb6, 0x1e, 0xda, 0xfa, 0x21, 0x63, 0xaf, 0x1f, 0x32, 0x76, 0x5f, 0x3e,
	0x64, 0xac, 0x83, 0xee, 0xc5, 0xb7, 0x03, 0x19, 0xcd, 0x0e, 0x42, 0xa7, 0x88, 0xea, 0xbc, 0xfe,
	0x0b, 0xeb, 0xe3, 0xeb, 0xd3, 0x71, 0x55, 0xe1, 0x3f, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd8,
	0x41, 0x00, 0xa1, 0xa5, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMachineAgentServiceClient is the client API for VirtualMachineAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineAgentServiceClient interface {
	BootVirtualMachine(ctx context.Context, in *BootVirtualMachineRequest, opts ...grpc.CallOption) (*BootVirtualMachineResponse, error)
	RebootVirtualMachine(ctx context.Context, in *RebootVirtualMachineRequest, opts ...grpc.CallOption) (*RebootVirtualMachineResponse, error)
	ShutdownVirtualMachine(ctx context.Context, in *ShutdownVirtualMachineRequest, opts ...grpc.CallOption) (*ShutdownVirtualMachineResponse, error)
	DeleteVirtualMachine(ctx context.Context, in *DeleteVirtualMachineRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type virtualMachineAgentServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineAgentServiceClient(cc *grpc.ClientConn) VirtualMachineAgentServiceClient {
	return &virtualMachineAgentServiceClient{cc}
}

func (c *virtualMachineAgentServiceClient) BootVirtualMachine(ctx context.Context, in *BootVirtualMachineRequest, opts ...grpc.CallOption) (*BootVirtualMachineResponse, error) {
	out := new(BootVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/BootVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentServiceClient) RebootVirtualMachine(ctx context.Context, in *RebootVirtualMachineRequest, opts ...grpc.CallOption) (*RebootVirtualMachineResponse, error) {
	out := new(RebootVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/RebootVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentServiceClient) ShutdownVirtualMachine(ctx context.Context, in *ShutdownVirtualMachineRequest, opts ...grpc.CallOption) (*ShutdownVirtualMachineResponse, error) {
	out := new(ShutdownVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/ShutdownVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentServiceClient) DeleteVirtualMachine(ctx context.Context, in *DeleteVirtualMachineRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/DeleteVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineAgentServiceServer is the server API for VirtualMachineAgentService service.
type VirtualMachineAgentServiceServer interface {
	BootVirtualMachine(context.Context, *BootVirtualMachineRequest) (*BootVirtualMachineResponse, error)
	RebootVirtualMachine(context.Context, *RebootVirtualMachineRequest) (*RebootVirtualMachineResponse, error)
	ShutdownVirtualMachine(context.Context, *ShutdownVirtualMachineRequest) (*ShutdownVirtualMachineResponse, error)
	DeleteVirtualMachine(context.Context, *DeleteVirtualMachineRequest) (*empty.Empty, error)
}

// UnimplementedVirtualMachineAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineAgentServiceServer struct {
}

func (*UnimplementedVirtualMachineAgentServiceServer) BootVirtualMachine(ctx context.Context, req *BootVirtualMachineRequest) (*BootVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BootVirtualMachine not implemented")
}
func (*UnimplementedVirtualMachineAgentServiceServer) RebootVirtualMachine(ctx context.Context, req *RebootVirtualMachineRequest) (*RebootVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebootVirtualMachine not implemented")
}
func (*UnimplementedVirtualMachineAgentServiceServer) ShutdownVirtualMachine(ctx context.Context, req *ShutdownVirtualMachineRequest) (*ShutdownVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutdownVirtualMachine not implemented")
}
func (*UnimplementedVirtualMachineAgentServiceServer) DeleteVirtualMachine(ctx context.Context, req *DeleteVirtualMachineRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVirtualMachine not implemented")
}

func RegisterVirtualMachineAgentServiceServer(s *grpc.Server, srv VirtualMachineAgentServiceServer) {
	s.RegisterService(&_VirtualMachineAgentService_serviceDesc, srv)
}

func _VirtualMachineAgentService_BootVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServiceServer).BootVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/BootVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServiceServer).BootVirtualMachine(ctx, req.(*BootVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineAgentService_RebootVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebootVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServiceServer).RebootVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/RebootVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServiceServer).RebootVirtualMachine(ctx, req.(*RebootVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineAgentService_ShutdownVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServiceServer).ShutdownVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/ShutdownVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServiceServer).ShutdownVirtualMachine(ctx, req.(*ShutdownVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineAgentService_DeleteVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServiceServer).DeleteVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService/DeleteVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServiceServer).DeleteVirtualMachine(ctx, req.(*DeleteVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineAgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.internal.n0core.provisioning.virtual_machine.VirtualMachineAgentService",
	HandlerType: (*VirtualMachineAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BootVirtualMachine",
			Handler:    _VirtualMachineAgentService_BootVirtualMachine_Handler,
		},
		{
			MethodName: "RebootVirtualMachine",
			Handler:    _VirtualMachineAgentService_RebootVirtualMachine_Handler,
		},
		{
			MethodName: "ShutdownVirtualMachine",
			Handler:    _VirtualMachineAgentService_ShutdownVirtualMachine_Handler,
		},
		{
			MethodName: "DeleteVirtualMachine",
			Handler:    _VirtualMachineAgentService_DeleteVirtualMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "n0core/pkg/api/provisioning/virtualmachine/agent.proto",
}
