// var websocket=new WebSocket(`ws://${window.location.hostname}:443`);

// function initWebSocket() {
//     console.log('Trying to open a WebSocket connection...');
//     websocket.onopen    = onOpen;
//     websocket.onclose   = onClose;
//     websocket.onmessage = onMessage;
// }

// function onOpen(event) {
//     console.log('Connection opened');
// }

// function onClose(event) {
//     console.log('Connection closed');
//     setTimeout(initWebSocket, 2000);
// }

// function onMessage(event) {
//     console.log(event)
// }

const rootHTMLElement = document.querySelector("#root")

rootHTMLElement.appendChild(wrapper("class=plusSignWrapper",{type:"click",func:()=>{document.querySelector(".dialogModal").showModal()}}))
               .appendChild(addPlusSign())
