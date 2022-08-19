#!/bin/bash
# Copyright (c) 2022 Proton AG
#
# This file is part of Proton Mail Bridge.
#
# Proton Mail Bridge is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Proton Mail Bridge is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.


BRIDGE_APP_VERSION=${BRIDGE_APP_VERSION:-2.2.1+git} # TODO get the version number from a unified location.
BUILD_CONFIG=${BRIDGE_GUI_BUILD_CONFIG:-Debug}
BUILD_DIR=$(echo "./cmake-build-${BUILD_CONFIG}" | tr '[:upper:]' '[:lower:]')
VCPKG_ROOT="../../../../extern/vcpkg"
VCPKG_EXE="${VCPKG_ROOT}/vcpkg"
VCPKG_BOOTSTRAP="${VCPKG_ROOT}/bootstrap-vcpkg.sh"

check_exit() {
  # shellcheck disable=SC2181
  if [ $? -ne 0 ]; then
    echo "Process failed: $1"
    rm -r "$BUILD_DIR"
    exit 1
  fi
}

git submodule update --init --recursive ${VCPKG_ROOT}
check_exit "Failed to initialize vcpkg as a submodule."

${VCPKG_BOOTSTRAP} -disableMetrics
check_exit "Failed to bootstrap vcpkg."

if [[ "$OSTYPE" == "darwin"* ]]; then
	if [[ "$(uname -m)" == "arm64" ]]; then
		${VCPKG_EXE} install grpc:arm64-osx 
		check_exit "Failed installing gRPC for macOS / Apple Silicon"
	fi
	${VCPKG_EXE} install grpc:x64-osx
    check_exit "Failed installing gRPC for macOS / Intel x64"
elif [[ "$OSTYPE" == "linux"* ]]; then
	${VCPKG_EXE} install grpc:x64-linux
	check_exit "Failed installing gRPC for Linux / Intel x64"

else
	echo "For Windows, use the build.ps1 Powershell script."
	exit 1
fi

${VCPKG_EXE} upgrade --no-dry-run

cmake  -DCMAKE_BUILD_TYPE="${BUILD_CONFIG}" -DBRIDGE_APP_VERSION="${BRIDGE_APP_VERSION}" -G Ninja -S . -B "${BUILD_DIR}"
check_exit "CMake failed"

cmake --build "${BUILD_DIR}"
check_exit "build failed"