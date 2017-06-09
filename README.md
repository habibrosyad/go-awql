# go-awql
[![Build Status](https://travis-ci.org/habibrosyad/go-awql.svg?branch=master)](https://travis-ci.org/habibrosyad/go-awql)
[![GoDoc](https://godoc.org/github.com/habibrosyad/go-awql?status.svg)](https://godoc.org/github.com/habibrosyad/go-awql)

A simple package to fetch AWQL query results. Currently only support querying against Adwords report, support for services will come later.

## Usage

Import the package.
```go
import github.com/habibrosyad/go-awql
```

Fill the authentication details.
```go
auth := &awql.Auth{
	ClientId:       "your-client-id.apps.googleusercontent.com",
	ClientSecret:   "your-client-secret",
	RefreshToken:   "your-refresh-token",
	AccessToken:    "your-access-token",
	CustomerId:     "your-customer-id",
	DeveloperToken: "your-developer-token",
}
```

Create a new client.
```go
client := awql.NewClient(auth)
```

Begin querying. This sample is for querying Adwords report.
```go
rows, err := c.Query("SELECT Date, ActiveViewCtr FROM ACCOUNT_PERFORMANCE_REPORT DURING TODAY");
if err != nil {
	panic(err) // Can not continue, do error handling.
}

defer rows.Close()

for rows.Next() {
	row, err := rows.Scan()
	if err != nil {
		// Do error handling.
	}
	// Do anthing with the returned row
}
```

The returned result from `rows.Scan()` is `map[string]interface{}`. To map the result into a struct you might want to use [mapstructure](https://github.com/mitchellh/mapstructure).


