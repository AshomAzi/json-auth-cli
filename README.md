# json-auth-cli
Json-auth-cli is a Go-based CLI for user registration via r (register) and s (signup) commands. It maps Go structs to JSON using the encoding/json package and ensures persistence by appending credentials to tasks.json using os.OpenFile. This tool efficiently bridges terminal input with permanent file storage.

## CLI User Registration
### Objectives Create a command-line interface to register a user and store credentials to a JSON file. 
You will learn to map Gostructs to JSON and perform file persistence. 
### Introduction 
Your tool will accept commands
like register and signup

* Credentials should be stored in tasks.json.
* Each user must have an ID, firstName, lastname, userName, password and date&Time joined.
* Usage: go run . r to register / go run . s to signup
* JSON Marshalling/Unmarshalling.
* CLI flag handling (os.Args or flag package).
