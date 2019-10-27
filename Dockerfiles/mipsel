FROM golang:1.11
SHELL ["/bin/bash", "-exo", "pipefail", "-c"]

RUN dpkg --add-architecture mipsel ;\
	apt-get update ;\
	DEBIAN_FRONTEND=noninteractive apt-get install --no-install-{recommends,suggests} -y \
		gcc-mipsel-linux-gnu {libc6-dev,libsensors5}:mipsel libsensors4-dev ;\
	apt-get clean ;\
	rm -vrf /var/lib/apt/lists/*

RUN  for f in $(find / -name libsensors.so.5); do \
	pushd "$(dirname "$f")" ;\
	ln -vs libsensors.so{.5,} || true ;\
	popd ;\
	done

CMD GOARCH=mipsle GOMIPS=softfloat CGO_ENABLED=1 CC=mipsel-linux-gnu-gcc exec go build -o /go/src/github.com/Al2Klimov/check_linux_sensors/check_linux_sensors.linux-mipsle-softfloat github.com/Al2Klimov/check_linux_sensors
