package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type recordingTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *recordingTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("recording").
func (v *recordingTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *recordingTable) Columns() []string {
	return []string{"id", "gid", "name", "artist_credit", "length", "comment", "edits_pending", "last_updated", "video"}
}

// NewStruct makes a new struct for that view or table.
func (v *recordingTable) NewStruct() reform.Struct {
	return new(Recording)
}

// NewRecord makes a new record for that table.
func (v *recordingTable) NewRecord() reform.Record {
	return new(Recording)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *recordingTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// RecordingTable represents recording view or table in SQL database.
var RecordingTable = &recordingTable{
	s: parse.StructInfo{Type: "Recording", SQLSchema: "", SQLName: "recording", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "ArtistCreditId", PKType: "", Column: "artist_credit"}, {Name: "Length", PKType: "", Column: "length"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "Video", PKType: "", Column: "video"}}, PKFieldIndex: 0},
	z: new(Recording).Values(),
}

// String returns a string representation of this struct or record.
func (s Recording) String() string {
	res := make([]string, 9)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "ArtistCreditId: " + reform.Inspect(s.ArtistCreditId, true)
	res[4] = "Length: " + reform.Inspect(s.Length, true)
	res[5] = "Comment: " + reform.Inspect(s.Comment, true)
	res[6] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[7] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[8] = "Video: " + reform.Inspect(s.Video, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Recording) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.ArtistCreditId,
		s.Length,
		s.Comment,
		s.EditsPending,
		s.LastUpdated,
		s.Video,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Recording) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.ArtistCreditId,
		&s.Length,
		&s.Comment,
		&s.EditsPending,
		&s.LastUpdated,
		&s.Video,
	}
}

// View returns View object for that struct.
func (s *Recording) View() reform.View {
	return RecordingTable
}

// Table returns Table object for that record.
func (s *Recording) Table() reform.Table {
	return RecordingTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Recording) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Recording) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Recording) HasPK() bool {
	return s.ID != RecordingTable.z[RecordingTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Recording) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = RecordingTable
	_ reform.Struct = new(Recording)
	_ reform.Table  = RecordingTable
	_ reform.Record = new(Recording)
	_ fmt.Stringer  = new(Recording)
)

func init() {
	parse.AssertUpToDate(&RecordingTable.s, new(Recording))
}
