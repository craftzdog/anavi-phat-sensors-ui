const sensors = require("./sensors");

let data = sensors.getBH1750();
console.log("BH1750:", data);

data = sensors.getBMP180();
console.log("BMP180:", data);

data = sensors.getHTU21D();
console.log("HTU21D:", data);

data = sensors.getMH_Z19();
console.log("MH_Z19:", data);
