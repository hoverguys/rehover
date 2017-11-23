class ConsoleWidget {
	constructor(parent) {
		let section = document.createElement("section");
		section.className = "console";
		// Console header
		let header = document.createElement("header");
		header.appendChild(document.createTextNode("Debug messages"));
		section.appendChild(header);
		// Console content
		this.msglist = document.createElement("article");
		section.appendChild(this.msglist);

		this.base = Date.now();

		parent.appendChild(section);
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