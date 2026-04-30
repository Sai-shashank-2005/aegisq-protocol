# Security Policy

## Scope

AegisQ Protocol is a **security-critical distributed system**.

Security includes, but is not limited to:

* Consensus correctness (safety + liveness)
* Transaction integrity
* Cryptographic soundness
* Deterministic execution
* State consistency across nodes

Any issue affecting these is considered a **security vulnerability**.

---

## Supported Versions

Security updates are provided only for actively maintained versions.

| Version         | Supported       |
| --------------- | --------------- |
| Latest (main)   | Supported       |
| Previous stable | Limited support |
| Older versions  | Not supported   |

Running unsupported versions is strongly discouraged.

---

## Vulnerability Classification

### Critical

* Consensus break (forking, double-finalization, invalid finality)
* State divergence across nodes
* Signature verification flaws (Dilithium / hashing)
* Determinism violations

### High

* Replay attacks (e.g., missing nonce)
* Transaction ordering inconsistencies
* Leader manipulation / predictability exploitation
* Quorum manipulation

### Medium

* API inconsistencies affecting validation
* Resource exhaustion affecting availability
* Observability inconsistencies

### Low

* Minor bugs without impact on correctness or security

---

## Reporting a Vulnerability

### Responsible Disclosure

Do not publicly disclose vulnerabilities before reporting.

Submit reports with the following details:

* Description of the issue
* Affected component(s)
* Steps to reproduce
* Impact analysis
* Suggested mitigation (optional)

Incomplete reports may be ignored.

---

## Reporting Channels

Send reports to:

* Email: [security@aegisq.local](mailto:security@aegisq.local)
* Or open a **private security issue** (if enabled)

Do not create public issues for security vulnerabilities.

---

## Response Process

* Initial acknowledgment: within 48 hours
* Triage and validation: within 3–7 days
* Resolution timeline depends on severity

You will be informed of:

* Acceptance or rejection
* Severity classification
* Planned remediation approach

---

## Handling of Vulnerabilities

Once confirmed:

* Issue is reproduced and verified
* Root cause is analyzed
* Fix is developed and tested
* Patch is released
* Disclosure is made (if appropriate)

Critical issues may result in immediate patch release without prior notice.

---

## Security Expectations

Contributors must:

* Preserve determinism
* Avoid introducing undefined behavior
* Validate all inputs strictly
* Ensure cryptographic correctness

Changes affecting consensus or cryptography require **explicit security reasoning**.

---

## Non-Goals

The following are **not considered vulnerabilities**:

* Theoretical issues without practical impact
* Issues requiring unrealistic assumptions
* Misuse of the system outside intended design

---

## Final Note

Security in AegisQ is not an afterthought.

Any contribution or change that weakens:

* Correctness
  n- Determinism
* Cryptographic guarantees

will be rejected.
