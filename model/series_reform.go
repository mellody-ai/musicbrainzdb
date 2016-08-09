package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type seriesTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *seriesTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("series").
func (v *seriesTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *seriesTable) Columns() []string {
	return []string{"id", "gid", "name", "comment", "type", "ordering_attribute", "ordering_type", "edits_pending", "last_updated"}
}

// NewStruct makes a new struct for that view or table.
func (v *seriesTable) NewStruct() reform.Struct {
	return new(Series)
}

// NewRecord makes a new record for that table.
func (v *seriesTable) NewRecord() reform.Record {
	return new(Series)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *seriesTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// SeriesTable represents series view or table in SQL database.
var SeriesTable = &seriesTable{
	s: parse.StructInfo{Type: "Series", SQLSchema: "", SQLName: "series", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "Type", PKType: "", Column: "type"}, {Name: "OrderingAttributeId", PKType: "", Column: "ordering_attribute"}, {Name: "OrderingTypeId", PKType: "", Column: "ordering_type"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}}, PKFieldIndex: 0},
	z: new(Series).Values(),
}

// String returns a string representation of this struct or record.
func (s Series) String() string {
	res := make([]string, 9)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "Comment: " + reform.Inspect(s.Comment, true)
	res[4] = "Type: " + reform.Inspect(s.Type, true)
	res[5] = "OrderingAttributeId: " + reform.Inspect(s.OrderingAttributeId, true)
	res[6] = "OrderingTypeId: " + reform.Inspect(s.OrderingTypeId, true)
	res[7] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[8] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Series) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.Comment,
		s.Type,
		s.OrderingAttributeId,
		s.OrderingTypeId,
		s.EditsPending,
		s.LastUpdated,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Series) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.Comment,
		&s.Type,
		&s.OrderingAttributeId,
		&s.OrderingTypeId,
		&s.EditsPending,
		&s.LastUpdated,
	}
}

// View returns View object for that struct.
func (s *Series) View() reform.View {
	return SeriesTable
}

// Table returns Table object for that record.
func (s *Series) Table() reform.Table {
	return SeriesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Series) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Series) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Series) HasPK() bool {
	return s.ID != SeriesTable.z[SeriesTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Series) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = SeriesTable
	_ reform.Struct = new(Series)
	_ reform.Table  = SeriesTable
	_ reform.Record = new(Series)
	_ fmt.Stringer  = new(Series)
)

func init() {
	parse.AssertUpToDate(&SeriesTable.s, new(Series))
}
