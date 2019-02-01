#!/bin/bash

set -exo pipefail

go generate

for f in Dockerfiles/*; do
	docker run -iv "$(pwd):/go/src/github.com/Al2Klimov/check_linux_sensors" "grandmaster/build-check_linux_sensors-$(basename "$f")"
done
