var ws = new WebSocket("ws://localhost:8000/wsChannel")

ws.onopen = function(evt){
    console.log("Connection opened")
}

ws.onclose = function(evt){
    console.log("Connection closed")
}

ws.onerror = function(evt){
    console.log("Connection error")
}

ws.onmessage = function(evt){
    console.log("Received message")
}

document.getElementById("sendButton").addEventListener("click",function(){
    ws.send("kalimera")
})