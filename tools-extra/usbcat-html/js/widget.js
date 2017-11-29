class Widget {
	constructor(parent, title = "Unknown widget") {
		this.container = document.createElement("section");
		this.container.className = "widget";
		parent.appendChild(this.container);

		let header = document.createElement("header");
		header.appendChild(document.createTextNode(title));
		this.container.appendChild(header);
	}
}