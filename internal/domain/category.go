package domain

type Category struct {
	id   int    `db:"id"`
	name string `db:"name"`
}
