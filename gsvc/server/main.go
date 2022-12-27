package main

import (
	"log"
	"os"
	"sync"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"gsvc/pkg/gateway"

	agmapi "gsvc/api/agm/restapi"
	agmop "gsvc/api/agm/restapi/operations"

	amapi "gsvc/api/am/restapi"
	amop "gsvc/api/am/restapi/operations"

	tsapi "gsvc/api/ts/restapi"
	tsop "gsvc/api/ts/restapi/operations"

	mcapi "gsvc/api/mc/restapi"
	mcop "gsvc/api/mc/restapi/operations"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go gateway.Start(8000)
	go agm(8001)
	go ts(8002)
	go am(8003)
	go mc(8004)
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

func mc(port int) {

	swaggerSpec, err := loads.Embedded(mcapi.SwaggerJSON, mcapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := mcop.NewMcAPI(swaggerSpec)
	server := mcapi.NewServer(api)
	defer server.Shutdown()

	server.Port = port

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "MindConnect API"
	parser.LongDescription = "\nMindConnect API provides following data ingestion functionalities:\n\n# Data Point Mappings\n\nCreating and managing mappings between an agent's data points and an entity's dynamic property to be able to upload TimeSeries data.\n\n\nEach agent has data points with unique ids. The mapping is between to this id to an entity's dynamic property set's property.\n\n- A data point can be mapped to many property of many property set of many\nentities.\n\n- A property cannot be mapped from more than one data point. \n\n- A propertyset can have mappings from many agents' many data points to its\nproperties.\n\n- The unit of the datapoint has to be same with the unit of the property.\n\n- The type of the datapoint has to be same with the type of the property.\n\n\nWhenever data source configuration of an agent is updated via Agent Management API; all mappings with __keepMapping__ attribute set gets their validity attribute updated and all mappings with __keepMapping__ attribute unset are deleted.\n\n\n# Exchange\n\nExchanging time series, events, files and data source configuration data. Combination of different data types can be uploaded via exchange endpoint within  multipart body. Maximum size of exchange body is 10MBs.\n\n# Diagnostic Activations\n\nManagement of Diagnostic Activations and querying Diagnostic Messages of time series, event, file and data source configuration requests.\n\n- Maximum 5 agents per tenant can be activated for data ingestion tracking.\n\n- For non-agents, the required permission allows to manage diagnostic activation resources of agents in the same tenant as in the token.\n\n- For agents, only the diagnostic activation related to the agent can be managed. Agents are forbidden to view/change the resources of other agents in the same tenant.\n\n- Agents are allowed to update activation for itself only. Users with sufficient scopes are allowed\n"
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
