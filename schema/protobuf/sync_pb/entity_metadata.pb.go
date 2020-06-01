// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity_metadata.proto

package sync_pb

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

// Sync proto to store entity metadata in model type storage.
type EntityMetadata struct {
	// A hash based on the client tag and model type.
	// Used for various map lookups. Should always be available.
	// Sent to the server as SyncEntity::client_defined_unique_tag.
	ClientTagHash *string `protobuf:"bytes,1,opt,name=client_tag_hash,json=clientTagHash" json:"client_tag_hash,omitempty"`
	// The entity's server-assigned ID.
	//
	// Prior to the item's first commit, we leave this value as an empty string.
	// The initial ID for a newly created item has to meet certain uniqueness
	// requirements, and we handle those on the sync thread.
	ServerId *string `protobuf:"bytes,2,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	// Whether or not the entity is deleted.
	IsDeleted *bool `protobuf:"varint,3,opt,name=is_deleted,json=isDeleted" json:"is_deleted,omitempty"`
	// A version number used to track in-progress commits. Each local change
	// increments this number.
	SequenceNumber *int64 `protobuf:"varint,4,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// The sequence number of the last item known to be successfully committed.
	AckedSequenceNumber *int64 `protobuf:"varint,5,opt,name=acked_sequence_number,json=ackedSequenceNumber" json:"acked_sequence_number,omitempty"`
	// The server version on which this item is based.
	//
	// If there are no local changes, this is the version of the entity as we see
	// it here.
	//
	// If there are local changes, this is the version of the entity on which
	// those changes are based.
	ServerVersion *int64 `protobuf:"varint,6,opt,name=server_version,json=serverVersion,def=-1" json:"server_version,omitempty"`
	// Entity creation and modification timestamps. Assigned by the client and
	// synced by the server, though the server usually doesn't bother to inspect
	// their values. They are encoded as milliseconds since the Unix epoch.
	CreationTime     *int64 `protobuf:"varint,7,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	ModificationTime *int64 `protobuf:"varint,8,opt,name=modification_time,json=modificationTime" json:"modification_time,omitempty"`
	// A hash of the current entity specifics value. Used to detect whether
	// entity's specifics value has changed without having to keep specifics in
	// memory.
	SpecificsHash *string `protobuf:"bytes,9,opt,name=specifics_hash,json=specificsHash" json:"specifics_hash,omitempty"`
	// A hash of the last specifics known by both the client and server. Used to
	// detect when local commits and remote updates are just for encryption. This
	// value will be the empty string only in the following cases: the entity is
	// in sync with the server, has never been synced, or is deleted.
	BaseSpecificsHash *string `protobuf:"bytes,10,opt,name=base_specifics_hash,json=baseSpecificsHash" json:"base_specifics_hash,omitempty"`
	// Used for positioning entities among their siblings. Relevant only for data
	// types that support positions (e.g bookmarks). Refer to its definition in
	// unique_position.proto for more information about its internal
	// representation.
	UniquePosition       *UniquePosition `protobuf:"bytes,11,opt,name=unique_position,json=uniquePosition" json:"unique_position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EntityMetadata) Reset()         { *m = EntityMetadata{} }
func (m *EntityMetadata) String() string { return proto.CompactTextString(m) }
func (*EntityMetadata) ProtoMessage()    {}
func (*EntityMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b14c157ddaae0a, []int{0}
}

func (m *EntityMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityMetadata.Unmarshal(m, b)
}
func (m *EntityMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityMetadata.Marshal(b, m, deterministic)
}
func (m *EntityMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityMetadata.Merge(m, src)
}
func (m *EntityMetadata) XXX_Size() int {
	return xxx_messageInfo_EntityMetadata.Size(m)
}
func (m *EntityMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_EntityMetadata proto.InternalMessageInfo

const Default_EntityMetadata_ServerVersion int64 = -1

func (m *EntityMetadata) GetClientTagHash() string {
	if m != nil && m.ClientTagHash != nil {
		return *m.ClientTagHash
	}
	return ""
}

func (m *EntityMetadata) GetServerId() string {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return ""
}

func (m *EntityMetadata) GetIsDeleted() bool {
	if m != nil && m.IsDeleted != nil {
		return *m.IsDeleted
	}
	return false
}

func (m *EntityMetadata) GetSequenceNumber() int64 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

func (m *EntityMetadata) GetAckedSequenceNumber() int64 {
	if m != nil && m.AckedSequenceNumber != nil {
		return *m.AckedSequenceNumber
	}
	return 0
}

func (m *EntityMetadata) GetServerVersion() int64 {
	if m != nil && m.ServerVersion != nil {
		return *m.ServerVersion
	}
	return Default_EntityMetadata_ServerVersion
}

func (m *EntityMetadata) GetCreationTime() int64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *EntityMetadata) GetModificationTime() int64 {
	if m != nil && m.ModificationTime != nil {
		return *m.ModificationTime
	}
	return 0
}

func (m *EntityMetadata) GetSpecificsHash() string {
	if m != nil && m.SpecificsHash != nil {
		return *m.SpecificsHash
	}
	return ""
}

func (m *EntityMetadata) GetBaseSpecificsHash() string {
	if m != nil && m.BaseSpecificsHash != nil {
		return *m.BaseSpecificsHash
	}
	return ""
}

func (m *EntityMetadata) GetUniquePosition() *UniquePosition {
	if m != nil {
		return m.UniquePosition
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityMetadata)(nil), "sync_pb.EntityMetadata")
}

