package domain

type Category struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	User    int    `db:"user_id"`
	Deleted bool   `db:"deleted"`
}
