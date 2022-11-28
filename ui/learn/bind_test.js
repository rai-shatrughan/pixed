var monitor = {
    threshold: 5,
    check: function (value) {
        if (value > this.threshold) {
            this.display("Value is High!");
        } else {
            this.display("Value is Low");
        }
    },
    display(message) {
        console.log(message);
    }
};


var mon = monitor.check.bind(monitor, 4);

mon();