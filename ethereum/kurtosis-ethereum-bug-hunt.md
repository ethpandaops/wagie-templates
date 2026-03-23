# ethpandaops/kurtosis-ethereum-bug-hunt

## Purpose

Launches or inspects a Kurtosis Ethereum enclave, then runs an iterative bug-hunt loop that accumulates findings and remaining gaps until the hunt converges or further probing has diminishing value.

## Key Inputs

- `goal`, `enclave_name`, `hunt_focus`
- `constraints`, `example_hint`
- `package_ref`, `config_mode`
- `devnet_name`, `client_type`, `image_hint`, `client_pairs`
- `max_iterations`, `launch`, `auto_cleanup`

## Key Outputs

- `resolved_network_name`, `resolved_network_group`
- `config`
- `effective_client_pairs`, `fallback_pair_added`
- `launch_summary`
- `report`, `findings`, `remaining_gaps`, `summary`
- `recommended_actions`
- `completeness`, `information_gain`
- `converged`, `iterations_used`, `iteration_log`

## Flow

```mermaid
flowchart TD
    A[Inputs] --> B[lifecycle<br/>uses kurtosis-ethereum-network-lifecycle]
    B --> C[hunt loop]

    subgraph C [hunt]
      C1[probe_round<br/>uses kurtosis-ethereum-bug-round]
      C2[accumulate_findings<br/>run accumulate-kurtosis-ethereum-bug-findings]
      C3[assess_progress<br/>uses kurtosis-ethereum-bug-assess]
      C4[log_iteration<br/>run log-kurtosis-ethereum-bug-iteration]
      C1 --> C2 --> C3 --> C4
    end

    C --> D[render_report<br/>run render-kurtosis-ethereum-bug-hunt-report]
    D --> E[Bug-hunt outputs]

    B -. finally if auto_cleanup && launch .-> F[cleanup_enclave<br/>uses kurtosis-enclave-cleanup]
```

## Notes

- `lifecycle` stays outside the loop so enclave setup and session ownership happen once.
- The loop carries typed `findings` and `gaps`, plus a compact `iteration_log`.
- Final report rendering happens once after convergence rather than appending markdown every round.
