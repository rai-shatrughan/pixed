// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"gsvc/api/mc/restapi"
	"gsvc/api/mc/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewMcAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

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
