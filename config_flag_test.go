package pdmq_test

import (
	"flag"

	"github.com/pdmqio/go-pdmq"
)

func ExampleConfigFlag() {
	cfg := pdmq.NewConfig()
	flagSet := flag.NewFlagSet("", flag.ExitOnError)

	flagSet.Var(&pdmq.ConfigFlag{cfg}, "consumer-opt", "option to pass through to pdmq.Consumer (may be given multiple times)")
	flagSet.PrintDefaults()

	err := flagSet.Parse([]string{
		"--consumer-opt=heartbeat_interval,1s",
		"--consumer-opt=max_attempts,10",
	})
	if err != nil {
		panic(err.Error())
	}
	println("HeartbeatInterval", cfg.HeartbeatInterval)
	println("MaxAttempts", cfg.MaxAttempts)
}
