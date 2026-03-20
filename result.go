package validators

// Result is the unified return type for all validators.
type Result struct {
	Valid    bool              `json:"valid"`
	Errors   []ValidationError `json:"errors,omitempty"`
	Metadata map[string]any    `json:"metadata,omitempty"`
}

// ValidationError carries a machine-readable code, a human-readable message,
// and structured context so both humans and AI agents can act on failures.
type ValidationError struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Field   string         `json:"field,omitempty"`
	Context map[string]any `json:"context,omitempty"`
}

// Helper to build a valid result with optional metadata.
func valid(metadata map[string]any) Result {
	return Result{Valid: true, Metadata: metadata}
}

// Helper to build an invalid result with a single error.
func invalid(code, message, field string, ctx map[string]any) Result {
	return Result{
		Valid: false,
		Errors: []ValidationError{
			{Code: code, Message: message, Field: field, Context: ctx},
		},
	}
}

// Error codes — email
const (
	ErrEmailFormatInvalid = "email.format.invalid"
	ErrEmailDomainNoMX    = "email.domain.no_mx"
)

// Error codes — IBAN
const (
	ErrIBANTooShort        = "iban.length.too_short"
	ErrIBANInvalidChars    = "iban.characters.invalid"
	ErrIBANCountryInvalid  = "iban.country.invalid"
	ErrIBANLengthMismatch  = "iban.length.mismatch"
	ErrIBANChecksumInvalid = "iban.checksum.invalid"
)

// Error codes — SWIFT/BIC
const (
	ErrSWIFTLengthInvalid = "swift.length.invalid"
	ErrSWIFTFormatInvalid = "swift.format.invalid"
)

// Error codes — phone
const (
	ErrPhoneTooShort      = "phone.length.too_short"
	ErrPhoneTooLong       = "phone.length.too_long"
	ErrPhoneInvalidChars  = "phone.characters.invalid"
	ErrPhoneFormatInvalid = "phone.format.invalid"
	ErrPhoneCountryInvalid = "phone.country.invalid"
)

// Error codes — credit card
const (
	ErrCreditCardTooShort      = "creditcard.length.too_short"
	ErrCreditCardTooLong       = "creditcard.length.too_long"
	ErrCreditCardInvalidChars  = "creditcard.characters.invalid"
	ErrCreditCardChecksumInvalid = "creditcard.checksum.invalid"
)

// Error codes — VAT
const (
	ErrVATTooShort        = "vat.length.too_short"
	ErrVATCountryInvalid  = "vat.country.invalid"
	ErrVATFormatInvalid   = "vat.format.invalid"
)

// Error codes — URL
const (
	ErrURLFormatInvalid  = "url.format.invalid"
	ErrURLSchemeInvalid  = "url.scheme.invalid"
	ErrURLHostMissing    = "url.host.missing"
)

// Error codes — UUID
const (
	ErrUUIDFormatInvalid   = "uuid.format.invalid"
	ErrUUIDVersionInvalid  = "uuid.version.invalid"
)

// Error codes — postal code
const (
	ErrPostalCodeCountryInvalid = "postalcode.country.invalid"
	ErrPostalCodeFormatInvalid  = "postalcode.format.invalid"
)

// Error codes — country code
const (
	ErrCountryCodeFormatInvalid = "countrycode.format.invalid"
	ErrCountryCodeUnknown       = "countrycode.unknown"
)

// Error codes — currency code
const (
	ErrCurrencyCodeFormatInvalid = "currencycode.format.invalid"
	ErrCurrencyCodeUnknown       = "currencycode.unknown"
)

// Error codes — date
const (
	ErrDateFormatInvalid = "date.format.invalid"
	ErrDateInvalid       = "date.value.invalid"
)

// Error codes — IPv4
const (
	ErrIPv4FormatInvalid = "ipv4.format.invalid"
	ErrIPv4OctetInvalid  = "ipv4.octet.invalid"
)

// Error codes — IPv6
const (
	ErrIPv6FormatInvalid = "ipv6.format.invalid"
)

// Error codes — domain
const (
	ErrDomainFormatInvalid = "domain.format.invalid"
	ErrDomainLabelInvalid  = "domain.label.invalid"
	ErrDomainTooLong       = "domain.length.too_long"
)

// Error codes — LEI
const (
	ErrLEILengthInvalid   = "lei.length.invalid"
	ErrLEIFormatInvalid   = "lei.format.invalid"
	ErrLEIChecksumInvalid = "lei.checksum.invalid"
)

// Error codes — semver
const (
	ErrSemVerFormatInvalid = "semver.format.invalid"
)

// Error codes — CIDR
const (
	ErrCIDRFormatInvalid = "cidr.format.invalid"
	ErrCIDRPrefixInvalid = "cidr.prefix.invalid"
)

// Error codes — MAC address
const (
	ErrMACFormatInvalid = "mac.format.invalid"
)

// Error codes — ISBN
const (
	ErrISBNFormatInvalid   = "isbn.format.invalid"
	ErrISBNChecksumInvalid = "isbn.checksum.invalid"
)

// Error codes — CRON
const (
	ErrCRONFormatInvalid = "cron.format.invalid"
	ErrCRONFieldInvalid  = "cron.field.invalid"
)

// Error codes — JWT
const (
	ErrJWTFormatInvalid  = "jwt.format.invalid"
	ErrJWTSegmentInvalid = "jwt.segment.invalid"
)

// Error codes — hex color
const (
	ErrHexColorFormatInvalid = "hexcolor.format.invalid"
)

// Error codes — latitude/longitude
const (
	ErrLatLonFormatInvalid = "latlon.format.invalid"
	ErrLatitudeOutOfRange  = "latlon.latitude.out_of_range"
	ErrLongitudeOutOfRange = "latlon.longitude.out_of_range"
)

// Error codes — Belgian National Number
const (
	ErrBNNLengthInvalid   = "bnn.length.invalid"
	ErrBNNFormatInvalid   = "bnn.format.invalid"
	ErrBNNDateInvalid     = "bnn.date.invalid"
	ErrBNNChecksumInvalid = "bnn.checksum.invalid"
)

// Error codes — Dutch BSN
const (
	ErrBSNLengthInvalid   = "bsn.length.invalid"
	ErrBSNFormatInvalid   = "bsn.format.invalid"
	ErrBSNChecksumInvalid = "bsn.checksum.invalid"
)
