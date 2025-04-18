// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/upgrade/v1beta1/upgrade.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Plan specifies information about a planned upgrade and when it should occur.
type Plan struct {
	// Sets the name for the upgrade. This name will be used by the upgraded
	// version of the software to apply any special "on-upgrade" commands during
	// the first BeginBlock method after the upgrade is applied. It is also used
	// to detect whether a software version can handle a given upgrade. If no
	// upgrade handler with this name has been set in the software, it will be
	// assumed that the software is out-of-date when the upgrade Time or Height is
	// reached and the software will exit.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated: Time based upgrades have been deprecated. Time based upgrade logic
	// has been removed from the SDK.
	// If this field is not empty, an error will be thrown.
	Time time.Time `protobuf:"bytes,2,opt,name=time,proto3,stdtime" json:"time"` // Deprecated: Do not use.
	// The height at which the upgrade must be performed.
	Height int64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	// Deprecated: UpgradedClientState field has been deprecated. IBC upgrade logic has been
	// moved to the IBC module in the sub module 02-client.
	// If this field is not empty, an error will be thrown.
	UpgradedClientState *types.Any `protobuf:"bytes,5,opt,name=upgraded_client_state,json=upgradedClientState,proto3" json:"upgraded_client_state,omitempty"` // Deprecated: Do not use.
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

// SoftwareUpgradeProposal is a gov Content type for initiating a software
// upgrade.
// Deprecated: This legacy proposal is deprecated in favor of Msg-based gov
// proposals, see MsgSoftwareUpgrade.
//
// Deprecated: Do not use.
type SoftwareUpgradeProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// plan of the proposal
	Plan Plan `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan"`
}

func (m *SoftwareUpgradeProposal) Reset()         { *m = SoftwareUpgradeProposal{} }
func (m *SoftwareUpgradeProposal) String() string { return proto.CompactTextString(m) }
func (*SoftwareUpgradeProposal) ProtoMessage()    {}
func (*SoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{1}
}
func (m *SoftwareUpgradeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SoftwareUpgradeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SoftwareUpgradeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SoftwareUpgradeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoftwareUpgradeProposal.Merge(m, src)
}
func (m *SoftwareUpgradeProposal) XXX_Size() int {
	return m.Size()
}
func (m *SoftwareUpgradeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SoftwareUpgradeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SoftwareUpgradeProposal proto.InternalMessageInfo

// CancelSoftwareUpgradeProposal is a gov Content type for cancelling a software
// upgrade.
// Deprecated: This legacy proposal is deprecated in favor of Msg-based gov
// proposals, see MsgCancelUpgrade.
//
// Deprecated: Do not use.
type CancelSoftwareUpgradeProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *CancelSoftwareUpgradeProposal) Reset()         { *m = CancelSoftwareUpgradeProposal{} }
func (m *CancelSoftwareUpgradeProposal) String() string { return proto.CompactTextString(m) }
func (*CancelSoftwareUpgradeProposal) ProtoMessage()    {}
func (*CancelSoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{2}
}
func (m *CancelSoftwareUpgradeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelSoftwareUpgradeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelSoftwareUpgradeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelSoftwareUpgradeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelSoftwareUpgradeProposal.Merge(m, src)
}
func (m *CancelSoftwareUpgradeProposal) XXX_Size() int {
	return m.Size()
}
func (m *CancelSoftwareUpgradeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelSoftwareUpgradeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CancelSoftwareUpgradeProposal proto.InternalMessageInfo

// ModuleVersion specifies a module and its consensus version.
type ModuleVersion struct {
	// name of the app module
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// consensus version of the app module
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ModuleVersion) Reset()         { *m = ModuleVersion{} }
func (m *ModuleVersion) String() string { return proto.CompactTextString(m) }
func (*ModuleVersion) ProtoMessage()    {}
func (*ModuleVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{3}
}
func (m *ModuleVersion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleVersion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleVersion.Merge(m, src)
}
func (m *ModuleVersion) XXX_Size() int {
	return m.Size()
}
func (m *ModuleVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleVersion proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Plan)(nil), "cosmos.upgrade.v1beta1.Plan")
	proto.RegisterType((*SoftwareUpgradeProposal)(nil), "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal")
	proto.RegisterType((*CancelSoftwareUpgradeProposal)(nil), "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal")
	proto.RegisterType((*ModuleVersion)(nil), "cosmos.upgrade.v1beta1.ModuleVersion")
}

func init() {
	proto.RegisterFile("cosmos/upgrade/v1beta1/upgrade.proto", fileDescriptor_ccf2a7d4d7b48dca)
}

