// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/explore.proto

package explore_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_go_proto "kythe.io/kythe/proto/common_go_proto"
import storage_go_proto "kythe.io/kythe/proto/storage_go_proto"
import xref_go_proto "kythe.io/kythe/proto/xref_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NodeData struct {
	Kind                 string                        `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Subkind              string                        `protobuf:"bytes,2,opt,name=subkind" json:"subkind,omitempty"`
	Locations            []*xref_go_proto.Location     `protobuf:"bytes,3,rep,name=locations" json:"locations,omitempty"`
	DefinitionAnchor     string                        `protobuf:"bytes,4,opt,name=definition_anchor,json=definitionAnchor" json:"definition_anchor,omitempty"`
	Code                 *common_go_proto.MarkedSource `protobuf:"bytes,5,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *NodeData) Reset()         { *m = NodeData{} }
func (m *NodeData) String() string { return proto.CompactTextString(m) }
func (*NodeData) ProtoMessage()    {}
func (*NodeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{0}
}
func (m *NodeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeData.Unmarshal(m, b)
}
func (m *NodeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeData.Marshal(b, m, deterministic)
}
func (dst *NodeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeData.Merge(dst, src)
}
func (m *NodeData) XXX_Size() int {
	return xxx_messageInfo_NodeData.Size(m)
}
func (m *NodeData) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeData.DiscardUnknown(m)
}

var xxx_messageInfo_NodeData proto.InternalMessageInfo

func (m *NodeData) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *NodeData) GetSubkind() string {
	if m != nil {
		return m.Subkind
	}
	return ""
}

func (m *NodeData) GetLocations() []*xref_go_proto.Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *NodeData) GetDefinitionAnchor() string {
	if m != nil {
		return m.DefinitionAnchor
	}
	return ""
}

func (m *NodeData) GetCode() *common_go_proto.MarkedSource {
	if m != nil {
		return m.Code
	}
	return nil
}

type GraphNode struct {
	NodeData             *NodeData `protobuf:"bytes,1,opt,name=node_data,json=nodeData" json:"node_data,omitempty"`
	Predecessors         *Tickets  `protobuf:"bytes,2,opt,name=predecessors" json:"predecessors,omitempty"`
	Successors           *Tickets  `protobuf:"bytes,3,opt,name=successors" json:"successors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GraphNode) Reset()         { *m = GraphNode{} }
func (m *GraphNode) String() string { return proto.CompactTextString(m) }
func (*GraphNode) ProtoMessage()    {}
func (*GraphNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{1}
}
func (m *GraphNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphNode.Unmarshal(m, b)
}
func (m *GraphNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphNode.Marshal(b, m, deterministic)
}
func (dst *GraphNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphNode.Merge(dst, src)
}
func (m *GraphNode) XXX_Size() int {
	return xxx_messageInfo_GraphNode.Size(m)
}
func (m *GraphNode) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphNode.DiscardUnknown(m)
}

var xxx_messageInfo_GraphNode proto.InternalMessageInfo

func (m *GraphNode) GetNodeData() *NodeData {
	if m != nil {
		return m.NodeData
	}
	return nil
}

func (m *GraphNode) GetPredecessors() *Tickets {
	if m != nil {
		return m.Predecessors
	}
	return nil
}

func (m *GraphNode) GetSuccessors() *Tickets {
	if m != nil {
		return m.Successors
	}
	return nil
}

type Graph struct {
	Nodes                map[string]*GraphNode `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Graph) Reset()         { *m = Graph{} }
func (m *Graph) String() string { return proto.CompactTextString(m) }
func (*Graph) ProtoMessage()    {}
func (*Graph) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{2}
}
func (m *Graph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Graph.Unmarshal(m, b)
}
func (m *Graph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Graph.Marshal(b, m, deterministic)
}
func (dst *Graph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Graph.Merge(dst, src)
}
func (m *Graph) XXX_Size() int {
	return xxx_messageInfo_Graph.Size(m)
}
func (m *Graph) XXX_DiscardUnknown() {
	xxx_messageInfo_Graph.DiscardUnknown(m)
}

var xxx_messageInfo_Graph proto.InternalMessageInfo

