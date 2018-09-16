// Code generated by protoc-gen-go.
// source: blockchain.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 区块头信息
// 	 version : 版本信息
// 	 parentHash :父哈希
// 	 txHash : 交易根哈希
// 	 stateHash :状态哈希
// 	 height : 区块高度
// 	 blockTime :区块产生时的时标
// 	 txCount : 区块上所有交易个数
// 	 difficulty :区块难度系数，
// 	 signature :交易签名
type Header struct {
	Version    int64      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte     `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte     `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	StateHash  []byte     `protobuf:"bytes,4,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height     int64      `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	BlockTime  int64      `protobuf:"varint,6,opt,name=blockTime" json:"blockTime,omitempty"`
	TxCount    int64      `protobuf:"varint,9,opt,name=txCount" json:"txCount,omitempty"`
	Hash       []byte     `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	Difficulty uint32     `protobuf:"varint,11,opt,name=difficulty" json:"difficulty,omitempty"`
	Signature  *Signature `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Header) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Header) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Header) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Header) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

func (m *Header) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Header) GetDifficulty() uint32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Header) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

//  参考Header解释
type Block struct {
	Version    int64          `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte         `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte         `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	StateHash  []byte         `protobuf:"bytes,4,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height     int64          `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	BlockTime  int64          `protobuf:"varint,6,opt,name=blockTime" json:"blockTime,omitempty"`
	Difficulty uint32         `protobuf:"varint,11,opt,name=difficulty" json:"difficulty,omitempty"`
	Signature  *Signature     `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`
	Txs        []*Transaction `protobuf:"bytes,7,rep,name=txs" json:"txs,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Block) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Block) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Block) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *Block) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Block) GetDifficulty() uint32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Block) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Block) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Blocks struct {
	Items []*Block `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Blocks) Reset()                    { *m = Blocks{} }
func (m *Blocks) String() string            { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()               {}
func (*Blocks) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Blocks) GetItems() []*Block {
	if m != nil {
		return m.Items
	}
	return nil
}

// 节点ID以及对应的Block
type BlockPid struct {
	Pid   string `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Block *Block `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *BlockPid) Reset()                    { *m = BlockPid{} }
func (m *BlockPid) String() string            { return proto.CompactTextString(m) }
func (*BlockPid) ProtoMessage()               {}
func (*BlockPid) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BlockPid) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *BlockPid) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

// resp
type BlockDetails struct {
	Items []*BlockDetail `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *BlockDetails) Reset()                    { *m = BlockDetails{} }
func (m *BlockDetails) String() string            { return proto.CompactTextString(m) }
func (*BlockDetails) ProtoMessage()               {}
func (*BlockDetails) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *BlockDetails) GetItems() []*BlockDetail {
	if m != nil {
		return m.Items
	}
	return nil
}

// resp
type Headers struct {
	Items []*Header `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Headers) Reset()                    { *m = Headers{} }
func (m *Headers) String() string            { return proto.CompactTextString(m) }
func (*Headers) ProtoMessage()               {}
func (*Headers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Headers) GetItems() []*Header {
	if m != nil {
		return m.Items
	}
	return nil
}

type HeadersPid struct {
	Pid     string   `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Headers *Headers `protobuf:"bytes,2,opt,name=headers" json:"headers,omitempty"`
}

func (m *HeadersPid) Reset()                    { *m = HeadersPid{} }
func (m *HeadersPid) String() string            { return proto.CompactTextString(m) }
func (*HeadersPid) ProtoMessage()               {}
func (*HeadersPid) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *HeadersPid) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *HeadersPid) GetHeaders() *Headers {
	if m != nil {
		return m.Headers
	}
	return nil
}

// 区块视图
// 	 head : 区块头信息
// 	 txCount :区块上交易个数
// 	 txHashes : 区块上交易的哈希列表
type BlockOverview struct {
	Head     *Header  `protobuf:"bytes,1,opt,name=head" json:"head,omitempty"`
	TxCount  int64    `protobuf:"varint,2,opt,name=txCount" json:"txCount,omitempty"`
	TxHashes [][]byte `protobuf:"bytes,3,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
}

