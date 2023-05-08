package sqlparser_test

import (
	"testing"

	"gopkg.inshopline.com/gsoul/component/go-sqlparser"
)

func TestPos_String(t *testing.T) {
	if got, want := (sqlparser.Pos{}).String(), `-`; got != want {
		t.Fatalf("String()=%q, want %q", got, want)
	}
}
