package main

type Item struct {
	ID        int
	Title     string
	Completed bool
}

type Tasks struct {
	Items          []Item
	Count          int
	CompletedCount int
}

func fetchTasks() ([]Item, error) {
	var items []Item
	rows, err := DB.Query("select id, title, completed from tasks order by position;")
	if err != nil {
		return []Item{}, err
	}
	defer rows.Close()
	for rows.Next() {
		item := Item{}
		err := rows.Scan(&item.ID, &item.Title, &item.Completed)
		if err != nil {
			return []Item{}, err
		}
		items = append(items, item)
	}

	return items, nil

}

func fetchTask(ID int) (Item, error) {
	var item Item

	err := DB.QueryRow(
		"select id, title, completed from tasks where id = (?)", ID,
	).Scan(
		&item.ID, &item.Title, &item.Completed,
	)

	if err != nil {
		return Item{}, err
	}

	return item, nil

}

func fetchCount() (int, error) {

	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM tasks;").Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil

}
