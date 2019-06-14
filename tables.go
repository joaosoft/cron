package cron

import "fmt"

func format(schema, table string) string {
	return fmt.Sprintf("%s.%s", schema, table)
}

var (
	cronTableJob = format(schemaCron, "job")
)
