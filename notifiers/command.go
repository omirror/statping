package notifiers

import (
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/notifier"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"strings"
	"time"
)

var _ notifier.Notifier = (*commandLine)(nil)

type commandLine struct {
	*notifications.Notification
}

func (c *commandLine) Select() *notifications.Notification {
	return c.Notification
}

var Command = &commandLine{&notifications.Notification{
	Method:      "command",
	Title:       "Shell Command",
	Description: "Shell Command allows you to run a customized shell/bash Command on the local machine it's running on.",
	Author:      "Hunter Long",
	AuthorUrl:   "https://github.com/hunterlong",
	Delay:       time.Duration(1 * time.Second),
	Icon:        "fas fa-terminal",
	Host:        "/bin/bash",
	Limits:      60,
	Form: []notifications.NotificationForm{{
		Type:        "text",
		Title:       "Shell or Bash",
		Placeholder: "/bin/bash",
		DbField:     "host",
		SmallText:   "You can use '/bin/sh', '/bin/bash' or even an absolute path for an application.",
	}, {
		Type:        "text",
		Title:       "Command to Run on OnSuccess",
		Placeholder: "curl google.com",
		DbField:     "var1",
		SmallText:   "This Command will run every time a service is receiving a Successful event.",
	}, {
		Type:        "text",
		Title:       "Command to Run on OnFailure",
		Placeholder: "curl offline.com",
		DbField:     "var2",
		SmallText:   "This Command will run every time a service is receiving a Failing event.",
	}}},
}

func runCommand(app string, cmd ...string) (string, string, error) {
	outStr, errStr, err := utils.Command(app, cmd...)
	return outStr, errStr, err
}

// OnFailure for commandLine will trigger failing service
func (c *commandLine) OnFailure(s *services.Service, f *failures.Failure) error {
	msg := c.GetValue("var2")
	tmpl := ReplaceVars(msg, s, f)
	_, _, err := runCommand(c.Host, tmpl)
	return err
}

// OnSuccess for commandLine will trigger successful service
func (c *commandLine) OnSuccess(s *services.Service) error {
	msg := c.GetValue("var1")
	tmpl := ReplaceVars(msg, s, nil)
	_, _, err := runCommand(c.Host, tmpl)
	return err
}

// OnTest for commandLine triggers when this notifier has been saved
func (c *commandLine) OnTest() error {
	cmds := strings.Split(c.Var1, " ")
	in, out, err := runCommand(c.Host, cmds...)
	utils.Log.Infoln(in)
	utils.Log.Infoln(out)
	return err
}
