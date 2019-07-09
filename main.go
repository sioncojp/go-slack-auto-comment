package main

import (
	"context"
	"errors"
	"os"
	"time"

	"os/signal"
	"syscall"

	"fmt"

	"github.com/sioncojp/go-slack-auto-comment/internal/pidfile"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	os.Exit(Run())
}

func init() {
	// logger初期化
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	s := logger.Sugar()
	log = Logger{logger, s}
}

// Run ... 実行
func Run() int {
	app := FlagSet()
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		log.sugar.Fatal(err)
		return 1
	}
	return 0
}

// run ... flag(cli)を初期化、graceful shutdownの設定を施して実行
func run(c *cli.Context) error {
	// create PIDFILE and defer remove
	if err := pidfile.Create(pid); err != nil {
		return fmt.Errorf("failed to remove the pidfile: %s: %s", pid, err)
	}
	defer pidfile.Remove(pid)

	if c.String("config-dir") == "" {
		return errors.New("required -c option")
	}

	conf, err := LoadToml(c.String("config-dir"), c.String("region"))
	if err != nil {
		return err
	}
	botID = conf.BotID

	exitCh := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	log.sugar.Info("start...")
	go Start(ctx, conf, exitCh)

	// receive syscall
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go signalHandler(sig, cancel)

	return <-exitCh
}

// Start ... start
func Start(ctx context.Context, c *Config, exitCh chan error) {
	slackListener := c.NewSlack()
	go slackListener.ListenAndResponse()

	// receive context
	select {
	case <-ctx.Done():
		log.sugar.Info("received done, exiting in 500 milliseconds")
		time.Sleep(500 * time.Millisecond)
		exitCh <- nil
		return
	}
}

// signalHandler ... Receive signal handler and do context.cancel
func signalHandler(sig chan os.Signal, cancel context.CancelFunc) {
	for {
		select {
		case s := <-sig:
			switch s {
			case syscall.SIGINT:
				log.sugar.Info("received SIGINT signal")
				log.sugar.Info("shutdown...")
				cancel()
			case syscall.SIGTERM:
				log.sugar.Info("received SIGTERM signal")
				log.sugar.Info("shutdown...")
				cancel()
			}
		}
	}
}
