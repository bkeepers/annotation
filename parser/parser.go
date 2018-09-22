package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Annotation is a code annotation that maps to the GitHub annotation object:
// https://developer.github.com/v3/checks/runs/#annotations-object
type Annotation struct {
	// Required. The path of the file to add an annotation to. For example, assets/css/main.css.
	Path *string `json:"path"`

	// Required. The start line of the annotation.
	StartLine *int `json:"start_line"`

	// Required. The end line of the annotation.
	EndLine *int `json:"end_line"`

	// The start column of the annotation.
	StartColumn *int `json:"start_column"`

	// The end column of the annotation.
	EndColumn *int `json:"end_column"`

	// Required. The level of the annotation. Can be one of notice, warning, or failure.
	AnnotationLevel *string `json:"annoation_level"`

	// Required. A short description of the feedback for these lines of code. The maximum size is 64 KB.
	Message *string `json:"message"`

	// The title that represents the annotation. The maximum size is 255 characters.
	Title *string `json:"title"`

	// Details about this annotation. The maximum size is 64 KB.
	RawDetails *string `json:"raw_details"`
}

// Parse reads input from a reader and returns Annotations
func Parse(reader *bufio.Reader) ([]Annotation, error) {
	var output []Annotation

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		annotation, err := ParseLine(string(line))
		if err != nil {
			return output, err
		}

		output = append(output, annotation)
	}

	return output, nil
}

// ParseLine parses the raw input and returns an Annotation struct
func ParseLine(line string) (Annotation, error) {
	parts := strings.Split(line, ":")

	if len(parts) != 4 {
		return Annotation{}, fmt.Errorf("%q is not a valid annotation", line)
	}

	startLine, err := strconv.Atoi(parts[1])
	if err != nil {
		return Annotation{}, err
	}

	startColumn, err := strconv.Atoi(parts[2])
	if err != nil {
		return Annotation{}, err
	}

	message := strings.TrimSpace(parts[3])

	annotation := Annotation{
		Path:        &parts[0],
		StartLine:   &startLine,
		StartColumn: &startColumn,
		Message:     &message,
	}

	return annotation, nil
}
