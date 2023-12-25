// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/election.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{0}
}

type ParameterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ParameterRequest) Reset() {
	*x = ParameterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterRequest) ProtoMessage() {}

func (x *ParameterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterRequest.ProtoReflect.Descriptor instead.
func (*ParameterRequest) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{1}
}

type SetupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetupReply) Reset() {
	*x = SetupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupReply) ProtoMessage() {}

func (x *SetupReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupReply.ProtoReflect.Descriptor instead.
func (*SetupReply) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{2}
}

type VoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoteReply) Reset() {
	*x = VoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteReply) ProtoMessage() {}

func (x *VoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteReply.ProtoReflect.Descriptor instead.
func (*VoteReply) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{3}
}

type BallotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BallotsRequest) Reset() {
	*x = BallotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BallotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BallotsRequest) ProtoMessage() {}

func (x *BallotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BallotsRequest.ProtoReflect.Descriptor instead.
func (*BallotsRequest) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{4}
}

type TallyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TallyRequest) Reset() {
	*x = TallyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TallyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TallyRequest) ProtoMessage() {}

func (x *TallyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TallyRequest.ProtoReflect.Descriptor instead.
func (*TallyRequest) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{5}
}

type AuthenticationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthenticationReply) Reset() {
	*x = AuthenticationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationReply) ProtoMessage() {}

func (x *AuthenticationReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationReply.ProtoReflect.Descriptor instead.
func (*AuthenticationReply) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{6}
}

type TallyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner uint32 `protobuf:"varint,1,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *TallyReply) Reset() {
	*x = TallyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TallyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TallyReply) ProtoMessage() {}

func (x *TallyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TallyReply.ProtoReflect.Descriptor instead.
func (*TallyReply) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{7}
}

func (x *TallyReply) GetWinner() uint32 {
	if x != nil {
		return x.Winner
	}
	return 0
}

type Keypair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public  []byte `protobuf:"bytes,1,opt,name=public,proto3" json:"public,omitempty"`
	Private []byte `protobuf:"bytes,2,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *Keypair) Reset() {
	*x = Keypair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keypair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keypair) ProtoMessage() {}

func (x *Keypair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keypair.ProtoReflect.Descriptor instead.
func (*Keypair) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{8}
}

func (x *Keypair) GetPublic() []byte {
	if x != nil {
		return x.Public
	}
	return nil
}

func (x *Keypair) GetPrivate() []byte {
	if x != nil {
		return x.Private
	}
	return nil
}

type AuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{9}
}

func (x *AuthenticationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generator       []byte   `protobuf:"bytes,1,opt,name=generator,proto3" json:"generator,omitempty"`
	Prime           []byte   `protobuf:"bytes,2,opt,name=prime,proto3" json:"prime,omitempty"`
	TallierKey      []byte   `protobuf:"bytes,3,opt,name=tallierKey,proto3" json:"tallierKey,omitempty"`
	VotingserverKey []byte   `protobuf:"bytes,4,opt,name=votingserverKey,proto3" json:"votingserverKey,omitempty"`
	Candidates      []uint32 `protobuf:"varint,5,rep,packed,name=candidates,proto3" json:"candidates,omitempty"`
}

func (x *SetupRequest) Reset() {
	*x = SetupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupRequest) ProtoMessage() {}

func (x *SetupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupRequest.ProtoReflect.Descriptor instead.
func (*SetupRequest) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{10}
}

func (x *SetupRequest) GetGenerator() []byte {
	if x != nil {
		return x.Generator
	}
	return nil
}

func (x *SetupRequest) GetPrime() []byte {
	if x != nil {
		return x.Prime
	}
	return nil
}

func (x *SetupRequest) GetTallierKey() []byte {
	if x != nil {
		return x.TallierKey
	}
	return nil
}

func (x *SetupRequest) GetVotingserverKey() []byte {
	if x != nil {
		return x.VotingserverKey
	}
	return nil
}

