package main

import (
	"fmt"
	"golang-gtcamp/week02/sql"
)

func main ()  {
	sqlClient, err := sql.InitClient()
	if err != nil {
		fmt.Println("initClient error:", err)
	}
	defer sqlClient.Close()
	err = sql.QueryUser(sqlClient, 2)
	if err != nil {
		fmt.Println("query sql error :", err)
	}
}
