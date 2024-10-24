package postgresql

import (
	"fyt/internal"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func newInt4(n int32) pgtype.Int4 {
	return pgtype.Int4{Int32: n, Valid: true}
}

func newText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: true}
}

func newBool(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}

func newDateFromString(s string) (pgtype.Timestamptz, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return pgtype.Timestamptz{}, internal.WrapErrorf(
			err,
			internal.ErrorCodeInvalidArgument,
			"newDateFrom String",
		)
	}
	return pgtype.Timestamptz{
		Time:  t,
		Valid: !t.IsZero(),
	}, nil
}

func newTimestamp(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: !t.IsZero(),
	}
}
