package model

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type artistCreditTable struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *artistCreditTable) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("artist_credit").
func (v *artistCreditTable) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *artistCreditTable) Columns() []string {
	return []string{"id", "name", "artist_count", "ref_count", "created"}
}

// NewStruct makes a new struct for that view or table.
func (v *artistCreditTable) NewStruct() reform.Struct {
	return new(ArtistCredit)
}

// NewRecord makes a new record for that table.
func (v *artistCreditTable) NewRecord() reform.Record {
	return new(ArtistCredit)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *artistCreditTable) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ArtistCreditTable represents artist_credit view or table in SQL database.
var ArtistCreditTable = &artistCreditTable{
	s: parse.StructInfo{Type: "ArtistCredit", SQLSchema: "", SQLName: "artist_credit", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "ArtistCount", PKType: "", Column: "artist_count"}, {Name: "RefCount", PKType: "", Column: "ref_count"}, {Name: "Created", PKType: "", Column: "created"}}, PKFieldIndex: 0},
	z: new(ArtistCredit).Values(),
}

// String returns a string representation of this struct or record.
func (s ArtistCredit) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "ArtistCount: " + reform.Inspect(s.ArtistCount, true)
	res[3] = "RefCount: " + reform.Inspect(s.RefCount, true)
	res[4] = "Created: " + reform.Inspect(s.Created, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ArtistCredit) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
		s.ArtistCount,
		s.RefCount,
		s.Created,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ArtistCredit) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
		&s.ArtistCount,
		&s.RefCount,
		&s.Created,
	}
}

// View returns View object for that struct.
func (s *ArtistCredit) View() reform.View {
	return ArtistCreditTable
}

// Table returns Table object for that record.
func (s *ArtistCredit) Table() reform.Table {
	return ArtistCreditTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *ArtistCredit) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *ArtistCredit) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *ArtistCredit) HasPK() bool {
	return s.ID != ArtistCreditTable.z[ArtistCreditTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *ArtistCredit) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = ArtistCreditTable
	_ reform.Struct = new(ArtistCredit)
	_ reform.Table  = ArtistCreditTable
	_ reform.Record = new(ArtistCredit)
	_ fmt.Stringer  = new(ArtistCredit)
)

func init() {
	parse.AssertUpToDate(&ArtistCreditTable.s, new(ArtistCredit))
}
