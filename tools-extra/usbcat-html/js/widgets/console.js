class ConsoleWidget extends Widget {
	constructor(parent, title) {
		super(parent, title)

		// Console content
		this.msglist = document.createElement("article");
		this.container.appendChild(this.msglist);

		this.base = Date.now();
	}

	writeLine(data) {
		let p = document.createElement("p");
		let timestamp = document.createElement("div");
		let datetime = (Date.now() - this.base) / 1000;
		timestamp.className = "timestamp";
		timestamp.appendChild(document.createTextNode(datetime.toFixed(3)));
		p.appendChild(timestamp);
		p.appendChild(document.createTextNode(data.Text));
		this.msglist.appendChild(p);
		this.msglist.scrollTop = this.msglist.scrollHeight;
	}
}