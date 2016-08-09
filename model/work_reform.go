package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type workTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *workTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("work").
func (v *workTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *workTable) Columns() []string {
	return []string{"id", "gid", "name", "type", "comment", "edits_pending", "last_updated", "language"}
}

// NewStruct makes a new struct for that view or table.
func (v *workTable) NewStruct() reform.Struct {
	return new(Work)
}

// NewRecord makes a new record for that table.
func (v *workTable) NewRecord() reform.Record {
	return new(Work)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *workTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// WorkTable represents work view or table in SQL database.
var WorkTable = &workTable{
	s: parse.StructInfo{Type: "Work", SQLSchema: "", SQLName: "work", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "Type", PKType: "", Column: "type"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "LanguageId", PKType: "", Column: "language"}}, PKFieldIndex: 0},
	z: new(Work).Values(),
}

// String returns a string representation of this struct or record.
func (s Work) String() string {
	res := make([]string, 8)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "Type: " + reform.Inspect(s.Type, true)
	res[4] = "Comment: " + reform.Inspect(s.Comment, true)
	res[5] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[6] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[7] = "LanguageId: " + reform.Inspect(s.LanguageId, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Work) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.Type,
		s.Comment,
		s.EditsPending,
		s.LastUpdated,
		s.LanguageId,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Work) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.Type,
		&s.Comment,
		&s.EditsPending,
		&s.LastUpdated,
		&s.LanguageId,
	}
}

// View returns View object for that struct.
func (s *Work) View() reform.View {
	return WorkTable
}

// Table returns Table object for that record.
func (s *Work) Table() reform.Table {
	return WorkTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Work) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Work) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Work) HasPK() bool {
	return s.ID != WorkTable.z[WorkTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Work) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = WorkTable
	_ reform.Struct = new(Work)
	_ reform.Table  = WorkTable
	_ reform.Record = new(Work)
	_ fmt.Stringer  = new(Work)
)

func init() {
	parse.AssertUpToDate(&WorkTable.s, new(Work))
}