var fileDescriptor_ccf2a7d4d7b48dca = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xbf, 0x6e, 0xd3, 0x40,
	0x18, 0xcf, 0xb5, 0x6e, 0x51, 0x2e, 0x42, 0x15, 0x26, 0xb4, 0xd7, 0xa8, 0x38, 0x96, 0xc5, 0x10,
	0x55, 0x8a, 0x4d, 0x5b, 0x26, 0x33, 0x20, 0x92, 0x11, 0x90, 0x8a, 0x0b, 0x1d, 0x58, 0xa2, 0x4b,
	0x7c, 0x71, 0xad, 0xda, 0x77, 0x96, 0xef, 0x12, 0xc8, 0x2b, 0x30, 0xf5, 0x11, 0x18, 0x11, 0x53,
	0x87, 0x3e, 0x44, 0xc4, 0x54, 0x31, 0x21, 0x21, 0xf1, 0x27, 0x19, 0xca, 0xc6, 0x2b, 0xa0, 0xbb,
	0xb3, 0xab, 0x08, 0x0a, 0x62, 0x60, 0x89, 0xbe, 0xef, 0xbb, 0xef, 0x97, 0xdf, 0x9f, 0x3b, 0xc3,
	0x3b, 0x03, 0xc6, 0x53, 0xc6, 0xbd, 0x51, 0x16, 0xe5, 0x38, 0x24, 0xde, 0x78, 0xa7, 0x4f, 0x04,
	0xde, 0x29, 0x7b, 0x37, 0xcb, 0x99, 0x60, 0xe6, 0xba, 0xde, 0x72, 0xcb, 0x69, 0xb1, 0xd5, 0xd8,
	0x8c, 0x18, 0x8b, 0x12, 0xe2, 0xa9, 0xad, 0xfe, 0x68, 0xe8, 0x61, 0x3a, 0xd1, 0x90, 0x46, 0x3d,
	0x62, 0x11, 0x53, 0xa5, 0x27, 0xab, 0x62, 0xda, 0xfc, 0x15, 0x20, 0xe2, 0x94, 0x70, 0x81, 0xd3,
	0xac, 0x58, 0xd8, 0xd4, 0x4c, 0x3d, 0x8d, 0x2c, 0x68, 0xf5, 0xd1, 0x0d, 0x9c, 0xc6, 0x94, 0x79,
	0xea, 0x57, 0x8f, 0x9c, 0x1f, 0x00, 0x1a, 0xfb, 0x09, 0xa6, 0xa6, 0x09, 0x0d, 0x8a, 0x53, 0x82,
	0x80, 0x0d, 0x5a, 0xd5, 0x40, 0xd5, 0xe6, 0x03, 0x68, 0xc8, 0x7f, 0x47, 0x4b, 0x36, 0x68, 0xd5,
	0x76, 0x1b, 0xae, 0xa6, 0x76, 0x4b, 0x6a, 0xf7, 0x59, 0x49, 0xdd, 0x59, 0x9b, 0x7e, 0x6e, 0x56,
	0x4e, 0xbe, 0x34, 0xc1, 0xdb, 0x8b, 0xd3, 0x6d, 0x80, 0x40, 0xa0, 0x80, 0xe6, 0x3a, 0x5c, 0x3d,
	0x22, 0x71, 0x74, 0x24, 0xd0, 0xb2, 0x0d, 0x5a, 0xcb, 0x41, 0xd1, 0x49, 0xb2, 0x98, 0x0e, 0x19,
	0x32, 0x34, 0x99, 0xac, 0xcd, 0xc7, 0xf0, 0x56, 0x11, 0x4e, 0xd8, 0x1b, 0x24, 0x31, 0xa1, 0xa2,
	0xc7, 0x05, 0x16, 0x04, 0xad, 0x28, 0xf6, 0xfa, 0x6f, 0xec, 0x0f, 0xe9, 0xa4, 0xb3, 0x84, 0x40,
	0x70, 0xb3, 0x84, 0x75, 0x15, 0xea, 0x40, 0x82, 0x7c, 0xf4, 0xfd, 0x4d, 0x13, 0xbc, 0xbe, 0x38,
	0xdd, 0x5e, 0xd3, 0x09, 0xb4, 0x79, 0x78, 0xec, 0x49, 0xa3, 0xce, 0x27, 0x00, 0x37, 0x0e, 0xd8,
	0x50, 0xbc, 0xc4, 0x39, 0x79, 0xae, 0x91, 0xfb, 0x39, 0xcb, 0x18, 0xc7, 0x89, 0x59, 0x87, 0x2b,
	0x22, 0x16, 0x49, 0x99, 0x82, 0x6e, 0x4c, 0x1b, 0xd6, 0x42, 0xc2, 0x07, 0x79, 0x9c, 0x89, 0x98,
	0x51, 0x95, 0x46, 0x35, 0x58, 0x1c, 0x99, 0xf7, 0xa1, 0x91, 0x25, 0x98, 0x2a, 0x97, 0xb5, 0xdd,
	0x2d, 0xf7, 0xea, 0xcb, 0x76, 0x25, 0x7f, 0xa7, 0x2a, 0xa3, 0x52, 0x31, 0x05, 0x0a, 0xe4, 0x3f,
	0x92, 0x52, 0xdf, 0x9f, 0xb5, 0x1b, 0x05, 0x2a, 0x62, 0xe3, 0x4b, 0x44, 0x97, 0x51, 0x41, 0xa8,
	0x90, 0x46, 0x9c, 0x05, 0x23, 0x7f, 0xd0, 0x8f, 0x80, 0xf3, 0x0e, 0xc0, 0xdb, 0x5d, 0x4c, 0x07,
	0x24, 0xf9, 0xcf, 0x1e, 0xfd, 0xa7, 0xff, 0x26, 0xb3, 0xb5, 0x20, 0xf3, 0xaf, 0x42, 0x10, 0x70,
	0x0e, 0xe1, 0xf5, 0x27, 0x2c, 0x1c, 0x25, 0xe4, 0x90, 0xe4, 0x5c, 0xe6, 0x78, 0xd5, 0x23, 0x44,
	0xf0, 0xda, 0x58, 0x1f, 0x2b, 0x55, 0x46, 0x50, 0xb6, 0xfe, 0x86, 0x54, 0xf4, 0xe1, 0xac, 0xbd,
	0x70, 0xc5, 0xf6, 0x5d, 0xf7, 0xde, 0x5e, 0xc7, 0x9f, 0x7e, 0xb3, 0x2a, 0xd3, 0x99, 0x05, 0xce,
	0x67, 0x16, 0xf8, 0x3a, 0xb3, 0xc0, 0xc9, 0xdc, 0xaa, 0x9c, 0xcf, 0xad, 0xca, 0xc7, 0xb9, 0x55,
	0x79, 0xb1, 0xa5, 0xd7, 0x79, 0x78, 0xec, 0xc6, 0xcc, 0x7b, 0x75, 0xf9, 0xe1, 0x8a, 0x49, 0x46,
	0x78, 0x7f, 0x55, 0xbd, 0xaf, 0xbd, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x6a, 0x37, 0x42,
	0xd7, 0x03, 0x00, 0x00,
}

