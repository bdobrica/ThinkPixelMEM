# ADR-0004: Temporal claims and immutable revisions

- Status: accepted
- Date: 2026-08-30
- Deciders: ThinkPixelMEM maintainers
- Supersedes: none

## Context

Facts change, observations arrive late, and corrections must remain auditable.

## Decision

A Claim is a stable logical identity with immutable, monotonically numbered revisions. Exactly one completed revision may be active. Corrections append a revision with reason, actor, evidence, prior revision, and integrity digest. Completed revisions cannot be updated.

Valid time (`valid_from`, `valid_until`) describes when a claim applies. `observed_at` describes when evidence observed it; `recorded_at` is the server commit time; `superseded_at` records the system transition. Intervals are half-open `[valid_from, valid_until)`. Null bounds mean unbounded. `valid_until` must exceed `valid_from`.

Statuses are `active`, `disputed`, `superseded`, `withdrawn`, `quarantined`, `expired`, and `deleted`. Contradictory credible evidence produces `disputed` unless deterministic policy establishes temporal succession. Supersession closes the old interval and links both claims.

## Consequences

Current and historical queries are distinct. Corrections never overwrite prior evidence. Privacy deletion may cryptographically erase or delete revision content while retaining only a policy-permitted non-sensitive deletion marker.

## Alternatives considered

Mutable rows and last-write-wins were rejected because they erase history and conceal contradictions.

## Verification

Database constraints, concurrency tests, interval property tests, and correction tests enforce these rules.