func init() {
	proto.RegisterFile("entity_metadata.proto", fileDescriptor_86b14c157ddaae0a)
}

var fileDescriptor_86b14c157ddaae0a = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xd1, 0x8b, 0xd3, 0x30,
	0x18, 0xc0, 0xe9, 0x55, 0xbd, 0xf5, 0x3b, 0xd7, 0x79, 0x39, 0x86, 0x45, 0x11, 0x8a, 0x72, 0x5a,
	0x39, 0x2c, 0x78, 0x8f, 0x3e, 0x89, 0x28, 0xcc, 0x07, 0x65, 0x74, 0xd3, 0xd7, 0x90, 0xa5, 0x9f,
	0x6b, 0x70, 0x49, 0xba, 0x24, 0x1d, 0xec, 0x2f, 0xf6, 0xdf, 0x90, 0x26, 0xdd, 0xb1, 0xed, 0xb1,
	0xbf, 0xdf, 0xef, 0x0b, 0xe5, 0x4b, 0x60, 0x8a, 0xca, 0x09, 0xb7, 0xa7, 0x12, 0x1d, 0xab, 0x99,
	0x63, 0x65, 0x6b, 0xb4, 0xd3, 0xe4, 0xd2, 0xee, 0x15, 0xa7, 0xed, 0xea, 0xc5, 0xb4, 0x53, 0x62,
	0xdb, 0x21, 0x6d, 0xb5, 0x15, 0x4e, 0x68, 0x15, 0xfc, 0xeb, 0x7f, 0x31, 0xa4, 0xdf, 0xfc, 0xe4,
	0x8f, 0x61, 0x90, 0xbc, 0x85, 0x09, 0xdf, 0x08, 0x54, 0x8e, 0x3a, 0xb6, 0xa6, 0x0d, 0xb3, 0x4d,
	0x16, 0xe5, 0x51, 0x91, 0x54, 0xe3, 0x80, 0x97, 0x6c, 0x3d, 0x63, 0xb6, 0x21, 0x2f, 0x21, 0xb1,
	0x68, 0x76, 0x68, 0xa8, 0xa8, 0xb3, 0x0b, 0x5f, 0x8c, 0x02, 0xf8, 0x5e, 0x93, 0x57, 0x00, 0xc2,
	0xd2, 0x1a, 0x37, 0xe8, 0xb0, 0xce, 0xe2, 0x3c, 0x2a, 0x46, 0x55, 0x22, 0xec, 0xd7, 0x00, 0xc8,
	0x3b, 0x98, 0x58, 0xdc, 0x76, 0xa8, 0x38, 0x52, 0xd5, 0xc9, 0x15, 0x9a, 0xec, 0x51, 0x1e, 0x15,
	0x71, 0x95, 0x1e, 0xf0, 0x4f, 0x4f, 0xc9, 0x3d, 0x4c, 0x19, 0xff, 0x8b, 0x35, 0x3d, 0xcf, 0x1f,
	0xfb, 0xfc, 0xc6, 0xcb, 0xc5, 0xe9, 0xcc, 0x7b, 0x48, 0x87, 0x1f, 0xdb, 0xa1, 0xb1, 0x42, 0xab,
	0xec, 0x49, 0x1f, 0x7f, 0xba, 0xf8, 0xf0, 0xb1, 0x1a, 0x07, 0xf3, 0x3b, 0x08, 0xf2, 0x06, 0xc6,
	0xdc, 0x20, 0xeb, 0x17, 0x42, 0x9d, 0x90, 0x98, 0x5d, 0xfa, 0x63, 0x9f, 0x1e, 0xe0, 0x52, 0x48,
	0x24, 0x77, 0x70, 0x2d, 0x75, 0x2d, 0xfe, 0x08, 0x7e, 0x14, 0x8e, 0x7c, 0xf8, 0xec, 0x58, 0xf8,
	0xf8, 0x16, 0x52, 0xdb, 0x22, 0xef, 0xa1, 0x0d, 0xcb, 0x4b, 0xc2, 0xf2, 0x1e, 0xa8, 0x5f, 0x5e,
	0x09, 0x37, 0x2b, 0x66, 0x91, 0x9e, 0xb5, 0xe0, 0xdb, 0xeb, 0x5e, 0x2d, 0x4e, 0xfa, 0xcf, 0x30,
	0x39, 0xbb, 0xc0, 0xec, 0x2a, 0x8f, 0x8a, 0xab, 0xfb, 0xe7, 0xe5, 0x70, 0xc3, 0xe5, 0x2f, 0xef,
	0xe7, 0x83, 0xae, 0xd2, 0xee, 0xe4, 0xfb, 0xcb, 0x1d, 0xdc, 0x6a, 0xb3, 0x2e, 0x79, 0x63, 0xb4,
	0x14, 0x9d, 0x2c, 0xb9, 0x96, 0xad, 0x56, 0xa8, 0x9c, 0xf5, 0x27, 0x84, 0xf7, 0xc0, 0xf5, 0x66,
	0x16, 0xcf, 0xa3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xed, 0xe1, 0xca, 0x4e, 0x02, 0x00,
	0x00,
}