func (m *BlockOverview) Reset()                    { *m = BlockOverview{} }
func (m *BlockOverview) String() string            { return proto.CompactTextString(m) }
func (*BlockOverview) ProtoMessage()               {}
func (*BlockOverview) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *BlockOverview) GetHead() *Header {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *BlockOverview) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

func (m *BlockOverview) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

// 区块详细信息
// 	 block : 区块信息
// 	 receipts :区块上所有交易的收据信息列表
type BlockDetail struct {
	Block          *Block         `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
	Receipts       []*ReceiptData `protobuf:"bytes,2,rep,name=receipts" json:"receipts,omitempty"`
	KV             []*KeyValue    `protobuf:"bytes,3,rep,name=KV" json:"KV,omitempty"`
	PrevStatusHash []byte         `protobuf:"bytes,4,opt,name=prevStatusHash,proto3" json:"prevStatusHash,omitempty"`
}

func (m *BlockDetail) Reset()                    { *m = BlockDetail{} }
func (m *BlockDetail) String() string            { return proto.CompactTextString(m) }
func (*BlockDetail) ProtoMessage()               {}
func (*BlockDetail) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *BlockDetail) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockDetail) GetReceipts() []*ReceiptData {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func (m *BlockDetail) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *BlockDetail) GetPrevStatusHash() []byte {
	if m != nil {
		return m.PrevStatusHash
	}
	return nil
}

type Receipts struct {
	Receipts []*Receipt `protobuf:"bytes,1,rep,name=receipts" json:"receipts,omitempty"`
}

func (m *Receipts) Reset()                    { *m = Receipts{} }
func (m *Receipts) String() string            { return proto.CompactTextString(m) }
func (*Receipts) ProtoMessage()               {}
func (*Receipts) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *Receipts) GetReceipts() []*Receipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

type PrivacyKV struct {
	PrivacyKVToken []*PrivacyKVToken `protobuf:"bytes,1,rep,name=privacyKVToken" json:"privacyKVToken,omitempty"`
}

func (m *PrivacyKV) Reset()                    { *m = PrivacyKV{} }
func (m *PrivacyKV) String() string            { return proto.CompactTextString(m) }
func (*PrivacyKV) ProtoMessage()               {}
func (*PrivacyKV) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *PrivacyKV) GetPrivacyKVToken() []*PrivacyKVToken {
	if m != nil {
		return m.PrivacyKVToken
	}
	return nil
}

type PrivacyKVToken struct {
	Token   string      `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	TxIndex int32       `protobuf:"varint,2,opt,name=txIndex" json:"txIndex,omitempty"`
	Txhash  []byte      `protobuf:"bytes,3,opt,name=txhash,proto3" json:"txhash,omitempty"`
	KV      []*KeyValue `protobuf:"bytes,4,rep,name=KV" json:"KV,omitempty"`
}

func (m *PrivacyKVToken) Reset()                    { *m = PrivacyKVToken{} }
func (m *PrivacyKVToken) String() string            { return proto.CompactTextString(m) }
func (*PrivacyKVToken) ProtoMessage()               {}
func (*PrivacyKVToken) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *PrivacyKVToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PrivacyKVToken) GetTxIndex() int32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *PrivacyKVToken) GetTxhash() []byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

func (m *PrivacyKVToken) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

type ReceiptsAndPrivacyKV struct {
	Receipts  *Receipts  `protobuf:"bytes,1,opt,name=receipts" json:"receipts,omitempty"`
	PrivacyKV *PrivacyKV `protobuf:"bytes,2,opt,name=privacyKV" json:"privacyKV,omitempty"`
}

