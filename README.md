# Web Interface for ANAVI pHAT & MH-Z19 Sensors

A web interface that displays [air quality data captured with ANAVI pHAT sensors and MH-Z19](https://www.crowdsupply.com/anavi-technology/infrared-phat).

Supported sensors:

- BH1750
- BMP180
- HTU21D
- MH_Z19

## Prerequisites

Read the instructions of [ANAVI pHAT](https://github.com/AnaviTechnology/anavi-examples) and [MH-Z19 Python module](https://pypi.org/project/mh-z19/).

Install dependencies:

```sh
apt install build-essential wiringpi
sudo apt install python-pip
sudo pip install mh-z19
```

Put your `service-account-key.json` for the Firebase in the project root directory.
See [the documentation](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) for more detail.

## How to build

HTU21D:

```sh
cd ./sensors/HTU21D
make
```

BMP180:

```sh
cd ./sensors/BMP180
make
```

BH1750:

```sh
cd ./sensors/BH1750
make
```

## How to set up

Configure root's crontab like so:

```cron
*/5 * * * * /usr/bin/python -m mh_z19                > /<path-to-project>/data/MH_Z19.json
*/5 * * * * /<path-to-project>/sensors/HTU21D/HTU21D > /<path-to-project>/data/HTU21D.json
*/5 * * * * /<path-to-project>/sensors/BMP180/BMP180 > /<path-to-project>/data/BMP180.json
*/5 * * * * /<path-to-project>/sensors/BH1750/BH1750 > /<path-to-project>/data/BH1750.json
```

## License

C programs for retrieving pHAT sensor data are provided under MIT License by [Anavi Technology](https://github.com/AnaviTechnology/anavi-examples).

The rest of the scripts are under MIT License by [Takuya Matsuyama](https://twitter.com/inkdrop_app).

The sound effect is provided by [freeSFX](https://www.freesfx.co.uk/).