func (m *Graph) GetNodes() map[string]*GraphNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NodeFilter struct {
	IncludedLanguages    []string                  `protobuf:"bytes,1,rep,name=included_languages,json=includedLanguages" json:"included_languages,omitempty"`
	IncludedFiles        []*storage_go_proto.VName `protobuf:"bytes,2,rep,name=included_files,json=includedFiles" json:"included_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *NodeFilter) Reset()         { *m = NodeFilter{} }
func (m *NodeFilter) String() string { return proto.CompactTextString(m) }
func (*NodeFilter) ProtoMessage()    {}
func (*NodeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{3}
}
func (m *NodeFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeFilter.Unmarshal(m, b)
}
func (m *NodeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeFilter.Marshal(b, m, deterministic)
}
func (dst *NodeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeFilter.Merge(dst, src)
}
func (m *NodeFilter) XXX_Size() int {
	return xxx_messageInfo_NodeFilter.Size(m)
}
func (m *NodeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_NodeFilter proto.InternalMessageInfo

func (m *NodeFilter) GetIncludedLanguages() []string {
	if m != nil {
		return m.IncludedLanguages
	}
	return nil
}

func (m *NodeFilter) GetIncludedFiles() []*storage_go_proto.VName {
	if m != nil {
		return m.IncludedFiles
	}
	return nil
}

type Tickets struct {
	Tickets              []string `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tickets) Reset()         { *m = Tickets{} }
func (m *Tickets) String() string { return proto.CompactTextString(m) }
func (*Tickets) ProtoMessage()    {}
func (*Tickets) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{4}
}
func (m *Tickets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tickets.Unmarshal(m, b)
}
func (m *Tickets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tickets.Marshal(b, m, deterministic)
}
func (dst *Tickets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tickets.Merge(dst, src)
}
func (m *Tickets) XXX_Size() int {
	return xxx_messageInfo_Tickets.Size(m)
}
func (m *Tickets) XXX_DiscardUnknown() {
	xxx_messageInfo_Tickets.DiscardUnknown(m)
}

var xxx_messageInfo_Tickets proto.InternalMessageInfo