func (m *ReceiptsAndPrivacyKV) Reset()                    { *m = ReceiptsAndPrivacyKV{} }
func (m *ReceiptsAndPrivacyKV) String() string            { return proto.CompactTextString(m) }
func (*ReceiptsAndPrivacyKV) ProtoMessage()               {}
func (*ReceiptsAndPrivacyKV) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *ReceiptsAndPrivacyKV) GetReceipts() *Receipts {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func (m *ReceiptsAndPrivacyKV) GetPrivacyKV() *PrivacyKV {
	if m != nil {
		return m.PrivacyKV
	}
	return nil
}

type ReceiptCheckTxList struct {
	Errs []string `protobuf:"bytes,1,rep,name=errs" json:"errs,omitempty"`
}

func (m *ReceiptCheckTxList) Reset()                    { *m = ReceiptCheckTxList{} }
func (m *ReceiptCheckTxList) String() string            { return proto.CompactTextString(m) }
func (*ReceiptCheckTxList) ProtoMessage()               {}
func (*ReceiptCheckTxList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *ReceiptCheckTxList) GetErrs() []string {
	if m != nil {
		return m.Errs
	}
	return nil
}

// 区块链状态
// 	 currentHeight : 区块最新高度
// 	 mempoolSize :内存池大小
// 	 msgQueueSize : 消息队列大小
type ChainStatus struct {
	CurrentHeight int64 `protobuf:"varint,1,opt,name=currentHeight" json:"currentHeight,omitempty"`
	MempoolSize   int64 `protobuf:"varint,2,opt,name=mempoolSize" json:"mempoolSize,omitempty"`
	MsgQueueSize  int64 `protobuf:"varint,3,opt,name=msgQueueSize" json:"msgQueueSize,omitempty"`
}

func (m *ChainStatus) Reset()                    { *m = ChainStatus{} }
func (m *ChainStatus) String() string            { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()               {}
func (*ChainStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *ChainStatus) GetCurrentHeight() int64 {
	if m != nil {
		return m.CurrentHeight
	}
	return 0
}

func (m *ChainStatus) GetMempoolSize() int64 {
	if m != nil {
		return m.MempoolSize
	}
	return 0
}

func (m *ChainStatus) GetMsgQueueSize() int64 {
	if m != nil {
		return m.MsgQueueSize
	}
	return 0
}

// 获取区块信息
// 	 start : 获取区块的开始高度
// 	 end :获取区块的结束高度
// 	 Isdetail : 是否需要获取区块的详细信息
// 	 pid : peer列表
type ReqBlocks struct {
	Start    int64    `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End      int64    `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	IsDetail bool     `protobuf:"varint,3,opt,name=isDetail" json:"isDetail,omitempty"`
	Pid      []string `protobuf:"bytes,4,rep,name=pid" json:"pid,omitempty"`
}

func (m *ReqBlocks) Reset()                    { *m = ReqBlocks{} }
func (m *ReqBlocks) String() string            { return proto.CompactTextString(m) }
func (*ReqBlocks) ProtoMessage()               {}
func (*ReqBlocks) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *ReqBlocks) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ReqBlocks) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *ReqBlocks) GetIsDetail() bool {
	if m != nil {
		return m.IsDetail
	}
	return false
}

func (m *ReqBlocks) GetPid() []string {
	if m != nil {
		return m.Pid
	}
	return nil
}

type MempoolSize struct {
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *MempoolSize) Reset()                    { *m = MempoolSize{} }
func (m *MempoolSize) String() string            { return proto.CompactTextString(m) }
func (*MempoolSize) ProtoMessage()               {}
func (*MempoolSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *MempoolSize) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ReplyBlockHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *ReplyBlockHeight) Reset()                    { *m = ReplyBlockHeight{} }
func (m *ReplyBlockHeight) String() string            { return proto.CompactTextString(m) }
func (*ReplyBlockHeight) ProtoMessage()               {}
func (*ReplyBlockHeight) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *ReplyBlockHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// 区块体信息
// 	 txs : 区块上所有交易列表
// 	 receipts :区块上所有交易的收据信息列表
type BlockBody struct {
	Txs      []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
	Receipts []*ReceiptData `protobuf:"bytes,2,rep,name=receipts" json:"receipts,omitempty"`
}

