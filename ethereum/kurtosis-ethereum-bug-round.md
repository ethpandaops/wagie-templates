# ethpandaops/kurtosis-ethereum-bug-round

## Purpose

Runs one focused Kurtosis Ethereum bug-hunt probe round against an enclave and returns only new or materially upgraded findings, plus the next gaps worth checking.

## Key Inputs

- `enclave_name`, `goal`, `hunt_focus`
- `config_summary`, `launch_summary`, `status_summary`
- `service_names`, `public_endpoints`
- `previous_findings`, `gaps`

## Key Outputs

- `findings`
- `summary`
- `suggested_gaps`

## Notes

- First round is broad baseline plus anomaly scan.
- Later rounds should prioritize `gaps` and avoid repeating already-established findings.
- Findings are evidence-backed and typed for downstream accumulation.
