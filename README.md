# AegisQ Protocol

**Post-Quantum Secure • Byzantine-Resilient • Deterministic Blockchain Core**

---

## Overview

**AegisQ** is a permissioned blockchain protocol engineered for deterministic consensus, strict equivocation prevention, and hardened Byzantine fault tolerance.

It serves as a secure, modular foundation for post-quantum data anchoring systems.

### Core Capabilities

* Byzantine Fault Tolerance (BFT)
* Deterministic Leader Rotation
* View-Based Failover
* Strict Equivocation Prevention
* Immediate Finality Guarantees
* Fork Resistance
* Post-Quantum Cryptography Integration Ready
* IPFS Storage Layer Ready

This repository contains the **core consensus and ledger engine**.

---

# Architecture — 10 Layer Design

AegisQ follows a strictly modular layered architecture:

| Layer | Component            | Responsibility                  |
| ----- | -------------------- | ------------------------------- |
| 1     | Crypto (PQC-Ready)   | Dilithium / Ed25519 abstraction |
| 2     | Identity             | Node identity & signing         |
| 3     | Transactions         | Signed metadata transactions    |
| 4     | Blocks               | Merkle-rooted signed blocks     |
| 5     | Ledger               | Chain integrity enforcement     |
| 6     | Validator Governance | Authorization & membership      |
| 7     | Deterministic Leader | Round-robin proposer            |
| 8     | View-Based Failover  | Leader rotation logic           |
| 9     | BFT Voting           | 2f+1 quorum enforcement         |
| 10    | Finality Engine      | Fork prevention & commit lock   |

Each layer is independently testable and designed for long-term extensibility.

---

# Security Guarantees

## Byzantine Fault Tolerance

For `n` validators:

```
f = (n - 1) / 3
```

The system tolerates up to **f Byzantine validators** while maintaining safety and liveness.

---

## Strict Equivocation Prevention

A validator cannot:

* Vote twice in the same view and phase
* Vote for two different blocks in the same view

All equivocation attempts are rejected at the voting layer.

---

## Fork Resistance

Under tolerated Byzantine conditions, two conflicting blocks cannot both reach quorum.

---

## Deterministic Finality

Once a block reaches commit quorum, it becomes irreversible.

No probabilistic confirmations.
No reorg assumptions.

---

# Stress Testing Results

| Transactions | Finalize Time | Total Execution Time |
| ------------ | ------------- | -------------------- |
| 1,000        | ~18ms         | ~0.2s                |
| 10,000       | ~107ms        | ~1.7s                |
| 50,000       | ~540ms        | ~8.7s                |

**Observations**

* Linear scalability
* Merkle hashing dominates compute cost
* Consensus overhead remains minimal
* No fork under tolerated Byzantine conditions

---

# Byzantine Attack Validation

Validated against:

* Double vote attack
* Equivocation attempt
* Fork injection exploit
* Unauthorized validator voting
* Malicious block tampering
* High-load Byzantine scenarios (50k transactions)

All tests pass within theoretical BFT tolerance bounds.

---

# Project Structure

```
core/
├── block/          # Block structure & integrity
├── consensus/      # BFT voting & finality engine
├── crypto/         # PQC + classical abstraction
├── identity/       # Node identities
├── ledger/         # Chain validation logic
├── scheduler/      # Leader rotation
├── simulation/     # Stress & Byzantine tests
├── transaction/    # Transaction structure
```

---

# Getting Started

```bash
git clone https://github.com/<your-username>/aegisq-protocol.git
cd aegisq-protocol
```

```bash
go clean -cache -testcache
go test ./... -v
```

---

# Version

`v1.0.0` — Hardened BFT Core Freeze

---

# License

This project is licensed under the **MIT License**.

See the `LICENSE` file for details.

---

<p align="center">
<b>Deterministic Trust. Engineered for Adversarial Environments.</b>
</p>
