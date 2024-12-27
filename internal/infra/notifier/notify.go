package notifier

import (
	"fmt"
	"github.com/gen2brain/beeep"
)

func Notify(title, message string) error {
	beeep.DefaultDuration = 500
	err := beeep.Notify(title, message, "")
	if err != nil {
		return fmt.Errorf("erro ao exibir notificação: %v", err)
	}
	return nil
}
