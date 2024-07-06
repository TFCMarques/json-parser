## JSON Parser:

This challenge is a JSON parser written in Go, designed to validate and parse JSON files or directories containing JSON files. It showcases the use of Go's standard library for file system operations, parsing techniques, and handling command-line arguments and was a practical way to get some hands-on experience with Go.

### How to Run

To run this parser, ensure you have Go installed on your machine. Then build the project by running:

```sh
go build json-parser .
```

Follow these steps to validate JSON files or directories:

- To parse and validate a single JSON file:
  ```sh
  go run json-parser -r <filename>
  ```
- To parse and validate all JSON files in a directory:
  ```sh
  go run json-parser -d <directory>
  ```

The parser will read the specified JSON file(s) and output whether each file contains valid JSON. If a directory is specified, it will recursively parse all JSON files within that directory.

For more information on the challenge, visit [codingchallenges.fyi/challenges/challenge-json-parser](https://codingchallenges.fyi/challenges/challenge-json-parser)
