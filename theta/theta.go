// Copyright (c) 2017 "Shun Yokota" All rights reserved
//
// Part of the source code is adapted from https://github.com/google/go-github
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package theta is a Go library for THETA API v2.
// THETA API v2 is compliant with Open Spherical Camera API from Google.
package theta

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL  = "http://192.168.1.1"
	defaultAPILevel = 1

	infoURL            = "/osc/info"
	stateURL           = "/osc/state"
	commandsExecuteURL = "/osc/commands/execute"
	commandStatusURL   = "/osc/commands/status"
)

var (
	// ErrClientIsNil is returned by Begin when the client is nil.
	ErrClientIsNil = errors.New("client is nil")
)

// Client manages communication with the THETA API.
type Client struct { // adapted from https://github.com/google/go-github
	client *http.Client // HTTP client used to communicate with the API.

	BaseURL  *url.URL // URL for a API requests.
	apiLevel int      // Theta API Level(1: v2.0, 2: v2.1).
	// sessionID of Theta API v2.0 (OSC v1.0). Deprecated in Theta API v2.1 (OSC v2.0).
	sessionID string

	common  service
	Info    *InfoServices
	Command *CommandServices
}

type service struct {
	client *Client
}

// NewClient returns a new THETA API client.
//
// adapted from https://github.com/google/go-github
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:   httpClient,
		BaseURL:  baseURL,
		apiLevel: defaultAPILevel,
	}
	c.common.client = c
	c.Info = (*InfoServices)(&c.common)
	c.Command = (*CommandServices)(&c.common)
	//c.Info = (*InfoServices)s
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) { // adapted from https://github.com/google/go-github
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	uri := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(uri.String())
	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Do sends an THETA API request and returns the THETA API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) { // adapted from https://github.com/google/go-github
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}
	return resp, err
}

// ErrorResponse reports one or more errors caused by an Theta API.
type ErrorResponse struct {
	Response *http.Response
	Code     string `json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
}

// CheckResponse checks the API response for errors, and returns them ifpresent.
func CheckResponse(r *http.Response) error { // adapted from https://github.com/google/go-github
	return nil
}

// Begin the session and set API Level to the Theta.
// Camera API version should be set to v2.1 with clientVersion in order to use
// RICOH THETA API v2.1, because the API version at the start point of the connection
// via wireless LAN is v2.0.
// When v2.1 is supported, API Version is set to v2.1 automatically. If you need to
// use v2.0, use StartSession and SetOptions to set to v2.0 manually.
func Begin(ctx context.Context, c *Client) error {
	if c == nil {
		return ErrClientIsNil
	}
	session, resp, err := c.Command.StartSession(ctx)
	if err != nil {
		return err
	}
	if session.Results != nil {
		if session.Results.SessionID != nil {
			fmt.Println(*session.Results.SessionID)
			c.sessionID = *session.Results.SessionID
		}
	}
	options := &Options{ClientVersion: Int(2)}
	if cmd, _, err := c.Command.SetOptions(ctx, options); err == nil {
		fmt.Println(*cmd.State)
		if cmd.Error != nil {
			fmt.Println(*cmd.Error)
		}
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v } // copied from https://github.com/google/go-github

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v } // copied from https://github.com/google/go-github

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v } // copied from https://github.com/google/go-github
