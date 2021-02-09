package main

import (
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageengine/pkg/engine"
	"github.com/sageflow/sageflow/pkg/envs"
	"github.com/sageflow/sageflow/pkg/logs"
	"github.com/sageflow/sageflow/pkg/configs"
)

func main() {
	// Set up log status file and load .env file.
	logs.SetStatusLogFile()
	envs.LoadEnvFile()

	// Connect to database.
	db := database.Connect()

	// TODO: Get from .env with a default value.
	const port = "3001"

	// Start a workflow engine gRPC server.
	eng := engine.NewEngine(db)

	// YAML
	if err := eng.ExecuteWorkflowString(configs.YAML, yamlString, &engine.Context{}); err != nil {
		panic(err)
	}

	// JSON
	if err := eng.ExecuteWorkflowString(configs.JSON, jsonString, &engine.Context{}); err != nil {
		panic(err)
	}

	// TOML
	if err := eng.ExecuteWorkflowString(configs.TOML, tomlString, &engine.Context{}); err != nil {
		panic(err)
	}
}

const yamlString = `
version: 1
kind: App

metadata:
  name: TeamCliff
  public_id: b780d792-bb57-4e29-a4a4-578d8db3bec9
  version: 1.0.0
  description: Lorem ipsum dolor sit amet
  category: accounting
  tags:
    - accounting
    - productivity
  avatar_id: TRgGRsfH635CSYGr2siYe7
  homepage_url: teamcliff.io
  resource_nouns:
    - email
  authors:
    - name: James Brown
      email: james@brown.com

auths:
  oauth2s:
    - label: OAuth2
      scopes:
        - send_email
      should_refresh_automatically: true
      fields:
        - key: email # Used to refer to this field
          label: Email # User friendly label. Some markdown formatting allowed
          tip: Enter your email here # User friendly description.
          is_required: true # If field has to be filled.
          is_administrative: true # If field meant for admin.
          input_kind: email # Corresponds somewhat to input[type] and select.
          default_value: null
          dropdown: null # For dropdown input types.
      apis:
        authorisation_request:
          code: null
          language: null
          form:
            method: GET
            url: "https://api.teamcliff.io/auth1"
            headers: null
            params:
              client_id: "{{env.CLIENT_ID}}"
              state: "{{env.STATE}}"
              redirect_uri: "{{env.REDIRECT_URI}}"
            body: null
        access_token_request:
          code: null
          language: null
          form:
            method: GET
            url: "https://api.teamcliff.io/auth2"
            headers: null
            params: null
            body:
              code: "{{auths.oauth2s[0].apis.authorisation_request.response.body.code}}"
              client_id: "{{env.CLIENT_ID}}"
              client_secret: "{{env.CLIENT_SECRET}}"
              redirect_uri: "{{env.REDIRECT_URI}}"
              grant_type: "authorization_code"
        refresh_token_request:
          code: null
          language: null
          form:
            method: GET
            url: "https://api.teamcliff.io/auth3"
            headers: null
            params:
              refresh_token: "{{auths.oauth2s[0].apis.authorisation_request.response.body.refresh_token}}"
              grant_type: "refresh_token"
            body: null
  api_keys:
    - label: API Key
      fields:
        - key: workspace
          label: Workspace
          tip: Select workspace
          is_required: true
          is_administrative: false
          input_kind: select
          default_value: null
          dropdown:
            kind: static # Options can be static or dynamic values
            allows_multiple: false # If multiple options can be selected
            allows_custom: true # If you can add you own custom entry
            options:
              - Google
              - Apple
      api:
        code: null
        language: null
        form:
          method: POST
          url: "https://api.teamcliff.io/auth"
          headers: null
          params: null
          body: null

operations:
  triggers:
    - key: new_email_subscriber
      label: New Email Subscriber
      tip: Triggers when there is a new email subscriber
      fields:
        - key: email
          label: Email
          tip: Enter your email here
          is_required: true
          is_write_op: false # If an operation that modifies or deletes resources.
          is_identification: true # If it is used to identify resource to operate on.
          resource_noun: email # A noun that best identifies what this trigger operates on.
          input_kind: select
          dropdown:
            kind: dynamic
            allows_multiple: false
            allows_custom: true
            options: "{{fields[0].email}}"
      apis:
        polls:
          - auth_kind: null
            code: null
            language: null
            form:
              method: GET
              url: "https://api.teamcliff.io/poll"
              headers: null
              params: null
              body: null
        rest_hooks:
          - auth_kind: null
            operations:
              subscribe:
                code: null
                language: null
                form:
                  method: POST
                  url: "https://api.teamcliff.io/hooks"
                  headers: null
                  params: null
                  body: null
              unsubscribe:
                code: null
                language: null
                form:
                  method: DELETE
                  url: "https://api.teamcliff.io/hooks"
                  headers: null
                  params: null
                  body: null
              list: # For checking recent items to provide sample data.
                code: null
                language: null
                form:
                  method: POST
                  url: "https://api.teamcliff.io/api"
                  headers: null
                  params: null
                  body: null

  actions:
    - action_kind: search
      key: email
      label: Email
      tip: Enter your email here
      is_required: true
      is_write_op: false # If an operation that modifies or deletes resources.
      is_identification: true # If it is used to identify resource to operate on.
      resource_noun: email # A noun that best identifies what this trigger operates on.
      input_kind: select
      dropdown:
        kind: dynamic
        allows_multiple: false
        allows_custom: true
        options: "{{fields[0].email}}"
      api:
        code: null
        language: null
        form:
          method: POST
          url: "https://api.teamcliff.io/auth"
          headers: null
          params: null
          body: null
`

