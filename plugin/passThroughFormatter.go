package plugin

import (
	"fmt"
	"strings"

	"github.com/sebsprenger/chatterschool/shared"
)

type MyPassThroughFormatter struct {
}

func (formatter MyPassThroughFormatter) Modify(msg shared.Message) shared.Message {
	fmt.Printf("Lauschangriff --> \"%s\" sagt \"%s\"\n", msg.Sender, msg.Text)

	// censoring
	msg.Text = strings.ReplaceAll(msg.Text, "scheisse", "s******e")
	msg.Text = strings.ReplaceAll(msg.Text, "scheiße", "s*****e")

	return msg
}
