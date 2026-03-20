# validators

A Go input validation library designed for both humans and AI agents. Every validation returns structured, machine-readable error codes alongside human-readable messages — so your UI, your logs, and your AI agent all get exactly what they need from a single call.

```go
result := validators.IBAN("DE0037040044053201300")
if !result.Valid {
    // Human: result.Errors[0].Message → "IBAN length for DE must be 22, got 21"
    // Agent: result.Errors[0].Code    → "iban.length.mismatch"
    //        result.Errors[0].Context → {"country_code": "DE", "expected_length": 22, "actual_length": 21}
}
```

## Install

```sh
go get github.com/laenen-partners/validators
```

## Validators

### Tier 1 — Core

| Validator | Function | Description |
|-----------|----------|-------------|
| Email | `Email(value, checkMX)` | RFC 5321 format, optional MX lookup |
| IBAN | `IBAN(value)` | Country length table, MOD-97 checksum |
| SWIFT/BIC | `SWIFT(value)` | 8 or 11 char bank identifier codes |
| Phone | `Phone(value)` | E.164 format, country-specific length rules (90+ countries) |
| Credit Card | `CreditCard(value)` | Luhn checksum, network detection (Visa, Mastercard, Amex, etc.) |
| VAT | `VAT(value)` | EU VAT numbers for 28 countries + Switzerland |
| URL | `URL(value)` | Scheme + host validation, metadata extraction |
| UUID | `UUID(value, version)` | Format and optional version check (v1–v5) |

### Tier 2 — Standard

| Validator | Function | Description |
|-----------|----------|-------------|
| Postal Code | `PostalCode(value, countryCode)` | Country-specific formats (40+ countries) |
| Country Code | `CountryCode(value)` | ISO 3166-1 alpha-2 and alpha-3 |
| Currency Code | `CurrencyCode(value)` | ISO 4217 three-letter codes |
| Date | `Date(value)` | ISO 8601 (YYYY-MM-DD), calendar correctness including leap years |
| IPv4 | `IPv4(value)` | Dotted-decimal, leading zero rejection |
| IPv6 | `IPv6(value)` | Full, compressed, and mixed formats |
| Domain | `Domain(value)` | RFC 1035 label rules, TLD validation |
| LEI | `LEI(value)` | Legal Entity Identifier, MOD-97 checksum |

### Tier 3 — Extended

| Validator | Function | Description |
|-----------|----------|-------------|
| Semantic Version | `SemVer(value)` | major.minor.patch with optional prerelease and build metadata |
| CIDR | `CIDR(value)` | IP + subnet prefix for IPv4 and IPv6 |
| MAC Address | `MAC(value)` | 48-bit in colon, hyphen, or dot notation |
| ISBN | `ISBN(value)` | ISBN-10 (mod-11) and ISBN-13 (mod-10) checksums |
| CRON | `CRON(value)` | 5-field cron expressions with range/step/list validation |
| JWT | `JWT(value)` | Three-part base64url structure, header/payload JSON check (no signature verification) |
| Hex Color | `HexColor(value)` | #RGB, #RGBA, #RRGGBB, #RRGGBBAA |
| Lat/Lon | `LatLon(lat, lon)` | Decimal degree range validation |
| Belgian National Number | `BelgianNationalNumber(value)` | 11-digit Rijksregisternummer with MOD-97 check |
| Dutch BSN | `DutchBSN(value)` | 9-digit Burgerservicenummer with 11-check |

## Structured Errors

Every validator returns a `Result`:

```go
type Result struct {
    Valid    bool              `json:"valid"`
    Errors   []ValidationError `json:"errors,omitempty"`
    Metadata map[string]any    `json:"metadata,omitempty"`
}

type ValidationError struct {
    Code    string         `json:"code"`              // Stable machine-readable identifier
    Message string         `json:"message"`            // Human-readable explanation
    Field   string         `json:"field,omitempty"`    // Which input field failed
    Context map[string]any `json:"context,omitempty"` // Structured key-value detail
}
```

### Who uses what

| Consumer | Uses | Ignores |
|----------|------|---------|
| Human / UI | `Message` | `Code`, `Context` |
| AI Agent | `Code`, `Context` | `Message` |
| Logging | `Code` for grouping, `Context` for detail | — |
| Tests | `Code` for assertions | `Message` |

### Error codes

Codes follow a `{validator}.{aspect}.{problem}` convention and are defined as exported constants:

```go
validators.ErrIBANChecksumInvalid  // "iban.checksum.invalid"
validators.ErrEmailFormatInvalid   // "email.format.invalid"
validators.ErrCreditCardChecksumInvalid // "creditcard.checksum.invalid"
```

Full list in [`result.go`](result.go).

### Context maps

Context provides the structured data an AI agent (or any programmatic consumer) needs to act on an error without parsing the message:

```go
r := validators.IBAN("DE893704004405320130")
// r.Errors[0].Context:
// {
//   "value":           "DE893704004405320130",
//   "country_code":    "DE",
//   "expected_length": 22,
//   "actual_length":   20
// }
```

### Metadata

On success, `Metadata` carries parsed information extracted during validation:

```go
r := validators.CreditCard("4111111111111111")
// r.Metadata: {"network": "visa", "length": 16}

r = validators.Email("user@example.com", false)
// r.Metadata: {"domain": "example.com"}

r = validators.SWIFT("DEUTDEFF500")
// r.Metadata: {"bank_code": "DEUT", "country_code": "DE", "location": "FF"}
```

## Empty values

All validators treat empty strings as valid. Use your framework's required-field check separately — validation and presence are different concerns.

## Usage examples

### Basic validation

```go
r := validators.Email("user@example.com", false)
if r.Valid {
    fmt.Println("Domain:", r.Metadata["domain"])
}
```

### With MX check

```go
r := validators.Email("user@example.com", true)
if !r.Valid {
    fmt.Println(r.Errors[0].Message)
}
```

### AI agent error handling

```go
r := validators.Phone("+321234")
if !r.Valid {
    err := r.Errors[0]
    switch err.Code {
    case validators.ErrPhoneTooShort:
        // err.Context["digits"] has the actual count
    case validators.ErrPhoneCountryInvalid:
        // err.Context["country_code"], ["expected_min"], ["expected_max"]
    case validators.ErrPhoneFormatInvalid:
        // missing + prefix
    }
}
```

### JSON serialization

The `Result` struct serializes directly to JSON — no adapter needed:

```go
r := validators.IBAN("INVALID")
data, _ := json.Marshal(r)
// {
//   "valid": false,
//   "errors": [{
//     "code": "iban.length.too_short",
//     "message": "IBAN is too short",
//     "field": "iban",
//     "context": {"value": "INVALID", "length": 7}
//   }]
// }
```

## Architecture decisions

See [`docs/adr/`](docs/adr/) for the reasoning behind:

- [ADR-001: Structured validation errors](docs/adr/001-structured-validation-errors.md) — why unified `Result` types with error codes and context maps
- [ADR-002: Validator scope](docs/adr/002-validator-scope.md) — what belongs in this library and why

## License

[MIT](LICENSE)
