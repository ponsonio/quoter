# quoter-test

##Design Considerations
The present implementation use a rule engine library, as this seems to fit in the proposed scenario
where rules correspond to laws, direct business rules and seems to be likely to change over time. As an advantages we can state:

* Business Rules are separated from code and easy to read and understand by an SME
* Easy to maintain and understand that a "code" implementation 


#ToDo
Due time constraints, the following aspect where not implemented, but it does mean those are less on none important;

1. Logging
2. Security
3. Automated test are not exhaustive, but test included shows the implementation proposal for them.
4. Only "British Columbia Mortgage Default Insurance Rates", has been considered 

##Running Locally

Code can be executed using:
````
go run main 
````

To create a request to the back end any http client tool can be use, in this case we use postman:

![img.png](img.png)

##Test

Tests can be run using:

````
go test ./...
````

UI
Install and Run : 
in the "ui" directory : 

To install:
npm install 

To Run: 
npm 