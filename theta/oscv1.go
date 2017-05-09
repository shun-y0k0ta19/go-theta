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
	"net/http"
)

// oscv1.go is descrived methods deprecated in Theta API v2.1(oscV2.0).

// StartSession begins the session. Issued the session ID. StartSession is deprecated
// in Theta API v2.1 (OSC v2.0).
func (s *CommandServices) StartSession(ctx context.Context) (*CommandResponse, *http.Response, error) {
	return s.commandsExecute(ctx, CommandRequest{Name: String("camera.startSession")})
}
