# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to the following versioning pattern:

Given a version number MAJOR.MINOR.PATCH, increment:

- MAJOR version when the **API** version is incremented. This may include backwards incompatible changes;
- MINOR version when **breaking changes** are introduced OR **new functionalities** are added in a backwards compatible manner;
- PATCH version when backwards compatible bug **fixes** are implemented.


## [Unreleased]
### Added
- IssuingEmbossingKit resource
- payerId and endToEndId parameter to BrcodePreview resource
- cashierBankCode and description parameter to StaticBrcode resource
### Changed
- nominalAmount and amount parameter to conditionally required to CreditNote resource
### Removed
- cardDesignId and envelopeDesignId attributes to IssuingEmbossingRequest resource

## [0.0.1] - 2023-03-27
### Added
- Full Stark Infra API v2 compatibility
