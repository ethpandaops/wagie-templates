# wagie-templates

## Scope

This repo is the companion library for specialized Wagie templates that do not belong in Wagie's core root-level library.

Current families:

- `ethereum/`: devnet and Kurtosis Ethereum workflows
- `code/`: code-review and code-quality workflows
- `research/`: iterative research workflows

Keep family directories shallow. Do not reintroduce a `templates/` wrapper here.

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
