package frontend

import (
    "net/http"
    "html/template"
    "fmt"
)

type DataTypes struct {
    Xlocation int32
    Ylocation int32
}

func (dt DataTypes) GetXY() (int32, int32) {
    return dt.Xlocation, dt.Ylocation
}

func Client(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/client.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println("There is an error parsing the html")
        return
    }

    ctx := struct{}{}
    err = tmpl.Execute(w, ctx)
    if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    fmt.Println("Error executing the template")
	    return
    }
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
    err = tmpl.Execute(w, ctx)
    if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    fmt.Println("Error executing the template")
	    return
    }
}
