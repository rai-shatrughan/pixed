package main

import (
	"log"
	"os"
	"sync"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	agmapi "gsvc/api/agm/restapi"
	agmop "gsvc/api/agm/restapi/operations"
	"gsvc/pkg/gateway"

	amapi "gsvc/api/am/restapi"
	amop "gsvc/api/am/restapi/operations"

	tsapi "gsvc/api/ts/restapi"
	tsop "gsvc/api/ts/restapi/operations"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go gateway.Start(8000)
	go agm(8001)
	go ts(8002)
	go am(8003)
	wg.Wait()
}

func agm(port int) {
	swaggerSpec, err := loads.Embedded(agmapi.SwaggerJSON, agmapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := agmop.NewAgmAPI(swaggerSpec)
	server := agmapi.NewServer(api)
	defer server.Shutdown()

	server.Port = port

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Agent Management API"
	parser.LongDescription = "API defining resources and operations for managing agents.\n\nGenerating a Boarding Configuration action is an asynchronous operation therefore it may take a while. \nIn case Boarding Configuration is not generated, try to read the configuration again after a couple of seconds.\n"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

func ts(port int) {
	swaggerSpec, err := loads.Embedded(tsapi.SwaggerJSON, tsapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := tsop.NewTsAPI(swaggerSpec)
	server := tsapi.NewServer(api)
	defer server.Shutdown()

	server.Port = port

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "IoT Time Series API"
	parser.LongDescription = "Create, update, and query time series data with a precision of 1 millisecond."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

func am(port int) {
	swaggerSpec, err := loads.Embedded(amapi.SwaggerJSON, amapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := amop.NewAmAPI(swaggerSpec)
	server := amapi.NewServer(api)
	defer server.Shutdown()

	server.Port = port

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Asset Management API"
	parser.LongDescription = "Service for configuring, reading and managing assets, asset types and aspect types."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
