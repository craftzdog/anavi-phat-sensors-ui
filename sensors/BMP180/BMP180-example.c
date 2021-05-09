#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <wiringPiI2C.h>
#include <wiringPi.h>
#include <string.h>

#include "BMP180.h"

int main()
{
	int fd = wiringPiI2CSetup(BMP180_I2CADDR);
	if (0 > fd)
	{
		fprintf(stderr, "ERROR: Unable to access BMP180 sensor module: %s\n", strerror (errno));
		exit(-1);
	}
	if (0 > begin(fd))
	{
		fprintf(stderr, "ERROR: BMP180 sensor module not found\n");
		exit(-1);
	}

	double temperature = 0;
	getTemperature(fd, &temperature);

	double pressure = 0;
	getPressure(fd, &pressure);

	printf("{ ");
	printf("\"temperature\": %0.1f, ", temperature);
	printf("\"pressure\": %0.2f ", pressure);
	printf("}");
	return 0;
}
