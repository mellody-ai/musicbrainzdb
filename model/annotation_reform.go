package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type annotationTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *annotationTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("annotation").
func (v *annotationTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *annotationTable) Columns() []string {
	return []string{"id", "editor", "text", "changelog", "created"}
}

// NewStruct makes a new struct for that view or table.
func (v *annotationTable) NewStruct() reform.Struct {
	return new(Annotation)
}

// NewRecord makes a new record for that table.
func (v *annotationTable) NewRecord() reform.Record {
	return new(Annotation)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *annotationTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// AnnotationTable represents annotation view or table in SQL database.
var AnnotationTable = &annotationTable{
	s: parse.StructInfo{Type: "Annotation", SQLSchema: "", SQLName: "annotation", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "Editor", PKType: "", Column: "editor"}, {Name: "Text", PKType: "", Column: "text"}, {Name: "Changelog", PKType: "", Column: "changelog"}, {Name: "Created", PKType: "", Column: "created"}}, PKFieldIndex: 0},
	z: new(Annotation).Values(),
}

// String returns a string representation of this struct or record.
func (s Annotation) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Editor: " + reform.Inspect(s.Editor, true)
	res[2] = "Text: " + reform.Inspect(s.Text, true)
	res[3] = "Changelog: " + reform.Inspect(s.Changelog, true)
	res[4] = "Created: " + reform.Inspect(s.Created, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Annotation) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Editor,
		s.Text,
		s.Changelog,
		s.Created,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Annotation) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Editor,
		&s.Text,
		&s.Changelog,
		&s.Created,
	}
}

// View returns View object for that struct.
func (s *Annotation) View() reform.View {
	return AnnotationTable
}

// Table returns Table object for that record.
func (s *Annotation) Table() reform.Table {
	return AnnotationTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Annotation) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Annotation) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Annotation) HasPK() bool {
	return s.ID != AnnotationTable.z[AnnotationTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Annotation) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = AnnotationTable
	_ reform.Struct = new(Annotation)
	_ reform.Table  = AnnotationTable
	_ reform.Record = new(Annotation)
	_ fmt.Stringer  = new(Annotation)
)

func init() {
	parse.AssertUpToDate(&AnnotationTable.s, new(Annotation))
}
