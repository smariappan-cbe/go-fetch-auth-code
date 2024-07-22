# Getting an auth code

This golang code sample demonstrates how to programmatically fetch an auth code using the sandbox app credentials obtained from recipient hub. It uses the gin web framework and templates to render a rudimentary HTML page with a button to start the consent flow process.

## key concepts
* Making a GET request for an auth code from the identity provider
* Redirecting to an example test bank "Mikomo" to simulate "consent flow" permissioning of account access
* Redirecting to recipient hub's auth code viewer page to receive the auth code after consenting to access

## overview of project
The main.go module includes a single route "/" which drops the user on a basic HTML page rendered from the content in ./templates/index.tmml. The page includes an anchor link that points to a sandbox identity provider URL. The value of your ClientId variable is used in this URL, so be sure it is properly set to your client id.

The go.mod file includes the application's golang dependencies, and the run.sh file is a shell script that consolidates the dependency fetching and application compilation, and binary execution into a single step for convenience.


## running the sample
1. Initialize the ClientId variable in main.go with the client id of your sandbox app
2. Open a terminal in the sandbox
3. Run the command "chmod +x run.sh"
4. Run the command "./run.sh"
5. A dilaog will show in code sandbox indicating "Port 8123 has been opened" with buttons to "Open Preview" and "Open Externally". Click "Open Externally" to open the application in the browser.
6. Follow the instructions in the browser once the application is opened.