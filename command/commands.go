package command

import (
	"github.com/innogames/slack-bot/bot"
	"github.com/innogames/slack-bot/client"
	"github.com/innogames/slack-bot/command/calendar"
	"github.com/innogames/slack-bot/command/cron"
	"github.com/innogames/slack-bot/command/custom"
	"github.com/innogames/slack-bot/command/games"
	"github.com/innogames/slack-bot/command/jenkins"
	"github.com/innogames/slack-bot/command/jira"
	"github.com/innogames/slack-bot/command/mqtt"
	"github.com/innogames/slack-bot/command/pullrequest"
	"github.com/innogames/slack-bot/command/queue"
	"github.com/innogames/slack-bot/config"
	"github.com/sirupsen/logrus"
)

// GetCommands returns the list of default command which are available
func GetCommands(slackClient client.SlackClient, cfg config.Config, logger *logrus.Logger) *bot.Commands {
	commands := &bot.Commands{}
	commands.AddCommand(
		// needs to be the first commands to store all executed commands
		NewRetryCommand(slackClient),

		NewMacroCommand(slackClient, cfg.Macros, logger),
		NewReplyCommand(slackClient),
		NewAddLinkCommand(slackClient),
		NewReactionCommand(slackClient),
		NewSendMessageCommand(slackClient),
		NewDelayCommand(slackClient),
		NewBotLogCommand(slackClient, cfg),
		NewRandomCommand(slackClient),
		NewHelpCommand(slackClient, commands),

		games.NewNumberGuesserCommand(slackClient),

		calendar.NewCalendarCommand(cfg.Calendars, logger),

		mqtt.NewMqttCommand(slackClient, cfg.Mqtt),

		cron.NewCronCommand(slackClient, logger, cfg.Crons),

		queue.NewQueueCommand(slackClient, logger),
		queue.NewListCommand(slackClient),

		custom.GetCommand(slackClient),
	)

	// jenkins
	commands.Merge(jira.GetCommands(cfg.Jira, slackClient, logger))

	// jenkins
	commands.Merge(jenkins.GetCommands(cfg.Jenkins, slackClient, logger))

	// pull-request
	commands.Merge(pullrequest.GetCommands(slackClient, cfg))

	return commands
}