# Go SDK for OpenFGA

[![Go Reference](https://pkg.go.dev/badge/github.com/openfga/go-sdk.svg)](https://pkg.go.dev/github.com/openfga/go-sdk)
[![Release](https://img.shields.io/github/v/release/openfga/go-sdk?sort=semver&color=green)](https://github.com/openfga/go-sdk/releases)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](./LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B4989%2Fgithub.com%2Fopenfga%2Fgo-sdk.svg?type=shield)](https://app.fossa.com/reports/1c1c164d-b721-47e0-9165-400029aad1e4)
[![Discord Server](https://img.shields.io/discord/759188666072825867?color=7289da&logo=discord "Discord Server")](https://discord.com/channels/759188666072825867/930524706854031421)
[![Twitter](https://img.shields.io/twitter/follow/openfga?color=%23179CF0&logo=twitter&style=flat-square "@openfga on Twitter")](https://twitter.com/openfga)

This is an autogenerated Go SDK for OpenFGA. It provides a wrapper around the [OpenFGA API definition](https://openfga.dev/api).

## Table of Contents

- [About OpenFGA](#about)
- [Resources](#resources)
- [Installation](#installation)
- [Getting Started](#getting-started)
  - [Initializing the API Client](#initializing-the-api-client)
  - [Get your Store ID](#get-your-store-id)
  - [Calling the API](#calling-the-api)
    - [List All Stores](#list-stores)
    - [Create a Store](#create-store)
    - [Get a Store](#get-store)
    - [Delete a Store](#delete-store)
    - [Write Authorization Model](#write-authorization-model)
    - [Read a Single Authorization Model](#read-a-single-authorization-model)
    - [Read Authorization Model IDs](#read-authorization-model-ids)
    - [Check](#check)
    - [Write Tuples](#write-tuples)
    - [Delete Tuples](#delete-tuples)
    - [Expand](#expand)
    - [Read Tuples](#read-tuples)
    - [Read Changes (Watch)](#read-changes-watch)
  - [API Endpoints](#api-endpoints)
  - [Models](#models)
- [Contributing](#contributing)
  - [Issues](#issues)
  - [Pull Requests](#pull-requests)
- [License](#license)

## <a id="about">About OpenFGA</a>

[OpenFGA](https://openfga.dev) is an open source Fine-Grained Authorization solution inspired by [Google's Zanzibar paper](https://research.google/pubs/pub48190/). It was created by the FGA team at [Auth0](https://auth0.com) based on [Auth0 Fine-Grained Authorization (FGA)](https://fga.dev), available under [a permissive license (Apache-2)](https://github.com/openfga/rfcs/blob/main/LICENSE) and welcomes community contributions.

OpenFGA is designed to make it easy for application builders to model their permission layer, and to add and integrate fine-grained authorization into their applications. OpenFGA’s design is optimized for reliability and low latency at a high scale.

It allows in-memory data storage for quick development, as well as pluggable database modules - with initial support for PostgreSQL.

It offers an [HTTP API](https://openfga.dev/api) and has SDKs for programming languages including [Node.js/JavaScript](https://github.com/openfga/js-sdk), [GoLang](https://github.com/openfga/go-sdk) and [.NET](https://github.com/openfga/dotnet-sdk).

More SDKs and integrations such as Rego are planned for the future.

## Resources

- [OpenFGA Documentation](https://openfga.dev/docs)
- [OpenFGA API Documentation](https://openfga.dev/api)
- [Twitter](https://twitter.com/openfga)
- [OpenFGA Discord Community](https://discord.gg/8naAwJfWN6)
- [Zanzibar Academy](https://zanzibar.academy)
- [Google's Zanzibar Paper (2019)](https://research.google/pubs/pub48190/)

## Installation

To install:

```
go get -u github.com/openfga/go-sdk
```

In your code, import the module and use it:

```go
import "github.com/openfga/go-sdk"

func Main() {
	configuration, err := openfga.NewConfiguration(openfga.UserConfiguration{})
}
```

You can then run

```shell
go mod tidy
```

to update `go.mod` and `go.sum` if you are using them.

## Getting Started

### Initializing the API Client

[Learn how to initialize your SDK](https://openfga.dev/docs/getting-started/setup-sdk-client)

Without an API Token

```golang
import (
    openfga "github.com/openfga/go-sdk"
    "os"
)

func main() {
    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiScheme:      os.Getenv("OPENFGA_API_SCHEME"), // optional, defaults to "https"
        ApiHost:        os.Getenv("OPENFGA_API_HOST"), // required, define without the scheme (e.g. api.fga.example instead of https://api.fga.example)
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)
}
```

With an API Token

```golang
import (
    openfga "github.com/openfga/go-sdk"
    "os"
)

func main() {
    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiScheme:      os.Getenv("OPENFGA_API_SCHEME"), // optional, defaults to "https"
        ApiHost:        os.Getenv("OPENFGA_API_HOST"), // required, define without the scheme (e.g. api.fga.example instead of https://api.fga.example)
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
        Credentials: &credentials.Credentials{
            Method: credentials.CredentialsMethodApiToken,
            Config: {
                ApiToken: os.Getenv("OPENFGA_API_TOKEN"), // will be passed as the "Authorization: Bearer ${ApiToken}" request header
            },
        },
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)
}
```


### Get your Store ID

You need your store id to call the OpenFGA API (unless it is to call the [CreateStore](#create-store) or [ListStores](#list-stores) methods).

If your server is configured with [authentication enabled](https://openfga.dev/docs/getting-started/setup-openfga#configuring-authentication), you also need to have your credentials ready.

### Calling the API

#### List Stores

[API Documentation](https://openfga.dev/api/docs/api#/Stores/ListStores)

```golang
configuration, err := openfga.NewConfiguration(openfga.Configuration{
    ApiHost: "api.fga.example"
})

if err != nil {
// .. Handle error
}

apiClient := openfga.NewAPIClient(configuration)
stores, response, err := apiClient.OpenFgaApi.ListStores(context.Background()).Execute();

// stores = [{ "id": "01FQH7V8BEG3GPQW93KTRFR8JB", "name": "FGA Demo Store", "created_at": "2022-01-01T00:00:00.000Z", "updated_at": "2022-01-01T00:00:00.000Z" }]
```

#### Create Store

[API Documentation](https://openfga.dev/api/docs/api#/Stores/CreateStore)

```golang
configuration, err := openfga.NewConfiguration(openfga.Configuration{
    ApiHost: "api.fga.example"
})

if err != nil {
// .. Handle error
}

apiClient := openfga.NewAPIClient(configuration)

store, _, err := apiClient.OpenFgaApi.CreateStore(context.Background()).Body(CreateStoreRequest{Name: PtrString("FGA Demo")}).Execute()
if err != nil {
    // handle error
}

// store.Id = "01FQH7V8BEG3GPQW93KTRFR8JB"

// store store.Id in database
// update the storeId of the current instance
apiClient.SetStoreId(*store.Id)
// continue calling the API normally
```

#### Get Store

[API Documentation](https://openfga.dev/api/docs/api#/Stores/GetStore)

> Requires a client initialized with a storeId

```golang
store, _, err := apiClient.OpenFgaApi.GetStore(context.Background()).Execute()
if err != nil {
    // handle error
}

// store = { "id": "01FQH7V8BEG3GPQW93KTRFR8JB", "name": "FGA Demo Store", "created_at": "2022-01-01T00:00:00.000Z", "updated_at": "2022-01-01T00:00:00.000Z" }
```

#### Delete Store

[API Documentation](https://openfga.dev/api/docs/api#/Stores/DeleteStore)

> Requires a client initialized with a storeId

```golang
_, _, err := apiClient.OpenFgaApi.DeleteStore(context.Background()).Execute()
if err != nil {
    // handle error
}
```

#### Write Authorization Model

[API Documentation](https://openfga.dev/api#/Authorization%20Models/WriteAuthorizationModel)

> Note: To learn how to build your authorization model, check the Docs at https://openfga.dev/docs.

> Learn more about [the OpenFGA configuration language](https://openfga.dev/docs/configuration-language).

```golang
body  := openfga.TypeDefinitions{TypeDefinitions: &[]openfga.TypeDefinition{
	{
		Type: "repo",
		Relations: map[string]openfga.Userset{
			"writer": {This: &map[string]interface{}{}},
			"reader": {Union: &openfga.Usersets{
				Child: &[]openfga.Userset{
					{This: &map[string]interface{}{}},
					{ComputedUserset: &openfga.ObjectRelation{
						Object:   openfga.PtrString(""),
						Relation: openfga.PtrString("writer"),
					}},
				},
			}},
		},
	},
}}
data, response, err := apiClient.OpenFgaApi.WriteAuthorizationModel(context.Background()).Body(body).Execute()

fmt.Printf("%s", data.AuthorizationModelId) // 1uHxCSuTP0VKPYSnkq1pbb1jeZw
```

#### Read a Single Authorization Model

[API Documentation](https://openfga.dev/api#/Authorization%20Models/ReadAuthorizationModel)

```golang
// Assuming `1uHxCSuTP0VKPYSnkq1pbb1jeZw` is an id of a single model
data, response, err := apiClient.OpenFgaApi.ReadAuthorizationModel(context.Background(), "1uHxCSuTP0VKPYSnkq1pbb1jeZw").Execute()

// data = {"authorization_model":{"id":"1uHxCSuTP0VKPYSnkq1pbb1jeZw","type_definitions":[{"type":"repo","relations":{"writer":{"this":{}},"reader":{ ... }}}]}} // JSON

fmt.Printf("%s", data.AuthorizationModel.Id) // 1uHxCSuTP0VKPYSnkq1pbb1jeZw
```

#### Read Authorization Model IDs

[API Documentation](https://openfga.dev/api#/Authorization%20Models/ReadAuthorizationModels)

```golang
data, response, err := apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background()).Execute()

// data = {"authorization_model_ids":["1uHxCSuTP0VKPYSnkq1pbb1jeZw","GtQpMohWezFmIbyXxVEocOCxxgq"]} // in JSON

fmt.Printf("%s", (*data.AuthorizationModelIds)[0]) // 1uHxCSuTP0VKPYSnkq1pbb1jeZw
```

#### Check

[API Documentation](https://openfga.dev/api#/Relationship%20Queries/Check)

> Provide a tuple and ask the OpenFGA API to check for a relationship

```golang
body := openfga.CheckRequest{
	TupleKey: &openfga.TupleKey{
		User: openfga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
		Relation: openfga.PtrString("admin"),
		Object: openfga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
	},
}
data, response, err := apiClient.OpenFgaApi.Check(context.Background()).Body(body).Execute()

// data = {"allowed":true,"resolution":""} // in JSON

fmt.Printf("%t", *data.Allowed) // True

```

#### Write Tuples

[API Documentation](https://openfga.dev/api#/Relationship%20Tuples/Write)

```golang
body := openfga.WriteRequest{
	Writes: &openfga.TupleKeys{
		TupleKeys: []openfga.TupleKey{
			{
				User: openfga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: openfga.PtrString("admin"),
				Object: openfga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
			},
		},
	},
}
_, response, err := apiClient.OpenFgaApi.Write(context.Background()).Body(body).Execute()

```

#### Delete Tuples

[API Documentation](https://openfga.dev/api#/Relationship%20Tuples/Write)

```golang
body := openfga.WriteRequest{
	Deletes: &openfga.TupleKeys{
		TupleKeys: []openfga.TupleKey{
			{
				User: openfga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
				Relation: openfga.PtrString("admin"),
				Object: openfga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
			},
		},
	},
}
_, response, err := apiClient.OpenFgaApi.Write(context.Background()).Body(body).Execute()

```

#### Expand

[API Documentation](https://openfga.dev/api#/Relationship%20Queries/Expand)

```golang
body := openfga.ExpandRequest{
	TupleKey: &openfga.TupleKey{
		Relation: openfga.PtrString("admin"),
		Object: openfga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
	},
}
data, response, err := apiClient.OpenFgaApi.Expand(context.Background()).Body(body).Execute()

// data = {"tree":{"root":{"name":"workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6#admin","leaf":{"users":{"users":["anne","beth"]}}}}} // JSON
```

#### Read Changes

[API Documentation](https://openfga.dev/api#/Relationship%20Tuples/Read)

```golang
// Find if a relationship tuple stating that a certain user is an admin on a certain workspace
body := openfga.ReadRequest{
    TupleKey: &openfga.TupleKey{
        User:     openfga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
        Relation: openfga.PtrString("admin"),
        Object:   openfga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
    },
}

// Find all relationship tuples where a certain user has a relationship as any relation to a certain workspace
body := openfga.ReadRequest{
    TupleKey: &openfga.TupleKey{
        User:     openfga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
        Object:   openfga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
    },
}

// Find all relationship tuples where a certain user is an admin on any workspace
body := openfga.ReadRequest{
    TupleKey: &openfga.TupleKey{
        User:     openfga.PtrString("81684243-9356-4421-8fbf-a4f8d36aa31b"),
        Relation: openfga.PtrString("admin"),
        Object:   openfga.PtrString("workspace:"),
    },
}

// Find all relationship tuples where any user has a relationship as any relation with a particular workspace
body := openfga.ReadRequest{
    TupleKey: &openfga.TupleKey{
        Object:   openfga.PtrString("workspace:675bcac4-ad38-4fb1-a19a-94a5648c91d6"),
    },
}

data, response, err := apiClient.OpenFgaApi.Read(context.Background()).Body(body).Execute()

// In all the above situations, the response will be of the form:
// data = {"tuples":[{"key":{"user":"...","relation":"...","object":"..."},"timestamp":"..."}]} // JSON
```

#### Read Changes (Watch)

[API Documentation](https://openfga.dev/api#/Relationship%20Tuples/ReadChanges)

```golang
data, response, err := apiClient.OpenFgaApi.ReadChanges(context.Background()).
    Type_("workspace").
    PageSize(25).
    ContinuationToken("eyJwayI6IkxBVEVTVF9OU0NPTkZJR19hdXRoMHN0b3JlIiwic2siOiIxem1qbXF3MWZLZExTcUoyN01MdTdqTjh0cWgifQ==").
    Execute()

// response.continuation_token = ...
// response.changes = [
//   { tuple_key: { user, relation, object }, operation: "write", timestamp: ... },
//   { tuple_key: { user, relation, object }, operation: "delete", timestamp: ... }
// ]
```


### API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OpenFgaApi* | [**Check**](docs/OpenFgaApi.md#check) | **Post** /stores/{store_id}/check | Check whether a user is authorized to access an object
*OpenFgaApi* | [**CreateStore**](docs/OpenFgaApi.md#createstore) | **Post** /stores | Create a store
*OpenFgaApi* | [**DeleteStore**](docs/OpenFgaApi.md#deletestore) | **Delete** /stores/{store_id} | Delete a store
*OpenFgaApi* | [**Expand**](docs/OpenFgaApi.md#expand) | **Post** /stores/{store_id}/expand | Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship
*OpenFgaApi* | [**GetStore**](docs/OpenFgaApi.md#getstore) | **Get** /stores/{store_id} | Get a store
*OpenFgaApi* | [**ListStores**](docs/OpenFgaApi.md#liststores) | **Get** /stores | Get all stores
*OpenFgaApi* | [**Read**](docs/OpenFgaApi.md#read) | **Post** /stores/{store_id}/read | Get tuples from the store that matches a query, without following userset rewrite rules
*OpenFgaApi* | [**ReadAssertions**](docs/OpenFgaApi.md#readassertions) | **Get** /stores/{store_id}/assertions/{authorization_model_id} | Read assertions for an authorization model ID
*OpenFgaApi* | [**ReadAuthorizationModel**](docs/OpenFgaApi.md#readauthorizationmodel) | **Get** /stores/{store_id}/authorization-models/{id} | Return a particular version of an authorization model
*OpenFgaApi* | [**ReadAuthorizationModels**](docs/OpenFgaApi.md#readauthorizationmodels) | **Get** /stores/{store_id}/authorization-models | Return all the authorization models for a particular store
*OpenFgaApi* | [**ReadChanges**](docs/OpenFgaApi.md#readchanges) | **Get** /stores/{store_id}/changes | Return a list of all the tuple changes
*OpenFgaApi* | [**Write**](docs/OpenFgaApi.md#write) | **Post** /stores/{store_id}/write | Add or delete tuples from the store
*OpenFgaApi* | [**WriteAssertions**](docs/OpenFgaApi.md#writeassertions) | **Put** /stores/{store_id}/assertions/{authorization_model_id} | Upsert assertions for an authorization model ID
*OpenFgaApi* | [**WriteAuthorizationModel**](docs/OpenFgaApi.md#writeauthorizationmodel) | **Post** /stores/{store_id}/authorization-models | Create a new authorization model


### Models

 - [Any](docs/Any.md)
 - [Assertion](docs/Assertion.md)
 - [AuthorizationModel](docs/AuthorizationModel.md)
 - [CheckRequest](docs/CheckRequest.md)
 - [CheckResponse](docs/CheckResponse.md)
 - [Computed](docs/Computed.md)
 - [ContextualTupleKeys](docs/ContextualTupleKeys.md)
 - [CreateStoreRequest](docs/CreateStoreRequest.md)
 - [CreateStoreResponse](docs/CreateStoreResponse.md)
 - [Difference](docs/Difference.md)
 - [ErrorCode](docs/ErrorCode.md)
 - [ExpandRequest](docs/ExpandRequest.md)
 - [ExpandResponse](docs/ExpandResponse.md)
 - [GetStoreResponse](docs/GetStoreResponse.md)
 - [InternalErrorCode](docs/InternalErrorCode.md)
 - [InternalErrorMessageResponse](docs/InternalErrorMessageResponse.md)
 - [Leaf](docs/Leaf.md)
 - [ListStoresResponse](docs/ListStoresResponse.md)
 - [Node](docs/Node.md)
 - [Nodes](docs/Nodes.md)
 - [NotFoundErrorCode](docs/NotFoundErrorCode.md)
 - [ObjectRelation](docs/ObjectRelation.md)
 - [PathUnknownErrorMessageResponse](docs/PathUnknownErrorMessageResponse.md)
 - [ReadAssertionsResponse](docs/ReadAssertionsResponse.md)
 - [ReadAuthorizationModelResponse](docs/ReadAuthorizationModelResponse.md)
 - [ReadAuthorizationModelsResponse](docs/ReadAuthorizationModelsResponse.md)
 - [ReadChangesResponse](docs/ReadChangesResponse.md)
 - [ReadRequest](docs/ReadRequest.md)
 - [ReadResponse](docs/ReadResponse.md)
 - [Status](docs/Status.md)
 - [Store](docs/Store.md)
 - [Tuple](docs/Tuple.md)
 - [TupleChange](docs/TupleChange.md)
 - [TupleKey](docs/TupleKey.md)
 - [TupleKeys](docs/TupleKeys.md)
 - [TupleOperation](docs/TupleOperation.md)
 - [TupleToUserset](docs/TupleToUserset.md)
 - [TypeDefinition](docs/TypeDefinition.md)
 - [TypeDefinitions](docs/TypeDefinitions.md)
 - [Users](docs/Users.md)
 - [Userset](docs/Userset.md)
 - [UsersetTree](docs/UsersetTree.md)
 - [UsersetTreeDifference](docs/UsersetTreeDifference.md)
 - [UsersetTreeTupleToUserset](docs/UsersetTreeTupleToUserset.md)
 - [Usersets](docs/Usersets.md)
 - [ValidationErrorMessageResponse](docs/ValidationErrorMessageResponse.md)
 - [WriteAssertionsRequest](docs/WriteAssertionsRequest.md)
 - [WriteAuthorizationModelResponse](docs/WriteAuthorizationModelResponse.md)
 - [WriteRequest](docs/WriteRequest.md)


## Contributing

### Issues

If you have found a bug or if you have a feature request, please report them on the [sdk-generator repo](https://github.com/openfga/sdk-generator/issues) issues section. Please do not report security vulnerabilities on the public GitHub issue tracker.

### Pull Requests

All changes made to this repo will be overwritten on the next generation, so we kindly ask that you send all pull requests related to the SDKs to the [sdk-generator repo](https://github.com/openfga/sdk-generator) instead.

## Author

[OpenFGA](https://github.com/openfga)

## License

This project is licensed under the Apache-2.0 license. See the [LICENSE](https://github.com/openfga/go-sdk/blob/main/LICENSE) file for more info.

The code in this repo was auto generated by [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) from a template based on the [go template](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go), licensed under the [Apache License 2.0](https://github.com/OpenAPITools/openapi-generator/blob/master/LICENSE).

This repo bundles some code from the [golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2) package. You can find the code [here](./oauth2) and corresponding [BSD-3 License](./oauth2/LICENSE).
