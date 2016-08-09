package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type releaseTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *releaseTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("release").
func (v *releaseTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *releaseTable) Columns() []string {
	return []string{"id", "gid", "name", "artist_credit", "release_group", "status", "packaging", "language", "script", "barcode", "comment", "edits_pending", "quality", "last_updated"}
}

// NewStruct makes a new struct for that view or table.
func (v *releaseTable) NewStruct() reform.Struct {
	return new(Release)
}

// NewRecord makes a new record for that table.
func (v *releaseTable) NewRecord() reform.Record {
	return new(Release)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *releaseTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ReleaseTable represents release view or table in SQL database.
var ReleaseTable = &releaseTable{
	s: parse.StructInfo{Type: "Release", SQLSchema: "", SQLName: "release", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "ArtistCreditId", PKType: "", Column: "artist_credit"}, {Name: "ReleaseGroupId", PKType: "", Column: "release_group"}, {Name: "StatusId", PKType: "", Column: "status"}, {Name: "PackagingId", PKType: "", Column: "packaging"}, {Name: "LanguageId", PKType: "", Column: "language"}, {Name: "ScriptId", PKType: "", Column: "script"}, {Name: "Barcode", PKType: "", Column: "barcode"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "Quality", PKType: "", Column: "quality"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}}, PKFieldIndex: 0},
	z: new(Release).Values(),
}

// String returns a string representation of this struct or record.
func (s Release) String() string {
	res := make([]string, 14)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "ArtistCreditId: " + reform.Inspect(s.ArtistCreditId, true)
	res[4] = "ReleaseGroupId: " + reform.Inspect(s.ReleaseGroupId, true)
	res[5] = "StatusId: " + reform.Inspect(s.StatusId, true)
	res[6] = "PackagingId: " + reform.Inspect(s.PackagingId, true)
	res[7] = "LanguageId: " + reform.Inspect(s.LanguageId, true)
	res[8] = "ScriptId: " + reform.Inspect(s.ScriptId, true)
	res[9] = "Barcode: " + reform.Inspect(s.Barcode, true)
	res[10] = "Comment: " + reform.Inspect(s.Comment, true)
	res[11] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[12] = "Quality: " + reform.Inspect(s.Quality, true)
	res[13] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Release) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.ArtistCreditId,
		s.ReleaseGroupId,
		s.StatusId,
		s.PackagingId,
		s.LanguageId,
		s.ScriptId,
		s.Barcode,
		s.Comment,
		s.EditsPending,
		s.Quality,
		s.LastUpdated,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Release) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.ArtistCreditId,
		&s.ReleaseGroupId,
		&s.StatusId,
		&s.PackagingId,
		&s.LanguageId,
		&s.ScriptId,
		&s.Barcode,
		&s.Comment,
		&s.EditsPending,
		&s.Quality,
		&s.LastUpdated,
	}
}

// View returns View object for that struct.
func (s *Release) View() reform.View {
	return ReleaseTable
}

// Table returns Table object for that record.
func (s *Release) Table() reform.Table {
	return ReleaseTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Release) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Release) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Release) HasPK() bool {
	return s.ID != ReleaseTable.z[ReleaseTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Release) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = ReleaseTable
	_ reform.Struct = new(Release)
	_ reform.Table  = ReleaseTable
	_ reform.Record = new(Release)
	_ fmt.Stringer  = new(Release)
)

func init() {
	parse.AssertUpToDate(&ReleaseTable.s, new(Release))
}
