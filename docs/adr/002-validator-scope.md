# ADR-002: Standard Validator Scope

- **Status:** Accepted
- **Date:** 2026-03-21
- **Author:** Pascal Laenen

## Context

We need to decide which validators belong in this library. The goal is a standard toolset that covers the most common input validation needs for business applications and SaaS products, without becoming a kitchen-sink utility.

### Inclusion criteria

A validator belongs here if:

1. **It has real logic** — checksums, format rules, cross-field checks, or lookup tables. Pure "is non-empty" or "is a string" checks do not qualify.
2. **It is domain-general** — useful across most business applications, not specific to one industry.
3. **It is self-contained** — no external API calls required for the core validation (optional enrichment like MX checks are acceptable).

## Decision

### Tier 1 — Core (implement first)

Already present or highest value for typical applications.

| Validator | Validates | Key logic |
|-----------|-----------|-----------|
| **Email** | Email address format + optional MX | RFC 5321 format, regex, optional DNS |
| **IBAN** | International Bank Account Number | Country length table, MOD-97 checksum |
| **SWIFT/BIC** | Bank identifier code | 8 or 11 char format, structure parsing |
| **Phone (E.164)** | International phone numbers | Country code, length per country, digit-only after + |
| **Credit Card (Luhn)** | Payment card numbers | Luhn checksum, network detection by IIN prefix |
| **VAT Number (EU)** | European VAT identification numbers | Per-country format rules and check digit algorithms |
| **URL** | Well-formed URLs | Scheme, host, port, path per RFC 3986 |
| **UUID** | Universally unique identifiers | Format and optional version validation per RFC 4122 |

### Tier 2 — Standard (implement second)

Commonly needed once an application handles international users, addresses, or financial data.

| Validator | Validates | Key logic |
|-----------|-----------|-----------|
| **Postal Code** | Postal/ZIP codes per country | Country-specific format patterns (US 5+4, UK alpha, BE 4-digit, DE 5-digit) |
| **Country Code (ISO 3166-1)** | 2-letter or 3-letter country codes | Lookup against ISO 3166-1 alpha-2/alpha-3 |
| **Currency Code (ISO 4217)** | 3-letter currency codes | Lookup against ISO 4217 |
| **Date (ISO 8601)** | Date strings | Calendar validity (leap years, month lengths), YYYY-MM-DD |
| **IPv4** | Dotted-decimal IP addresses | Octet range 0-255, no leading zeros |
| **IPv6** | IPv6 addresses | Full, compressed (::), and mixed formats per RFC 4291 |
| **Domain Name** | Domain labels and structure | Label length, total length, TLD presence, RFC 1035 |
| **LEI** | Legal Entity Identifier | 20-char alphanumeric, MOD-97 check digits |

### Tier 3 — Extended (implement as needed)

Useful but more situational.

| Validator | Validates | Key logic |
|-----------|-----------|-----------|
| **Semantic Version** | SemVer strings | major.minor.patch + pre-release/build metadata |
| **CIDR** | IP + subnet prefix | Valid IP + prefix length in range |
| **MAC Address** | Hardware addresses | 48-bit in colon/hyphen/dot notation |
| **ISBN** | Book numbers | ISBN-10 (mod-11) and ISBN-13 (mod-10) checksums |
| **CRON Expression** | Cron schedules | 5-6 field validation with valid ranges |
| **JWT (structure)** | JSON Web Token format | 3-part base64url structure, no signature verification |
| **Hex Color** | CSS hex color codes | 3, 4, 6, or 8 hex chars with # prefix |
| **Latitude/Longitude** | Geographic coordinates | Range validation (-90/90, -180/180) |
| **Belgian National Number** | Belgian Rijksregisternummer | 11-digit, birth date encoding, MOD-97 check |
| **Dutch BSN** | Dutch Burgerservicenummer | 9-digit, modulo 11 check |

### Explicitly excluded

| Category | Why excluded |
|----------|-------------|
| **Password strength** | Policy-specific, not format validation. Strength rules vary per product. |
| **Name validation** | Beyond "contains valid Unicode letters" there are no universal rules. Names are too culturally diverse to validate meaningfully. |
| **Address validation** | Requires external APIs (Google Maps, postal services). Not self-contained. |
| **SSN (US)** | Too jurisdiction-specific for a general library. National ID validators (Belgian, Dutch) are included as Tier 3 because of the project's European context. |
| **Regex validation** | `regexp.Compile` already does this. No added value wrapping it. |

## Consequences

- Tier 1 sets the initial release scope. Each validator follows the structured error pattern from ADR-001.
- Tier 2 and 3 are implemented incrementally based on actual need, not speculatively.
- New validators must meet the three inclusion criteria. This prevents scope creep.
- Each validator lives in its own file (`phone.go`, `creditcard.go`, etc.) with a corresponding test file.
