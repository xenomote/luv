# Link Validation Utility (`luv`)

A command line tool for validating links in free-form text

## Overview

`luv` performs the following steps in order:

1. Resolve a list of **targets** to files
2. Select the appropriate **extractor** per file
3. Parse extracted free-form text for **links**
4. Select the appropriate **validation** to apply per link
5. Validate links and collate **results**
6. Report results

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
