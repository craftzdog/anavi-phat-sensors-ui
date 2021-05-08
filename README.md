# ANAVI pHAT Sensors Recorder for Firestore

Store [air quality data captured with ANAVI pHAT sensors and MH-Z19](https://www.crowdsupply.com/anavi-technology/infrared-phat) to Firestore.
Written in C and Go.

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

## How to use

### Check sensor data

```sh
go run ./test-sensors.go
```

### Record sensor data

```sh
go run ./write.go
```

### Check Firestore data

```sh
go run ./read.go
```

## License

C programs for retrieving pHAT sensor data are provided under MIT License by [Anavi Technology](https://github.com/AnaviTechnology/anavi-examples).

The rest of the scripts are under MIT License by [Takuya Matsuyama](https://twitter.com/inkdrop_app).
