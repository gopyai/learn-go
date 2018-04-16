Please notice that antlr4.bat contains:

`java org.antlr.v4.Tool -Dlanguage=Go %*`


To run this example:

1. Within parser directory, initially there is `JSON.g4` file.
2. Run: `antlr4 JSON.g4`. It will generate bunch of files.
3. Run: `cd ..`
4. Run: `go run main.go input.json`