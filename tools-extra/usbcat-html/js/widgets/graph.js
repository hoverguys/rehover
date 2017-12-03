class GraphWidget extends Widget {
	constructor(parent, title, options = {}, maxLength = 100) {
		super(parent, title);

		let canvas = document.createElement("canvas");
		this.container.appendChild(canvas);

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
		this.maxLength = maxLength;

		this.reset();
	}

	addSet(setid, options = {}) {
		this.setmap[setid] = Object.keys(this.chart.data.datasets).length;
		if (!("data" in options)) {
			options.data = [];
		}
		this.chart.data.datasets.push(options);
		this.chart.update(0);
	}
	addPoints(x, points) {
		this.chart.data.labels.push(x);
		for (let pointid in points) {
			if (!(pointid in this.setmap)) {
				this.addSet(pointid, { label: pointid });
			}
			this.chart.data.datasets[this.setmap[pointid]].data.push(points[pointid]);
		}
		if (this.chart.data.labels.length > this.maxLength) {
			this._pop();
		}
		this.chart.update(0);
	}
	_pop() {
		this.chart.data.labels.shift();
		this.chart.data.datasets.forEach((dataset) => {
			dataset.data.shift();
		});
	}
	reset() {
		this.chart.data.labels = [];
		this.chart.data.datasets.forEach((dataset) => {
			dataset.data = [];
		});
		this.chart.update(0);
	}
}