func (x *SetupRequest) GetCandidates() []uint32 {
	if x != nil {
		return x.Candidates
	}
	return nil
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Com string `protobuf:"bytes,2,opt,name=com,proto3" json:"com,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{11}
}

func (x *Commit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Commit) GetCom() string {
	if x != nil {
		return x.Com
	}
	return ""
}

type CommitReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommitReply) Reset() {
	*x = CommitReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitReply) ProtoMessage() {}

func (x *CommitReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitReply.ProtoReflect.Descriptor instead.
func (*CommitReply) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{12}
}

func (x *CommitReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ParameterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generator       []byte   `protobuf:"bytes,1,opt,name=generator,proto3" json:"generator,omitempty"`
	Prime           []byte   `protobuf:"bytes,2,opt,name=prime,proto3" json:"prime,omitempty"`
	TallierKey      []byte   `protobuf:"bytes,3,opt,name=tallierKey,proto3" json:"tallierKey,omitempty"`
	VotingserverKey []byte   `protobuf:"bytes,4,opt,name=votingserverKey,proto3" json:"votingserverKey,omitempty"`
	Candidates      []uint32 `protobuf:"varint,5,rep,packed,name=candidates,proto3" json:"candidates,omitempty"`
}

func (x *ParameterReply) Reset() {
	*x = ParameterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterReply) ProtoMessage() {}

func (x *ParameterReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterReply.ProtoReflect.Descriptor instead.
func (*ParameterReply) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{13}
}

func (x *ParameterReply) GetGenerator() []byte {
	if x != nil {
		return x.Generator
	}
	return nil
}

func (x *ParameterReply) GetPrime() []byte {
	if x != nil {
		return x.Prime
	}
	return nil
}

func (x *ParameterReply) GetTallierKey() []byte {
	if x != nil {
		return x.TallierKey
	}
	return nil
}

func (x *ParameterReply) GetVotingserverKey() []byte {
	if x != nil {
		return x.VotingserverKey
	}
	return nil
}

func (x *ParameterReply) GetCandidates() []uint32 {
	if x != nil {
		return x.Candidates
	}
	return nil
}

type GetCommitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCommitRequest) Reset() {
	*x = GetCommitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommitRequest) ProtoMessage() {}

func (x *GetCommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommitRequest.ProtoReflect.Descriptor instead.
func (*GetCommitRequest) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{14}
}

func (x *GetCommitRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Ballot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Candidate uint32 `protobuf:"varint,2,opt,name=candidate,proto3" json:"candidate,omitempty"`
}

func (x *Ballot) Reset() {
	*x = Ballot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_election_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ballot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ballot) ProtoMessage() {}

func (x *Ballot) ProtoReflect() protoreflect.Message {
	mi := &file_proto_election_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ballot.ProtoReflect.Descriptor instead.
func (*Ballot) Descriptor() ([]byte, []int) {
	return file_proto_election_proto_rawDescGZIP(), []int{15}
}

func (x *Ballot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ballot) GetCandidate() uint32 {
	if x != nil {
		return x.Candidate
	}
	return 0
}

var File_proto_election_proto protoreflect.FileDescriptor

var file_proto_election_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12,
	0x0a, 0x10, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x0b, 0x0a, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a,
	0x0e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x54, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x15, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x0a, 0x54, 0x61, 0x6c, 0x6c, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x07,
	0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x6c, 0x6c, 0x69,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x61, 0x6c,
	0x6c, 0x69, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x22, 0x2a, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6f, 0x6d, 0x22, 0x1d, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xae, 0x01, 0x0a,
	0x0e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x61, 0x6c, 0x6c, 0x69, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x22, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x36, 0x0a, 0x06, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x32, 0xd4, 0x02, 0x0a, 0x0d, 0x42, 0x75,
	0x6c, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x29, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x32, 0x3c, 0x0a, 0x09, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x12, 0x2f, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x6a,
	0x0a, 0x07, 0x54, 0x61, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61,
	0x69, 0x72, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x54, 0x61, 0x6c, 0x6c, 0x79,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61,
	0x6c, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x88, 0x01, 0x0a, 0x0c, 0x56,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79,
	0x70, 0x61, 0x69, 0x72, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0c, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x62, 0x73, 0x63, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_election_proto_rawDescOnce sync.Once
	file_proto_election_proto_rawDescData = file_proto_election_proto_rawDesc
)

func file_proto_election_proto_rawDescGZIP() []byte {
	file_proto_election_proto_rawDescOnce.Do(func() {
		file_proto_election_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_election_proto_rawDescData)
	})
	return file_proto_election_proto_rawDescData
}

var file_proto_election_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proto_election_proto_goTypes = []interface{}{
	(*RegisterReply)(nil),         // 0: proto.RegisterReply
	(*ParameterRequest)(nil),      // 1: proto.ParameterRequest
	(*SetupReply)(nil),            // 2: proto.SetupReply
	(*VoteReply)(nil),             // 3: proto.VoteReply
	(*BallotsRequest)(nil),        // 4: proto.BallotsRequest
	(*TallyRequest)(nil),          // 5: proto.TallyRequest
	(*AuthenticationReply)(nil),   // 6: proto.AuthenticationReply
	(*TallyReply)(nil),            // 7: proto.TallyReply
	(*Keypair)(nil),               // 8: proto.Keypair
	(*AuthenticationRequest)(nil), // 9: proto.AuthenticationRequest
	(*SetupRequest)(nil),          // 10: proto.SetupRequest
	(*Commit)(nil),                // 11: proto.Commit
	(*CommitReply)(nil),           // 12: proto.CommitReply
	(*ParameterReply)(nil),        // 13: proto.ParameterReply
	(*GetCommitRequest)(nil),      // 14: proto.GetCommitRequest
	(*Ballot)(nil),                // 15: proto.Ballot
}
var file_proto_election_proto_depIdxs = []int32{
	10, // 0: proto.BulletinBoard.Setup:input_type -> proto.SetupRequest
	11, // 1: proto.BulletinBoard.Register:input_type -> proto.Commit
	15, // 2: proto.BulletinBoard.Vote:input_type -> proto.Ballot
	1,  // 3: proto.BulletinBoard.GetParameters:input_type -> proto.ParameterRequest
	14, // 4: proto.BulletinBoard.GetCommit:input_type -> proto.GetCommitRequest
	4,  // 5: proto.BulletinBoard.GetBallots:input_type -> proto.BallotsRequest
	11, // 6: proto.Registrar.Register:input_type -> proto.Commit
	8,  // 7: proto.Tallier.Setup:input_type -> proto.Keypair
	5,  // 8: proto.Tallier.Tally:input_type -> proto.TallyRequest
	8,  // 9: proto.VotingServer.Setup:input_type -> proto.Keypair
	9,  // 10: proto.VotingServer.Authenticate:input_type -> proto.AuthenticationRequest
	2,  // 11: proto.BulletinBoard.Setup:output_type -> proto.SetupReply
	0,  // 12: proto.BulletinBoard.Register:output_type -> proto.RegisterReply
	3,  // 13: proto.BulletinBoard.Vote:output_type -> proto.VoteReply
	13, // 14: proto.BulletinBoard.GetParameters:output_type -> proto.ParameterReply
	11, // 15: proto.BulletinBoard.GetCommit:output_type -> proto.Commit
	15, // 16: proto.BulletinBoard.GetBallots:output_type -> proto.Ballot
	12, // 17: proto.Registrar.Register:output_type -> proto.CommitReply
	2,  // 18: proto.Tallier.Setup:output_type -> proto.SetupReply
	7,  // 19: proto.Tallier.Tally:output_type -> proto.TallyReply
	2,  // 20: proto.VotingServer.Setup:output_type -> proto.SetupReply
	6,  // 21: proto.VotingServer.Authenticate:output_type -> proto.AuthenticationReply
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_election_proto_init() }
func file_proto_election_proto_init() {
	if File_proto_election_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_election_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_proto_election_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterRequest); i {
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
		file_proto_election_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupReply); i {
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
		file_proto_election_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteReply); i {
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
		file_proto_election_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BallotsRequest); i {
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
		file_proto_election_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TallyRequest); i {
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
		file_proto_election_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationReply); i {
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
		file_proto_election_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TallyReply); i {
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
		file_proto_election_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keypair); i {
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
		file_proto_election_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationRequest); i {
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
		file_proto_election_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupRequest); i {
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
		file_proto_election_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_proto_election_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitReply); i {
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
		file_proto_election_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterReply); i {
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
		file_proto_election_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommitRequest); i {
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
		file_proto_election_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ballot); i {
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
			RawDescriptor: file_proto_election_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_proto_election_proto_goTypes,
		DependencyIndexes: file_proto_election_proto_depIdxs,
		MessageInfos:      file_proto_election_proto_msgTypes,
	}.Build()
	File_proto_election_proto = out.File
	file_proto_election_proto_rawDesc = nil
	file_proto_election_proto_goTypes = nil
	file_proto_election_proto_depIdxs = nil
}
