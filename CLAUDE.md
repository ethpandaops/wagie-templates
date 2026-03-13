# wagie-templates

This repo is the specialized template companion to Wagie core.

Use `@docs/agent-guides/project-templates.md` before editing templates.

## Working Rules

- Keep this repo for non-core templates only. If a template is broadly reusable across domains, it likely belongs back in Wagie core.
- Preserve the shallow family layout at repo root: `ethereum/`, `code/`, `research/`.
- Prefer composition with `uses:` before inventing new task logic.
- Write retrieval-friendly metadata: precise name, dense `Use when ...` description, controlled tags, stable inputs and outputs.
- Keep contracts explicit and reusable. Avoid undocumented side effects.
- Do not restate injected inputs or output schemas inside prompts; refer to inputs naturally and let declared outputs define the contract.
- Validate from a Wagie checkout or another consumer that loads this repo. This repo is content-only and does not have its own standalone harness.
- Use `make test` for the standard local validation path. It loads core templates from `../wagie` unless `WAGIE_CORE_DIR` is set.
