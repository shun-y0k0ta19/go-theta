// Copyright (c) 2017 "Shun Yokota" All rights reserved
//
// Part of the source code is adapted from https://github.com/google/go-github
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theta

import (
	"context"
	"fmt"
	"net/http"
)

// Info represents a Theta information.
type Info struct {
	Manufacturer     *string    `json:"manufacturer"`
	Model            *string    `json:"model"`
	SerialNumber     *string    `json:"serialNumber"`
	FirmwareVerision *string    `json:"firmwareVersion"`
	SupportURL       *string    `json:"supportUrl"`
	GPS              *bool      `json:"gps"`
	Gyro             *bool      `json:"gyro"`
	Uptime           *int       `json:"uptime"`
	API              []string   `json:"api"`
	Endpoints        *Endpoints `json:"endpoints"`
}

func (i Info) String() string {
	return Stringify(i)
}

// Endpoints reprents endpoints of the Theta API server.
type Endpoints struct {
	HTTPPort        *int   `json:"httpPort"`
	HTTPUpdatesPort *int   `json:"httpUpdatesPort"`
	APILevel        []int `json:"apiLevel"`
}

func (e Endpoints) String() string {
	return Stringify(e)
}

// InfoServices handles communication with the info of connected Theta.
type InfoServices service

// Get the Theta information.
func (s *InfoServices) Get(ctx context.Context) (*Info, *http.Response, error) {
	var info *Info
	ep := info.Endpoints
	fmt.Println(ep.HTTPPort)
	return info, nil, nil
}
