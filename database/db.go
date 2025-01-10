package database

import (
	"database/sql"
	"fmt"
	"os"
)

// SQLファイルを実行する関数
func ExecuteSQLFile(db *sql.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to execute SQL: %w", err)
	}
	return nil
}

// データベースに名前を保存
func SaveName(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO names (name) VALUES (?)", name)
	return err
}

// 名前リストを取得
func GetNames(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM names")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}
