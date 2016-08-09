package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type artistTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *artistTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("artist").
func (v *artistTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *artistTable) Columns() []string {
	return []string{"id", "gid", "name", "sort_name", "begin_date_year", "begin_date_month", "begin_date_day", "end_date_year", "end_date_month", "end_date_day", "type", "area", "gender", "comment", "edits_pending", "last_updated", "ended", "begin_area", "end_area"}
}

// NewStruct makes a new struct for that view or table.
func (v *artistTable) NewStruct() reform.Struct {
	return new(Artist)
}

// NewRecord makes a new record for that table.
func (v *artistTable) NewRecord() reform.Record {
	return new(Artist)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *artistTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ArtistTable represents artist view or table in SQL database.
var ArtistTable = &artistTable{
	s: parse.StructInfo{Type: "Artist", SQLSchema: "", SQLName: "artist", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "SortName", PKType: "", Column: "sort_name"}, {Name: "BeginDateYear", PKType: "", Column: "begin_date_year"}, {Name: "BeginDateMonth", PKType: "", Column: "begin_date_month"}, {Name: "BeginDateDay", PKType: "", Column: "begin_date_day"}, {Name: "EndDateYear", PKType: "", Column: "end_date_year"}, {Name: "EndDateMonth", PKType: "", Column: "end_date_month"}, {Name: "EndDateDay", PKType: "", Column: "end_date_day"}, {Name: "TypeId", PKType: "", Column: "type"}, {Name: "AreaId", PKType: "", Column: "area"}, {Name: "GenderId", PKType: "", Column: "gender"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "Ended", PKType: "", Column: "ended"}, {Name: "BeginAreaId", PKType: "", Column: "begin_area"}, {Name: "EndArea", PKType: "", Column: "end_area"}}, PKFieldIndex: 0},
	z: new(Artist).Values(),
}

// String returns a string representation of this struct or record.
func (s Artist) String() string {
	res := make([]string, 19)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "SortName: " + reform.Inspect(s.SortName, true)
	res[4] = "BeginDateYear: " + reform.Inspect(s.BeginDateYear, true)
	res[5] = "BeginDateMonth: " + reform.Inspect(s.BeginDateMonth, true)
	res[6] = "BeginDateDay: " + reform.Inspect(s.BeginDateDay, true)
	res[7] = "EndDateYear: " + reform.Inspect(s.EndDateYear, true)
	res[8] = "EndDateMonth: " + reform.Inspect(s.EndDateMonth, true)
	res[9] = "EndDateDay: " + reform.Inspect(s.EndDateDay, true)
	res[10] = "TypeId: " + reform.Inspect(s.TypeId, true)
	res[11] = "AreaId: " + reform.Inspect(s.AreaId, true)
	res[12] = "GenderId: " + reform.Inspect(s.GenderId, true)
	res[13] = "Comment: " + reform.Inspect(s.Comment, true)
	res[14] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[15] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[16] = "Ended: " + reform.Inspect(s.Ended, true)
	res[17] = "BeginAreaId: " + reform.Inspect(s.BeginAreaId, true)
	res[18] = "EndArea: " + reform.Inspect(s.EndArea, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Artist) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.Name,
		s.SortName,
		s.BeginDateYear,
		s.BeginDateMonth,
		s.BeginDateDay,
		s.EndDateYear,
		s.EndDateMonth,
		s.EndDateDay,
		s.TypeId,
		s.AreaId,
		s.GenderId,
		s.Comment,
		s.EditsPending,
		s.LastUpdated,
		s.Ended,
		s.BeginAreaId,
		s.EndArea,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Artist) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.Name,
		&s.SortName,
		&s.BeginDateYear,
		&s.BeginDateMonth,
		&s.BeginDateDay,
		&s.EndDateYear,
		&s.EndDateMonth,
		&s.EndDateDay,
		&s.TypeId,
		&s.AreaId,
		&s.GenderId,
		&s.Comment,
		&s.EditsPending,
		&s.LastUpdated,
		&s.Ended,
		&s.BeginAreaId,
		&s.EndArea,
	}
}

// View returns View object for that struct.
func (s *Artist) View() reform.View {
	return ArtistTable
}

// Table returns Table object for that record.
func (s *Artist) Table() reform.Table {
	return ArtistTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Artist) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Artist) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Artist) HasPK() bool {
	return s.ID != ArtistTable.z[ArtistTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Artist) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = ArtistTable
	_ reform.Struct = new(Artist)
	_ reform.Table  = ArtistTable
	_ reform.Record = new(Artist)
	_ fmt.Stringer  = new(Artist)
)

func init() {
	parse.AssertUpToDate(&ArtistTable.s, new(Artist))
}
