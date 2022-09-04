package collector

import (
	"github.com/pedrolopesme/open-rba/internal/core/domains"
)

func validate(data *domains.AuthenticationData) error {
	if data.UserID == "" {
		return ErrEmptyUserId
	}

	if data.IP == "" {
		return ErrEmptyIP
	}

	return nil
}
