class Logger {
	// host : string (Host to connect to)
	constructor(host) {
		// Make empty list for subscribers
		this.subs = {};

		// Connect to websocket
		this.conn = new WebSocket(`ws://${host}/ws`);
		// Bind ws events
		this.conn.onopen = () => { console.log("Connected!"); };
		this.conn.onmessage = this._onMessage.bind(this);
		this.conn.onclose = (evt) => {
			console.log("Disconnected!");
			// TODO Do something !!1
		};
	}

	// Handler for incoming messages from websocket
	// evt : MessageEvent (Message from ws)
	_onMessage(evt) {
		let data = JSON.parse(evt.data);

		// Check for subs
		if (data.Type in this.subs) {
			// Send to each sub
			this.subs[data.Type].forEach((fn) => fn(data));
		} else {
			// Maybe log?
			console.warn(`Received msg of type "${data.Type}" but no listeners`, data);
		}
	}

	// Subscribe to all messages of a certain type
	// type : string (type of messages to listen for)
	// fn   : ({Type: str, Text: str}) => void (function to call with the message data)
	subscribe(type, fn) {
		// First sub for type? Create the sub list for the type
		if (!(type in this.subs)) {
			this.subs[type] = [];
		}
		// Append to list
		this.subs[type].push(fn);
	}

	// TODO Unsub?
}