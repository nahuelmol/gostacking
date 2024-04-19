package py

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

func PySocket(w http.ResponseWriter, r *http.Request) {
    conn, err := UPGRADER.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("cannot upgrade\n")
    }
    defer conn.Close()

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

            _, is := body["gethead"] //pyclient wants head
            if is {
                location := stack.MEGASTACK.Gethead()
                x := string(location.X())
                y := string(location.Y())
                toPyC := "{ 'x': " +x+", 'y': " +y+ "}" 
                conn.WriteMessage(websocket.TextMessage, []byte(toPyC))
            } else {
                fmt.Println("there is not data in the stack")
            }


        case websocket.BinaryMessage:
            buffer := make([]byte, 1024)
            n := len(p)
            buff := string(buffer[:n])
            fmt.Println("buff:", buff)
        default:
            fmt.Println("messag is not binary neither text")
        }
    }
}