const jsonString = `
{
	"version": 1,
	"kind": "App",
	"metadata": {
	  "name": "TeamCliff",
	  "public_id": "b780d792-bb57-4e29-a4a4-578d8db3bec9",
	  "version": "1.0.0",
	  "description": "Lorem ipsum dolor sit amet",
	  "category": "accounting",
	  "tags": [
		"accounting",
		"productivity"
	  ],
	  "avatar_id": "TRgGRsfH635CSYGr2siYe7",
	  "homepage_url": "teamcliff.io",
	  "resource_nouns": [
		"email"
	  ],
	  "authors": [
		{
		  "name": "James Brown",
		  "email": "james@brown.com"
		}
	  ]
	},
	"auths": {
	  "oauth2s": [
		{
		  "label": "OAuth2",
		  "scopes": [
			"send_email"
		  ],
		  "should_refresh_automatically": true,
		  "fields": [
			{
			  "key": "email",
			  "label": "Email",
			  "tip": "Enter your email here",
			  "is_required": true,
			  "is_administrative": true,
			  "input_kind": "email",
			  "default_value": null,
			  "dropdown": null
			}
		  ],
		  "apis": {
			"authorisation_request": {
			  "code": null,
			  "language": null,
			  "form": {
				"method": "GET",
				"url": "https://api.teamcliff.io/auth1",
				"headers": null,
				"params": {
				  "client_id": "{{env.CLIENT_ID}}",
				  "state": "{{env.STATE}}",
				  "redirect_uri": "{{env.REDIRECT_URI}}"
				},
				"body": null
			  }
			},
			"access_token_request": {
			  "code": null,
			  "language": null,
			  "form": {
				"method": "GET",
				"url": "https://api.teamcliff.io/auth2",
				"headers": null,
				"params": null,
				"body": {
				  "code": "{{auths.oauth2s[0].apis.authorisation_request.response.body.code}}",
				  "client_id": "{{env.CLIENT_ID}}",
				  "client_secret": "{{env.CLIENT_SECRET}}",
				  "redirect_uri": "{{env.REDIRECT_URI}}",
				  "grant_type": "authorization_code"
				}
			  }
			},
			"refresh_token_request": {
			  "code": null,
			  "language": null,
			  "form": {
				"method": "GET",
				"url": "https://api.teamcliff.io/auth3",
				"headers": null,
				"params": {
				  "refresh_token": "{{auths.oauth2s[0].apis.authorisation_request.response.body.refresh_token}}",
				  "grant_type": "refresh_token"
				},
				"body": null
			  }
			}
		  }
		}
	  ],
	  "api_keys": [
		{
		  "label": "API Key",
		  "fields": [
			{
			  "key": "workspace",
			  "label": "Workspace",
			  "tip": "Select workspace",
			  "is_required": true,
			  "is_administrative": false,
			  "input_kind": "select",
			  "default_value": null,
			  "dropdown": {
				"kind": "static",
				"allows_multiple": false,
				"allows_custom": true,
				"options": [
				  "Google",
				  "Apple"
				]
			  }
			}
		  ],
		  "api": {
			"code": null,
			"language": null,
			"form": {
			  "method": "POST",
			  "url": "https://api.teamcliff.io/auth",
			  "headers": null,
			  "params": null,
			  "body": null
			}
		  }
		}
	  ]
	},
	"operations": {
	  "triggers": [
		{
		  "key": "new_email_subscriber",
		  "label": "New Email Subscriber",
		  "tip": "Triggers when there is a new email subscriber",
		  "fields": [
			{
			  "key": "email",
			  "label": "Email",
			  "tip": "Enter your email here",
			  "is_required": true,
			  "is_write_op": false,
			  "is_identification": true,
			  "resource_noun": "email",
			  "input_kind": "select",
			  "dropdown": {
				"kind": "dynamic",
				"allows_multiple": false,
				"allows_custom": true,
				"options": "{{fields[0].email}}"
			  }
			}
		  ],
		  "apis": {
			"polls": [
			  {
				"auth_kind": null,
				"code": null,
				"language": null,
				"form": {
				  "method": "GET",
				  "url": "https://api.teamcliff.io/poll",
				  "headers": null,
				  "params": null,
				  "body": null
				}
			  }
			],
			"rest_hooks": [
			  {
				"auth_kind": null,
				"operations": {
				  "subscribe": {
					"code": null,
					"language": null,
					"form": {
					  "method": "POST",
					  "url": "https://api.teamcliff.io/hooks",
					  "headers": null,
					  "params": null,
					  "body": null
					}
				  },
				  "unsubscribe": {
					"code": null,
					"language": null,
					"form": {
					  "method": "DELETE",
					  "url": "https://api.teamcliff.io/hooks",
					  "headers": null,
					  "params": null,
					  "body": null
					}
				  },
				  "list": {
					"code": null,
					"language": null,
					"form": {
					  "method": "POST",
					  "url": "https://api.teamcliff.io/api",
					  "headers": null,
					  "params": null,
					  "body": null
					}
				  }
				}
			  }
			]
		  }
		}
	  ],
	  "actions": [
		{
		  "action_kind": "search",
		  "key": "email",
		  "label": "Email",
		  "tip": "Enter your email here",
		  "is_required": true,
		  "is_write_op": false,
		  "is_identification": true,
		  "resource_noun": "email",
		  "input_kind": "select",
		  "dropdown": {
			"kind": "dynamic",
			"allows_multiple": false,
			"allows_custom": true,
			"options": "{{fields[0].email}}"
		  },
		  "api": {
			"code": null,
			"language": null,
			"form": {
			  "method": "POST",
			  "url": "https://api.teamcliff.io/auth",
			  "headers": null,
			  "params": null,
			  "body": null
			}
		  }
		}
	  ]
	}
  }
`

