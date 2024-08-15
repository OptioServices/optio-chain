// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package tokenfactory

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_Denom                    protoreflect.MessageDescriptor
	fd_Denom_denom              protoreflect.FieldDescriptor
	fd_Denom_description        protoreflect.FieldDescriptor
	fd_Denom_ticker             protoreflect.FieldDescriptor
	fd_Denom_precision          protoreflect.FieldDescriptor
	fd_Denom_url                protoreflect.FieldDescriptor
	fd_Denom_maxSupply          protoreflect.FieldDescriptor
	fd_Denom_supply             protoreflect.FieldDescriptor
	fd_Denom_canChangeMaxSupply protoreflect.FieldDescriptor
	fd_Denom_owner              protoreflect.FieldDescriptor
)

func init() {
	file_optio_tokenfactory_denom_proto_init()
	md_Denom = File_optio_tokenfactory_denom_proto.Messages().ByName("Denom")
	fd_Denom_denom = md_Denom.Fields().ByName("denom")
	fd_Denom_description = md_Denom.Fields().ByName("description")
	fd_Denom_ticker = md_Denom.Fields().ByName("ticker")
	fd_Denom_precision = md_Denom.Fields().ByName("precision")
	fd_Denom_url = md_Denom.Fields().ByName("url")
	fd_Denom_maxSupply = md_Denom.Fields().ByName("maxSupply")
	fd_Denom_supply = md_Denom.Fields().ByName("supply")
	fd_Denom_canChangeMaxSupply = md_Denom.Fields().ByName("canChangeMaxSupply")
	fd_Denom_owner = md_Denom.Fields().ByName("owner")
}

var _ protoreflect.Message = (*fastReflection_Denom)(nil)

type fastReflection_Denom Denom

func (x *Denom) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Denom)(x)
}

