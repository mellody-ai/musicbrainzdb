package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type labelTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *labelTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("label").
func (v *labelTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *labelTable) Columns() []string {
	return []string{"id", "gid", "name", "begin_date_year", "begin_date_month", "begin_date_day", "end_date_year", "end_date_month", "end_date_day", "label_code", "type", "area", "comment", "edits_pending", "last_updated", "ended"}
}

// NewStruct makes a new struct for that view or table.
func (v *labelTable) NewStruct() reform.Struct {
	return new(Label)
}

// NewRecord makes a new record for that table.
func (v *labelTable) NewRecord() reform.Record {
	return new(Label)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *labelTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// LabelTable represents label view or table in SQL database.
var LabelTable = &labelTable{
	s: parse.StructInfo{Type: "Label", SQLSchema: "", SQLName: "label", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "BeginDateYear", PKType: "", Column: "begin_date_year"}, {Name: "BeginDateMonth", PKType: "", Column: "begin_date_month"}, {Name: "BeginDateDay", PKType: "", Column: "begin_date_day"}, {Name: "EndDateYear", PKType: "", Column: "end_date_year"}, {Name: "EndDateMonth", PKType: "", Column: "end_date_month"}, {Name: "EndDateDay", PKType: "", Column: "end_date_day"}, {Name: "LabelCode", PKType: "", Column: "label_code"}, {Name: "TypeId", PKType: "", Column: "type"}, {Name: "AreaId", PKType: "", Column: "area"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "Ended", PKType: "", Column: "ended"}}, PKFieldIndex: 0},
	z: new(Label).Values(),
}

// String returns a string representation of this struct or record.
func (s Label) String() string {
	res := make([]string, 16)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "BeginDateYear: " + reform.Inspect(s.BeginDateYear, true)
	res[4] = "BeginDateMonth: " + reform.Inspect(s.BeginDateMonth, true)
	res[5] = "BeginDateDay: " + reform.Inspect(s.BeginDateDay, true)
	res[6] = "EndDateYear: " + reform.Inspect(s.EndDateYear, true)
	res[7] = "EndDateMonth: " + reform.Inspect(s.EndDateMonth, true)
	res[8] = "EndDateDay: " + reform.Inspect(s.EndDateDay, true)
	res[9] = "LabelCode: " + reform.Inspect(s.LabelCode, true)
	res[10] = "TypeId: " + reform.Inspect(s.TypeId, true)
	res[11] = "AreaId: " + reform.Inspect(s.AreaId, true)
	res[12] = "Comment: " + reform.Inspect(s.Comment, true)
	res[13] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[14] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[15] = "Ended: " + reform.Inspect(s.Ended, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Label) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.BeginDateYear,
		s.BeginDateMonth,
		s.BeginDateDay,
		s.EndDateYear,
		s.EndDateMonth,
		s.EndDateDay,
		s.LabelCode,
		s.TypeId,
		s.AreaId,
		s.Comment,
		s.EditsPending,
		s.LastUpdated,
		s.Ended,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Label) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.BeginDateYear,
		&s.BeginDateMonth,
		&s.BeginDateDay,
		&s.EndDateYear,
		&s.EndDateMonth,
		&s.EndDateDay,
		&s.LabelCode,
		&s.TypeId,
		&s.AreaId,
		&s.Comment,
		&s.EditsPending,
		&s.LastUpdated,
		&s.Ended,
	}
}

// View returns View object for that struct.
func (s *Label) View() reform.View {
	return LabelTable
}

// Table returns Table object for that record.
func (s *Label) Table() reform.Table {
	return LabelTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Label) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Label) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Label) HasPK() bool {
	return s.ID != LabelTable.z[LabelTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Label) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = LabelTable
	_ reform.Struct = new(Label)
	_ reform.Table  = LabelTable
	_ reform.Record = new(Label)
	_ fmt.Stringer  = new(Label)
)

func init() {
	parse.AssertUpToDate(&LabelTable.s, new(Label))
}