func (m *BlockBody) Reset()                    { *m = BlockBody{} }
func (m *BlockBody) String() string            { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()               {}
func (*BlockBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

func (m *BlockBody) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *BlockBody) GetReceipts() []*ReceiptData {
	if m != nil {
		return m.Receipts
	}
	return nil
}

//  区块追赶主链状态，用于判断本节点区块是否已经同步好
type IsCaughtUp struct {
	Iscaughtup bool `protobuf:"varint,1,opt,name=Iscaughtup" json:"Iscaughtup,omitempty"`
}

func (m *IsCaughtUp) Reset()                    { *m = IsCaughtUp{} }
func (m *IsCaughtUp) String() string            { return proto.CompactTextString(m) }
func (*IsCaughtUp) ProtoMessage()               {}
func (*IsCaughtUp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

func (m *IsCaughtUp) GetIscaughtup() bool {
	if m != nil {
		return m.Iscaughtup
	}
	return false
}

//  ntp时钟状态
type IsNtpClockSync struct {
	Isntpclocksync bool `protobuf:"varint,1,opt,name=isntpclocksync" json:"isntpclocksync,omitempty"`
}

func (m *IsNtpClockSync) Reset()                    { *m = IsNtpClockSync{} }
func (m *IsNtpClockSync) String() string            { return proto.CompactTextString(m) }
func (*IsNtpClockSync) ProtoMessage()               {}
func (*IsNtpClockSync) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{20} }

func (m *IsNtpClockSync) GetIsntpclocksync() bool {
	if m != nil {
		return m.Isntpclocksync
	}
	return false
}

type BlockChainQuery struct {
	Driver    string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	FuncName  string `protobuf:"bytes,2,opt,name=funcName" json:"funcName,omitempty"`
	StateHash []byte `protobuf:"bytes,3,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Param     []byte `protobuf:"bytes,4,opt,name=param,proto3" json:"param,omitempty"`
}

func (m *BlockChainQuery) Reset()                    { *m = BlockChainQuery{} }
func (m *BlockChainQuery) String() string            { return proto.CompactTextString(m) }
func (*BlockChainQuery) ProtoMessage()               {}
func (*BlockChainQuery) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{21} }

func (m *BlockChainQuery) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *BlockChainQuery) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *BlockChainQuery) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *BlockChainQuery) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

//  通过block hash记录block的操作类型及add/del：1/2
type BlockSequence struct {
	Hash []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Type int64  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
}

func (m *BlockSequence) Reset()                    { *m = BlockSequence{} }
func (m *BlockSequence) String() string            { return proto.CompactTextString(m) }
func (*BlockSequence) ProtoMessage()               {}
func (*BlockSequence) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{22} }

func (m *BlockSequence) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockSequence) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

// resp
type BlockSequences struct {
	Items []*BlockSequence `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *BlockSequences) Reset()                    { *m = BlockSequences{} }
func (m *BlockSequences) String() string            { return proto.CompactTextString(m) }
func (*BlockSequences) ProtoMessage()               {}
func (*BlockSequences) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{23} }

func (m *BlockSequences) GetItems() []*BlockSequence {
	if m != nil {
		return m.Items
	}
	return nil
}

