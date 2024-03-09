# Link Validation Utility (`luv`)

A command line tool for validating links in free-form text

## Overview

`luv` performs the following steps in order:

1. Resolve a list of **targets** to files
2. Select and apply the appropriate **extractor** per file
3. Parse extracted free-form text for **links**
4. Select and apply the appropriate **validation** per link

## Concepts

### Link

Links are textual references to a resource which may be followed using
some mechanism to reach that resource

### Validation

Validations are methods which must be applied to a given link to check
the link still points to a resource

### Target

Targets are file names or directory names

### Extractor

Extractors take text of a known type and return the plain text elements

For example, a Go file may contain comments and strings, a markdown document
contains paragraphs, table contents, list items etc
