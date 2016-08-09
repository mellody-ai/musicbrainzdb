package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type releaseGroupTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *releaseGroupTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("release_group").
func (v *releaseGroupTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *releaseGroupTable) Columns() []string {
	return []string{"id", "gid", "name", "artist_credit", "type", "comment", "edits_pending", "last_updated"}
}

// NewStruct makes a new struct for that view or table.
func (v *releaseGroupTable) NewStruct() reform.Struct {
	return new(ReleaseGroup)
}

// NewRecord makes a new record for that table.
func (v *releaseGroupTable) NewRecord() reform.Record {
	return new(ReleaseGroup)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *releaseGroupTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ReleaseGroupTable represents release_group view or table in SQL database.
var ReleaseGroupTable = &releaseGroupTable{
	s: parse.StructInfo{Type: "ReleaseGroup", SQLSchema: "", SQLName: "release_group", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "ArtistCreditId", PKType: "", Column: "artist_credit"}, {Name: "Type", PKType: "", Column: "type"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}}, PKFieldIndex: 0},
	z: new(ReleaseGroup).Values(),
}

// String returns a string representation of this struct or record.
func (s ReleaseGroup) String() string {
	res := make([]string, 8)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "ArtistCreditId: " + reform.Inspect(s.ArtistCreditId, true)
	res[4] = "Type: " + reform.Inspect(s.Type, true)
	res[5] = "Comment: " + reform.Inspect(s.Comment, true)
	res[6] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[7] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ReleaseGroup) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.ArtistCreditId,
		s.Type,
		s.Comment,
		s.EditsPending,
		s.LastUpdated,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ReleaseGroup) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.ArtistCreditId,
		&s.Type,
		&s.Comment,
		&s.EditsPending,
		&s.LastUpdated,
	}
}

// View returns View object for that struct.
func (s *ReleaseGroup) View() reform.View {
	return ReleaseGroupTable
}

// Table returns Table object for that record.
func (s *ReleaseGroup) Table() reform.Table {
	return ReleaseGroupTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *ReleaseGroup) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *ReleaseGroup) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *ReleaseGroup) HasPK() bool {
	return s.ID != ReleaseGroupTable.z[ReleaseGroupTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *ReleaseGroup) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = ReleaseGroupTable
	_ reform.Struct = new(ReleaseGroup)
	_ reform.Table  = ReleaseGroupTable
	_ reform.Record = new(ReleaseGroup)
	_ fmt.Stringer  = new(ReleaseGroup)
)

func init() {
	parse.AssertUpToDate(&ReleaseGroupTable.s, new(ReleaseGroup))
}
