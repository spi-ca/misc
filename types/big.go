package types

import (
	"database/sql/driver"
	"github.com/ericlagergren/decimal"
	"github.com/pkg/errors"
)

// Decimal is a DECIMAL in sql. Its zero value is valid for use with both
// Value and Scan.
//
// Although decimal can represent NaN and Infinity it will return an error
// if an attempt to store these values in the database is made.
//
// Because it cannot be nil, when Big is nil Value() will return "0"
// It will error if an attempt to Scan() a "null" value into it.
type Decimal struct {
	decimal.Big
}

// NullDecimal is the same as Decimal, but allows the Big pointer to be nil.
// See docmentation for Decimal for more details.
//
// When going into a database, if Big is nil it's value will be "null".
type NullDecimal struct {
	*decimal.Big
}

// NewDecimal creates a new decimal from a decimal
func NewDecimal(d *decimal.Big) (num Decimal) {
	if d != nil {
		num.Big.Copy(d)
	}
	return
}

// NewNullDecimal creates a new null decimal from a decimal
func NewNullDecimal(d *decimal.Big) NullDecimal {
	return NullDecimal{Big: d}
}

// Value implements driver.Valuer.
func (d Decimal) Value() (driver.Value, error) {
	return decimalValue(&d.Big, false)
}

// Scan implements sql.Scanner.
func (d *Decimal) Scan(val any) (err error) {
	_, err = decimalScan(&d.Big, val, false)
	return
}

// Value implements driver.Valuer.
func (n NullDecimal) Value() (driver.Value, error) {
	return decimalValue(n.Big, true)
}

// Scan implements sql.Scanner.
func (n *NullDecimal) Scan(val any) error {
	newD, err := decimalScan(n.Big, val, true)
	if err != nil {
		return err
	}

	n.Big = newD
	return nil
}

func decimalValue(d *decimal.Big, canNull bool) (driver.Value, error) {
	if canNull && d == nil {
		return nil, nil
	}

	if d.IsNaN(0) {
		return nil, errors.New("refusing to allow NaN into database")
	}
	if d.IsInf(0) {
		return nil, errors.New("refusing to allow infinity into database")
	}

	return d.String(), nil
}

func decimalScan(d *decimal.Big, val any, canNull bool) (*decimal.Big, error) {
	if val == nil {
		if !canNull {
			return nil, errors.New("null cannot be scanned into decimal")
		}

		return nil, nil
	}

	if d == nil {
		d = new(decimal.Big)
	}

	switch t := val.(type) {
	case float64:
		d.SetFloat64(t)
		return d, nil
	case string:
		if _, ok := d.SetString(t); !ok {
			if err := d.Context.Err(); err != nil {
				return nil, err
			}
			return nil, errors.Errorf("invalid decimal syntax: %q", t)
		}
		return d, nil
	case []byte:
		if err := d.UnmarshalText(t); err != nil {
			return nil, err
		}
		return d, nil
	default:
		return nil, errors.Errorf("cannot scan decimal value: %#v", val)
	}
}
