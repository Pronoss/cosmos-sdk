// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/app/v1alpha1/config.proto

package v1alpha1

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Config represents the configuration for a Cosmos SDK ABCI app.
// It is intended that all state machine logic including the version of
// baseapp and tx handlers (and possibly even Tendermint) that an app needs
// can be described in a config object. For compatibility, the framework should
// allow a mixture of declarative and imperative app wiring, however, apps
// that strive for the maximum ease of maintainability should be able to describe
// their state machine with a config object alone.
type Config struct {
	// modules are the module configurations for the app.
	Modules []*ModuleConfig `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
	// golang_bindings specifies explicit interface to implementation type bindings which
	// depinject uses to resolve interface inputs to provider functions.  The scope of this
	// field's configuration is global (not module specific).
	GolangBindings []*GolangBinding `protobuf:"bytes,2,rep,name=golang_bindings,json=golangBindings,proto3" json:"golang_bindings,omitempty"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_5af1d229673256fa, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetModules() []*ModuleConfig {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *Config) GetGolangBindings() []*GolangBinding {
	if m != nil {
		return m.GolangBindings
	}
	return nil
}

// ModuleConfig is a module configuration for an app.
type ModuleConfig struct {
	// name is the unique name of the module within the app. It should be a name
	// that persists between different versions of a module so that modules
	// can be smoothly upgraded to new versions.
	//
	// For example, for the module cosmos.bank.module.v1.Module, we may chose
	// to simply name the module "bank" in the app. When we upgrade to
	// cosmos.bank.module.v2.Module, the app-specific name "bank" stays the same
	// and the framework knows that the v2 module should receive all the same state
	// that the v1 module had. Note: modules should provide info on which versions
	// they can migrate from in the ModuleDescriptor.can_migration_from field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// config is the config object for the module. Module config messages should
	// define a ModuleDescriptor using the cosmos.app.v1alpha1.is_module extension.
	Config *types.Any `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// golang_bindings specifies explicit interface to implementation type bindings which
	// depinject uses to resolve interface inputs to provider functions.  The scope of this
	// field's configuration is module specific.
	GolangBindings []*GolangBinding `protobuf:"bytes,3,rep,name=golang_bindings,json=golangBindings,proto3" json:"golang_bindings,omitempty"`
}

func (m *ModuleConfig) Reset()         { *m = ModuleConfig{} }
func (m *ModuleConfig) String() string { return proto.CompactTextString(m) }
func (*ModuleConfig) ProtoMessage()    {}
func (*ModuleConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5af1d229673256fa, []int{1}
}
func (m *ModuleConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleConfig.Merge(m, src)
}
func (m *ModuleConfig) XXX_Size() int {
	return m.Size()
}
func (m *ModuleConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleConfig proto.InternalMessageInfo

func (m *ModuleConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModuleConfig) GetConfig() *types.Any {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ModuleConfig) GetGolangBindings() []*GolangBinding {
	if m != nil {
		return m.GolangBindings
	}
	return nil
}

// GolangBinding is an explicit interface type to implementing type binding for dependency injection.
type GolangBinding struct {
	// interface_type is the interface type which will be bound to a specific implementation type
	InterfaceType string `protobuf:"bytes,1,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	// implementation is the implementing type which will be supplied when an input of type interface is requested
	Implementation string `protobuf:"bytes,2,opt,name=implementation,proto3" json:"implementation,omitempty"`
}

func (m *GolangBinding) Reset()         { *m = GolangBinding{} }
func (m *GolangBinding) String() string { return proto.CompactTextString(m) }
func (*GolangBinding) ProtoMessage()    {}
func (*GolangBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_5af1d229673256fa, []int{2}
}
func (m *GolangBinding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GolangBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GolangBinding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GolangBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GolangBinding.Merge(m, src)
}
func (m *GolangBinding) XXX_Size() int {
	return m.Size()
}
func (m *GolangBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_GolangBinding.DiscardUnknown(m)
}

var xxx_messageInfo_GolangBinding proto.InternalMessageInfo

func (m *GolangBinding) GetInterfaceType() string {
	if m != nil {
		return m.InterfaceType
	}
	return ""
}

func (m *GolangBinding) GetImplementation() string {
	if m != nil {
		return m.Implementation
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "cosmos.app.v1alpha1.Config")
	proto.RegisterType((*ModuleConfig)(nil), "cosmos.app.v1alpha1.ModuleConfig")
	proto.RegisterType((*GolangBinding)(nil), "cosmos.app.v1alpha1.GolangBinding")
}

func init() { proto.RegisterFile("cosmos/app/v1alpha1/config.proto", fileDescriptor_5af1d229673256fa) }

var fileDescriptor_5af1d229673256fa = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4f, 0x02, 0x31,
	0x1c, 0xc5, 0x29, 0x18, 0x0c, 0x45, 0x30, 0xa9, 0x0e, 0xa7, 0xc3, 0xe5, 0x24, 0xd1, 0x60, 0x62,
	0xda, 0x80, 0xa3, 0x93, 0x30, 0x38, 0x18, 0x97, 0x8b, 0x93, 0x83, 0xa4, 0xdc, 0x95, 0x5a, 0xbd,
	0x6b, 0x1b, 0x5a, 0x4c, 0xee, 0x5b, 0x18, 0x77, 0xbf, 0x8f, 0x23, 0xa3, 0xa3, 0x81, 0x2f, 0x62,
	0x6c, 0x3d, 0x82, 0x86, 0xc9, 0xad, 0xfd, 0xe7, 0xf7, 0x7f, 0x7d, 0xaf, 0x79, 0x30, 0x4a, 0x94,
	0xc9, 0x95, 0x21, 0x54, 0x6b, 0xf2, 0xdc, 0xa3, 0x99, 0x7e, 0xa0, 0x3d, 0x92, 0x28, 0x39, 0x11,
	0x1c, 0xeb, 0xa9, 0xb2, 0x0a, 0xed, 0x79, 0x02, 0x53, 0xad, 0x71, 0x49, 0x1c, 0x1e, 0x70, 0xa5,
	0x78, 0xc6, 0x88, 0x43, 0xc6, 0xb3, 0x09, 0xa1, 0xb2, 0xf0, 0x7c, 0xe7, 0x15, 0xc0, 0xfa, 0xd0,
	0x09, 0xa0, 0x0b, 0xb8, 0x9d, 0xab, 0x74, 0x96, 0x31, 0x13, 0x80, 0xa8, 0xd6, 0x6d, 0xf6, 0x8f,
	0xf0, 0x06, 0x31, 0x7c, 0xe3, 0x18, 0xbf, 0x13, 0x97, 0x1b, 0xe8, 0x1a, 0xee, 0x72, 0x95, 0x51,
	0xc9, 0x47, 0x63, 0x21, 0x53, 0x21, 0xb9, 0x09, 0xaa, 0x4e, 0xa4, 0xb3, 0x51, 0xe4, 0xca, 0xb1,
	0x03, 0x8f, 0xc6, 0x6d, 0xbe, 0x7e, 0x35, 0x9d, 0x37, 0x00, 0x77, 0xd6, 0x9f, 0x41, 0x08, 0x6e,
	0x49, 0x9a, 0xb3, 0x00, 0x44, 0xa0, 0xdb, 0x88, 0xdd, 0x19, 0x9d, 0xc1, 0xba, 0x4f, 0x1e, 0x54,
	0x23, 0xd0, 0x6d, 0xf6, 0xf7, 0xb1, 0x4f, 0x89, 0xcb, 0x94, 0xf8, 0x52, 0x16, 0xf1, 0x0f, 0xb3,
	0xc9, 0x5f, 0xed, 0xdf, 0xfe, 0xee, 0x61, 0xeb, 0x17, 0x80, 0x8e, 0x61, 0x5b, 0x48, 0xcb, 0xa6,
	0x13, 0x9a, 0xb0, 0x91, 0x2d, 0x74, 0xe9, 0xb4, 0xb5, 0x9a, 0xde, 0x16, 0x9a, 0xa1, 0x13, 0xd8,
	0x16, 0xb9, 0xce, 0x58, 0xce, 0xa4, 0xa5, 0x56, 0x28, 0xe9, 0xac, 0x37, 0xe2, 0x3f, 0xd3, 0xc1,
	0xf0, 0x7d, 0x11, 0x82, 0xf9, 0x22, 0x04, 0x9f, 0x8b, 0x10, 0xbc, 0x2c, 0xc3, 0xca, 0x7c, 0x19,
	0x56, 0x3e, 0x96, 0x61, 0xe5, 0xee, 0xd4, 0x9b, 0x35, 0xe9, 0x13, 0x16, 0x8a, 0xa4, 0x4c, 0x0b,
	0xf9, 0xc8, 0x12, 0xfb, 0x5d, 0x08, 0x1f, 0x74, 0x55, 0x8b, 0x71, 0xdd, 0xfd, 0xc3, 0xf9, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xfa, 0x9f, 0xcf, 0x34, 0x02, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GolangBindings) > 0 {
		for iNdEx := len(m.GolangBindings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GolangBindings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Modules) > 0 {
		for iNdEx := len(m.Modules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Modules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ModuleConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GolangBindings) > 0 {
		for iNdEx := len(m.GolangBindings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GolangBindings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GolangBinding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GolangBinding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GolangBinding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Implementation) > 0 {
		i -= len(m.Implementation)
		copy(dAtA[i:], m.Implementation)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Implementation)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.InterfaceType) > 0 {
		i -= len(m.InterfaceType)
		copy(dAtA[i:], m.InterfaceType)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.InterfaceType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Modules) > 0 {
		for _, e := range m.Modules {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if len(m.GolangBindings) > 0 {
		for _, e := range m.GolangBindings {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *ModuleConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.GolangBindings) > 0 {
		for _, e := range m.GolangBindings {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *GolangBinding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterfaceType)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Implementation)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Modules = append(m.Modules, &ModuleConfig{})
			if err := m.Modules[len(m.Modules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GolangBindings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GolangBindings = append(m.GolangBindings, &GolangBinding{})
			if err := m.GolangBindings[len(m.GolangBindings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func (m *ModuleConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ModuleConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &types.Any{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GolangBindings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GolangBindings = append(m.GolangBindings, &GolangBinding{})
			if err := m.GolangBindings[len(m.GolangBindings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func (m *GolangBinding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: GolangBinding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GolangBinding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Implementation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Implementation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
