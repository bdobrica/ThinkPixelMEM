# Architecture and trust boundaries

## System context

```mermaid
flowchart LR
    CLIENT[Users and applications]
    AR[ThinkPixelAR]
    WS[ThinkPixelWS]
    TG[ThinkPixelTG]
    AG[ThinkPixelAG]
    LLMGW[ThinkPixelLLMGW]
    GR[ThinkPixelGR]
    MP[ThinkPixelMP]
    MEM[ThinkPixelMEM API and workers]
    PG[(PostgreSQL canonical state)]
    Q[(Qdrant projection)]

    CLIENT -->|explicit writes and admin API| MEM
    AR -->|Run evidence and ingestion| MEM
    WS -->|generation evidence| MEM
    TG -->|verified outcomes| MEM
    AG -->|MemoryGrant| MEM
    MEM -->|ContextPack| AR
    MEM -->|extract, embed, rerank| LLMGW
    MEM -->|write and retrieval inspection| GR
    MEM -->|procedure candidates| MP
    MEM --> PG
    MEM --> Q
```

## Trust boundaries

```mermaid
flowchart TB
    subgraph HOSTILE[Untrusted content boundary]
        USERS[User and application text]
        SOURCES[Repositories, web, documents, imported memory]
        AGENTS[Agent and model output]
    end

    subgraph TRUSTED[Authenticated control boundary]
        API[MEM API authorization and validation]
        WORKERS[Leased workers]
        POLICY[AG grant or standalone authorizer]
        RISK[Deterministic policy and optional GR]
    end

    subgraph CANON[Canonical data boundary]
        PG[(PostgreSQL)]
        AUDIT[(Audit and outbox)]
    end

    subgraph DERIVED[Disposable projection boundary]
        Q[(Qdrant)]
        PROFILES[Profiles and caches]
    end

    USERS --> API
    SOURCES --> API
    AGENTS --> API
    POLICY --> API
    API --> RISK --> PG
    PG --> AUDIT
    PG --> WORKERS --> DERIVED
    DERIVED -->|candidate IDs only| API
```

Authoritative identity, scope, classification, residency, trust, and evidence metadata cross into MEM only through authenticated adapters. Content remains untrusted at every boundary. Projection filtering is defense in depth; the API reauthorizes and rereads canonical state before disclosure.

