package tools

import (
	"fmt"

	"github.com/Str1kez/SportiqSubscriptionService/internal/config"
)

func EventStatusChecker(config *config.EventStatus, eventStatus string) error {
	switch eventStatus {
	case config.Planned:
		return nil
	case config.Completed:
		return nil
	case config.Deleted:
		return nil
	case config.Underway:
		return nil
	default:
		return fmt.Errorf("event status [%s] not in event status config", eventStatus)
	}
}