const tomlString = `
version = 1
kind = "App"

[metadata]
name = "TeamCliff"
public_id = "b780d792-bb57-4e29-a4a4-578d8db3bec9"
version = "1.0.0"
description = "Lorem ipsum dolor sit amet"
category = "accounting"
tags = [ "accounting", "productivity" ]
avatar_id = "TRgGRsfH635CSYGr2siYe7"
homepage_url = "teamcliff.io"
resource_nouns = [ "email" ]

  [[metadata.authors]]
  name = "James Brown"
  email = "james@brown.com"

[[auths.oauth2s]]
label = "OAuth2"
scopes = [ "send_email" ]
should_refresh_automatically = true

  [[auths.oauth2s.fields]]
  key = "email"
  label = "Email"
  tip = "Enter your email here"
  is_required = true
  is_administrative = true
  input_kind = "email"

[auths.oauth2s.apis.authorisation_request]
  [auths.oauth2s.apis.authorisation_request.form]
  method = "GET"
  url = "https://api.teamcliff.io/auth1"

    [auths.oauth2s.apis.authorisation_request.form.params]
    client_id = "{{env.CLIENT_ID}}"
    state = "{{env.STATE}}"
    redirect_uri = "{{env.REDIRECT_URI}}"

[auths.oauth2s.apis.access_token_request]
  [auths.oauth2s.apis.access_token_request.form]
  method = "GET"
  url = "https://api.teamcliff.io/auth2"

    [auths.oauth2s.apis.access_token_request.form.body]
    code = "{{auths.oauth2s[0].apis.authorisation_request.response.body.code}}"
    client_id = "{{env.CLIENT_ID}}"
    client_secret = "{{env.CLIENT_SECRET}}"
    redirect_uri = "{{env.REDIRECT_URI}}"
    grant_type = "authorization_code"

[auths.oauth2s.apis.refresh_token_request]
  [auths.oauth2s.apis.refresh_token_request.form]
  method = "GET"
  url = "https://api.teamcliff.io/auth3"

    [auths.oauth2s.apis.refresh_token_request.form.params]
    refresh_token = "{{auths.oauth2s[0].apis.authorisation_request.response.body.refresh_token}}"
    grant_type = "refresh_token"

[[auths.api_keys]]
label = "API Key"

  [[auths.api_keys.fields]]
  key = "workspace"
  label = "Workspace"
  tip = "Select workspace"
  is_required = true
  is_administrative = false
  input_kind = "select"

    [auths.api_keys.fields.dropdown]
    kind = "static"
    allows_multiple = false
    allows_custom = true
    options = [ "Google", "Apple" ]

  [auths.api_keys.api]
    [auths.api_keys.api.form]
    method = "POST"
    url = "https://api.teamcliff.io/auth"

[[operations.triggers]]
key = "new_email_subscriber"
label = "New Email Subscriber"
tip = "Triggers when there is a new email subscriber"

  [[operations.triggers.fields]]
  key = "email"
  label = "Email"
  tip = "Enter your email here"
  is_required = true
  is_write_op = false
  is_identification = true
  resource_noun = "email"
  input_kind = "select"

    [operations.triggers.fields.dropdown]
    kind = "dynamic"
    allows_multiple = false
    allows_custom = true
    options = "{{fields[0].email}}"

[[operations.triggers.apis.polls]]
  [operations.triggers.apis.polls.form]
  method = "GET"
  url = "https://api.teamcliff.io/poll"

[[operations.triggers.apis.rest_hooks]]
[operations.triggers.apis.rest_hooks.operations.subscribe]
  [operations.triggers.apis.rest_hooks.operations.subscribe.form]
  method = "POST"
  url = "https://api.teamcliff.io/hooks"

[operations.triggers.apis.rest_hooks.operations.unsubscribe]
  [operations.triggers.apis.rest_hooks.operations.unsubscribe.form]
  method = "DELETE"
  url = "https://api.teamcliff.io/hooks"

[operations.triggers.apis.rest_hooks.operations.list]
  [operations.triggers.apis.rest_hooks.operations.list.form]
  method = "POST"
  url = "https://api.teamcliff.io/api"

[[operations.actions]]
action_kind = "search"
key = "email"
label = "Email"
tip = "Enter your email here"
is_required = true
is_write_op = false
is_identification = true
resource_noun = "email"
input_kind = "select"

  [operations.actions.dropdown]
  kind = "dynamic"
  allows_multiple = false
  allows_custom = true
  options = "{{fields[0].email}}"

  [operations.actions.api]
    [operations.actions.api.form]
    method = "POST"
    url = "https://api.teamcliff.io/auth"

`
