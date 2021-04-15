# Our Source

## Problem Statement

### Ideal

Ideally we have a wholistic understanding of our source code across our repositories so we can make sensible default decisions and be aware of deviations based on data not rhetoric.

### Reality

In reality we don't have a convenient way of aggregating repository statistics across our repositories and rely on rhetoric or manual analysis.

### Consequences

* analyse our repositories periodically to understand the technology landscape
* inability or lack of confidence when making sensible default decisions
* lack of visibility when there is drift from our sensible defaults

### Proposal

Write a service that will crawl, generate statistics for and aggregate the statistics of GitHub repositories.

## Decisions

We are using [adr-tools](https://github.com/npryce/adr-tools) to capture project decisions.

[Decision Register](documentation/decisions)
