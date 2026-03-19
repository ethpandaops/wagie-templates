# wagie-templates

This repo is the specialized template companion to Wagie core.

Use `@docs/agent-guides/project-templates.md` before editing templates.

## Working Rules

- Keep this repo for non-core templates only. If a template is broadly reusable across domains, it likely belongs back in Wagie core.
- Preserve the shallow family layout at repo root: `ethereum/`, `code/`, `research/`.
- Prefer composition with `uses:` before inventing new task logic.
- Write retrieval-friendly metadata: precise name, dense `Use when ...` description, short retrieval-oriented tags, stable inputs and outputs.
- Tags should be search terms a user would look for, not taxonomy labels. No `type:*`, `flow:*`, or `cap:*` prefixes.
- Match worker capabilities and skills correctly: `'reasoning' in capabilities` for judgment tasks, `'tool-use' in capabilities` for CLI/tool tasks, `'writing' in skills` for content generation, `'coding' in skills` for code editing. Combine when needed.
- Keep contracts explicit and reusable. Avoid undocumented side effects.
- Do not restate injected inputs or output schemas inside prompts; refer to inputs naturally and let declared outputs define the contract.
- Use `make test` to validate. Core templates are loaded from wagie's embedded Go module — no sibling checkout needed.
