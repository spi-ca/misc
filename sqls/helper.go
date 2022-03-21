package sqls

import "github.com/jmoiron/sqlx"

// MapRows is a convert function,that converts sqlx.Row to specified interface.
func MapRows[T any](rows *sqlx.Rows, itemGetter func() T) (items []T, err error) {
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		item := itemGetter()
		if err = rows.StructScan(item); err != nil {
			return
		}
		items = append(items, item)
	}

	err = rows.Err()
	return
}
