/*
A simple package to fetch AWQL query results. Currently only support querying against Adwords report, support for services will come later.

Usage

Fill the authentication details.

	auth := &awql.Auth{
		ClientId:       "your-client-id.apps.googleusercontent.com",
		ClientSecret:   "your-client-secret",
		RefreshToken:   "your-refresh-token",
		AccessToken:    "your-access-token",
		CustomerId:     "your-customer-id",
		DeveloperToken: "your-developer-token",
	}

Create a new client.

	client := awql.NewClient(auth)

Begin querying. This sample is for querying Adwords report.

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
		// Do anthing with the returned row.
		// Result is automatically mapped to its respective field in the AWQL definition.
		// E.g. to access the previous result you can use row["Date"] and row["ActiveViewCtr"].
	}

*/
package awql
