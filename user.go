package main

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func getUsers() ([]User, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func createUser(name string) (int64, error) {
	var id int64
	err := conn.QueryRow(context.Background(), "INSERT INTO users (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func updateUser(id int64, name string) error {
	_, err := conn.Exec(context.Background(), "UPDATE users SET name=$1 WHERE id=$2", name, id)
	return err
}

func deleteUser(id int64) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM users WHERE id=$1", id)
	return err
}
