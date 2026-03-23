# ethpandaops/kurtosis-ethereum-reference

## Purpose

Fetches narrowly scoped ethereum-package reference material for Kurtosis workflows.

## Key Inputs

- `section`
- `query`
- `package_ref`

## Key Outputs

- `reference`
- `reference_summary`
- `extracted_items`

## Flow

```mermaid
flowchart TD
    A[Inputs] --> B[fetch_reference<br/>run fetch-kurtosis-ethereum-reference]
    B --> C[Reference outputs]
```

## Notes

- `query` only matters for sections that require a search term or example name.
- The `reference` artifact is now explicitly mapped through the task output contract.
