const ws = new WebSocket("ws://localhost:8080/wsmobile");

ws.onopen = () => {
    console.log("connected");
    ws.send("hello");
};

ws.onmessage = (e) => {
    console.log(e.data);
};
