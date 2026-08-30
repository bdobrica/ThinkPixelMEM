# ADR-0003: Trust dimensions and authoritative metadata

- Status: accepted
- Date: 2026-08-30
- Deciders: ThinkPixelMEM maintainers
- Supersedes: none

## Context

Extraction confidence is not evidence trust. Models can confidently infer false content and can be induced to forge scope or provenance.

## Decision

`confidence` measures support for derived content on `[0,1]`; it never describes the source. `source_trust` is infrastructure-assigned from the authenticated source vocabulary. `source_kind` distinguishes observation from inference. These dimensions are stored and scored independently.

Tenant, principal, MemorySpace, source/evidence identity, classification, residency, trusted timestamps, and source trust are authoritative metadata. They come from authenticated infrastructure and cannot be accepted from extractor output. Extractors may propose normalized content, entities, topics, relationships, confidence, importance, and poison-risk signals.

## Consequences

Extractor schemas have no writable authoritative fields. Candidate validation joins authoritative metadata from the ingestion envelope and rejects conflicts. A trusted source may still contain hostile content; source trust does not imply instruction trust.

## Alternatives considered

A single confidence score and model-selected provenance were rejected because neither establishes origin or authority.

## Verification

Schema validation and integration tests must reject forged tenant, space, classification, evidence, and `verified-tool-output` claims.

