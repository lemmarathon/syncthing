// Copyright (C) 2014 Jakob Borg and Contributors (see the CONTRIBUTORS file).
// All rights reserved. Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package protocol

import (
	"bytes"
	"io"

	"github.com/calmh/syncthing/xdr"
)

/*

IndexMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                     Length of Repository                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                 Repository (variable length)                  \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Number of Files                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\               Zero or more FileInfo Structures                \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct IndexMessage {
	string Repository<64>;
	FileInfo Files<>;
}

*/

func (o IndexMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o IndexMessage) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o IndexMessage) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o IndexMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.Repository) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.Repository)
	xw.WriteUint32(uint32(len(o.Files)))
	for i := range o.Files {
		o.Files[i].encodeXDR(xw)
	}
	return xw.Tot(), xw.Error()
}

func (o *IndexMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *IndexMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *IndexMessage) decodeXDR(xr *xdr.Reader) error {
	o.Repository = xr.ReadStringMax(64)
	_FilesSize := int(xr.ReadUint32())
	o.Files = make([]FileInfo, _FilesSize)
	for i := range o.Files {
		(&o.Files[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

FileInfo Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Name                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Name (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Flags                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                      Modified (64 bits)                       +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                       Version (64 bits)                       +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                    Local Version (64 bits)                    +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Blocks                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\               Zero or more BlockInfo Structures               \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct FileInfo {
	string Name<1024>;
	unsigned int Flags;
	hyper Modified;
	unsigned hyper Version;
	unsigned hyper LocalVersion;
	BlockInfo Blocks<>;
}

*/

func (o FileInfo) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o FileInfo) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o FileInfo) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o FileInfo) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.Name) > 1024 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.Name)
	xw.WriteUint32(o.Flags)
	xw.WriteUint64(uint64(o.Modified))
	xw.WriteUint64(o.Version)
	xw.WriteUint64(o.LocalVersion)
	xw.WriteUint32(uint32(len(o.Blocks)))
	for i := range o.Blocks {
		o.Blocks[i].encodeXDR(xw)
	}
	return xw.Tot(), xw.Error()
}

func (o *FileInfo) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *FileInfo) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *FileInfo) decodeXDR(xr *xdr.Reader) error {
	o.Name = xr.ReadStringMax(1024)
	o.Flags = xr.ReadUint32()
	o.Modified = int64(xr.ReadUint64())
	o.Version = xr.ReadUint64()
	o.LocalVersion = xr.ReadUint64()
	_BlocksSize := int(xr.ReadUint32())
	o.Blocks = make([]BlockInfo, _BlocksSize)
	for i := range o.Blocks {
		(&o.Blocks[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

BlockInfo Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Size                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Hash                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Hash (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct BlockInfo {
	unsigned int Size;
	opaque Hash<64>;
}

*/

func (o BlockInfo) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o BlockInfo) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o BlockInfo) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o BlockInfo) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteUint32(o.Size)
	if len(o.Hash) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteBytes(o.Hash)
	return xw.Tot(), xw.Error()
}

func (o *BlockInfo) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *BlockInfo) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *BlockInfo) decodeXDR(xr *xdr.Reader) error {
	o.Size = xr.ReadUint32()
	o.Hash = xr.ReadBytesMax(64)
	return xr.Error()
}

/*

RequestMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                     Length of Repository                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                 Repository (variable length)                  \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Name                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Name (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                       Offset (64 bits)                        +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Size                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct RequestMessage {
	string Repository<64>;
	string Name<1024>;
	unsigned hyper Offset;
	unsigned int Size;
}

*/

func (o RequestMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o RequestMessage) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o RequestMessage) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o RequestMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.Repository) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.Repository)
	if len(o.Name) > 1024 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.Name)
	xw.WriteUint64(o.Offset)
	xw.WriteUint32(o.Size)
	return xw.Tot(), xw.Error()
}

func (o *RequestMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *RequestMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *RequestMessage) decodeXDR(xr *xdr.Reader) error {
	o.Repository = xr.ReadStringMax(64)
	o.Name = xr.ReadStringMax(1024)
	o.Offset = xr.ReadUint64()
	o.Size = xr.ReadUint32()
	return xr.Error()
}

