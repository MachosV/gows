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
    element = document.createElement("p")
    element.innerText = evt.data
    document.getElementById("chatPanel").appendChild(element)
}

document.getElementById("sendButton").addEventListener("click",function(){
    var sendData = document.getElementById("inputArea").value
    document.getElementById("inputArea").value = ""
    ws.send(sendData)
})