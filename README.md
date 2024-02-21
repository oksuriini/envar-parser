# Envar Parser

This golang tool parses a file for environmental variables.

Envars need to be in format "ENVAR_KEY=ENVAR_VALUE" or "ENVAR_KEY = ENVAR_VALUE"

The program splits the key and value by looking for equals "="-mark, and assumes key is always on the left side of the pair.

If there is an error in parsing the file, the program prints an error message.

## Module info

Module has two functions "GetOneEnvar" and "SetAllEnvars", and additional private helper functions.

### GetOneEnvar

GetOneEnvar takes in two arguments "envName" and "fileName".

"envName" is the name of the environmental variable that is desired from the environmental variable file.

"fileName" is the name of the file in root directory, which holds all of the environmental variables.

Function returns the value of environmental variable.

If function fails to parse the file or doesn't find the desired environmental variable, then it returns empty string and error.

### SetAllEnvars

SetAllEnvars takes in "fileName" as the only argument.

"fileName" is the name of the file in root directory, which holds all of the environmental variables.

Function sets all of the found environmental variables and returns nil on successful run.

If function fails to parse file or fails to set an environmental variable, then it returns an error.

## Usage

Create a file to project's root directory where you store your envars, eg. "dotenvar", ".env", "envarfile.txt", "whateverfloatsyourboat" etc.

Then import this repo to your project by running:

`go get github.com/oksuriini/envarParser`

Import the module to your project:

`import envfile "github.com/oksuriini/envarParser"`

After which you can use the module in your project:

`func main() {
env, err := envfile.GetOneEnvar("EXAMPLE_KEY", "envarfile")
if err != nil {
fmt.Printf("error %v encountered",err)
return
}
fmt.Println(env) // --> Prints "EXAMPLE_VALUE"
}`