/*

ClusterConfigMessage Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                     Length of Client Name                     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                 Client Name (variable length)                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                   Length of Client Version                    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\               Client Version (variable length)                \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Number of Repositories                     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\              Zero or more Repository Structures               \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Number of Options                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Option Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct ClusterConfigMessage {
	string ClientName<64>;
	string ClientVersion<64>;
	Repository Repositories<64>;
	Option Options<64>;
}

*/

func (o ClusterConfigMessage) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o ClusterConfigMessage) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o ClusterConfigMessage) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o ClusterConfigMessage) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.ClientName) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.ClientName)
	if len(o.ClientVersion) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.ClientVersion)
	if len(o.Repositories) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteUint32(uint32(len(o.Repositories)))
	for i := range o.Repositories {
		o.Repositories[i].encodeXDR(xw)
	}
	if len(o.Options) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteUint32(uint32(len(o.Options)))
	for i := range o.Options {
		o.Options[i].encodeXDR(xw)
	}
	return xw.Tot(), xw.Error()
}

func (o *ClusterConfigMessage) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *ClusterConfigMessage) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *ClusterConfigMessage) decodeXDR(xr *xdr.Reader) error {
	o.ClientName = xr.ReadStringMax(64)
	o.ClientVersion = xr.ReadStringMax(64)
	_RepositoriesSize := int(xr.ReadUint32())
	if _RepositoriesSize > 64 {
		return xdr.ErrElementSizeExceeded
	}
	o.Repositories = make([]Repository, _RepositoriesSize)
	for i := range o.Repositories {
		(&o.Repositories[i]).decodeXDR(xr)
	}
	_OptionsSize := int(xr.ReadUint32())
	if _OptionsSize > 64 {
		return xdr.ErrElementSizeExceeded
	}
	o.Options = make([]Option, _OptionsSize)
	for i := range o.Options {
		(&o.Options[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

Repository Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of ID                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     ID (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Number of Nodes                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                 Zero or more Node Structures                  \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Repository {
	string ID<64>;
	Node Nodes<64>;
}

*/

func (o Repository) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Repository) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Repository) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o Repository) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.ID) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.ID)
	if len(o.Nodes) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteUint32(uint32(len(o.Nodes)))
	for i := range o.Nodes {
		o.Nodes[i].encodeXDR(xw)
	}
	return xw.Tot(), xw.Error()
}

func (o *Repository) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Repository) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Repository) decodeXDR(xr *xdr.Reader) error {
	o.ID = xr.ReadStringMax(64)
	_NodesSize := int(xr.ReadUint32())
	if _NodesSize > 64 {
		return xdr.ErrElementSizeExceeded
	}
	o.Nodes = make([]Node, _NodesSize)
	for i := range o.Nodes {
		(&o.Nodes[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

Node Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of ID                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     ID (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Flags                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
+                  Max Local Version (64 bits)                  +
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Node {
	opaque ID<32>;
	unsigned int Flags;
	unsigned hyper MaxLocalVersion;
}

*/

func (o Node) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Node) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Node) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o Node) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.ID) > 32 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteBytes(o.ID)
	xw.WriteUint32(o.Flags)
	xw.WriteUint64(o.MaxLocalVersion)
	return xw.Tot(), xw.Error()
}

func (o *Node) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Node) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Node) decodeXDR(xr *xdr.Reader) error {
	o.ID = xr.ReadBytesMax(32)
	o.Flags = xr.ReadUint32()
	o.MaxLocalVersion = xr.ReadUint64()
	return xr.Error()
}

/*

Option Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of Key                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     Key (variable length)                     \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Length of Value                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                    Value (variable length)                    \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Option {
	string Key<64>;
	string Value<1024>;
}

*/

func (o Option) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Option) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Option) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o Option) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.Key) > 64 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.Key)
	if len(o.Value) > 1024 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteString(o.Value)
	return xw.Tot(), xw.Error()
}

func (o *Option) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Option) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Option) decodeXDR(xr *xdr.Reader) error {
	o.Key = xr.ReadStringMax(64)
	o.Value = xr.ReadStringMax(1024)
	return xr.Error()
}
