# json-auth-cli
Json-auth-cli is a Go-based CLI for user registration via r (register) and s (signup) commands. It maps Go structs to JSON using the encoding/json package and ensures persistence by appending credentials to tasks.json using os.OpenFile. This tool efficiently bridges terminal input with permanent file storage.
