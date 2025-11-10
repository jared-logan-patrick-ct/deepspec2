# Code Specification Schemas

This directory contains JSON Schema definitions for the universal code specification format used by deepspec.

## Schema Files

### `code-spec.schema.json`
Main schema defining the structure of code specifications. This schema is language-agnostic but includes language-specific enums where needed.

**Key Features:**
- **Simple & Practical**: Validates existing spec files with minimal changes
- **Language Support**: Extensible to multiple programming languages
- **Flexible Body Representation**: Supports both detailed `body_statements` and simple `body_summary`
- **Type System**: Covers structs, classes, interfaces, enums, unions, and type aliases
- **Metadata**: Tracks code metrics (LOC, complexity, etc.)

## Schema Version

Current version: **1.0.0**

## Usage

### Validation
Use any JSON Schema validator to validate specification files:

```bash
# Using ajv-cli
npx ajv-cli validate -s schemas/code-spec.schema.json -d "specs/*.json"

# Using check-jsonschema
pip install check-jsonschema
check-jsonschema --schemafile schemas/code-spec.schema.json specs/*.json
```

### Integration
Reference the schema in your spec files:

```json
{
  "$schema": "../schemas/code-spec.schema.json",
  "spec_version": "1.0.0",
  "language": "go",
  ...
}
```

## Schema Structure

### Required Fields
- `spec_version`: Must be "1.0.0"
- `language`: Programming language (go, python, javascript, etc.)
- `module`: Module/package/namespace name
- `file`: Source file name

### Optional Fields
- `language_version`: Minimum language version required
- `description`: High-level file description
- `imports`: External dependencies
- `types`: Type definitions
- `functions`: Functions and methods
- `variables`: Module-level variables
- `constants`: Module-level constants
- `external_references`: Cross-file dependencies
- `metadata`: Code metrics and notes

## Design Principles

1. **KISS (Keep It Simple, Stupid)**: Schema validates existing specs with minimal modifications
2. **Backward Compatible**: All current spec files conform with minor additions
3. **Extensible**: Easy to add new languages and statement types
4. **Practical**: Focuses on what's needed for code generation, not theoretical completeness
5. **Self-Documenting**: Extensive descriptions in schema definitions

## Visibility Modifiers

The schema supports both universal and language-specific visibility:
- **Universal**: `public`, `private`, `protected`, `internal`
- **Go-specific**: `exported`, `unexported`

This allows specifications to be language-agnostic while preserving original semantics.

## Statement Types

Supported statement types:
- `assignment`: Variable assignment
- `declaration`: Type or variable declaration
- `variable_declaration`: Variable declaration with initialization
- `method_call`: Method invocation
- `function_call`: Function invocation
- `return`: Return statement
- `if_statement`: Conditional with true/false branches
- `for_loop`: Loop construct
- `switch_statement`: Switch/case with multiple branches
- `type_assertion`: Type assertion/cast
- `append`: Array/slice append operation
- `function_body`: Inline function/closure body

## Future Enhancements

Potential additions without breaking current specs:
- **Intent Layer**: Semantic purpose annotations
- **Pattern Tags**: Design pattern identifiers
- **Type Mappings**: Cross-language type equivalents
- **Dependency Graph**: Structural relationships
- **Execution Context**: Runtime requirements

These can be added as optional fields without invalidating existing specifications.
