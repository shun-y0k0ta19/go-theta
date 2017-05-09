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

// Options represents Theta options.
type Options struct {
	Aparture               *float64  `json:"aparture,omitempty"`
	ApartureSupport        []float64 `json:"apartureSupport,omitempty"`
	AutoBracket            *Bracket  `json:"_autoBracket,omitempty"`
	AutoBracketSupport     []int     `json:"_autoBracketSupport,omitempty"`
	CaptureInterval        *int      `json:"_captureInterval,omitempty"`
	CaptureIntervalSupport []int     `json:"_captureIntervalSupport,omitempty"`
	ClientVersion          *int      `json:"clientVersion,omitempty"`
}

// Bracket represents an bracket parameters.
type Bracket struct {
	BracketNumber     int `json:"_bracketNumber,omitempty"`
	BracketParameters struct {
		ShutterSpeed     float64 `json:"shutterSpeed,omitempty"`
		ISO              int     `json:"iso,omitempty"`
		ColorTemperature int     `json:"_colorTemperature,omitempty"`
	} `json:"_bracketParameters,omitempty"`
}

// SetOptions sets options to the Theta.
func (s *CommandServices) SetOptions(ctx context.Context, options *Options) (*CommandResponse, *http.Response, error) {
	parameters := Parameters{Options: options}
	fmt.Printf("APILevel: %d\n", s.client.apiLevel)
	if s.client.apiLevel == 1 {
		parameters.SessionID = &s.client.sessionID
	}
	fmt.Printf("sessionID: %s\n", *parameters.SessionID)
	body := CommandRequest{
		Name:       String("camera.setOptions"),
		Parameters: &parameters,
	}
	return s.commandsExecute(ctx, body)
}
