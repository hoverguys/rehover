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

		parent.appendChild(section);
	}

	writeLine(data) {
		let p = document.createElement("p");
		p.appendChild(document.createTextNode(data.Text));
		this.msglist.appendChild(p);
	}
}