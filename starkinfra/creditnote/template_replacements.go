package creditnote

//	CreditNote.TemplateReplacements type
//
//	The TemplateReplacements type carries the custom values used to fill the placeholders of the contract
//	template of a CreditNote. Its keys are free: they depend on the template.
//
//	Every key, at any level, goes through strcase.ToLowerCamel before it is sent (see the TemplateReplacements field of
//	CreditNote), and keys with nil values are not sent. String values in RFC3339 date format are rewritten as a date
//	(ex: "2026-01-02T00:00:00Z" is sent as "2026-01-02") and numbers go through float64 (precision is lost above 2^53),
//	so send placeholders as strings that do not look like an RFC3339 date or a large number.
//	A map[string]interface{} literal can be assigned to it directly.
//
//	ex: CreditNote.TemplateReplacements{"guarantorName": "Jon Snow", "guarantorTaxId": "012.345.678-90"}

type TemplateReplacements map[string]interface{}
