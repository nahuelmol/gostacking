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

func Home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println("there is an error parsing html to template")
        return
    }

    ctx := DataTypes {
        Xlocation: 33,
        Ylocation: 44,
    }

    err = tmpl.Execute(w, ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println("error executing the template")
        return
    }

}
