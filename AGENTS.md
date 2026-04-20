# wagie-templates

## Scope

This repo is the companion library for specialized Wagie templates that do not belong in Wagie's core root-level library.

Current families:

- `ethereum/`: devnet and Kurtosis Ethereum workflows
- `code/`: code-review and code-quality workflows
- `research/`: iterative research workflows
- `experiments/`: autonomous experiment loops (metric-driven hillclimb over any git-tracked artifact)

Keep family directories shallow. Do not reintroduce a `templates/` wrapper here.

## Entry Points vs Building Blocks

Within a family, some templates are user-facing entry points and others are internal building blocks composed by those entry points. Current roles:

- `ethereum/`: entry points are `devnet-debug`, `kurtosis-ethereum-bug-hunt`, `kurtosis-ethereum-network-lifecycle`, `kurtosis-ethereum-config`. The rest (`devnet-context`, `devnet-baseline`, `devnet-finality-assessment`, `devnet-notes`, `devnet-topology-profile`, `devnet-report`, `kurtosis-ethereum-devnet-config`, `kurtosis-ethereum-reference`, `kurtosis-enclave-*`) are building blocks, most also callable standalone.
- `code/`: entry points are `code-review-committee`, `code-review-fix-loop`, `code-review-adversarial`, `code-verification`. `code-diff` and `code-reviewer` are building blocks.
- `research/`: entry point is `deep-research`. `research-plan`, `research-findings`, `research-coverage-assess`, `research-verify` are building blocks, callable standalone for one-shot use.
- `experiments/`: entry point is `experiment-loop`. The setup task inside the template handles discovery (resolving target files, setup/benchmark/correctness commands, metric name from the repo + goal) so callers only need to supply `repo_url`, `goal`, and `result_branch` in the common case. All discovered parameters can be overridden explicitly.

When authoring new templates, flag the role in the description if it is not a top-level entry point.

## Boundary

Templates in this repo should be specialized, domain-coupled, or operator-specific.

Keep generic core templates in the main Wagie repo, especially:

- atomic cognitive primitives: decide, evaluate, extract, transform, summarize, generate
- orchestration patterns: map-reduce, review-loop, self-consistency-jury
- structural glue: promote-reject, evaluation-aggregate

If a template can stand as a root-level composable primitive for many unrelated workflows, it probably belongs in core, not here.

## Repo Rules

- Prefer targeted validation while iterating, then run the smallest meaningful broader check before finishing.
- Keep instruction files concise and scoped; put detailed topic guidance in referenced docs instead of growing this file.
- Prefer the smallest family that owns the workflow. Avoid dumping unrelated templates into a catch-all bucket.
- Cross-family dependencies should be rare and intentional. Prefer depending on core Wagie templates over coupling families together without a strong reason.

## Core Commands

```bash
make test
make tidy
```

`make test` validates domain templates against core templates from wagie's embedded Go module.

## Template Work

When editing `**/*.yaml`, also read:

- `docs/agent-guides/project-templates.md`

Template authoring is retrieval-sensitive. Names, descriptions, tags, inputs, `uses:`, task names, and output contracts all affect how Wagie finds and composes templates.
