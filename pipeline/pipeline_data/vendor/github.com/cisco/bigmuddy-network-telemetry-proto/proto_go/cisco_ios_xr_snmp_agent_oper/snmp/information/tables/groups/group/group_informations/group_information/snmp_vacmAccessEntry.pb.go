// Code generated by protoc-gen-go.
// source: snmp_vacmAccessEntry.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_tables_groups_group_group_informations_group_information is a generated protocol buffer package.

It is generated from these files:
	snmp_vacmAccessEntry.proto

It has these top-level messages:
	SnmpVacmAccessEntry_KEYS
	SnmpVacmAccessEntry
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_tables_groups_group_group_informations_group_information

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SNMP vacmAccessTable Information
type SnmpVacmAccessEntry_KEYS struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Modelnumber string `protobuf:"bytes,2,opt,name=modelnumber" json:"modelnumber,omitempty"`
	Level       string `protobuf:"bytes,3,opt,name=level" json:"level,omitempty"`
}

func (m *SnmpVacmAccessEntry_KEYS) Reset()                    { *m = SnmpVacmAccessEntry_KEYS{} }
func (m *SnmpVacmAccessEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpVacmAccessEntry_KEYS) ProtoMessage()               {}
func (*SnmpVacmAccessEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SnmpVacmAccessEntry_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SnmpVacmAccessEntry_KEYS) GetModelnumber() string {
	if m != nil {
		return m.Modelnumber
	}
	return ""
}

func (m *SnmpVacmAccessEntry_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type SnmpVacmAccessEntry struct {
	// Read view name
	VacmAccessReadViewName string `protobuf:"bytes,50,opt,name=vacm_access_read_view_name,json=vacmAccessReadViewName" json:"vacm_access_read_view_name,omitempty"`
	// Write view name
	VacmAccessWriteViewName string `protobuf:"bytes,51,opt,name=vacm_access_write_view_name,json=vacmAccessWriteViewName" json:"vacm_access_write_view_name,omitempty"`
	// Notify view name
	VacmAccessNotifyViewName string `protobuf:"bytes,52,opt,name=vacm_access_notify_view_name,json=vacmAccessNotifyViewName" json:"vacm_access_notify_view_name,omitempty"`
	// Status of this view configuration
	VacmAccessStatus uint32 `protobuf:"varint,53,opt,name=vacm_access_status,json=vacmAccessStatus" json:"vacm_access_status,omitempty"`
}

func (m *SnmpVacmAccessEntry) Reset()                    { *m = SnmpVacmAccessEntry{} }
func (m *SnmpVacmAccessEntry) String() string            { return proto.CompactTextString(m) }
func (*SnmpVacmAccessEntry) ProtoMessage()               {}
func (*SnmpVacmAccessEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpVacmAccessEntry) GetVacmAccessReadViewName() string {
	if m != nil {
		return m.VacmAccessReadViewName
	}
	return ""
}

func (m *SnmpVacmAccessEntry) GetVacmAccessWriteViewName() string {
	if m != nil {
		return m.VacmAccessWriteViewName
	}
	return ""
}

func (m *SnmpVacmAccessEntry) GetVacmAccessNotifyViewName() string {
	if m != nil {
		return m.VacmAccessNotifyViewName
	}
	return ""
}

func (m *SnmpVacmAccessEntry) GetVacmAccessStatus() uint32 {
	if m != nil {
		return m.VacmAccessStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpVacmAccessEntry_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.tables.groups.group.group_informations.group_information.snmp_vacmAccessEntry_KEYS")
	proto.RegisterType((*SnmpVacmAccessEntry)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.tables.groups.group.group_informations.group_information.snmp_vacmAccessEntry")
}

func init() { proto.RegisterFile("snmp_vacmAccessEntry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0x07, 0x70, 0xfa, 0xfb, 0xa9, 0x60, 0x44, 0x90, 0x30, 0xb4, 0x4e, 0x0f, 0x63, 0xa7, 0x1d,
	0xa4, 0x07, 0xa7, 0x17, 0x11, 0xc1, 0xc3, 0x4e, 0xc2, 0x0e, 0x1d, 0x28, 0x9e, 0x1e, 0x69, 0xfb,
	0x5a, 0x02, 0x4d, 0x52, 0x92, 0xb4, 0x75, 0xff, 0xbc, 0x48, 0xdf, 0x64, 0x0d, 0xb8, 0x4b, 0xdb,
	0x7c, 0xf3, 0xfd, 0x34, 0x0f, 0xc2, 0xa6, 0x4e, 0xab, 0x06, 0x3a, 0x91, 0xab, 0xd7, 0x3c, 0x47,
	0xe7, 0x56, 0xda, 0xdb, 0x6d, 0xd2, 0x58, 0xe3, 0x0d, 0x2f, 0x73, 0xe9, 0x72, 0x03, 0xd2, 0x38,
	0xf8, 0xb2, 0x40, 0x45, 0x51, 0xa1, 0xf6, 0x60, 0x1a, 0xb4, 0xc9, 0xb0, 0x4e, 0xa4, 0x2e, 0x8d,
	0x55, 0xc2, 0x4b, 0xa3, 0x13, 0x2f, 0xb2, 0x1a, 0x5d, 0x52, 0x59, 0xd3, 0x36, 0xbf, 0xaf, 0xdd,
	0x13, 0x82, 0x96, 0xfb, 0x1b, 0xcd, 0x2b, 0x76, 0x7d, 0x68, 0x0a, 0x78, 0x5b, 0x7d, 0x6e, 0x38,
	0x67, 0x47, 0x5a, 0x28, 0x8c, 0xa3, 0x59, 0xb4, 0x38, 0x4d, 0xe9, 0x9b, 0xcf, 0xd8, 0x99, 0x32,
	0x05, 0xd6, 0xba, 0x55, 0x19, 0xda, 0xf8, 0x1f, 0x6d, 0x85, 0x11, 0x9f, 0xb0, 0xe3, 0x1a, 0x3b,
	0xac, 0xe3, 0xff, 0xb4, 0xb7, 0x5b, 0xcc, 0xbf, 0x23, 0x36, 0x39, 0x74, 0x12, 0x7f, 0x62, 0xd3,
	0x21, 0x02, 0x41, 0x19, 0x58, 0x14, 0x05, 0x74, 0x12, 0x7b, 0xa0, 0xa3, 0xef, 0xe9, 0x1f, 0x97,
	0x23, 0x4a, 0x51, 0x14, 0xef, 0x12, 0xfb, 0xf5, 0x30, 0xcc, 0x33, 0xbb, 0x09, 0x6d, 0x6f, 0xa5,
	0xc7, 0x00, 0x2f, 0x09, 0x5f, 0x8d, 0xf8, 0x63, 0x28, 0xec, 0xf5, 0x0b, 0xbb, 0x0d, 0xb5, 0x36,
	0x5e, 0x96, 0xdb, 0x80, 0x3f, 0x10, 0x8f, 0x47, 0xbe, 0xa6, 0xc6, 0xde, 0xdf, 0x31, 0x1e, 0x7a,
	0xe7, 0x85, 0x6f, 0x5d, 0xfc, 0x38, 0x8b, 0x16, 0xe7, 0xe9, 0xc5, 0xa8, 0x36, 0x94, 0x67, 0x27,
	0x74, 0xb1, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x62, 0x34, 0x5a, 0xf6, 0x01, 0x00,
	0x00,
}
