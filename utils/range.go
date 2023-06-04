package utils

import (
	"errors"
	"strconv"
	"strings"
)

type rangeInfo struct {
	Start int64
	End   int64
}

func ParseRangeHeader(rangeHeader string, fileSize int64) (rangeInfo, error) {
	parts := strings.Split(rangeHeader, "=")
	if len(parts) != 2 {
		return rangeInfo{}, errors.New("Invalid Range header format")
	}

	rangeParts := strings.Split(parts[1], "-")
	if len(rangeParts) != 2 {
		return rangeInfo{}, errors.New("Invalid Range header format")
	}

	var start, end int64
	var err error

	start, err = strconv.ParseInt(rangeParts[0], 10, 64)
	if err != nil || start >= fileSize {
		return rangeInfo{}, errors.New("Invalid Range header format")
	}

	end, err = strconv.ParseInt(rangeParts[1], 10, 64)
	if err != nil || end >= fileSize {
		return rangeInfo{}, errors.New("Invalid Range header format")
	}

	if start > end {
		return rangeInfo{}, errors.New("Invalid Range header format")
	}

	return rangeInfo{Start: start, End: end}, nil
}
