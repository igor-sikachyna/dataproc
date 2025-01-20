// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: igorsikachyna/dataproc/v1/query.proto

package dataproc

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryGetCodeRequest is the request type for the Query/GetCode RPC
// method.
type QueryGetCodeRequest struct {
	// index defines the index of the code to retrieve.
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetCodeRequest) Reset()         { *m = QueryGetCodeRequest{} }
func (m *QueryGetCodeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCodeRequest) ProtoMessage()    {}
func (*QueryGetCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd51cf3aab79481, []int{0}
}
func (m *QueryGetCodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCodeRequest.Merge(m, src)
}
func (m *QueryGetCodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCodeRequest proto.InternalMessageInfo

func (m *QueryGetCodeRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

// QueryGetCodeResponse is the response type for the Query/GetCode RPC
// method.
type QueryGetCodeResponse struct {
	// Code defines the code at the requested index.
	Code *StoredCode `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (m *QueryGetCodeResponse) Reset()         { *m = QueryGetCodeResponse{} }
func (m *QueryGetCodeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCodeResponse) ProtoMessage()    {}
func (*QueryGetCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd51cf3aab79481, []int{1}
}
func (m *QueryGetCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCodeResponse.Merge(m, src)
}
func (m *QueryGetCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCodeResponse proto.InternalMessageInfo

func (m *QueryGetCodeResponse) GetCode() *StoredCode {
	if m != nil {
		return m.Code
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetCodeRequest)(nil), "igorsikachyna.dataproc.v1.QueryGetCodeRequest")
	proto.RegisterType((*QueryGetCodeResponse)(nil), "igorsikachyna.dataproc.v1.QueryGetCodeResponse")
}

func init() {
	proto.RegisterFile("igorsikachyna/dataproc/v1/query.proto", fileDescriptor_fcd51cf3aab79481)
}

var fileDescriptor_fcd51cf3aab79481 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x4c, 0xcf, 0x2f,
	0x2a, 0xce, 0xcc, 0x4e, 0x4c, 0xce, 0xa8, 0xcc, 0x4b, 0xd4, 0x4f, 0x49, 0x2c, 0x49, 0x2c, 0x28,
	0xca, 0x4f, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x44, 0x51, 0xa6, 0x07, 0x53, 0xa6, 0x57, 0x66, 0x28, 0x85, 0xc7, 0x84, 0x92,
	0xca, 0x82, 0xd4, 0x62, 0x88, 0x09, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89,
	0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x30, 0x59, 0xe9,
	0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0x88, 0x9d, 0x68, 0x96, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7,
	0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0x49, 0x9b, 0x4b, 0x38, 0x10, 0xa4, 0xc8, 0x3d, 0xb5,
	0xc4, 0x39, 0x3f, 0x25, 0x35, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x84, 0x8b, 0x35,
	0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x51, 0x0a, 0xe4,
	0x12, 0x41, 0x55, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xc9, 0xc5, 0x02, 0xe2, 0x83,
	0x15, 0x73, 0x1b, 0xa9, 0xea, 0xe1, 0xf4, 0xa6, 0x5e, 0x70, 0x49, 0x7e, 0x51, 0x6a, 0x0a, 0x58,
	0x33, 0x58, 0x8b, 0xd1, 0x2a, 0x46, 0x2e, 0x56, 0xb0, 0x99, 0x42, 0x0b, 0x18, 0xb9, 0xd8, 0xa1,
	0x06, 0x0b, 0xe9, 0xe1, 0x31, 0x02, 0x8b, 0x73, 0xa5, 0xf4, 0x89, 0x56, 0x0f, 0x71, 0xb1, 0x92,
	0x49, 0xc7, 0xf3, 0x0d, 0x5a, 0x8c, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0xd2, 0x14, 0x52, 0xd7, 0xc7,
	0x1d, 0xf8, 0xc9, 0xf9, 0x29, 0xa9, 0xfa, 0xd5, 0x60, 0xef, 0xd7, 0x3a, 0xd9, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78,
	0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x52, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0x2e, 0xd8, 0x30, 0x5d, 0x4c, 0xd3, 0x92, 0xd8, 0xc0, 0x21, 0x6e, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xb6, 0x62, 0x78, 0x94, 0x2d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// GetCode returns the code at the requested index.
	GetCode(ctx context.Context, in *QueryGetCodeRequest, opts ...grpc.CallOption) (*QueryGetCodeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetCode(ctx context.Context, in *QueryGetCodeRequest, opts ...grpc.CallOption) (*QueryGetCodeResponse, error) {
	out := new(QueryGetCodeResponse)
	err := c.cc.Invoke(ctx, "/igorsikachyna.dataproc.v1.Query/GetCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// GetCode returns the code at the requested index.
	GetCode(context.Context, *QueryGetCodeRequest) (*QueryGetCodeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetCode(ctx context.Context, req *QueryGetCodeRequest) (*QueryGetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/igorsikachyna.dataproc.v1.Query/GetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetCode(ctx, req.(*QueryGetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "igorsikachyna.dataproc.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCode",
			Handler:    _Query_GetCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "igorsikachyna/dataproc/v1/query.proto",
}

func (m *QueryGetCodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Code != nil {
		{
			size, err := m.Code.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != nil {
		l = m.Code.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetCodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Code == nil {
				m.Code = &StoredCode{}
			}
			if err := m.Code.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
