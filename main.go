package main

import (
	"os"

	"github.com/Bo0mer/logger"
	"github.com/nlopes/slack"
)

func main() {
	logger := logger.NewLogger(os.Stdout)
	slacker := &slacker{
		Token: os.Getenv("SLACK_TOKEN"),
	}

	commandsChan := slacker.Start()

	for c := range commandsChan {
		logger.Printf("Got command: %q\n", c)
	}
}

type Command string

const (
	CommandMoveUp       Command = "up"
	CommandMoveDown     Command = "down"
	CommandMoveLeft     Command = "left"
	CommandMoveRight    Command = "right"
	CommandMoveForward  Command = "forward"
	CommandMoveBackward Command = "backward"
	CommandHover        Command = "hover"
)

type slacker struct {
	Token string

	commandsChan chan Command
}

func (s *slacker) Start() <-chan Command {
	s.commandsChan = make(chan Command)

	slackClient := slack.New(s.Token)
	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()
	go func() {
		for msg := range rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				s.handleMessage(ev)
			}
		}
	}()

	return s.commandsChan
}

func (s *slacker) handleMessage(msg *slack.MessageEvent) {
	switch msg.Msg.Text {
	case "up":
		s.commandsChan <- CommandMoveUp
	case "down":
		s.commandsChan <- CommandMoveDown
	case "forward":
		s.commandsChan <- CommandMoveForward
	case "backward":
		s.commandsChan <- CommandMoveBackward
	case "left":
		s.commandsChan <- CommandMoveLeft
	case "right":
		s.commandsChan <- CommandMoveRight
	case "hover", "hold", "stop", "pause", "panic":
		s.commandsChan <- CommandHover
	default:
		s.commandsChan <- CommandHover
	}
}
