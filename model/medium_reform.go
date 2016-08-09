package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type mediumTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *mediumTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("medium").
func (v *mediumTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *mediumTable) Columns() []string {
	return []string{"id", "release", "position", "format", "name", "edits_pending", "last_updated", "track_count"}
}

// NewStruct makes a new struct for that view or table.
func (v *mediumTable) NewStruct() reform.Struct {
	return new(Medium)
}

// NewRecord makes a new record for that table.
func (v *mediumTable) NewRecord() reform.Record {
	return new(Medium)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *mediumTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// MediumTable represents medium view or table in SQL database.
var MediumTable = &mediumTable{
	s: parse.StructInfo{Type: "Medium", SQLSchema: "", SQLName: "medium", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "ReleaseId", PKType: "", Column: "release"}, {Name: "Position", PKType: "", Column: "position"}, {Name: "FormatId", PKType: "", Column: "format"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "TrackCount", PKType: "", Column: "track_count"}}, PKFieldIndex: 0},
	z: new(Medium).Values(),
}

// String returns a string representation of this struct or record.
func (s Medium) String() string {
	res := make([]string, 8)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "ReleaseId: " + reform.Inspect(s.ReleaseId, true)
	res[2] = "Position: " + reform.Inspect(s.Position, true)
	res[3] = "FormatId: " + reform.Inspect(s.FormatId, true)
	res[4] = "Name: " + reform.Inspect(s.Name, true)
	res[5] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[6] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[7] = "TrackCount: " + reform.Inspect(s.TrackCount, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Medium) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.ReleaseId,
		s.Position,
		s.FormatId,
		s.Name,
		s.EditsPending,
		s.LastUpdated,
		s.TrackCount,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Medium) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.ReleaseId,
		&s.Position,
		&s.FormatId,
		&s.Name,
		&s.EditsPending,
		&s.LastUpdated,
		&s.TrackCount,
	}
}

// View returns View object for that struct.
func (s *Medium) View() reform.View {
	return MediumTable
}

// Table returns Table object for that record.
func (s *Medium) Table() reform.Table {
	return MediumTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Medium) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Medium) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Medium) HasPK() bool {
	return s.ID != MediumTable.z[MediumTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Medium) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = MediumTable
	_ reform.Struct = new(Medium)
	_ reform.Table  = MediumTable
	_ reform.Record = new(Medium)
	_ fmt.Stringer  = new(Medium)
)

func init() {
	parse.AssertUpToDate(&MediumTable.s, new(Medium))
}
