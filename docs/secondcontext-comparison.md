# SecondContext concept comparison

The comparison baseline is the public [SecondContext repository](https://github.com/bdobrica/SecondContext) as reviewed on 2026-08-30. SecondContext is a context-augmented assistant prototype; ThinkPixelMEM is vendor-neutral memory infrastructure.

| SecondContext concept | ThinkPixelMEM disposition | Reason |
| --- | --- | --- |
| PostgreSQL canonical structured state | Migrate | Retain, with tenant/space isolation, immutable revisions, provenance, retention, and outbox guarantees. |
| Qdrant dense+sparse retrieval | Migrate with changes | Keep as rebuildable projection; add strict metadata filters, versioned embeddings, deletion generations, and canonical reauthorization. |
| LLM extraction with JSON validation | Migrate with changes | Place behind `MemoryExtractor`/LLMGW; extractor cannot assign authoritative metadata and writes candidates, not canonical truth. |
| Hybrid retrieval and transparent scoring | Migrate | Preserve explainability; add temporal, trust, confidence, dispute, classification, poison, and grant signals. |
| Memories, entities, topics, beliefs/claims | Reshape | Use typed Episodes, temporal Claims/revisions, Relationships, Outcomes, Lessons, and entity refs. Topics remain derived metadata. |
| Person/topic models | Reshape as Profiles | Profiles are schema-defined projections with field-level evidence and sensitive-inference policy. |
| Outcome memories and feedback loops | Migrate | Model as Outcomes/Lessons with verified provenance and no automatic authority. |
| Salience/utility/belief-impact scoring | Experimental | Keep as versioned optional ranking features pending evaluation; never canonical truth. |
| Goal-conditioned scenario generation | Defer | Belongs to an agent/answering layer, not core memory infrastructure. |
| OpenAI-compatible `/v1/responses` and answer generation | Do not migrate | MEM returns ContextPacks and does not act as an assistant. |
| Session/message persistence | Do not migrate | ThinkPixelAR owns canonical Session and execution history. |
| Provider keys/direct OpenAI client | Do not migrate in reference deployment | ThinkPixelLLMGW owns provider access; standalone adapters remain ports. |
| Automatic prompt augmentation | Do not migrate as trusted behavior | AR/harness consumes structured, explicitly untrusted ContextPacks under an AG grant. |
| Debug endpoints | Migrate with changes | Provide authorized inspect/explain APIs with bounded, classified output and audit. |
| Social-role strategies | Experimental/restricted | May become explicit ProfileSchema or ProcedureCandidate; sensitive traits are deny-by-default and procedures require MP qualification. |

No SecondContext data migration is implied by Phase 0. A future importer must preserve original IDs, source trust, conversion provenance, and quarantine state; imported claims cannot become verified observations.

