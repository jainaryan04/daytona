// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetAvailableEphemeralPort() (uint16, error) {
	var ephemeralPort uint16
	for ephemeralPort = 50000; ephemeralPort < 60000; ephemeralPort++ {
		if IsPortAvailable(ephemeralPort) {
			log.Debug("EPHEMERAL PORT: " + strconv.FormatUint(uint64(ephemeralPort), 10))
			return ephemeralPort, nil
		}
	}
	return 0, errors.New("no more ephemeral ports available")
}

func IsPortAvailable(port uint16) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		_ = ln.Close()
		return true
	}
	return false
}
