package model

import (
	"fmt"
	database "message_queuing_system/database"
	"strconv"
)

type Models struct{}
type RequestData struct {
	User_id             int      `json:"user_id"`
	Product_name        string   `json:"product_name"`
	Product_description string   `json:"product_description"`
	Product_images      []string `json:"product_images"`
	Product_price       float32  `json:"product_price"`
}

func (m *Models) IsAuthenticClient(id int) (bool, error) {
	count := 0
	err := database.Database_Conn.QueryRow("select COUNT(id) from Users where id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *Models) InsertIntoDatabase(requestData RequestData) (int64, error) {
	res, err := database.Database_Conn.Exec("insert into Products(product_name,product_description,product_price) values(?,?,?)", requestData.Product_name, requestData.Product_description, requestData.Product_price)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	id_str := strconv.FormatInt(id, 10)
	query := "insert into Image_Store(product_id,product_images) values"
	for i, image := range requestData.Product_images {
		if i == len(requestData.Product_images)-1 {
			query += "('" + id_str + "','" + image + "')"
		} else {
			query += "('" + id_str + "','" + image + "'),"
		}
	}
	_, err = database.Database_Conn.Exec(query)
	if err != nil {
		return 0, err
	}
	return id, nil
}
