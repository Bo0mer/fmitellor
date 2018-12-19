package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Bo0mer/logger"
	"github.com/nlopes/slack"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	logger := logger.NewLogger(os.Stdout)
	slacker := &slacker{
		Token: os.Getenv("SLACK_TOKEN"),
	}

	drone := &drone{Port: "8080", Logger: logger}
	drone.Start()

	incoming := slacker.Start()
	outgoing := drone.CommandsChan()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case c := <-incoming:
			logger.Printf("Executing command %q\n", c)
			outgoing <- c
		case <-sigChan:
			drone.Stop()
			return
		}
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

type drone struct {
	Port   string
	Logger *logger.Logger

	driver       *tello.Driver
	commandsChan chan Command
}

func (d *drone) Start() {
	d.commandsChan = make(chan Command)
	d.driver = tello.NewDriver(d.Port)

	work := func() {
		d.Logger.Printf("Drone is ready, taking off...\n")
		d.driver.TakeOff()
		d.driver.SetFastMode()
		go d.executeCommands()
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{d.driver},
		work,
	)

	go robot.Start()
}

func (d *drone) Stop() {
	d.driver.Land()
}

func (d *drone) CommandsChan() chan<- Command {
	return d.commandsChan
}

func (d *drone) executeCommands() {
	for c := range d.commandsChan {
		var err error
		switch c {
		case CommandHover:
			d.driver.Hover()
		case CommandMoveUp:
			err = d.driver.Up(10)
		case CommandMoveDown:
			err = d.driver.Down(10)
		case CommandMoveForward:
			err = d.driver.Forward(10)
		case CommandMoveBackward:
			err = d.driver.Backward(10)
		case CommandMoveLeft:
			err = d.driver.Left(10)
		case CommandMoveRight:
			err = d.driver.Right(10)
		}

		if err != nil {
			d.Logger.Printf("executing command %q has failed: %v; entering hover mode\n", c, err)
			d.driver.Hover()
		}
	}
}

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
