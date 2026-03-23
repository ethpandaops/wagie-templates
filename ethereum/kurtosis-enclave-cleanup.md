# ethpandaops/kurtosis-enclave-cleanup

## Purpose

Removes a Kurtosis enclave and returns a minimal cleanup result.

## Key Inputs

- `enclave_name`
- `force`

## Key Outputs

- `removed`
- `summary`

## Flow

```mermaid
flowchart TD
    A[Inputs] --> B[cleanup_enclave<br/>run cleanup-kurtosis-enclave]
    B --> C[Cleanup outputs]
```

## Notes

- This is a thin operational wrapper with no downstream branching.