func (m *Tickets) GetTickets() []string {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type TypeHierarchyRequest struct {
	TypeTicket           string      `protobuf:"bytes,1,opt,name=type_ticket,json=typeTicket" json:"type_ticket,omitempty"`
	NodeFilter           *NodeFilter `protobuf:"bytes,2,opt,name=node_filter,json=nodeFilter" json:"node_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TypeHierarchyRequest) Reset()         { *m = TypeHierarchyRequest{} }
func (m *TypeHierarchyRequest) String() string { return proto.CompactTextString(m) }
func (*TypeHierarchyRequest) ProtoMessage()    {}
func (*TypeHierarchyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{5}
}
func (m *TypeHierarchyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeHierarchyRequest.Unmarshal(m, b)
}
func (m *TypeHierarchyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeHierarchyRequest.Marshal(b, m, deterministic)
}
func (dst *TypeHierarchyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeHierarchyRequest.Merge(dst, src)
}
func (m *TypeHierarchyRequest) XXX_Size() int {
	return xxx_messageInfo_TypeHierarchyRequest.Size(m)
}
func (m *TypeHierarchyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeHierarchyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TypeHierarchyRequest proto.InternalMessageInfo

func (m *TypeHierarchyRequest) GetTypeTicket() string {
	if m != nil {
		return m.TypeTicket
	}
	return ""
}

func (m *TypeHierarchyRequest) GetNodeFilter() *NodeFilter {
	if m != nil {
		return m.NodeFilter
	}
	return nil
}

type TypeHierarchyReply struct {
	TypeTicket           string   `protobuf:"bytes,1,opt,name=type_ticket,json=typeTicket" json:"type_ticket,omitempty"`
	Graph                *Graph   `protobuf:"bytes,2,opt,name=graph" json:"graph,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeHierarchyReply) Reset()         { *m = TypeHierarchyReply{} }
func (m *TypeHierarchyReply) String() string { return proto.CompactTextString(m) }
func (*TypeHierarchyReply) ProtoMessage()    {}
func (*TypeHierarchyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{6}
}
func (m *TypeHierarchyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeHierarchyReply.Unmarshal(m, b)
}
func (m *TypeHierarchyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeHierarchyReply.Marshal(b, m, deterministic)
}
func (dst *TypeHierarchyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeHierarchyReply.Merge(dst, src)
}
func (m *TypeHierarchyReply) XXX_Size() int {
	return xxx_messageInfo_TypeHierarchyReply.Size(m)
}
func (m *TypeHierarchyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeHierarchyReply.DiscardUnknown(m)
}

var xxx_messageInfo_TypeHierarchyReply proto.InternalMessageInfo

func (m *TypeHierarchyReply) GetTypeTicket() string {
	if m != nil {
		return m.TypeTicket
	}
	return ""
}

func (m *TypeHierarchyReply) GetGraph() *Graph {
	if m != nil {
		return m.Graph
	}
	return nil
}

type CallersRequest struct {
	FunctionTicket       string      `protobuf:"bytes,1,opt,name=function_ticket,json=functionTicket" json:"function_ticket,omitempty"`
	NodeFilter           *NodeFilter `protobuf:"bytes,2,opt,name=node_filter,json=nodeFilter" json:"node_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CallersRequest) Reset()         { *m = CallersRequest{} }
func (m *CallersRequest) String() string { return proto.CompactTextString(m) }
func (*CallersRequest) ProtoMessage()    {}
func (*CallersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{7}
}
func (m *CallersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallersRequest.Unmarshal(m, b)
}
func (m *CallersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallersRequest.Marshal(b, m, deterministic)
}
func (dst *CallersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallersRequest.Merge(dst, src)
}
func (m *CallersRequest) XXX_Size() int {
	return xxx_messageInfo_CallersRequest.Size(m)
}
func (m *CallersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallersRequest proto.InternalMessageInfo

func (m *CallersRequest) GetFunctionTicket() string {
	if m != nil {
		return m.FunctionTicket
	}
	return ""
}

func (m *CallersRequest) GetNodeFilter() *NodeFilter {
	if m != nil {
		return m.NodeFilter
	}
	return nil
}

type CallersReply struct {
	FunctionTicket       string   `protobuf:"bytes,1,opt,name=function_ticket,json=functionTicket" json:"function_ticket,omitempty"`
	Graph                *Graph   `protobuf:"bytes,2,opt,name=graph" json:"graph,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallersReply) Reset()         { *m = CallersReply{} }
func (m *CallersReply) String() string { return proto.CompactTextString(m) }
func (*CallersReply) ProtoMessage()    {}
func (*CallersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{8}
}
func (m *CallersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallersReply.Unmarshal(m, b)
}
func (m *CallersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallersReply.Marshal(b, m, deterministic)
}
func (dst *CallersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallersReply.Merge(dst, src)
}
func (m *CallersReply) XXX_Size() int {
	return xxx_messageInfo_CallersReply.Size(m)
}
func (m *CallersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CallersReply.DiscardUnknown(m)
}

var xxx_messageInfo_CallersReply proto.InternalMessageInfo

func (m *CallersReply) GetFunctionTicket() string {
	if m != nil {
		return m.FunctionTicket
	}
	return ""
}

func (m *CallersReply) GetGraph() *Graph {
	if m != nil {
		return m.Graph
	}
	return nil
}

type CalleesRequest struct {
	FunctionTicket       string      `protobuf:"bytes,1,opt,name=function_ticket,json=functionTicket" json:"function_ticket,omitempty"`
	NodeFilter           *NodeFilter `protobuf:"bytes,2,opt,name=node_filter,json=nodeFilter" json:"node_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CalleesRequest) Reset()         { *m = CalleesRequest{} }
func (m *CalleesRequest) String() string { return proto.CompactTextString(m) }
func (*CalleesRequest) ProtoMessage()    {}
func (*CalleesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{9}
}
func (m *CalleesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalleesRequest.Unmarshal(m, b)
}
func (m *CalleesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalleesRequest.Marshal(b, m, deterministic)
}
func (dst *CalleesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalleesRequest.Merge(dst, src)
}
func (m *CalleesRequest) XXX_Size() int {
	return xxx_messageInfo_CalleesRequest.Size(m)
}
func (m *CalleesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalleesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalleesRequest proto.InternalMessageInfo

func (m *CalleesRequest) GetFunctionTicket() string {
	if m != nil {
		return m.FunctionTicket
	}
	return ""
}

func (m *CalleesRequest) GetNodeFilter() *NodeFilter {
	if m != nil {
		return m.NodeFilter
	}
	return nil
}

type CalleesReply struct {
	FunctionTicket       string   `protobuf:"bytes,1,opt,name=function_ticket,json=functionTicket" json:"function_ticket,omitempty"`
	Graph                *Graph   `protobuf:"bytes,2,opt,name=graph" json:"graph,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalleesReply) Reset()         { *m = CalleesReply{} }
func (m *CalleesReply) String() string { return proto.CompactTextString(m) }
func (*CalleesReply) ProtoMessage()    {}
func (*CalleesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{10}
}
func (m *CalleesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalleesReply.Unmarshal(m, b)
}
func (m *CalleesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalleesReply.Marshal(b, m, deterministic)
}
func (dst *CalleesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalleesReply.Merge(dst, src)
}
func (m *CalleesReply) XXX_Size() int {
	return xxx_messageInfo_CalleesReply.Size(m)
}
func (m *CalleesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CalleesReply.DiscardUnknown(m)
}

var xxx_messageInfo_CalleesReply proto.InternalMessageInfo

func (m *CalleesReply) GetFunctionTicket() string {
	if m != nil {
		return m.FunctionTicket
	}
	return ""
}

func (m *CalleesReply) GetGraph() *Graph {
	if m != nil {
		return m.Graph
	}
	return nil
}

type ParametersRequest struct {
	FunctionTickets      []string `protobuf:"bytes,1,rep,name=function_tickets,json=functionTickets" json:"function_tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParametersRequest) Reset()         { *m = ParametersRequest{} }
func (m *ParametersRequest) String() string { return proto.CompactTextString(m) }
func (*ParametersRequest) ProtoMessage()    {}
func (*ParametersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{11}
}
func (m *ParametersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParametersRequest.Unmarshal(m, b)
}
func (m *ParametersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParametersRequest.Marshal(b, m, deterministic)
}
func (dst *ParametersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParametersRequest.Merge(dst, src)
}
func (m *ParametersRequest) XXX_Size() int {
	return xxx_messageInfo_ParametersRequest.Size(m)
}
func (m *ParametersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParametersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParametersRequest proto.InternalMessageInfo

func (m *ParametersRequest) GetFunctionTickets() []string {
	if m != nil {
		return m.FunctionTickets
	}
	return nil
}

type ParametersReply struct {
	FunctionToParameters  map[string]*Tickets  `protobuf:"bytes,1,rep,name=function_to_parameters,json=functionToParameters" json:"function_to_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FunctionToReturnValue map[string]string    `protobuf:"bytes,2,rep,name=function_to_return_value,json=functionToReturnValue" json:"function_to_return_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeData              map[string]*NodeData `protobuf:"bytes,3,rep,name=node_data,json=nodeData" json:"node_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *ParametersReply) Reset()         { *m = ParametersReply{} }
func (m *ParametersReply) String() string { return proto.CompactTextString(m) }
func (*ParametersReply) ProtoMessage()    {}
func (*ParametersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{12}
}
func (m *ParametersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParametersReply.Unmarshal(m, b)
}
func (m *ParametersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParametersReply.Marshal(b, m, deterministic)
}
func (dst *ParametersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParametersReply.Merge(dst, src)
}
func (m *ParametersReply) XXX_Size() int {
	return xxx_messageInfo_ParametersReply.Size(m)
}
func (m *ParametersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ParametersReply.DiscardUnknown(m)
}

var xxx_messageInfo_ParametersReply proto.InternalMessageInfo

func (m *ParametersReply) GetFunctionToParameters() map[string]*Tickets {
	if m != nil {
		return m.FunctionToParameters
	}
	return nil
}

func (m *ParametersReply) GetFunctionToReturnValue() map[string]string {
	if m != nil {
		return m.FunctionToReturnValue
	}
	return nil
}

func (m *ParametersReply) GetNodeData() map[string]*NodeData {
	if m != nil {
		return m.NodeData
	}
	return nil
}

type ParentsRequest struct {
	Tickets              []string `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParentsRequest) Reset()         { *m = ParentsRequest{} }
func (m *ParentsRequest) String() string { return proto.CompactTextString(m) }
func (*ParentsRequest) ProtoMessage()    {}
func (*ParentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{13}
}
func (m *ParentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentsRequest.Unmarshal(m, b)
}
func (m *ParentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentsRequest.Marshal(b, m, deterministic)
}
func (dst *ParentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentsRequest.Merge(dst, src)
}
func (m *ParentsRequest) XXX_Size() int {
	return xxx_messageInfo_ParentsRequest.Size(m)
}
func (m *ParentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParentsRequest proto.InternalMessageInfo

func (m *ParentsRequest) GetTickets() []string {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type ParentsReply struct {
	InputToParents       map[string]*Tickets  `protobuf:"bytes,1,rep,name=input_to_parents,json=inputToParents" json:"input_to_parents,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeData             map[string]*NodeData `protobuf:"bytes,2,rep,name=node_data,json=nodeData" json:"node_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ParentsReply) Reset()         { *m = ParentsReply{} }
func (m *ParentsReply) String() string { return proto.CompactTextString(m) }
func (*ParentsReply) ProtoMessage()    {}
func (*ParentsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{14}
}
func (m *ParentsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentsReply.Unmarshal(m, b)
}
func (m *ParentsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentsReply.Marshal(b, m, deterministic)
}
func (dst *ParentsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentsReply.Merge(dst, src)
}
func (m *ParentsReply) XXX_Size() int {
	return xxx_messageInfo_ParentsReply.Size(m)
}
func (m *ParentsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ParentsReply proto.InternalMessageInfo

func (m *ParentsReply) GetInputToParents() map[string]*Tickets {
	if m != nil {
		return m.InputToParents
	}
	return nil
}

func (m *ParentsReply) GetNodeData() map[string]*NodeData {
	if m != nil {
		return m.NodeData
	}
	return nil
}

type ChildrenRequest struct {
	Tickets              []string `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChildrenRequest) Reset()         { *m = ChildrenRequest{} }
func (m *ChildrenRequest) String() string { return proto.CompactTextString(m) }
func (*ChildrenRequest) ProtoMessage()    {}
func (*ChildrenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{15}
}
func (m *ChildrenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChildrenRequest.Unmarshal(m, b)
}
func (m *ChildrenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChildrenRequest.Marshal(b, m, deterministic)
}
func (dst *ChildrenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChildrenRequest.Merge(dst, src)
}
func (m *ChildrenRequest) XXX_Size() int {
	return xxx_messageInfo_ChildrenRequest.Size(m)
}
func (m *ChildrenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChildrenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChildrenRequest proto.InternalMessageInfo

func (m *ChildrenRequest) GetTickets() []string {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type ChildrenReply struct {
	InputToChildren      map[string]*Tickets  `protobuf:"bytes,1,rep,name=input_to_children,json=inputToChildren" json:"input_to_children,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeData             map[string]*NodeData `protobuf:"bytes,2,rep,name=node_data,json=nodeData" json:"node_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChildrenReply) Reset()         { *m = ChildrenReply{} }
func (m *ChildrenReply) String() string { return proto.CompactTextString(m) }
func (*ChildrenReply) ProtoMessage()    {}
func (*ChildrenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_explore_3783c8ec273b81e8, []int{16}
}
func (m *ChildrenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChildrenReply.Unmarshal(m, b)
}
func (m *ChildrenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChildrenReply.Marshal(b, m, deterministic)
}
func (dst *ChildrenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChildrenReply.Merge(dst, src)
}
func (m *ChildrenReply) XXX_Size() int {
	return xxx_messageInfo_ChildrenReply.Size(m)
}
func (m *ChildrenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChildrenReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChildrenReply proto.InternalMessageInfo

func (m *ChildrenReply) GetInputToChildren() map[string]*Tickets {
	if m != nil {
		return m.InputToChildren
	}
	return nil
}

func (m *ChildrenReply) GetNodeData() map[string]*NodeData {
	if m != nil {
		return m.NodeData
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeData)(nil), "kythe.proto.NodeData")
	proto.RegisterType((*GraphNode)(nil), "kythe.proto.GraphNode")
	proto.RegisterType((*Graph)(nil), "kythe.proto.Graph")
	proto.RegisterMapType((map[string]*GraphNode)(nil), "kythe.proto.Graph.NodesEntry")
	proto.RegisterType((*NodeFilter)(nil), "kythe.proto.NodeFilter")
	proto.RegisterType((*Tickets)(nil), "kythe.proto.Tickets")
	proto.RegisterType((*TypeHierarchyRequest)(nil), "kythe.proto.TypeHierarchyRequest")
	proto.RegisterType((*TypeHierarchyReply)(nil), "kythe.proto.TypeHierarchyReply")
	proto.RegisterType((*CallersRequest)(nil), "kythe.proto.CallersRequest")
	proto.RegisterType((*CallersReply)(nil), "kythe.proto.CallersReply")
	proto.RegisterType((*CalleesRequest)(nil), "kythe.proto.CalleesRequest")
	proto.RegisterType((*CalleesReply)(nil), "kythe.proto.CalleesReply")
	proto.RegisterType((*ParametersRequest)(nil), "kythe.proto.ParametersRequest")
	proto.RegisterType((*ParametersReply)(nil), "kythe.proto.ParametersReply")
	proto.RegisterMapType((map[string]*Tickets)(nil), "kythe.proto.ParametersReply.FunctionToParametersEntry")
	proto.RegisterMapType((map[string]string)(nil), "kythe.proto.ParametersReply.FunctionToReturnValueEntry")
	proto.RegisterMapType((map[string]*NodeData)(nil), "kythe.proto.ParametersReply.NodeDataEntry")
	proto.RegisterType((*ParentsRequest)(nil), "kythe.proto.ParentsRequest")
	proto.RegisterType((*ParentsReply)(nil), "kythe.proto.ParentsReply")
	proto.RegisterMapType((map[string]*Tickets)(nil), "kythe.proto.ParentsReply.InputToParentsEntry")
	proto.RegisterMapType((map[string]*NodeData)(nil), "kythe.proto.ParentsReply.NodeDataEntry")
	proto.RegisterType((*ChildrenRequest)(nil), "kythe.proto.ChildrenRequest")
	proto.RegisterType((*ChildrenReply)(nil), "kythe.proto.ChildrenReply")
	proto.RegisterMapType((map[string]*Tickets)(nil), "kythe.proto.ChildrenReply.InputToChildrenEntry")
	proto.RegisterMapType((map[string]*NodeData)(nil), "kythe.proto.ChildrenReply.NodeDataEntry")
}

func init() { proto.RegisterFile("kythe/proto/explore.proto", fileDescriptor_explore_3783c8ec273b81e8) }

var fileDescriptor_explore_3783c8ec273b81e8 = []byte{
	// 987 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x5f, 0x6f, 0xdb, 0x54,
	0x14, 0xaf, 0x9b, 0x86, 0x36, 0x27, 0x6d, 0xd2, 0x5e, 0xb2, 0x92, 0x66, 0x83, 0x16, 0xef, 0x61,
	0xa1, 0x65, 0xa9, 0x94, 0x4e, 0x50, 0x78, 0x40, 0x82, 0xae, 0xdd, 0x90, 0xca, 0x54, 0x79, 0x63,
	0x43, 0x42, 0x28, 0xba, 0xb3, 0x4f, 0x12, 0x2b, 0x8e, 0xaf, 0x77, 0x6d, 0x57, 0xcb, 0x97, 0xe0,
	0xab, 0xf0, 0xc8, 0x97, 0xe0, 0x81, 0x67, 0xde, 0xf8, 0x26, 0xe8, 0xde, 0x6b, 0x3b, 0xbe, 0xa9,
	0x9d, 0x0d, 0xc1, 0x04, 0x6f, 0xf7, 0x9e, 0x3f, 0xbf, 0xdf, 0xf9, 0xe7, 0x73, 0x0d, 0x7b, 0x93,
	0x59, 0x34, 0xc6, 0xe3, 0x80, 0xb3, 0x88, 0x1d, 0xe3, 0xeb, 0xc0, 0x63, 0x1c, 0x7b, 0xf2, 0x46,
	0xea, 0x52, 0xa5, 0x2e, 0x9d, 0x76, 0xde, 0xce, 0x66, 0xd3, 0x29, 0xf3, 0x13, 0x8d, 0x86, 0x10,
	0x46, 0x8c, 0xd3, 0x51, 0xea, 0xb4, 0x9b, 0x57, 0xbd, 0xe6, 0x38, 0x54, 0x72, 0xf3, 0x77, 0x03,
	0x36, 0x9e, 0x30, 0x07, 0x1f, 0xd2, 0x88, 0x12, 0x02, 0x6b, 0x13, 0xd7, 0x77, 0xda, 0xc6, 0x81,
	0xd1, 0xad, 0x59, 0xf2, 0x4c, 0xda, 0xb0, 0x1e, 0xc6, 0x2f, 0xa5, 0x78, 0x55, 0x8a, 0xd3, 0x2b,
	0x39, 0x81, 0x9a, 0xc7, 0x6c, 0x1a, 0xb9, 0xcc, 0x0f, 0xdb, 0x95, 0x83, 0x4a, 0xb7, 0xde, 0xbf,
	0xd5, 0xcb, 0x05, 0xda, 0xbb, 0x4c, 0xb4, 0xd6, 0xdc, 0x8e, 0x1c, 0xc1, 0x8e, 0x83, 0x43, 0xd7,
	0x77, 0xc5, 0x75, 0x40, 0x7d, 0x7b, 0xcc, 0x78, 0x7b, 0x4d, 0x02, 0x6f, 0xcf, 0x15, 0x5f, 0x4b,
	0x39, 0x79, 0x00, 0x6b, 0x36, 0x73, 0xb0, 0x5d, 0x3d, 0x30, 0xba, 0xf5, 0xfe, 0x81, 0x06, 0x9e,
	0x24, 0xfe, 0x1d, 0xe5, 0x13, 0x74, 0x9e, 0xb2, 0x98, 0xdb, 0x68, 0x49, 0x6b, 0xf3, 0x17, 0x03,
	0x6a, 0x8f, 0x38, 0x0d, 0xc6, 0x22, 0x2f, 0xd2, 0x87, 0x9a, 0xcf, 0x1c, 0x1c, 0x38, 0x34, 0xa2,
	0x32, 0xb1, 0xc5, 0x28, 0xd3, 0xec, 0xad, 0x0d, 0x3f, 0xad, 0xc3, 0x29, 0x6c, 0x06, 0x1c, 0x1d,
	0xb4, 0x31, 0x0c, 0x19, 0x0f, 0x65, 0xe2, 0xf5, 0x7e, 0x4b, 0x73, 0x7b, 0xe6, 0xda, 0x13, 0x8c,
	0x42, 0x4b, 0xb3, 0x24, 0x0f, 0x00, 0xc2, 0xd8, 0x4e, 0xfd, 0x2a, 0x4b, 0xfc, 0x72, 0x76, 0xe6,
	0xcf, 0x06, 0x54, 0x65, 0xc4, 0xe4, 0x04, 0xaa, 0x22, 0x8a, 0xb0, 0x6d, 0xc8, 0x7a, 0x7e, 0xa8,
	0xb9, 0x4a, 0x13, 0x19, 0x6f, 0x78, 0xee, 0x47, 0x7c, 0x66, 0x29, 0xdb, 0xce, 0x15, 0xc0, 0x5c,
	0x48, 0xb6, 0xa1, 0x32, 0xc1, 0x59, 0xd2, 0x43, 0x71, 0x24, 0x9f, 0x42, 0xf5, 0x9a, 0x7a, 0x31,
	0x26, 0x79, 0xec, 0xde, 0x04, 0x15, 0xee, 0x96, 0x32, 0xfa, 0x72, 0xf5, 0xd4, 0x30, 0xaf, 0x15,
	0xe2, 0x85, 0xeb, 0x45, 0xc8, 0xc9, 0x7d, 0x20, 0xae, 0x6f, 0x7b, 0xb1, 0x83, 0xce, 0xc0, 0xa3,
	0xfe, 0x28, 0xa6, 0xa3, 0x24, 0xc2, 0x9a, 0xb5, 0x93, 0x6a, 0x2e, 0x53, 0x05, 0xf9, 0x02, 0x1a,
	0x99, 0xf9, 0xd0, 0xf5, 0x50, 0xd4, 0x4f, 0x24, 0x43, 0x34, 0xde, 0xe7, 0x4f, 0xe8, 0x14, 0xad,
	0xad, 0xd4, 0xf2, 0x42, 0x18, 0x9a, 0x77, 0x61, 0x3d, 0xa9, 0x8f, 0x98, 0xbb, 0x48, 0x1d, 0x13,
	0xa6, 0xf4, 0x6a, 0xbe, 0x82, 0xd6, 0xb3, 0x59, 0x80, 0x8f, 0x5d, 0xe4, 0x94, 0xdb, 0xe3, 0x99,
	0x85, 0xaf, 0x62, 0x0c, 0x23, 0xb2, 0x0f, 0xf5, 0x68, 0x16, 0xe0, 0x40, 0xd9, 0x25, 0x05, 0x00,
	0x21, 0x52, 0x98, 0xe4, 0x14, 0xea, 0x72, 0x14, 0x86, 0x32, 0xad, 0xa4, 0x1a, 0x1f, 0xdc, 0x18,
	0x06, 0x95, 0xb5, 0x05, 0x7e, 0x76, 0x36, 0x07, 0x40, 0x16, 0x28, 0x03, 0x6f, 0xf6, 0x66, 0xc2,
	0x2e, 0x54, 0x47, 0xa2, 0xbc, 0x09, 0x15, 0xb9, 0x59, 0x78, 0x4b, 0x19, 0x98, 0x21, 0x34, 0xce,
	0xa8, 0xe7, 0x21, 0x0f, 0xd3, 0x6c, 0xee, 0x41, 0x73, 0x18, 0xfb, 0xb6, 0xfc, 0x4c, 0x34, 0x82,
	0x46, 0x2a, 0xfe, 0xc7, 0x59, 0x51, 0xd8, 0xcc, 0x48, 0x45, 0x3e, 0x6f, 0x4d, 0xf9, 0xf7, 0xf3,
	0xc2, 0xff, 0x22, 0x2f, 0x7c, 0x77, 0x79, 0x7d, 0x05, 0x3b, 0x57, 0x94, 0xd3, 0x29, 0x46, 0xb9,
	0x96, 0x7d, 0x02, 0xdb, 0x0b, 0x3c, 0xe9, 0xec, 0x36, 0x75, 0xa2, 0xd0, 0xfc, 0x75, 0x0d, 0x9a,
	0x79, 0x00, 0x11, 0xa6, 0x07, 0xbb, 0x73, 0x77, 0x36, 0x08, 0x32, 0x75, 0xb2, 0x0c, 0x3e, 0xd3,
	0xc2, 0x59, 0xf0, 0xee, 0x5d, 0xa4, 0x0c, 0x6c, 0xae, 0x51, 0x5b, 0xa2, 0x35, 0x2c, 0x50, 0x91,
	0x00, 0xda, 0x79, 0x36, 0x8e, 0x51, 0xcc, 0xfd, 0x41, 0xba, 0x27, 0x04, 0xdf, 0xe7, 0x6f, 0xc9,
	0x67, 0x49, 0xd7, 0xe7, 0xc2, 0x53, 0x11, 0xde, 0x1a, 0x16, 0xe9, 0xc8, 0xa3, 0xfc, 0x26, 0x56,
	0xef, 0xc5, 0xe1, 0x52, 0x8a, 0x74, 0x33, 0x2b, 0xd4, 0x6c, 0x3d, 0x77, 0x7e, 0x82, 0xbd, 0xd2,
	0x6c, 0x0b, 0xd6, 0xdf, 0xa1, 0xbe, 0xfe, 0x8a, 0xd7, 0xf1, 0x7c, 0xf9, 0x75, 0x1e, 0x43, 0xa7,
	0x3c, 0xb9, 0x02, 0xfc, 0x56, 0x1e, 0xbf, 0x96, 0x47, 0xb2, 0x60, 0x4b, 0xcb, 0xa1, 0xc0, 0xf9,
	0x48, 0x0f, 0xae, 0xe4, 0x69, 0xca, 0xad, 0xe6, 0x43, 0x68, 0x5c, 0x51, 0x8e, 0x7e, 0x94, 0x8d,
	0x5d, 0xf9, 0xa6, 0xfc, 0x63, 0x15, 0x36, 0x33, 0x63, 0x31, 0x62, 0x2f, 0x60, 0xdb, 0xf5, 0x83,
	0x38, 0x4a, 0xe6, 0x4b, 0x28, 0x92, 0xe1, 0xba, 0xbf, 0xd8, 0x89, 0xcc, 0xa9, 0xf7, 0xad, 0xf0,
	0x90, 0x85, 0x16, 0x32, 0xd5, 0x8c, 0x86, 0xab, 0x09, 0xc9, 0xc3, 0x7c, 0x6f, 0xd5, 0xf8, 0xdc,
	0x2b, 0x47, 0x2c, 0x6b, 0xec, 0x0b, 0x78, 0xbf, 0x80, 0xec, 0x5f, 0x68, 0xe9, 0xbb, 0x68, 0xc4,
	0x11, 0x34, 0xcf, 0xc6, 0xae, 0xe7, 0x70, 0xf4, 0xdf, 0xdc, 0x89, 0x3f, 0x57, 0x61, 0x6b, 0x6e,
	0x2d, 0x5a, 0xf1, 0x23, 0xec, 0x64, 0xad, 0xb0, 0x13, 0x4d, 0xd2, 0x8b, 0x63, 0x8d, 0x5b, 0x73,
	0x4b, 0x9b, 0x91, 0x0a, 0x55, 0x05, 0x9b, 0xae, 0x2e, 0x25, 0xe7, 0x37, 0xdb, 0xd1, 0x5d, 0x02,
	0x5a, 0xd6, 0x8f, 0x1f, 0xa0, 0x55, 0xc4, 0xf7, 0xff, 0x6c, 0x48, 0xff, 0xb7, 0x0a, 0x34, 0xce,
	0xd5, 0x6f, 0xf3, 0x53, 0xe4, 0xd7, 0xae, 0x8d, 0xe4, 0x7b, 0xd8, 0xd2, 0xde, 0x6d, 0xf2, 0xb1,
	0x1e, 0x58, 0xc1, 0x6f, 0x44, 0x67, 0x7f, 0x99, 0x49, 0xe0, 0xcd, 0xcc, 0x15, 0x72, 0x06, 0xeb,
	0xc9, 0xc3, 0x49, 0x6e, 0xeb, 0x65, 0xd5, 0xde, 0xf0, 0xce, 0x5e, 0xb1, 0x52, 0x07, 0xc1, 0x42,
	0x10, 0x5c, 0x06, 0x82, 0x19, 0xc8, 0x25, 0x40, 0x6e, 0xa7, 0x7f, 0x54, 0xba, 0x4e, 0x15, 0xd4,
	0x9d, 0x65, 0xeb, 0x56, 0x85, 0x94, 0x7e, 0xd0, 0xb7, 0x8b, 0xbf, 0xde, 0xa2, 0x90, 0xf2, 0x9f,
	0xb6, 0xb9, 0x42, 0x2e, 0x60, 0x23, 0x9b, 0xc3, 0x3b, 0x25, 0x43, 0xa7, 0x60, 0x3a, 0xe5, 0x23,
	0x69, 0xae, 0x7c, 0x73, 0x17, 0xf6, 0x6d, 0x36, 0xed, 0x8d, 0x18, 0x1b, 0x79, 0xd8, 0x73, 0xf0,
	0x3a, 0x62, 0xcc, 0x0b, 0xf3, 0x2e, 0x57, 0xc6, 0xcb, 0xf7, 0xe4, 0xe1, 0xe4, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x22, 0x27, 0x4a, 0xda, 0x3c, 0x0d, 0x00, 0x00,
}
