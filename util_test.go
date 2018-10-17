package gopgstats

import (
	"testing"
)

func TestParseDSN1(t *testing.T) {
	result := parseDSN("username=foo dbname=bar")
	if result["username"] != "foo" {
		t.Errorf("error: Expected %q, got %q", "foo", result["username"])
	}
	if result["dbname"] != "bar" {
		t.Errorf("error: Expected %q, got %q", "foo", result["dbname"])
	}
}

func TestParseDSN2(t *testing.T) {
	result := parseDSN("username=foo")
	if result["username"] != "foo" {
		t.Errorf("error: Expected %q, got %q", "foo", result["username"])
	}
}

func TestMakeNewDSN(t *testing.T) {
	tmp := DsnForDatabase("username=foo", "newdb")
	result := parseDSN(tmp)
	if result["dbname"] != "newdb" {
		t.Errorf("error: Expected %q, got %q", "newdb", result["dbname"])
	}

	tmp = DsnForDatabase("username=foo dbname=olddb", "newdb")
	result = parseDSN(tmp)
	if result["dbname"] != "newdb" {
		t.Errorf("error: Expected %q, got %q", "newdb", result["dbname"])
	}
}