func (this *Plan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Plan)
	if !ok {
		that2, ok := that.(Plan)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	if !this.UpgradedClientState.Equal(that1.UpgradedClientState) {
		return false
	}
	return true
}
func (this *SoftwareUpgradeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SoftwareUpgradeProposal)
	if !ok {
		that2, ok := that.(SoftwareUpgradeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.Plan.Equal(&that1.Plan) {
		return false
	}
	return true
}
func (this *CancelSoftwareUpgradeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CancelSoftwareUpgradeProposal)
	if !ok {
		that2, ok := that.(CancelSoftwareUpgradeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	return true
}
func (this *ModuleVersion) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ModuleVersion)
	if !ok {
		that2, ok := that.(ModuleVersion)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpgradedClientState != nil {
		{
			size, err := m.UpgradedClientState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintUpgrade(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintUpgrade(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SoftwareUpgradeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SoftwareUpgradeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SoftwareUpgradeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintUpgrade(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CancelSoftwareUpgradeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelSoftwareUpgradeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelSoftwareUpgradeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModuleVersion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleVersion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleVersion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpgrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpgrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovUpgrade(uint64(l))
	if m.Height != 0 {
		n += 1 + sovUpgrade(uint64(m.Height))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	if m.UpgradedClientState != nil {
		l = m.UpgradedClientState.Size()
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func (m *SoftwareUpgradeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = m.Plan.Size()
	n += 1 + l + sovUpgrade(uint64(l))
	return n
}

func (m *CancelSoftwareUpgradeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func (m *ModuleVersion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovUpgrade(uint64(m.Version))
	}
	return n
}

func sovUpgrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpgrade(x uint64) (n int) {
	return sovUpgrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpgradedClientState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpgradedClientState == nil {
				m.UpgradedClientState = &types.Any{}
			}
			if err := m.UpgradedClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SoftwareUpgradeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SoftwareUpgradeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SoftwareUpgradeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CancelSoftwareUpgradeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CancelSoftwareUpgradeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelSoftwareUpgradeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ModuleVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ModuleVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUpgrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpgrade
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthUpgrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpgrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpgrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpgrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpgrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpgrade = fmt.Errorf("proto: unexpected end of group")
)