// 平行链区块详细信息
// 	 blockdetail : 区块详细信息
// 	 sequence :区块序列号
type ParaChainBlockDetail struct {
	Blockdetail *BlockDetail `protobuf:"bytes,1,opt,name=blockdetail" json:"blockdetail,omitempty"`
	Sequence    int64        `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *ParaChainBlockDetail) Reset()                    { *m = ParaChainBlockDetail{} }
func (m *ParaChainBlockDetail) String() string            { return proto.CompactTextString(m) }
func (*ParaChainBlockDetail) ProtoMessage()               {}
func (*ParaChainBlockDetail) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{24} }

func (m *ParaChainBlockDetail) GetBlockdetail() *BlockDetail {
	if m != nil {
		return m.Blockdetail
	}
	return nil
}

func (m *ParaChainBlockDetail) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*Header)(nil), "types.Header")
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*Blocks)(nil), "types.Blocks")
	proto.RegisterType((*BlockPid)(nil), "types.BlockPid")
	proto.RegisterType((*BlockDetails)(nil), "types.BlockDetails")
	proto.RegisterType((*Headers)(nil), "types.Headers")
	proto.RegisterType((*HeadersPid)(nil), "types.HeadersPid")
	proto.RegisterType((*BlockOverview)(nil), "types.BlockOverview")
	proto.RegisterType((*BlockDetail)(nil), "types.BlockDetail")
	proto.RegisterType((*Receipts)(nil), "types.Receipts")
	proto.RegisterType((*PrivacyKV)(nil), "types.PrivacyKV")
	proto.RegisterType((*PrivacyKVToken)(nil), "types.PrivacyKVToken")
	proto.RegisterType((*ReceiptsAndPrivacyKV)(nil), "types.ReceiptsAndPrivacyKV")
	proto.RegisterType((*ReceiptCheckTxList)(nil), "types.ReceiptCheckTxList")
	proto.RegisterType((*ChainStatus)(nil), "types.ChainStatus")
	proto.RegisterType((*ReqBlocks)(nil), "types.ReqBlocks")
	proto.RegisterType((*MempoolSize)(nil), "types.MempoolSize")
	proto.RegisterType((*ReplyBlockHeight)(nil), "types.ReplyBlockHeight")
	proto.RegisterType((*BlockBody)(nil), "types.BlockBody")
	proto.RegisterType((*IsCaughtUp)(nil), "types.IsCaughtUp")
	proto.RegisterType((*IsNtpClockSync)(nil), "types.IsNtpClockSync")
	proto.RegisterType((*BlockChainQuery)(nil), "types.BlockChainQuery")
	proto.RegisterType((*BlockSequence)(nil), "types.BlockSequence")
	proto.RegisterType((*BlockSequences)(nil), "types.BlockSequences")
	proto.RegisterType((*ParaChainBlockDetail)(nil), "types.ParaChainBlockDetail")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 997 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x56, 0x5d, 0x6f, 0xe3, 0x44,
	0x17, 0x96, 0xf3, 0xd1, 0x26, 0x27, 0x69, 0xb6, 0xef, 0x28, 0x2f, 0xb2, 0x2a, 0x04, 0xd9, 0xa1,
	0x42, 0xd1, 0xb2, 0x0a, 0x52, 0x83, 0x60, 0x2f, 0x40, 0x82, 0x66, 0x2f, 0x5a, 0x0a, 0x4b, 0x77,
	0x12, 0x7a, 0xc1, 0xdd, 0xd4, 0x9e, 0xc6, 0xa3, 0x26, 0xb6, 0xd7, 0x33, 0x0e, 0x31, 0x7f, 0x87,
	0x3b, 0xc4, 0x8f, 0x44, 0x73, 0x66, 0x9c, 0xd8, 0xd9, 0x5d, 0x24, 0x24, 0x6e, 0xb8, 0xca, 0x3c,
	0xe7, 0x63, 0xce, 0xc7, 0x9c, 0xf3, 0x38, 0x70, 0x7a, 0xbf, 0x4a, 0x82, 0xc7, 0x20, 0xe2, 0x32,
	0x9e, 0xa4, 0x59, 0xa2, 0x13, 0xd2, 0xd6, 0x45, 0x2a, 0xd4, 0xd9, 0xff, 0x74, 0xc6, 0x63, 0xc5,
	0x03, 0x2d, 0x13, 0xa7, 0x39, 0xeb, 0x07, 0xc9, 0x7a, 0x5d, 0x22, 0xfa, 0x67, 0x03, 0x8e, 0xae,
	0x04, 0x0f, 0x45, 0x46, 0x7c, 0x38, 0xde, 0x88, 0x4c, 0xc9, 0x24, 0xf6, 0xbd, 0x91, 0x37, 0x6e,
	0xb2, 0x12, 0x92, 0x8f, 0x00, 0x52, 0x9e, 0x89, 0x58, 0x5f, 0x71, 0x15, 0xf9, 0x8d, 0x91, 0x37,
	0xee, 0xb3, 0x8a, 0x84, 0x7c, 0x00, 0x47, 0x7a, 0x8b, 0xba, 0x26, 0xea, 0x1c, 0x22, 0x1f, 0x42,
	0x57, 0x69, 0xae, 0x05, 0xaa, 0x5a, 0xa8, 0xda, 0x0b, 0x8c, 0x57, 0x24, 0xe4, 0x32, 0xd2, 0x7e,
	0x1b, 0xc3, 0x39, 0x64, 0xbc, 0xb0, 0x9c, 0x85, 0x5c, 0x0b, 0xff, 0x08, 0x55, 0x7b, 0x81, 0xc9,
	0x52, 0x6f, 0x67, 0x49, 0x1e, 0x6b, 0xbf, 0x6b, 0xb3, 0x74, 0x90, 0x10, 0x68, 0x45, 0x26, 0x10,
	0x60, 0x20, 0x3c, 0x9b, 0xcc, 0x43, 0xf9, 0xf0, 0x20, 0x83, 0x7c, 0xa5, 0x0b, 0xbf, 0x37, 0xf2,
	0xc6, 0x27, 0xac, 0x22, 0x21, 0x13, 0xe8, 0x2a, 0xb9, 0x8c, 0xb9, 0xce, 0x33, 0xe1, 0x77, 0x46,
	0xde, 0xb8, 0x77, 0x71, 0x3a, 0xc1, 0xd6, 0x4d, 0xe6, 0xa5, 0x9c, 0xed, 0x4d, 0xe8, 0xef, 0x0d,
	0x68, 0x5f, 0x9a, 0x5c, 0xfe, 0x23, 0xdd, 0xfa, 0x97, 0xeb, 0x27, 0xe7, 0xd0, 0xd4, 0x5b, 0xe5,
	0x1f, 0x8f, 0x9a, 0xe3, 0xde, 0x05, 0x71, 0x96, 0x8b, 0xfd, 0x8c, 0x31, 0xa3, 0xa6, 0xcf, 0xe1,
	0x08, 0x9b, 0xa4, 0x08, 0x85, 0xb6, 0xd4, 0x62, 0xad, 0x7c, 0x0f, 0x3d, 0xfa, 0xce, 0x03, 0xb5,
	0xcc, 0xaa, 0xe8, 0xb7, 0xd0, 0x41, 0x7c, 0x2b, 0x43, 0x72, 0x0a, 0xcd, 0x54, 0x86, 0xd8, 0xd1,
	0x2e, 0x33, 0x47, 0x73, 0x03, 0x96, 0x83, 0x8d, 0x7c, 0xeb, 0x06, 0x54, 0xd1, 0x17, 0xd0, 0x47,
	0xfc, 0x52, 0x68, 0x2e, 0x57, 0x8a, 0x8c, 0xeb, 0x51, 0x49, 0xd5, 0xc7, 0xda, 0x94, 0xb1, 0x27,
	0x70, 0x6c, 0xa7, 0x5f, 0x91, 0x4f, 0xea, 0x4e, 0x27, 0xce, 0xc9, 0xaa, 0x4b, 0xfb, 0x2b, 0x00,
	0x67, 0xff, 0xee, 0x6c, 0xc7, 0x70, 0x1c, 0x59, 0xbd, 0xcb, 0x77, 0x50, 0xbb, 0x46, 0xb1, 0x52,
	0x4d, 0x23, 0x38, 0xc1, 0x7c, 0x7e, 0xda, 0x88, 0x6c, 0x23, 0xc5, 0xaf, 0xe4, 0x29, 0xb4, 0x8c,
	0x0e, 0x6f, 0x7b, 0x2b, 0x3c, 0xaa, 0xaa, 0xb3, 0xdf, 0xa8, 0xcf, 0xfe, 0x19, 0x74, 0xec, 0x14,
	0x09, 0xe5, 0x37, 0x47, 0xcd, 0x71, 0x9f, 0xed, 0x30, 0xfd, 0xc3, 0x83, 0x5e, 0xa5, 0xf4, 0x7d,
	0x47, 0xbd, 0xf7, 0x76, 0x94, 0x4c, 0xa0, 0x93, 0x89, 0x40, 0xc8, 0x54, 0x9b, 0x42, 0xaa, 0x4d,
	0x64, 0x56, 0xfc, 0x92, 0x6b, 0xce, 0x76, 0x36, 0xe4, 0x63, 0x68, 0xdc, 0xdc, 0x61, 0xe4, 0xde,
	0xc5, 0x13, 0x67, 0x79, 0x23, 0x8a, 0x3b, 0xbe, 0xca, 0x05, 0x6b, 0xdc, 0xdc, 0x91, 0x4f, 0x61,
	0x90, 0x66, 0x62, 0x33, 0xd7, 0x5c, 0xe7, 0xaa, 0x32, 0xe1, 0x07, 0x52, 0xfa, 0x25, 0x74, 0x58,
	0x79, 0xe9, 0xb3, 0x4a, 0x12, 0xf6, 0x51, 0x06, 0xf5, 0x24, 0xf6, 0x09, 0xd0, 0xef, 0xa1, 0x7b,
	0x9b, 0xc9, 0x0d, 0x0f, 0x8a, 0x9b, 0x3b, 0xf2, 0x8d, 0x09, 0xe6, 0xc0, 0x22, 0x79, 0x14, 0xb1,
	0x73, 0xff, 0xbf, 0x73, 0xbf, 0xad, 0x29, 0xd9, 0x81, 0x31, 0x2d, 0x60, 0x50, 0xb7, 0x20, 0x43,
	0x68, 0x6b, 0x77, 0x8f, 0x79, 0x6a, 0x0b, 0xec, 0x73, 0x5c, 0xc7, 0xa1, 0xd8, 0xe2, 0x73, 0xb4,
	0x59, 0x09, 0xed, 0x8a, 0x47, 0xb5, 0x15, 0x47, 0x3a, 0xb2, 0x6d, 0x6a, 0xbd, 0xb7, 0x4d, 0x54,
	0xc1, 0xb0, 0x2c, 0xff, 0xbb, 0x38, 0xdc, 0x57, 0xf4, 0x59, 0xad, 0x15, 0x5e, 0xc5, 0xbd, 0x34,
	0xaf, 0x3c, 0xc6, 0x04, 0xba, 0xbb, 0x8a, 0xdc, 0x18, 0x9e, 0x1e, 0x56, 0xce, 0xf6, 0x26, 0x74,
	0x0c, 0xc4, 0xdd, 0x32, 0x8b, 0x44, 0xf0, 0xb8, 0xd8, 0xfe, 0x20, 0x15, 0xd2, 0xa9, 0xc8, 0x32,
	0xdb, 0xf9, 0x2e, 0xc3, 0x33, 0x2d, 0xa0, 0x37, 0x33, 0x1f, 0x19, 0xfb, 0x60, 0xe4, 0x1c, 0x4e,
	0x82, 0x3c, 0x43, 0x62, 0xb3, 0xd4, 0x64, 0x99, 0xb0, 0x2e, 0x24, 0x23, 0xe8, 0xad, 0xc5, 0x3a,
	0x4d, 0x92, 0xd5, 0x5c, 0xfe, 0x26, 0xdc, 0xe4, 0x56, 0x45, 0x84, 0x42, 0x7f, 0xad, 0x96, 0xaf,
	0x73, 0x91, 0x0b, 0x34, 0x69, 0xa2, 0x49, 0x4d, 0x46, 0x39, 0x74, 0x99, 0x78, 0xe3, 0x68, 0x65,
	0x08, 0x6d, 0xa5, 0x79, 0x56, 0x06, 0xb4, 0xc0, 0xac, 0xa3, 0x88, 0x43, 0x17, 0xc0, 0x1c, 0xcd,
	0x5a, 0x48, 0x65, 0xc7, 0x1e, 0x2f, 0xed, 0xb0, 0x1d, 0x2e, 0x97, 0xb7, 0x85, 0xe5, 0x99, 0x23,
	0x7d, 0x0a, 0xbd, 0x1f, 0x2b, 0x59, 0x11, 0x68, 0x29, 0x93, 0x8d, 0x8d, 0x81, 0x67, 0xfa, 0x0c,
	0x4e, 0x99, 0x48, 0x57, 0x05, 0xe6, 0xe1, 0xea, 0xdb, 0x33, 0xb3, 0x57, 0x65, 0x66, 0x93, 0x31,
	0x9a, 0x5d, 0x26, 0x61, 0x51, 0x12, 0xa7, 0xf7, 0xb7, 0xc4, 0xf9, 0x4f, 0xd7, 0x8e, 0x3e, 0x07,
	0xb8, 0x56, 0x33, 0x9e, 0x2f, 0x23, 0xfd, 0x73, 0x6a, 0xc8, 0xfe, 0x5a, 0x05, 0x88, 0xf2, 0x14,
	0x93, 0xe9, 0xb0, 0x8a, 0x84, 0xbe, 0x80, 0xc1, 0xb5, 0x7a, 0xa5, 0xd3, 0x99, 0xc9, 0x6a, 0x5e,
	0xc4, 0x81, 0xd9, 0x4a, 0xa9, 0x62, 0x9d, 0x06, 0xd8, 0xd6, 0x22, 0x0e, 0x9c, 0xd7, 0x81, 0x94,
	0x16, 0xf0, 0x04, 0x4b, 0xc1, 0xc7, 0x7f, 0x9d, 0x8b, 0xac, 0x30, 0x55, 0x87, 0x99, 0xdc, 0x88,
	0xcc, 0xed, 0x84, 0x43, 0xa6, 0xe5, 0x0f, 0x79, 0x1c, 0xbc, 0xe2, 0x6b, 0xfb, 0xd4, 0x5d, 0xb6,
	0xc3, 0xf5, 0x2f, 0x5c, 0xf3, 0xf0, 0x0b, 0x37, 0x84, 0x76, 0xca, 0x33, 0xbe, 0x76, 0xcc, 0x60,
	0x01, 0xfd, 0xca, 0xf1, 0xe4, 0x5c, 0xbc, 0xc9, 0x45, 0x1c, 0xe0, 0xb3, 0xa0, 0xbf, 0x67, 0x3f,
	0xf3, 0xe8, 0x4a, 0xa0, 0xb5, 0x28, 0xd2, 0x72, 0xb6, 0xf0, 0x4c, 0xbf, 0x86, 0x41, 0xcd, 0xd1,
	0xf0, 0x49, 0x8d, 0xe1, 0x87, 0x55, 0xe2, 0x2b, 0xad, 0x4a, 0xa2, 0x8f, 0x60, 0x78, 0xcb, 0x33,
	0x8e, 0x05, 0x57, 0xc9, 0xf3, 0x0b, 0xe8, 0x21, 0x43, 0x86, 0x76, 0xa8, 0xec, 0x2e, 0xbe, 0xeb,
	0x03, 0x53, 0x35, 0x33, 0x4d, 0x51, 0x2e, 0x80, 0xcb, 0x71, 0x87, 0x2f, 0xcf, 0x7f, 0xa1, 0x4b,
	0xa9, 0x57, 0xfc, 0x7e, 0x32, 0x9d, 0x4e, 0x82, 0xf8, 0x73, 0xfc, 0x17, 0x37, 0x9d, 0xee, 0x7e,
	0xf1, 0xea, 0xfb, 0x23, 0xfc, 0xbb, 0x36, 0xfd, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x14, 0x32, 0xbd,
	0xe0, 0xea, 0x09, 0x00, 0x00,
}
