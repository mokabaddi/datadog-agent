package main

import (
	_ "expvar"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/DataDog/datadog-agent/cmd/agent/app"

	// register core checks
	_ "github.com/DataDog/datadog-agent/pkg/collector/check/core/embed"
	_ "github.com/DataDog/datadog-agent/pkg/collector/check/core/network"
	_ "github.com/DataDog/datadog-agent/pkg/collector/check/core/system"
)

func main() {
	// go_expvar server
	go http.ListenAndServe("127.0.0.1:5000", http.DefaultServeMux)

	// Invoke the Agent
	if err := app.AgentCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
