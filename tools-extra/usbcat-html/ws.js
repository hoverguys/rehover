let conn = new WebSocket("ws://" + document.location.host + "/ws");
conn.onmessage = (evt) => {
	console.log(evt);
};