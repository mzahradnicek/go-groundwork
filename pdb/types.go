package pdb

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math"
	"strings"

	"github.com/lib/pq"
)

var (
	ErrTypeAssertionArrayByte          = errors.New("Type assertion .([]byte) failed.")
	ErrTypeAssertionMapStringInterface = errors.New("Type assertion .(map[string]interface{}) failed.")
)

type NullBool struct {
	sql.NullBool
}

func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nb.Bool)
}

func (nb *NullBool) UnmarshalJSON(b []byte) error {
	if strings.ToLower(string(b)) == "null" {
		nb.Valid = false
		return nil
	}

	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = (err == nil)
	return err
}

func (v NullBool) CompareHash() []byte {
	if !v.Valid {
		return []byte("null")
	}

	if v.Bool {
		return []byte("1")
	} else {
		return []byte("0")
	}
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nf.Float64)
}

func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	if strings.ToLower(string(b)) == "null" {
		nf.Valid = false
		return nil
	}

	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = (err == nil)
	return err
}

func (v NullFloat64) CompareHash() []byte {
	if !v.Valid {
		return []byte{}
	}

	fbits := math.Float64bits(v.Float64)

	return []byte{
		byte(0xff & fbits),
		byte(0xff & (fbits >> 8)),
		byte(0xff & (fbits >> 16)),
		byte(0xff & (fbits >> 24)),
		byte(0xff & (fbits >> 32)),
		byte(0xff & (fbits >> 40)),
		byte(0xff & (fbits >> 48)),
		byte(0xff & (fbits >> 56)),
	}
}

type NullInt64 struct {
	sql.NullInt64
}

func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ni.Int64)
}

func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if strings.ToLower(s) == "null" {
		ni.Valid = false
		return nil
	}

	err := json.Unmarshal([]byte(s), &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

func (v NullInt64) CompareHash() []byte {
	if !v.Valid {
		return []byte{}
	}

	return []byte{
		byte(0xff & v.Int64),
		byte(0xff & (v.Int64 >> 8)),
		byte(0xff & (v.Int64 >> 16)),
		byte(0xff & (v.Int64 >> 24)),
		byte(0xff & (v.Int64 >> 32)),
		byte(0xff & (v.Int64 >> 40)),
		byte(0xff & (v.Int64 >> 48)),
		byte(0xff & (v.Int64 >> 56)),
	}
}

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	if strings.ToLower(string(b)) == "null" {
		ns.Valid = false
		return nil
	}

	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

func (v NullString) CompareHash() []byte {
	if !v.Valid {
		return []byte{}
	}

	return []byte(v.String)
}

type NullTime struct {
	pq.NullTime
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nt.Time)
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {
	if strings.ToLower(string(b)) == "null" {
		nt.Valid = false
		return nil
	}

	err := json.Unmarshal(b, &nt.Time)
	nt.Valid = (err == nil)
	return err
}

func (v NullTime) CompareHash() []byte {
	if !v.Valid {
		return []byte{}
	}

	return []byte(v.Time.String())
}

type JsonB map[string]interface{}

func (b JsonB) Value() (driver.Value, error) {
	j, err := json.Marshal(b)
	if string(j) == "null" {
		return []byte("{}"), nil
	}
	return j, err
}

func (b *JsonB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return ErrTypeAssertionArrayByte
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*b, ok = i.(map[string]interface{})
	if !ok {
		return ErrTypeAssertionMapStringInterface
	}

	return nil
}
