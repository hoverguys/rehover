class GraphWidget {
	constructor(parent, options = {}) {
		let section = document.createElement("section");
		section.className = "visgraph";

		// Create vis dataset
		this.dataset = new vis.DataSet();
		this.graph = new vis.Graph2d(parent, this.dataset, {

		});

		this.base = Date.now();
	}
}