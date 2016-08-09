package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type placeTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *placeTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("place").
func (v *placeTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *placeTable) Columns() []string {
	return []string{"id", "gid", "name", "type", "area", "address", "begin_date_year", "begin_date_month", "begin_date_day", "end_date_year", "end_date_month", "end_date_day", "comment", "edits_pending", "last_updated", "ended"}
}

// NewStruct makes a new struct for that view or table.
func (v *placeTable) NewStruct() reform.Struct {
	return new(Place)
}

// NewRecord makes a new record for that table.
func (v *placeTable) NewRecord() reform.Record {
	return new(Place)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *placeTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// PlaceTable represents place view or table in SQL database.
var PlaceTable = &placeTable{
	s: parse.StructInfo{Type: "Place", SQLSchema: "", SQLName: "place", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "TypeId", PKType: "", Column: "type"}, {Name: "AreaId", PKType: "", Column: "area"}, {Name: "Address", PKType: "", Column: "address"}, {Name: "BeginDateYear", PKType: "", Column: "begin_date_year"}, {Name: "BeginDateMonth", PKType: "", Column: "begin_date_month"}, {Name: "BeginDateDay", PKType: "", Column: "begin_date_day"}, {Name: "EndDateYear", PKType: "", Column: "end_date_year"}, {Name: "EndDateMonth", PKType: "", Column: "end_date_month"}, {Name: "EndDateDay", PKType: "", Column: "end_date_day"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "Ended", PKType: "", Column: "ended"}}, PKFieldIndex: 0},
	z: new(Place).Values(),
}

// String returns a string representation of this struct or record.
func (s Place) String() string {
	res := make([]string, 16)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "TypeId: " + reform.Inspect(s.TypeId, true)
	res[4] = "AreaId: " + reform.Inspect(s.AreaId, true)
	res[5] = "Address: " + reform.Inspect(s.Address, true)
	res[6] = "BeginDateYear: " + reform.Inspect(s.BeginDateYear, true)
	res[7] = "BeginDateMonth: " + reform.Inspect(s.BeginDateMonth, true)
	res[8] = "BeginDateDay: " + reform.Inspect(s.BeginDateDay, true)
	res[9] = "EndDateYear: " + reform.Inspect(s.EndDateYear, true)
	res[10] = "EndDateMonth: " + reform.Inspect(s.EndDateMonth, true)
	res[11] = "EndDateDay: " + reform.Inspect(s.EndDateDay, true)
	res[12] = "Comment: " + reform.Inspect(s.Comment, true)
	res[13] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[14] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[15] = "Ended: " + reform.Inspect(s.Ended, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Place) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.TypeId,
		s.AreaId,
		s.Address,
		s.BeginDateYear,
		s.BeginDateMonth,
		s.BeginDateDay,
		s.EndDateYear,
		s.EndDateMonth,
		s.EndDateDay,
		s.Comment,
		s.EditsPending,
		s.LastUpdated,
		s.Ended,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Place) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.TypeId,
		&s.AreaId,
		&s.Address,
		&s.BeginDateYear,
		&s.BeginDateMonth,
		&s.BeginDateDay,
		&s.EndDateYear,
		&s.EndDateMonth,
		&s.EndDateDay,
		&s.Comment,
		&s.EditsPending,
		&s.LastUpdated,
		&s.Ended,
	}
}

// View returns View object for that struct.
func (s *Place) View() reform.View {
	return PlaceTable
}

// Table returns Table object for that record.
func (s *Place) Table() reform.Table {
	return PlaceTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Place) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Place) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Place) HasPK() bool {
	return s.ID != PlaceTable.z[PlaceTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Place) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = PlaceTable
	_ reform.Struct = new(Place)
	_ reform.Table  = PlaceTable
	_ reform.Record = new(Place)
	_ fmt.Stringer  = new(Place)
)

func init() {
	parse.AssertUpToDate(&PlaceTable.s, new(Place))
}
