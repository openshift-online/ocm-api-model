# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the OpenShift Cluster Manager (OCM) API model repository, which contains the specification of the OCM API using a custom DSL (Domain Specific Language). The project is structured around three main submodules:

1. **`model/`** - Contains API definitions in `.model` files using a Go-like DSL
2. **`clientapi/`** - Contains generated Go types and serialization code for the OCM SDK
3. **`metamodel_generator/`** - Contains tooling to generate code from the model definitions

## Architecture

### Model Language
The API specification is written using a custom DSL similar to Go, with `.model` files located in service-specific directories:
- `model/clusters_mgmt/v1/` - Cluster management service definitions
- `model/accounts_mgmt/v1/` - Account management service definitions  
- `model/addons_mgmt/v1/` - Add-ons management service definitions
- `model/aro_hcp/v1alpha1/` - ARO HCP specific definitions
- And other service directories

### Generated Code Structure
- **Classes and Structs**: Represent data types (classes have built-in ID/HREF, structs don't)
- **Resources**: Define API endpoints and collections with methods and locators
- **Enums**: Define enumerated values
- **JSON Serialization**: Automatically generated with snake_case conversion

### Class References
The model supports cross-service references using `@ref` annotations:
```
@ref(path = "/clusters_mgmt/v1/cluster")
class Cluster {
}
```

## Common Development Commands

### Core Workflow
```bash
# Update generated code (clientapi and openapi) from model changes
make update

# Verify that generated code matches the model
make verify

# Check model syntax and structure
make check

# Clean all generated files
make clean
```

### Generation Commands
```bash
# Generate only Go client API code
make clientapi

# Generate only OpenAPI specifications
make openapi

# Build the metamodel generator binary
make metamodel
```

### Linting
```bash
# Check model files for proper indentation (tabs only)
make lint
```

## Typical Development Workflow

1. **Modify API Model**: Edit `.model` files in the appropriate service directory under `model/`
2. **Generate Code**: Run `make update` to regenerate `clientapi/` and `openapi/` from model changes
3. **Verify Changes**: Run `make verify` to ensure generated code matches expectations
4. **Check Syntax**: Run `make check` to validate model syntax

## Key Files and Directories

- `Makefile` - Contains all build and generation commands
- `model/` - Source of truth for API definitions
- `clientapi/` - Generated Go types and builders (submodule)
- `metamodel_generator/` - Code generation tooling (submodule)
- `openapi/` - Generated OpenAPI specifications
- `CHANGES.md` - Change log for API versions

## Model File Conventions

- Use CamelCase for attribute names in model files
- JSON representation automatically converts to snake_case
- Classes represent objects with identity (automatic ID/HREF attributes)
- Structs represent value objects without identity
- Resources define API collections and endpoints
- Methods define HTTP operations (List, Get, Add, Delete)
- Locators define relationships between resources

## Dependencies

The project requires:
- Go 1.23+ (see `metamodel_generator/go.mod`)
- `goimports` for Go code formatting
- `shasum` or `sha256sum` for verification

## Important Notes

- The `clientapi/` directory is generated - never edit files there directly
- The `openapi/` directory is generated - never edit files there directly  
- Always run `make verify` after making changes to ensure consistency
- The metamodel generator creates both Go types and OpenAPI specifications from the same model source