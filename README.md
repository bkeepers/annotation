# Annotation

Convert linter output into annotations for the GitHub [Checks API](https://developer.github.com/v3/checks/runs/#annotations-object)

```sh
$ golint ./...                        
parser/parser.go:68:35: error strings should not be capitalized or end with punctuation or a newline

$ golint ./... | annotation | jsonpp
[
  {
    "path": "parser/parser.go",
    "start_line": 68,
    "end_line": null,
    "start_column": 35,
    "end_column": null,
    "annoation_level": null,
    "message": "error strings should not be capitalized or end with punctuation or a newline",
    "title": null,
    "raw_details": null
  }
]
```
