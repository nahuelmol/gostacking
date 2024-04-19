package frontend

import (
    "net/http"
    "html/template"
    "fmt"
)

type DataTypes struct {
    xlocation int32
    ylocation int32
}

func Home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println("there is an error parsing html to template")
        return
    }

    ctx := DataTypes {
        xlocation: 33,
        ylocation: 44,
    }

    err = tmpl.Execute(w, ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println("error executing the template")
        return
    }

}
