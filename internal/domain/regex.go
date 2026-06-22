package domain

import "regexp"

var ownerNameRegex = regexp.MustCompile(`^[a-zA-Z0-9]{2,255}$`)