func (x *Denom) slowProtoReflect() protoreflect.Message {
	mi := &file_optio_tokenfactory_denom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Denom_messageType fastReflection_Denom_messageType
var _ protoreflect.MessageType = fastReflection_Denom_messageType{}

type fastReflection_Denom_messageType struct{}

func (x fastReflection_Denom_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Denom)(nil)
}
func (x fastReflection_Denom_messageType) New() protoreflect.Message {
	return new(fastReflection_Denom)
}
func (x fastReflection_Denom_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Denom
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Denom) Descriptor() protoreflect.MessageDescriptor {
	return md_Denom
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Denom) Type() protoreflect.MessageType {
	return _fastReflection_Denom_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Denom) New() protoreflect.Message {
	return new(fastReflection_Denom)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Denom) Interface() protoreflect.ProtoMessage {
	return (*Denom)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Denom) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_Denom_denom, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_Denom_description, value) {
			return
		}
	}
	if x.Ticker != "" {
		value := protoreflect.ValueOfString(x.Ticker)
		if !f(fd_Denom_ticker, value) {
			return
		}
	}
	if x.Precision != int32(0) {
		value := protoreflect.ValueOfInt32(x.Precision)
		if !f(fd_Denom_precision, value) {
			return
		}
	}
	if x.Url != "" {
		value := protoreflect.ValueOfString(x.Url)
		if !f(fd_Denom_url, value) {
			return
		}
	}
	if x.MaxSupply != int32(0) {
		value := protoreflect.ValueOfInt32(x.MaxSupply)
		if !f(fd_Denom_maxSupply, value) {
			return
		}
	}
	if x.Supply != int32(0) {
		value := protoreflect.ValueOfInt32(x.Supply)
		if !f(fd_Denom_supply, value) {
			return
		}
	}
	if x.CanChangeMaxSupply != false {
		value := protoreflect.ValueOfBool(x.CanChangeMaxSupply)
		if !f(fd_Denom_canChangeMaxSupply, value) {
			return
		}
	}
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_Denom_owner, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Denom) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "optio.tokenfactory.Denom.denom":
		return x.Denom != ""
	case "optio.tokenfactory.Denom.description":
		return x.Description != ""
	case "optio.tokenfactory.Denom.ticker":
		return x.Ticker != ""
	case "optio.tokenfactory.Denom.precision":
		return x.Precision != int32(0)
	case "optio.tokenfactory.Denom.url":
		return x.Url != ""
	case "optio.tokenfactory.Denom.maxSupply":
		return x.MaxSupply != int32(0)
	case "optio.tokenfactory.Denom.supply":
		return x.Supply != int32(0)
	case "optio.tokenfactory.Denom.canChangeMaxSupply":
		return x.CanChangeMaxSupply != false
	case "optio.tokenfactory.Denom.owner":
		return x.Owner != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: optio.tokenfactory.Denom"))
		}
		panic(fmt.Errorf("message optio.tokenfactory.Denom does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Denom) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "optio.tokenfactory.Denom.denom":
		x.Denom = ""
	case "optio.tokenfactory.Denom.description":
		x.Description = ""
	case "optio.tokenfactory.Denom.ticker":
		x.Ticker = ""
	case "optio.tokenfactory.Denom.precision":
		x.Precision = int32(0)
	case "optio.tokenfactory.Denom.url":
		x.Url = ""
	case "optio.tokenfactory.Denom.maxSupply":
		x.MaxSupply = int32(0)
	case "optio.tokenfactory.Denom.supply":
		x.Supply = int32(0)
	case "optio.tokenfactory.Denom.canChangeMaxSupply":
		x.CanChangeMaxSupply = false
	case "optio.tokenfactory.Denom.owner":
		x.Owner = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: optio.tokenfactory.Denom"))
		}
		panic(fmt.Errorf("message optio.tokenfactory.Denom does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Denom) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "optio.tokenfactory.Denom.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "optio.tokenfactory.Denom.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	case "optio.tokenfactory.Denom.ticker":
		value := x.Ticker
		return protoreflect.ValueOfString(value)
	case "optio.tokenfactory.Denom.precision":
		value := x.Precision
		return protoreflect.ValueOfInt32(value)
	case "optio.tokenfactory.Denom.url":
		value := x.Url
		return protoreflect.ValueOfString(value)
	case "optio.tokenfactory.Denom.maxSupply":
		value := x.MaxSupply
		return protoreflect.ValueOfInt32(value)
	case "optio.tokenfactory.Denom.supply":
		value := x.Supply
		return protoreflect.ValueOfInt32(value)
	case "optio.tokenfactory.Denom.canChangeMaxSupply":
		value := x.CanChangeMaxSupply
		return protoreflect.ValueOfBool(value)
	case "optio.tokenfactory.Denom.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: optio.tokenfactory.Denom"))
		}
		panic(fmt.Errorf("message optio.tokenfactory.Denom does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Denom) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "optio.tokenfactory.Denom.denom":
		x.Denom = value.Interface().(string)
	case "optio.tokenfactory.Denom.description":
		x.Description = value.Interface().(string)
	case "optio.tokenfactory.Denom.ticker":
		x.Ticker = value.Interface().(string)
	case "optio.tokenfactory.Denom.precision":
		x.Precision = int32(value.Int())
	case "optio.tokenfactory.Denom.url":
		x.Url = value.Interface().(string)
	case "optio.tokenfactory.Denom.maxSupply":
		x.MaxSupply = int32(value.Int())
	case "optio.tokenfactory.Denom.supply":
		x.Supply = int32(value.Int())
	case "optio.tokenfactory.Denom.canChangeMaxSupply":
		x.CanChangeMaxSupply = value.Bool()
	case "optio.tokenfactory.Denom.owner":
		x.Owner = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: optio.tokenfactory.Denom"))
		}
		panic(fmt.Errorf("message optio.tokenfactory.Denom does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Denom) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "optio.tokenfactory.Denom.denom":
		panic(fmt.Errorf("field denom of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.description":
		panic(fmt.Errorf("field description of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.ticker":
		panic(fmt.Errorf("field ticker of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.precision":
		panic(fmt.Errorf("field precision of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.url":
		panic(fmt.Errorf("field url of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.maxSupply":
		panic(fmt.Errorf("field maxSupply of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.supply":
		panic(fmt.Errorf("field supply of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.canChangeMaxSupply":
		panic(fmt.Errorf("field canChangeMaxSupply of message optio.tokenfactory.Denom is not mutable"))
	case "optio.tokenfactory.Denom.owner":
		panic(fmt.Errorf("field owner of message optio.tokenfactory.Denom is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: optio.tokenfactory.Denom"))
		}
		panic(fmt.Errorf("message optio.tokenfactory.Denom does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Denom) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "optio.tokenfactory.Denom.denom":
		return protoreflect.ValueOfString("")
	case "optio.tokenfactory.Denom.description":
		return protoreflect.ValueOfString("")
	case "optio.tokenfactory.Denom.ticker":
		return protoreflect.ValueOfString("")
	case "optio.tokenfactory.Denom.precision":
		return protoreflect.ValueOfInt32(int32(0))
	case "optio.tokenfactory.Denom.url":
		return protoreflect.ValueOfString("")
	case "optio.tokenfactory.Denom.maxSupply":
		return protoreflect.ValueOfInt32(int32(0))
	case "optio.tokenfactory.Denom.supply":
		return protoreflect.ValueOfInt32(int32(0))
	case "optio.tokenfactory.Denom.canChangeMaxSupply":
		return protoreflect.ValueOfBool(false)
	case "optio.tokenfactory.Denom.owner":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: optio.tokenfactory.Denom"))
		}
		panic(fmt.Errorf("message optio.tokenfactory.Denom does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Denom) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in optio.tokenfactory.Denom", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Denom) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Denom) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Denom) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Denom) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Denom)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Description)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Ticker)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Precision != 0 {
			n += 1 + runtime.Sov(uint64(x.Precision))
		}
		l = len(x.Url)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxSupply != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxSupply))
		}
		if x.Supply != 0 {
			n += 1 + runtime.Sov(uint64(x.Supply))
		}
		if x.CanChangeMaxSupply {
			n += 2
		}
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Denom)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0x4a
		}
		if x.CanChangeMaxSupply {
			i--
			if x.CanChangeMaxSupply {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x40
		}
		if x.Supply != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Supply))
			i--
			dAtA[i] = 0x38
		}
		if x.MaxSupply != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxSupply))
			i--
			dAtA[i] = 0x30
		}
		if len(x.Url) > 0 {
			i -= len(x.Url)
			copy(dAtA[i:], x.Url)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Url)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Precision != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Precision))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Ticker) > 0 {
			i -= len(x.Ticker)
			copy(dAtA[i:], x.Ticker)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Ticker)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Description) > 0 {
			i -= len(x.Description)
			copy(dAtA[i:], x.Description)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Description)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Denom)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Denom: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Denom: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Description = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Ticker = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
				}
				x.Precision = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Precision |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Url = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
				}
				x.MaxSupply = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxSupply |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
				}
				x.Supply = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Supply |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 8:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CanChangeMaxSupply", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.CanChangeMaxSupply = bool(v != 0)
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: optio/tokenfactory/denom.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Denom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom              string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Description        string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Ticker             string `protobuf:"bytes,3,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Precision          int32  `protobuf:"varint,4,opt,name=precision,proto3" json:"precision,omitempty"`
	Url                string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	MaxSupply          int32  `protobuf:"varint,6,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	Supply             int32  `protobuf:"varint,7,opt,name=supply,proto3" json:"supply,omitempty"`
	CanChangeMaxSupply bool   `protobuf:"varint,8,opt,name=canChangeMaxSupply,proto3" json:"canChangeMaxSupply,omitempty"`
	Owner              string `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Denom) Reset() {
	*x = Denom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optio_tokenfactory_denom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Denom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Denom) ProtoMessage() {}

// Deprecated: Use Denom.ProtoReflect.Descriptor instead.
func (*Denom) Descriptor() ([]byte, []int) {
	return file_optio_tokenfactory_denom_proto_rawDescGZIP(), []int{0}
}

func (x *Denom) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Denom) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Denom) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *Denom) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *Denom) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Denom) GetMaxSupply() int32 {
	if x != nil {
		return x.MaxSupply
	}
	return 0
}

func (x *Denom) GetSupply() int32 {
	if x != nil {
		return x.Supply
	}
	return 0
}

func (x *Denom) GetCanChangeMaxSupply() bool {
	if x != nil {
		return x.CanChangeMaxSupply
	}
	return false
}

func (x *Denom) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

var File_optio_tokenfactory_denom_proto protoreflect.FileDescriptor

var file_optio_tokenfactory_denom_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x22, 0x83, 0x02, 0x0a, 0x05, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x63, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0xab, 0x01, 0x0a, 0x16, 0x63,
	0x6f, 0x6d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x0a, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0xa2, 0x02, 0x03, 0x4f, 0x54, 0x58, 0xaa, 0x02, 0x12, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0xca, 0x02, 0x12, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x5c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0xe2, 0x02, 0x1e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x5c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x13, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x3a, 0x3a, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_optio_tokenfactory_denom_proto_rawDescOnce sync.Once
	file_optio_tokenfactory_denom_proto_rawDescData = file_optio_tokenfactory_denom_proto_rawDesc
)

func file_optio_tokenfactory_denom_proto_rawDescGZIP() []byte {
	file_optio_tokenfactory_denom_proto_rawDescOnce.Do(func() {
		file_optio_tokenfactory_denom_proto_rawDescData = protoimpl.X.CompressGZIP(file_optio_tokenfactory_denom_proto_rawDescData)
	})
	return file_optio_tokenfactory_denom_proto_rawDescData
}

var file_optio_tokenfactory_denom_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_optio_tokenfactory_denom_proto_goTypes = []interface{}{
	(*Denom)(nil), // 0: optio.tokenfactory.Denom
}
var file_optio_tokenfactory_denom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_optio_tokenfactory_denom_proto_init() }
func file_optio_tokenfactory_denom_proto_init() {
	if File_optio_tokenfactory_denom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_optio_tokenfactory_denom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Denom); i {
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
			RawDescriptor: file_optio_tokenfactory_denom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_optio_tokenfactory_denom_proto_goTypes,
		DependencyIndexes: file_optio_tokenfactory_denom_proto_depIdxs,
		MessageInfos:      file_optio_tokenfactory_denom_proto_msgTypes,
	}.Build()
	File_optio_tokenfactory_denom_proto = out.File
	file_optio_tokenfactory_denom_proto_rawDesc = nil
	file_optio_tokenfactory_denom_proto_goTypes = nil
	file_optio_tokenfactory_denom_proto_depIdxs = nil
}
