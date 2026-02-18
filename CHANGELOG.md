# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-02-18

### Added
- Initial public release
- Core interfaces: DB, Logger, Metrics
- Core domain types: Alert, Issue, MoveMapping, ChannelProcessingState
- Alert severity levels: panic, error, warning, resolved, info
- Comprehensive alert validation and cleaning
- Webhook support with interactive buttons, input forms, and access control
- Escalation system for unresolved issues
- WebhookCallback for handling webhook responses
- Field specifications for additional key-value pairs
- Testing utilities: dbtests package for DB implementation testing
- No-op implementations: NoopLogger, NoopMetrics
- InMemoryFifoQueue for testing purposes
- Extensive validation constants and limits
- Package documentation and examples

### Security
- All fields are validated and sanitized to prevent injection attacks
- Maximum length limits enforced on all string fields
- URL validation for webhooks and links
- Slack mention and channel ID validation

[Unreleased]: https://github.com/peteraglen/slack-manager-common/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/peteraglen/slack-manager-common/releases/tag/v0.1.0
