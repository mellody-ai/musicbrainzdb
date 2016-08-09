package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type eventTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *eventTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("event").
func (v *eventTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *eventTable) Columns() []string {
	return []string{"id", "gid", "name", "begin_date_year", "begin_date_month", "begin_date_day", "end_date_year", "end_date_month", "end_date_day", "time", "type", "cancelled", "setlist", "comment", "edits_pending", "last_updated", "ended"}
}

// NewStruct makes a new struct for that view or table.
func (v *eventTable) NewStruct() reform.Struct {
	return new(Event)
}

// NewRecord makes a new record for that table.
func (v *eventTable) NewRecord() reform.Record {
	return new(Event)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *eventTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// EventTable represents event view or table in SQL database.
var EventTable = &eventTable{
	s: parse.StructInfo{Type: "Event", SQLSchema: "", SQLName: "event", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "BeginDateYear", PKType: "", Column: "begin_date_year"}, {Name: "BeginDateMonth", PKType: "", Column: "begin_date_month"}, {Name: "BeginDateDay", PKType: "", Column: "begin_date_day"}, {Name: "EndDateYear", PKType: "", Column: "end_date_year"}, {Name: "EndDateMonth", PKType: "", Column: "end_date_month"}, {Name: "EndDateDay", PKType: "", Column: "end_date_day"}, {Name: "Time", PKType: "", Column: "time"}, {Name: "TypeId", PKType: "", Column: "type"}, {Name: "Cancelled", PKType: "", Column: "cancelled"}, {Name: "Setlist", PKType: "", Column: "setlist"}, {Name: "Comment", PKType: "", Column: "comment"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "Ended", PKType: "", Column: "ended"}}, PKFieldIndex: 0},
	z: new(Event).Values(),
}

// String returns a string representation of this struct or record.
func (s Event) String() string {
	res := make([]string, 17)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "BeginDateYear: " + reform.Inspect(s.BeginDateYear, true)
	res[4] = "BeginDateMonth: " + reform.Inspect(s.BeginDateMonth, true)
	res[5] = "BeginDateDay: " + reform.Inspect(s.BeginDateDay, true)
	res[6] = "EndDateYear: " + reform.Inspect(s.EndDateYear, true)
	res[7] = "EndDateMonth: " + reform.Inspect(s.EndDateMonth, true)
	res[8] = "EndDateDay: " + reform.Inspect(s.EndDateDay, true)
	res[9] = "Time: " + reform.Inspect(s.Time, true)
	res[10] = "TypeId: " + reform.Inspect(s.TypeId, true)
	res[11] = "Cancelled: " + reform.Inspect(s.Cancelled, true)
	res[12] = "Setlist: " + reform.Inspect(s.Setlist, true)
	res[13] = "Comment: " + reform.Inspect(s.Comment, true)
	res[14] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[15] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[16] = "Ended: " + reform.Inspect(s.Ended, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Event) Values() []interface{} {
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
		s.Time,
		s.TypeId,
		s.Cancelled,
		s.Setlist,
		s.Comment,
		s.EditsPending,
		s.LastUpdated,
		s.Ended,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Event) Pointers() []interface{} {
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
		&s.Time,
		&s.TypeId,
		&s.Cancelled,
		&s.Setlist,
		&s.Comment,
		&s.EditsPending,
		&s.LastUpdated,
		&s.Ended,
	}
}

// View returns View object for that struct.
func (s *Event) View() reform.View {
	return EventTable
}

// Table returns Table object for that record.
func (s *Event) Table() reform.Table {
	return EventTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Event) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Event) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Event) HasPK() bool {
	return s.ID != EventTable.z[EventTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Event) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = EventTable
	_ reform.Struct = new(Event)
	_ reform.Table  = EventTable
	_ reform.Record = new(Event)
	_ fmt.Stringer  = new(Event)
)

func init() {
	parse.AssertUpToDate(&EventTable.s, new(Event))
}
