const submitButton = document.getElementById("winner-button");
const winnerInput = document.getElementById("winner");

if (window['WebSocket']) {
    const conn = new WebSocket("ws://" + document.location.host + "/ws");
    submitButton.onclick = event = function () {
        conn.send(winnerInput.value);
    }
}