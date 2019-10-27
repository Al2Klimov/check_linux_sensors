FROM golang:1.11
SHELL ["/bin/bash", "-exo", "pipefail", "-c"]

RUN apt-get update ;\
	DEBIAN_FRONTEND=noninteractive apt-get install --no-install-{recommends,suggests} -y \
		gcc libc6-dev libsensors4-dev ;\
	apt-get clean ;\
	rm -vrf /var/lib/apt/lists/*

CMD exec go build -o /go/src/github.com/Al2Klimov/check_linux_sensors/check_linux_sensors.linux-amd64 github.com/Al2Klimov/check_linux_sensors
