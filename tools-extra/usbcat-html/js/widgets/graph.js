class GraphWidget extends Widget {
	constructor(parent, title, options = {}) {
		super(parent, title);

		let canvas = document.createElement("canvas");
		parent.appendChild(canvas);

		if (!("options" in options)) {
			options.options = {
				responsive: true
			};
		}

		if (!("data" in options)) {
			options.data = {};
		}

		let ctx = canvas.getContext("2d");
		this.chart = new Chart(ctx, options);
		this.setmap = {};
	}

	addSet(setid, options = {}) {
		this.setmap[setid] = Object.keys(this.chart.data.datasets).length;
		if (!("data" in options)) {
			options.data = [];
		}
		this.chart.data.datasets.push(options);
		this.chart.update();
	}
	addPoints(x, points) {
		this.chart.data.labels.push(x);
		for (let pointid in points) {
			if (!(pointid in this.setmap)) {
				this.addSet(pointid, { label: pointid });
			}
			this.chart.data.datasets[this.setmap[pointid]].data.push(points[pointid]);
		}
		this.chart.update();
	}
}