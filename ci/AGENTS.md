# ci/

Templates for fetching CI failures, grouping pre-fingerprinted failures, and publishing structured triage state back to GitHub issues.

These are CI-specific building blocks. They avoid domain ownership rules: adapters, fingerprinting, investigation, and issue rendering live in the family folder that owns the artifact format.

Current building blocks:

- `github-actions-failure-fetch`: read-only GitHub Actions run/job/artifact fetch via `gh`.
- `failure-cluster`: exact grouping of already-fingerprinted normalized failures.
- `ci-investigate-issue-context`: read-only open/recently-closed issue context lookup for semantic reuse before domain fingerprinting.
- `ci-investigate-issue-state-check`: read-only marker/body state lookup for idempotent ci-investigate issues.
- `ci-investigate-worklist`: status-policy mapping from issue state-check results to investigate/resume/recurrence/no-op work items.
- `ci-investigate-issue-upsert-by-fingerprint`: body-only GitHub issue create/update/dedup plumbing for ci-investigate issue bodies.
- `discord-webhook-notify`: best-effort final notification to a Discord incoming webhook after durable issue publication.

## Boundary

Reusable CI templates may assume a small normalized failure shape for grouping, but they must not encode assertoor, Ethereum client, test-harness, or repository-specific root-cause rules.

```yaml
NormalizedFailure:
  job_id: string
  job_name: string
  shard_id: string              # human-readable shard, matrix axis, or job-local id
  test_name: string             # the failing test or check
  failed_step: string           # specific step/task within the test
  category: string              # timeout | assertion | exception | infra-startup | crash | unknown
  error_signature: string       # short canonical line for fingerprinting, ids stripped
  log_excerpts: [string]        # short failure-relevant evidence excerpts
  artifact_paths: [{name, path}]
  prior_ai_summary: string      # optional CI-side hint, may be empty
  prior_ai_confidence: number
  run_id: string
  run_url: string
```

Family adapters can add domain fields, but generic `ci/` templates should only require the stable fields they consume.

## Marker Schema Versioning

Canonical issue bodies carry `<!-- ci-investigate:schema-version=N -->`. State-check reads by version and the issue upsert task writes by version. Bump the version when the marker shape changes. The current version is **1**.

## Worker Requirements

Tasks in this family assume required CLIs are installed on the worker. GitHub-facing templates accept an optional `gh_token` secret and otherwise use ambient `gh` authentication. Where external tools or network calls are required, the task declares `selection.capabilities: [tool-use]`.
