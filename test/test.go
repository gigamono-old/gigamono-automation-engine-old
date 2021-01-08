package main

import (
	"github.com/sageflow/sagedb"
	"github.com/sageflow/sageengine"
	"github.com/sageflow/sageutils"
)

func main() {
	// Set up log status file and load .env file.
	sageutils.SetStatusLogFile()
	sageutils.LoadEnvFile()

	// Connect to database.
	db := sagedb.Connect()

	// Start a workflow engine gRPC server.
	engine := sageengine.NewEngine(db)
	engine.ExecuteWorkflowYAML(yamlString, sageengine.Context{})
	engine.ExecuteWorkflowJSON(jsonString, sageengine.Context{})
}

const yamlString = `
version: 1
kind: Workflow

metadata:
  name: TeamCliff > Sorbel
  description: Lorem ipsum dolor sit amet
  execution_contexts: [protected]
  authors:
    - name: James Brown
      email: james@brown.com

tasks:
  - kind: trigger
    name: teamcliff
    index: 0
    dependencies: null
    position: [200, 150]
    execution_context: protected
    app_name: TeamCliff
    app_id: 24f495c3-7fc6-4bea-bf7a-adb0a955a54f
    account_id: 216e424e-0f1a-4a53-bdc2-27d6d277c347
    fields:
      email: james@brown.com
      name: James Brown

  - kind: action
    name: sorbel
    index: 1
    dependencies: [0] # Task 1 depends on Task 0
    position: [400, 150]
    execution_context: protected
    app_name: Sorbel
    app_id: ca41c8c5-dd47-4474-92ef-4b786b435663
    account_id: 647a954b-ed59-4ad5-990e-03f4177c57da
    fields:
      email: "$(tasks[0].fields.email)" # reference by task index
      name: "$(tasks.sorbel.fields.email)" # reference by task name
`

const jsonString = `
{
	"version": 1,
	"kind": "Workflow",
	"metadata": {
	  "name": "TeamCliff > Sorbel",
	  "description": "Lorem ipsum dolor sit amet",
	  "execution_contexts": [
		"protected"
	  ],
	  "authors": [
		{
		  "name": "James Brown",
		  "email": "james@brown.com"
		}
	  ]
	},
	"tasks": [
	  {
		"kind": "trigger",
		"name": "teamcliff",
		"index": 0,
		"dependencies": null,
		"position": [
		  200,
		  150
		],
		"execution_context": "protected",
		"app_name": "TeamCliff",
		"app_id": "24f495c3-7fc6-4bea-bf7a-adb0a955a54f",
		"account_id": "216e424e-0f1a-4a53-bdc2-27d6d277c347",
		"fields": {
		  "email": "james@brown.com",
		  "name": "James Brown"
		}
	  },
	  {
		"kind": "action",
		"name": "sorbel",
		"index": 1,
		"dependencies": [
		  0
		],
		"position": [
		  400,
		  150
		],
		"execution_context": "protected",
		"app_name": "Sorbel",
		"app_id": "ca41c8c5-dd47-4474-92ef-4b786b435663",
		"account_id": "647a954b-ed59-4ad5-990e-03f4177c57da",
		"fields": {
		  "email": "$(tasks[0].fields.email)",
		  "name": "$(tasks.sorbel.fields.email)"
		}
	  }
	]
}
`
