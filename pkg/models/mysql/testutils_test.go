package mysql

import (
	"database/sql"
	"os"
	"testing"
)

func newTestDB(t *testing.T) (*sql.DB, func()) {

	db, err := sql.Open("mysql", "test_web:pass@/test_snippetbox?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}

	scriptfiles := []string{
		"./testdata/setup.sql",
		"./testdata/setup2.sql",
		"./testdata/setup3.sql",
		"./testdata/setup4.sql",
	}

	for _, scriptfile := range scriptfiles {
		script, err := os.ReadFile(scriptfile)
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
	}

	return db, func() {
		scriptfiles := []string{
			"./testdata/teardown.sql",
			"./testdata/teardown2.sql",
		}

		for _, scriptfile := range scriptfiles {
			script, err := os.ReadFile(scriptfile)
			if err != nil {
				t.Fatal(err)
			}

			_, err = db.Exec(string(script))
			if err != nil {
				t.Fatal(err)
			}
		}

		db.Close()
	}
}
