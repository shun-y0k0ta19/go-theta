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

// CommandRequest represents a Commands request from Theta API.
type CommandRequest struct {
	Name       *string     `json:"name,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

func (c CommandRequest) String() string {
	return Stringify(c)
}

// Parameters represents a command request parameters.
type Parameters struct {
	Options *Options `json:"options,omitempty"`

	// Deprecated in Theta API v2.1 (OSC v2.0).
	SessionID *string `json:"sessionId,omitempty"`
}

func (p Parameters) String() string {
	return Stringify(p)
}

// CommandResponse represents a Commands response from Theta API.
type CommandResponse struct {
	Name     *string   `json:"name"`
	State    *string   `json:"state"`
	ID       *string   `json:"id"`
	Results  *Results  `json:"results"`
	Error    *Error    `json:"error"`
	Progress *Progress `json:"progress"`
}

func (c CommandResponse) String() string {
	return Stringify(c)
}

// Results in command request.
type Results struct {
	Timeout *int `json:"timeout"`

	FileURI *string `json:"fileUri"`

	Entries      *Entries `json:"entries"`
	TotalEntries *int     `json:"totalEntries"`

	EXIF *EXIF `json:"exif"`
	XMP  *XMP  `json:"XMP"`

	Options *Options `json:"options"`

	// Deprecated in Theta API v2.1 (OSC v2.0).
	SessionID         *string `json:"sessionId"`
	ContinuationToken *string `json:"continuationToken"`
}

func (r Results) String() string {
	return Stringify(r)
}

// Entries represents file list of acquired still image files.
type Entries struct {
	Name                     *string `json:"name"`
	FileURL                  *string `json:"fileUrl"`
	Size                     *int    `json:"size"`
	DateTimeZone             *string `json:"dataTimeZone"`
	DateTime                 *string `json:"dateTime"`
	Width                    *int    `json:"width"`
	Height                   *int    `json:"height"`
	RecordTime               *int    `json:"_recordTime"`
	Thumbnail                *string `json:"_thumbnail"`
	ThumbSize                *string `json:"_thmubSize"`
	IntervalCaptureGroupID   *string `json:"_intervalCaptureGroupId"`
	CompositeShootingGroupID *string `json:"_compositeShootingGroupId"`
	AutoBracketGroupID       *string `json:"_autoBracketGroupId"`
	IsProcessed              *bool   `json:"isProcessed"`
	PreviewURL               *string `json:"previewUrl"`

	// Deprecated in Theta API v2.1 (OSC v2.0).
	RecordTimev20 *int    `json:"recordTime"`
	URI           *string `json:"uri"`

	// Not supported by Theta.
	Lat *int `json:"lat"`
	Lng *int `json:"lng"`
}

func (e Entries) String() string {
	return Stringify(e)
}

// EXIF represents exif information.
type EXIF struct {
	EXIFVersion       *string  `json:"ExifVersion"`
	ImageDescription  *string  `json:"ImageDescription"`
	DateTime          *string  `json:"DateTime"`
	ImageWidth        *int     `json:"ImageWidth"`
	ImageLength       *int     `json:"ImageLength"`
	ColorSpace        *int     `json:"ColorSpace"`
	Compression       *int     `json:"Compression"`
	Orientation       *int     `json:"Orientation"`
	Flash             *int     `json:"Flash"`
	FocalLength       *float64 `json:"FocalLength"`
	WhiteBalance      *int     `json:"WhiteBalance"`
	ExposureTime      *float64 `json:"ExposureTime"`
	ISOSpeedRatings   *int     `json:"ISOSpeedRatings"`
	ApertureValue     *float64 `json:"ApertureValue"`
	BrightnessValue   *float64 `json:"BrightnessValue"`
	ExposureBiasValue *float64 `json:"ExposureBiasValue"`
	GPSLatitudeRef    *string  `json:"GPSLatitudeRef"`
	GPSLatitude       *float64 `json:"GPSLatitude"`
	GPSLongitudeRef   *string  `json:"GPSLongitudeRef"`
	GPSLongitude      *float64 `json:"GPSLongitude"`
	Make              *string  `json:"Make"`
	Model             *string  `json:"Model"`
	Software          *string  `json:"Software"`
	Copyright         *string  `json:"Copyright"`
}

func (e EXIF) String() string {
	return Stringify(e)
}

// XMP represents exif information.
type XMP struct {
	ProjectionType               *string `json:"ProjectionType"`
	UsePanoramaViewer            *bool   `json:"UsePanoramaViewer"`
	CroppedAreaImageWidthPixels  *int    `json:"CroppedAreaImageWidthPixels"`
	CroppedAreaImageHeightPixels *int    `json:"CroppedAreaImageHeightPixels"`
	FullPanoWidthPixels          *int    `json:"FullPanoWidthPixels"`
	FullPanoHeightPixels         *int    `json:"FullPanoHeightPixels"`
	CroppedAreaLeftPixels        *int    `json:"CroppedAreaLeftPixels"`
	CroppedAreaTopPixels         *int    `json:"CroppedAreaTopPixels"`
}

func (x XMP) String() string {
	return Stringify(x)
}

// Error in command request.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) String() string {
	return Stringify(e)
}

// Progress represents parameters in progress.
type Progress struct {
	SessionID string `json:"sessionId"`
	Timeout   int    `json:"timeout"`
}

func (p Progress) String() string {
	return Stringify(p)
}

// CommandServices handles communication with the command methods of Theta API.
type CommandServices service

func (s *CommandServices) commandsExecute(ctx context.Context, body interface{}) (*CommandResponse, *http.Response, error) {
	req, err := s.client.NewRequest("POST", commandsExecuteURL, body)
	if err != nil {
		return nil, nil, err
	}
	commandResponse := new(CommandResponse)
	resp, err := s.client.Do(ctx, req, commandResponse)
	if err != nil {
		return nil, resp, err
	}
	fmt.Printf("Name: %v\nState: %v\n\n", *(commandResponse.Name), *(commandResponse.State))
	return commandResponse, resp, nil
}
