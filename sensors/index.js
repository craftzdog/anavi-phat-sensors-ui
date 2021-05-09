const { execSync } = require("child_process");
const path = require("path");

function getSensorData(name) {
  try {
    const resJSON = execSync(path.join(__dirname, `${name}/${name}`), {
      encoding: "utf8",
    });
    const res = JSON.parse(resJSON);
    return res;
  } catch (e) {
    console.error(`Failed to get ${name} data:`, e);
    return null;
  }
}

function getBH1750() {
  return getSensorData("BH1750");
}

function getBMP180() {
  return getSensorData("BMP180");
}

function getHTU21D() {
  return getSensorData("HTU21D");
}

function getMH_Z19() {
  try {
    const resJSON = execSync(`/usr/bin/python -m mh_z19`, {
      encoding: "utf8",
    });
    const res = JSON.parse(resJSON);
    return res;
  } catch (e) {
    console.error("Failed to get MH_Z19 data:", e);
    return null;
  }
}

module.exports = {
  getSensorData,
  getBH1750,
  getBMP180,
  getHTU21D,
  getMH_Z19,
};
