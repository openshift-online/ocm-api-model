# AGENTS.md

This file provides guidance to AI coding assistants when working with this repository.

## Project Overview

OCM API Model — contains the specification of the OpenShift Cluster Manager API using a custom model definition language. This model is processed by `ocm-api-metamodel` to generate Go SDKs, OpenAPI specs, and documentation.

## Build & Test Commands

```bash
make check           # Validate model definitions
make openapi         # Generate OpenAPI specifications
make clientapi       # Generate Go client API code
make lint            # Run linters
make verify          # Verify generated code is up to date
make clean           # Remove generated files
```

## Architecture

- **model/**: API model definitions organized by service area
  - **model/accounts_mgmt/**: Account management resources
  - **model/clusters_mgmt/**: Cluster management resources
  - **model/addons_mgmt/**: Add-on management resources
  - **model/service_logs/**: Service log resources
- **openapi/**: Generated OpenAPI specifications
- **clientapi/**: Generated Go client code

## Key Conventions

- Model files use a custom DSL processed by the metamodel tool
- Each service area has its own versioned model directory
- API changes require model definition updates, then regeneration
- Keep model files concise and well-documented
