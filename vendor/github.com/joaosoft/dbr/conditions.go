package dbr

import (
	"fmt"
)

type conditions struct {
	list []*condition
	db   *db
}

func newConditions(db *db) *conditions {
	return &conditions{
		db: db,
		list: make([]*condition, 0),
	}
}

func (c conditions) Build(db ...*db) (query string, err error) {

	if len(c.list) == 0 {
		return "", nil
	}

	if len(db) > 0 {
		c.db = db[0]
	}

	lenC := len(c.list)
	for i, item := range c.list {
		condition, err := item.Build(db...)
		if err != nil {
			return "", err
		}

		query += condition

		if i+1 < lenC {
			query += fmt.Sprintf(" %s ", c.list[i+1].operator)
		}
	}

	return query, nil
}
