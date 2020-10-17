package common

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"strconv"
)

type NullInt64 sql.NullInt64

func (n *NullInt64) Scan(value interface{}) error {
	if value == nil {
		n.Int64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	n.Int64 = value.(int64)
	return nil
}

func (n NullInt64) Value() driver.Value {
	if !n.Valid {
		return nil
	}
	return n.Int64
}

func (n *NullInt64) UnmarshalParam(src string) error {
	tmp, _ := strconv.Atoi(src)
	if tmp > 0 {
		n.Int64 = int64(tmp)
		n.Valid = true
	} else {
		n.Int64 = 0
		n.Valid = false
	}
	return nil
}

func (n NullInt64) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if n.Int64 == 0 {
		buf.WriteString(`null`)
	} else {
		buf.WriteString(strconv.Itoa(int(n.Int64)))
	}
	return buf.Bytes(), nil
}

type NullString sql.NullString

func (n *NullString) Scan(value interface{}) error {
	if value == nil {
		n.String, n.Valid = "", false
		return nil
	}
	n.Valid = true
	n.String = value.(string)
	return nil
}

func (n NullString) Value() driver.Value {
	if !n.Valid {
		return nil
	}
	return n.String
}

func (n *NullString) UnmarshalParam(src string) error {
	if src != "" {
		n.String = src
		n.Valid = true
	} else {
		n.String = ""
		n.Valid = false
	}
	return nil
}

func (n NullString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(n.String) == 0 {
		buf.WriteString(`null`)
	} else {
		buf.WriteString(`"` + n.String + `"`)
	}
	return buf.Bytes(), nil
}
