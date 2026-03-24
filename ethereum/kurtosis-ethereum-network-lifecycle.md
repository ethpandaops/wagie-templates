# ethpandaops/kurtosis-ethereum-network-lifecycle

## Purpose

Owns the full Kurtosis Ethereum network lifecycle in one session: config design, optional launch, enclave inspection, optional observation, summary, and optional cleanup.

## Key Inputs

- `goal`, `enclave_name`
- `constraints`, `example_hint`
- `package_ref`, `config_mode`
- `observation_window`, `success_criteria`
- `devnet_name`, `client_type`, `image_hint`, `client_pairs`
- `launch`, `auto_cleanup`

## Key Outputs

- `resolved_network_name`, `resolved_network_group`
- `config`, `config_summary`, `inferred_features`, `devnet_assumptions`
- `effective_client_pairs`, `fallback_pair_added`
- `launch_summary`, `run_command`, `args_file_path`
- `status_summary`, `service_names`, `public_endpoints`, `suggested_next_commands`
- `stable`, `validation_summary`, `observation_summary`
- `finalized_epoch`, `head_epoch`, `participation_rate`
- `proposed_blocks`, `missed_proposals`, `missed_attestations`
- `observation_evidence`, `recommended_actions`
- `summary`

## Flow

```mermaid
flowchart TD
    A[Inputs] --> B[design_config<br/>uses kurtosis-ethereum-config]
    B --> C{launch?}
    C -->|true| D[launch_network<br/>run launch-kurtosis-ethereum-package]
    C -->|false| E[inspect_enclave<br/>uses kurtosis-enclave-context]
    D --> E

    E --> F{observe?}
    F -->|observation_window > 0 and launch| G[observe_network<br/>run validate-kurtosis-ethereum-network]
    F -->|skip| H[summarize_result<br/>run summarize-kurtosis-ethereum-network]
    G --> H
    H --> I[Lifecycle outputs]

    H -. finally if auto_cleanup && launch .-> J[cleanup_enclave<br/>uses kurtosis-enclave-cleanup]
```

## Notes

- Inspection always runs, even when launch is skipped, so the template can summarize an existing enclave.
- Observation only runs when launch is enabled and the observation window is non-zero.
