<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) [#<issue-number>](LINK) message

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## Unreleased

### Improvements

- (ci) [#5](https://github.com/MalteHerrmann/ginkgo-parser/5) Add linter to CI.

## [v0.1.0](https://github.com/MalteHerrmann/ginkgo-parser/releases/tag/v0.1.0) - 2023-08-24

### Features

- (report) Parses a given JSON report created by the Ginkgo BDD tool
- (md) Creates a markdown file listing the ran tests as bullet points
