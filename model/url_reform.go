package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type urlTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *urlTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("url").
func (v *urlTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *urlTable) Columns() []string {
	return []string{"id", "gid", "url", "edits_pending", "last_updated"}
}

// NewStruct makes a new struct for that view or table.
func (v *urlTable) NewStruct() reform.Struct {
	return new(Url)
}

// NewRecord makes a new record for that table.
func (v *urlTable) NewRecord() reform.Record {
	return new(Url)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *urlTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// UrlTable represents url view or table in SQL database.
var UrlTable = &urlTable{
	s: parse.StructInfo{Type: "Url", SQLSchema: "", SQLName: "url", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Url", PKType: "", Column: "url"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}}, PKFieldIndex: 0},
	z: new(Url).Values(),
}

// String returns a string representation of this struct or record.
func (s Url) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Url: " + reform.Inspect(s.Url, true)
	res[3] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[4] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Url) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Url,
		s.EditsPending,
		s.LastUpdated,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Url) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Url,
		&s.EditsPending,
		&s.LastUpdated,
	}
}

// View returns View object for that struct.
func (s *Url) View() reform.View {
	return UrlTable
}

// Table returns Table object for that record.
func (s *Url) Table() reform.Table {
	return UrlTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Url) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Url) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Url) HasPK() bool {
	return s.ID != UrlTable.z[UrlTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Url) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = UrlTable
	_ reform.Struct = new(Url)
	_ reform.Table  = UrlTable
	_ reform.Record = new(Url)
	_ fmt.Stringer  = new(Url)
)

func init() {
	parse.AssertUpToDate(&UrlTable.s, new(Url))
}
