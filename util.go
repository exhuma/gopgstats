package gopgstats

import "strings"

func parseDSN(dsn string) map[string]string {
	args := strings.Split(dsn, " ")
	parsed := make(map[string]string)
	for _, arg := range args {
		keyvalue := strings.Split(arg, "=")
		if len(keyvalue) == 2 {
			parsed[keyvalue[0]] = keyvalue[1]
		}
	}
	return parsed
}

func unparseDSN(items map[string]string) string {
	args := []string{}
	for key, value := range items {
		keyvalue := [2]string{key, value}
		args = append(args, strings.Join(keyvalue[:], "="))
	}
	return strings.Join(args, " ")
}

// Given an existing DSN, returns a new DSN where only the database name is
// different
func DsnForDatabase(dsn string, dbname string) string {
	items := parseDSN(dsn)
	items["dbname"] = dbname
	return unparseDSN(items)
}
