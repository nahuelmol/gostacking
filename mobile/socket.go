package mobile

import (
    "fmt"
    "encoding/json"
    "net/http"

    "molinahuel/qgis/back/stack"
    "github.com/gorilla/websocket"
)

var UPGRADER = websocket.Upgrader {
    ReadBufferSize: 1024,
    WriteBufferSize:1024,
}

func MobileSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := UPGRADER.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("cannot upgrade\n")
    }
    defer conn.Close()
    //make the user to send his postion throuhg an specfic ws://

    for {
        msgType, p, err := conn.ReadMessage()
        if err != nil {
            continue
        }
        switch msgType {
        case websocket.TextMessage:
            json_string := string(p)
            var body map[string]interface{}
            json.Unmarshal([]byte(json_string), &body)

            x, xexists := body["xlocation"].(int32)
            y, yexists := body["ylocation"].(int32)
            if xexists && yexists {
                var location stack.Location 
                location.SetXY(x,y)
                stack.MEGASTACK.Push(&location)
            }

        case websocket.BinaryMessage:
            buffer := make([]byte, 1024)
            n := len(p)
            fmt.Println("buffer: ", string(buffer[:n]))
        default:
            fmt.Println("messag is not binary neither text")
        }
    }

}



