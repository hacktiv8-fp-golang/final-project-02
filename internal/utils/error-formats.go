package utils

import "strings"

func ParseError(err error) Error {
	if strings.Contains(err.Error(), "no record found") {
		return NotFound("Data not found")
	}

	return InternalServerError("Something went wrong")
}