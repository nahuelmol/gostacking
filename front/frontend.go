package frontend

import (
    "net/http"
    "html/template"
    "fmt"
    "time"
)

type DataTypes struct {
    Xlocation int32
    Ylocation int32
}

func (dt DataTypes) GetXY() (int32, int32) {
    return dt.Xlocation, dt.Ylocation
}

func Home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println("There is an error parsing the html")
        return
    }

    ctx := DataTypes {
        Xlocation: 33,
        Ylocation: 44,
    }

    for i := 0; i < 5; i++ {
	time.Sleep(2 * time.Second)
	err = tmpl.Execute(w, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error executing the template")
		return
	}
	}
}
