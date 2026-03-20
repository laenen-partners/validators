# ADR-001: Structured Validation Errors for Human and AI Agent Consumption

- **Status:** Accepted
- **Date:** 2026-03-21
- **Author:** Pascal Laenen

## Context

This library provides input validators (email, IBAN, SWIFT, and more to come) intended for use by both traditional application UIs and AI agents. The current implementation returns validation results as per-validator structs with free-text error strings:

```go
type IBANResult struct {
    Valid       bool
    Error       string
    CountryCode string
}
```

This creates three problems:

1. **AI agents must parse natural language** to understand what went wrong. An agent receiving `"Invalid IBAN length for DE"` has to string-match to extract the country code and infer the error type. This is fragile, ambiguous, and breaks if messages change.
2. **Each validator returns a different type**, forcing consumers to handle N different shapes. Adding a new validator means adding a new result type.
3. **Only the first error is returned**, so consumers fix one issue, re-validate, hit the next — a frustrating loop for both humans and agents.

## Decision

### Unified result type

All validators return the same `Result` type:

```go
type Result struct {
    Valid    bool              `json:"valid"`
    Errors   []ValidationError `json:"errors,omitempty"`
    Metadata map[string]any    `json:"metadata,omitempty"`
}
```

- `Errors` is a slice — validators report all problems found in a single pass.
- `Metadata` carries parsed information (domain, country code, bank code) that was previously scattered across per-validator result structs.

### Structured validation errors

Each error carries a machine-readable code, a human-readable message, and structured context:

```go
type ValidationError struct {
    Code    string         `json:"code"`
    Message string         `json:"message"`
    Field   string         `json:"field,omitempty"`
    Context map[string]any `json:"context,omitempty"`
}
```

| Field     | Purpose                          | Consumer              |
|-----------|----------------------------------|-----------------------|
| `Code`    | Stable, hierarchical identifier  | Agents, tests, logs   |
| `Message` | Human-readable explanation       | UIs, humans           |
| `Field`   | Which input field failed         | Forms, agents         |
| `Context` | Structured key-value detail      | Agents, observability |

### Error code convention

Codes follow a `{validator}.{aspect}.{problem}` pattern and are defined as exported constants:

```
email.format.invalid
email.domain.no_mx
iban.length.too_short
iban.length.mismatch
iban.characters.invalid
iban.country.invalid
iban.checksum.invalid
swift.length.invalid
swift.format.invalid
```

Codes are the machine contract. They are stable across releases. Messages may change freely.

### Context carries the "why" data

Instead of an agent parsing `"IBAN length for DE must be 22, got 19"`, it receives:

```json
{
  "code": "iban.length.mismatch",
  "message": "IBAN length for DE must be 22, got 19",
  "context": {
    "country_code": "DE",
    "expected_length": 22,
    "actual_length": 19
  }
}
```

The agent can programmatically read `expected_length` and `actual_length` without any parsing.

## Consequences

### Positive

- **AI agents get structured data** — no NLP required to understand errors, branch on types, or extract parameters.
- **Humans get readable messages** — the `Message` field remains a plain English sentence.
- **Tests assert on codes, not strings** — tests are decoupled from message wording.
- **Observability improves** — error codes can be used as metric labels and log grouping keys.
- **All errors returned at once** — one validation pass surfaces everything wrong with an input.
- **Uniform API** — every validator returns `Result`, reducing cognitive load and boilerplate for consumers.

### Negative

- **Per-validator result types are removed** — consumers lose compile-time access to fields like `CountryCode`. They access `result.Metadata["country_code"]` instead (type assertion required).
- **`map[string]any` is untyped** — context keys are stringly typed. We accept this tradeoff because the alternative (per-error context structs) would reintroduce the type proliferation problem. Documentation and constants mitigate this.

### Neutral

- **No backwards compatibility** — this is a greenfield library with no downstream consumers yet. A clean break is appropriate.

## Alternatives Considered

### Keep per-validator result types, add error codes

Would solve the code problem but not the "N different result shapes" problem. AI agents would still need per-validator handling logic.

### Implement the `error` interface on `ValidationError`

Rejected. Validation errors are **data**, not control flow. They are expected outcomes, not exceptional conditions. Implementing `error` would encourage `if err != nil` patterns that conflate "the input is invalid" with "something went wrong", and would lose the multi-error capability.

### Proto/gRPC definitions for errors

Premature. There is no service boundary yet. If this library is later exposed via an API, the `Result` struct serializes cleanly to JSON/proto without an adapter.
