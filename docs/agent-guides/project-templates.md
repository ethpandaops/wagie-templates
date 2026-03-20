# Wagie Templates Guide

Use this guide when authoring or reviewing YAML templates in `wagie-templates`.

This repo is not the core Wagie template library. It is the home for specialized families that are better loaded as an external companion library.

## Library Boundary

Keep templates here when they are:

- domain-specific, such as Ethereum devnet and Kurtosis workflows
- tool-specific, such as code-review automation
- operator-specific or opinionated enough that they should not shape Wagie core
- tightly coupled to one family's concepts, inputs, or operational surfaces

Keep templates in Wagie core when they are:

- generic atomic capabilities like classify, extract, summarize, transform, evaluate, or answer
- generic orchestration patterns like review loops, consensus, routing, map-reduce, or promotion logic
- reusable building blocks that would improve many unrelated domains

If a template would make sense as a core root-level composable primitive, it probably does not belong here.

## Family Placement

The repo uses shallow root-level families:

- `ethereum/`
- `code/`
- `research/`

Place templates by the domain that owns their meaning, not by incidental tooling.

Examples:

- a devnet-to-Kurtosis repro workflow still belongs in `ethereum/`
- a code-review pipeline belongs in `code/`
- iterative finding accumulation belongs in `research/`

Do not add an extra `templates/` wrapper. Do not create deep category trees unless the repo grows enough to justify them.

## What Good Looks Like

A good template in this repo is:

- valid under the Wagie spec and consumer validation path
- clearly specialized enough to justify living outside core
- easy for Wagie to retrieve and compose
- smaller because of composition, not larger
- explicit about inputs, outputs, and control flow

## Start From Nearby Templates

Before writing a template:

1. Inspect 2-4 nearby templates in the same family.
2. Check whether the workflow really belongs in this repo rather than Wagie core.
3. Prefer updating a nearby template or reusing a building block over creating an overlapping workflow.

Use these categories consistently:

- `type:atomic`: one focused operation
- `type:building-block`: reusable sub-workflow
- `type:pattern`: orchestration shape such as loop, matrix, routing, or verification
- `type:domain`: end-to-end workflow for a business or operational use case

## Compose First

Default to `uses:` when an existing template already solves part of the problem.

Prefer:

- `uses:` for stable reusable blocks
- small typed data flow between tasks
- thin wrappers around well-named building blocks
- depending on Wagie core for generic primitives instead of copying them into this repo

Avoid:

- duplicating a core template here with family-specific naming
- large opaque `run` tasks that hide orchestration semantics
- adding extra review or evaluation steps that the combined library already has

## Write For Retrieval

Retrieval quality depends on more than tags: name, description, tags, inputs, `uses:`, task names, instructions, and outputs all matter.

### Descriptions

- Start with `Use when ...`
- Keep it to one or two tight sentences
- Include:
  - trigger or situation
  - desired outcome
  - important constraint, tool, or workflow shape when material

Good:

`Use when an active Ethereum devnet issue must be investigated and reproduced in Kurtosis with config choices grounded in the source devnet profile.`

Weak:

`Workflow for Ethereum debugging.`

### Tags

Use a controlled, faceted set of tags. Prefer accuracy over volume.

Always cover:

- one `type:*` tag
- any real `flow:*` tags the template itself owns
- one or more `cap:*` tags when capability is clear

Add family or integration tags only when they help retrieval or classification.

Do not:

- pile on synonyms
- invent new tag families without a clear library-level reason
- add `flow:session` unless the template itself owns the session boundary

## Design The Contract Before The Tasks

Define:

1. what the caller provides
2. what the template guarantees on success
3. which outputs downstream templates will consume

Input rules:

- every input declares a `type`
- arrays and objects use explicit `schema`
- add `required`, `default`, `options`, and `description` when they improve predictability
- keep caller-facing inputs simpler than internal task wiring

Output rules:

- every public output should be intentionally consumable
- every `run` task with outputs should define `outputs.schema`
- keep output keys stable and descriptive

## Task Authoring

For `run` tasks:

- use `worker.match` only when routing matters
- add `quality-gate` when malformed or empty output would poison downstream steps
- add `retry` for plausible transient failure
- add `timeout` when runtime should be bounded
- write instructions that constrain the worker to the exact contract you need

### Prompt Hygiene

Wagie already injects task inputs and output expectations into worker context. Do not restate that machinery in the prompt unless a task has a truly unusual requirement.

Prefer:

- referring to fields by their natural names, such as `finding.description` or `repo`
- concise task instructions focused on reasoning, policy, or transformation logic
- relying on declared `outputs` and `outputs.schema` for the response contract

Avoid:

- dumping input values into the prompt with direct `${{ inputs.* }}` interpolation
- sections that only restate already-injected inputs
- repeating output format instructions or JSON shape requirements that are already declared in `outputs.schema`

The schema is the source of truth for outputs. Keep prompts aligned with it, but do not duplicate it.

For composed tasks:

- map only the child inputs that are needed
- re-export only the outputs the parent needs
- preserve established field names when possible

For loops and matrices:

- keep state minimal and typed
- use `fail-fast: false` only when partial progress is useful
- add explicit convergence, threshold, or stopping logic

## Validation

This repo is content-only. Validate changes from a Wagie checkout or another consumer that loads this library.

Check templates in this order:

1. Compare against neighboring templates in the same family.
2. Check that the template still belongs in this repo and has not drifted into a generic core primitive.
3. Validate from the consuming Wagie environment.
4. Run the smallest relevant retrieval, validation, or library-loading tests there.

Run retrieval-focused validation when the change affects descriptions, tags, `uses:`, task structure, or public contracts.

## Keep The Public Contract Small

A template should expose the smallest useful contract.

Prefer:

- clear input names
- stable outputs
- obvious control flow
- explicit descriptions on unusual requirements

Avoid hidden reliance on:

- undocumented side effects
- implicit worker behavior
- brittle output shapes that only make sense to one caller
