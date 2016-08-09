package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type areaTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *areaTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("area").
func (v *areaTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *areaTable) Columns() []string {
	return []string{"id", "gid", "name", "type", "edits_pending", "last_updated", "begin_date_year", "begin_date_month", "begin_date_day", "end_date_year", "end_date_month", "end_date_day", "ended", "comment"}
}

// NewStruct makes a new struct for that view or table.
func (v *areaTable) NewStruct() reform.Struct {
	return new(Area)
}

// NewRecord makes a new record for that table.
func (v *areaTable) NewRecord() reform.Record {
	return new(Area)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *areaTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// AreaTable represents area view or table in SQL database.
var AreaTable = &areaTable{
	s: parse.StructInfo{Type: "Area", SQLSchema: "", SQLName: "area", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "TypeId", PKType: "", Column: "type"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "BeginDateYear", PKType: "", Column: "begin_date_year"}, {Name: "BeginDateMonth", PKType: "", Column: "begin_date_month"}, {Name: "BeginDateDay", PKType: "", Column: "begin_date_day"}, {Name: "EndDateYear", PKType: "", Column: "end_date_year"}, {Name: "EndDateMonth", PKType: "", Column: "end_date_month"}, {Name: "EndDateDay", PKType: "", Column: "end_date_day"}, {Name: "Ended", PKType: "", Column: "ended"}, {Name: "Comment", PKType: "", Column: "comment"}}, PKFieldIndex: 0},
	z: new(Area).Values(),
}

// String returns a string representation of this struct or record.
func (s Area) String() string {
	res := make([]string, 14)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "TypeId: " + reform.Inspect(s.TypeId, true)
	res[4] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[5] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[6] = "BeginDateYear: " + reform.Inspect(s.BeginDateYear, true)
	res[7] = "BeginDateMonth: " + reform.Inspect(s.BeginDateMonth, true)
	res[8] = "BeginDateDay: " + reform.Inspect(s.BeginDateDay, true)
	res[9] = "EndDateYear: " + reform.Inspect(s.EndDateYear, true)
	res[10] = "EndDateMonth: " + reform.Inspect(s.EndDateMonth, true)
	res[11] = "EndDateDay: " + reform.Inspect(s.EndDateDay, true)
	res[12] = "Ended: " + reform.Inspect(s.Ended, true)
	res[13] = "Comment: " + reform.Inspect(s.Comment, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Area) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.TypeId,
		s.EditsPending,
		s.LastUpdated,
		s.BeginDateYear,
		s.BeginDateMonth,
		s.BeginDateDay,
		s.EndDateYear,
		s.EndDateMonth,
		s.EndDateDay,
		s.Ended,
		s.Comment,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Area) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.TypeId,
		&s.EditsPending,
		&s.LastUpdated,
		&s.BeginDateYear,
		&s.BeginDateMonth,
		&s.BeginDateDay,
		&s.EndDateYear,
		&s.EndDateMonth,
		&s.EndDateDay,
		&s.Ended,
		&s.Comment,
	}
}

// View returns View object for that struct.
func (s *Area) View() reform.View {
	return AreaTable
}

// Table returns Table object for that record.
func (s *Area) Table() reform.Table {
	return AreaTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Area) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Area) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Area) HasPK() bool {
	return s.ID != AreaTable.z[AreaTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Area) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = AreaTable
	_ reform.Struct = new(Area)
	_ reform.Table  = AreaTable
	_ reform.Record = new(Area)
	_ fmt.Stringer  = new(Area)
)

func init() {
	parse.AssertUpToDate(&AreaTable.s, new(Area))
}
