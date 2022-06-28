package mapper

import (
	"fmt"
	"log"
)


type Apps struct {
	ID			string		`json:"id"`
	App_Name 	string	 	`json:"appName"`
	Url 		string 		`json:"url"`
	Country 	string 		`json:"country"`
}


func GetApps() (apps []Apps) {
	DB:= GetCon()
	rows, err := DB.Query("select * from apps")

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var app Apps
		err = rows.Scan(&app.ID, &app.App_Name, &app.Url, &app.Country)
		if err!=nil {
			log.Printf("查询出错")
		}else{
			apps = append(apps, app)
		}
	}
	return apps
}
