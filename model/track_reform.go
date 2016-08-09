package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type trackTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *trackTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("track").
func (v *trackTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *trackTable) Columns() []string {
	return []string{"id", "gid", "recording", "medium", "position", "number", "name", "artist_credit", "length", "edits_pending", "last_updated", "is_data_track"}
}

// NewStruct makes a new struct for that view or table.
func (v *trackTable) NewStruct() reform.Struct {
	return new(Track)
}

// NewRecord makes a new record for that table.
func (v *trackTable) NewRecord() reform.Record {
	return new(Track)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *trackTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// TrackTable represents track view or table in SQL database.
var TrackTable = &trackTable{
	s: parse.StructInfo{Type: "Track", SQLSchema: "", SQLName: "track", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "GID", PKType: "", Column: "gid"}, {Name: "RecordingId", PKType: "", Column: "recording"}, {Name: "MediumId", PKType: "", Column: "medium"}, {Name: "Position", PKType: "", Column: "position"}, {Name: "Number", PKType: "", Column: "number"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "ArtistCreditId", PKType: "", Column: "artist_credit"}, {Name: "Length", PKType: "", Column: "length"}, {Name: "EditsPending", PKType: "", Column: "edits_pending"}, {Name: "LastUpdated", PKType: "", Column: "last_updated"}, {Name: "IsDataTrack", PKType: "", Column: "is_data_track"}}, PKFieldIndex: 0},
	z: new(Track).Values(),
}

// String returns a string representation of this struct or record.
func (s Track) String() string {
	res := make([]string, 12)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GID: " + reform.Inspect(s.GID, true)
	res[2] = "RecordingId: " + reform.Inspect(s.RecordingId, true)
	res[3] = "MediumId: " + reform.Inspect(s.MediumId, true)
	res[4] = "Position: " + reform.Inspect(s.Position, true)
	res[5] = "Number: " + reform.Inspect(s.Number, true)
	res[6] = "Name: " + reform.Inspect(s.Name, true)
	res[7] = "ArtistCreditId: " + reform.Inspect(s.ArtistCreditId, true)
	res[8] = "Length: " + reform.Inspect(s.Length, true)
	res[9] = "EditsPending: " + reform.Inspect(s.EditsPending, true)
	res[10] = "LastUpdated: " + reform.Inspect(s.LastUpdated, true)
	res[11] = "IsDataTrack: " + reform.Inspect(s.IsDataTrack, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Track) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GID,
		s.RecordingId,
		s.MediumId,
		s.Position,
		s.Number,
		s.Name,
		s.ArtistCreditId,
		s.Length,
		s.EditsPending,
		s.LastUpdated,
		s.IsDataTrack,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Track) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GID,
		&s.RecordingId,
		&s.MediumId,
		&s.Position,
		&s.Number,
		&s.Name,
		&s.ArtistCreditId,
		&s.Length,
		&s.EditsPending,
		&s.LastUpdated,
		&s.IsDataTrack,
	}
}

// View returns View object for that struct.
func (s *Track) View() reform.View {
	return TrackTable
}

// Table returns Table object for that record.
func (s *Track) Table() reform.Table {
	return TrackTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Track) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Track) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Track) HasPK() bool {
	return s.ID != TrackTable.z[TrackTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Track) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = TrackTable
	_ reform.Struct = new(Track)
	_ reform.Table  = TrackTable
	_ reform.Record = new(Track)
	_ fmt.Stringer  = new(Track)
)

func init() {
	parse.AssertUpToDate(&TrackTable.s, new(Track))
}
