# Contributing to AegisQ Protocol

## Scope

This repository is **not open for casual contributions**.

Only **authorized contributors and core team members** are allowed to contribute. External pull requests will be ignored unless explicitly requested.

---

## Contribution Principles

All contributions must follow these principles:

* Deterministic behavior is **non-negotiable**
* No breaking of consensus guarantees (safety + liveness)
* Code must be **auditable, testable, and reproducible**
* Avoid unnecessary abstractions — keep systems **explicit and traceable**
* Every change must have a **clear system-level justification**

---

## Workflow

### 1. Issue First

Before making any change:

* Create an issue describing:

  * Problem
  * Root cause
  * Proposed solution
  * Impact on system

No issue → No contribution

---

### 2. Branching

Use structured branch naming:

```
feature/<component>-<description>
fix/<component>-<issue>
refactor/<component>-<reason>
```

Examples:

```
feature/consensus-locking-mechanism
fix/tx-nonce-validation
refactor/hash-determinism
```

---

### 3. Commit Standards

Commits must be **precise and meaningful**:

```
<type>: <component> - <message>
```

Types:

* feat
* fix
* refactor
* test
* docs

Example:

```
fix: consensus - prevent cross-view voting without locking
```

Invalid examples:

```
update code
fix bug
changes
```

---

### 4. Code Requirements

#### General

* No dead code
* No commented-out logic
* No magic numbers
* Explicit error handling required

#### Consensus Layer (Critical)

Any change affecting:

* Voting
* Quorum
* Block validation
* State transitions

Must include:

* Formal reasoning (why it is safe)
* Edge case analysis
* Failure scenario consideration

---

### 5. Testing Requirements

Every contribution must include:

* Unit tests (mandatory)
* Edge case tests
* Failure scenario validation

For consensus-related changes:

* Simulate adversarial scenarios
* Validate deterministic outcomes across runs

---

### 6. Security Review

Mandatory for:

* Cryptographic logic
* Transaction validation
* Consensus flow changes

Checklist:

* Can this break determinism?
* Can this be exploited?
* Does this introduce inconsistency across nodes?

---

### 7. Pull Request Process

PR must include:

* Linked issue
* Clear explanation of change
* Before vs After behavior
* Risk assessment
* Test evidence

PRs without this will be rejected.

---

## Code Style

### Go (Core)

* Follow standard Go conventions
* Keep functions small and deterministic
* Avoid hidden state mutations

### Frontend

* Keep UI purely observational
* No business logic in frontend
* Data must reflect backend truth only

---

## Forbidden Practices

* Non-deterministic hashing
* Implicit state changes
* Skipping validation checks
* Mixing consensus logic with UI/backend shortcuts
* Temporary fixes that remain in production

---

## Security Mindset

This is a **security-critical system**.

Always evaluate:

* Failure modes
* Attack vectors
* Network partition scenarios
* Malicious validator behavior

---

## Review Policy

Maintainers will:

* Reject low-quality contributions
* Challenge assumptions
* Require justification for design decisions

---

## Access to Contribute

To contribute:

* Demonstrate system understanding
* Provide evidence of prior relevant work
* Be prepared for deep technical review

---

## Final Note

AegisQ Protocol is designed to explore **distributed system correctness under adversarial conditions**.

Contributions must improve at least one of the following:

* Correctness
* Observability
* Security

If not, the contribution will be rejected.
