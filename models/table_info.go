package models

import "github.com/lilysnc/pocketbasepg/tools/types"

type TableInfoRow struct {
	// the `db:"pk"` tag has special semantic so we cannot rename
	// the original field without specifying a custom mapper
	PK int

	Index        int           `db:"ordinal_position"`
	Name         string        `db:"column_name"`
	Type         string        `db:"data_type"`
	NotNull      bool          `db:"nullable"`
	DefaultValue types.JsonRaw `db:"column_default"`
}
