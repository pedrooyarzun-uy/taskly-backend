package domain

type Category struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	User    int    `db:"user_id"`
	Deleted bool   `db:"deleted"`
}

type AllCategories struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	TotalTasks int    `db:"total_tasks"`
}
