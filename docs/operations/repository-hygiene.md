# Repository hygiene

ThinkPixelMEM source control must not become a persistence path for runtime memory or credentials. The root `.gitignore` excludes common local forms of test-memory exports, environment and key files, Qdrant storage and snapshots, PostgreSQL data directories, SQLite databases, caches, and related sidecar files.

Run the tracked-file enforcement gate with:

```sh
make hygiene-check
```

The gate inspects paths returned by `git ls-files`, so it also rejects forbidden artifacts that were force-added despite ignore rules. It rejects test-memory datasets, secret-bearing file types, Qdrant dumps, and local PostgreSQL or SQLite artifacts. Text files are additionally inspected for private-key headers and recognizable OpenAI, Anthropic, Google, AWS, Hugging Face, and GitHub credential formats. It reports only the path and finding category; credential values are never printed.

`.env.example` and `.env.template` may be tracked for names and safe placeholders. They are still subject to content inspection. Tests that exercise redaction should use unmistakably synthetic canaries assembled in code rather than realistic credential strings.

This check is part of `make verify` and therefore runs in CI. It is a focused guardrail, not a replacement for provider-side secret scanning, credential rotation after exposure, or review of data provenance. If a secret is ever committed, revoke it immediately and follow the repository's incident process; removing it in a later commit does not remove it from Git history.
