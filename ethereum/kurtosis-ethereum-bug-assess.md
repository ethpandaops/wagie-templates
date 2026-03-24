# ethpandaops/kurtosis-ethereum-bug-assess

## Purpose

Evaluates accumulated Kurtosis Ethereum bug-hunt findings and decides whether another probe round is worthwhile.

## Key Inputs

- `goal`, `hunt_focus`
- `findings`
- `prior_gaps`, `suggested_gaps`
- `previous_gap_count`, `latest_round_summary`

## Key Outputs

- `completeness`
- `information_gain`
- `gaps`
- `summary`

## Notes

- Merges prior and newly suggested gaps into a small deduplicated next-round set.
- Uses both current evidence quality and recent information gain to determine whether the hunt should continue.
