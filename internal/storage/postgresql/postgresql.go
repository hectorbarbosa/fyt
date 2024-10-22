package postgresql

import "github.com/jackc/pgx/v5/pgtype"

func newText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: true}
}

func newBool(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}
