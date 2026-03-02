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
* Post-Quantum Cryptography Integration
* IPFS Storage Layer Ready

This repository contains the **core consensus and ledger engine**.

---

# Architecture — 10 Layer Design

AegisQ follows a strictly modular layered architecture:

| Layer | Component            | Responsibility                |
| ----- | -------------------- | ----------------------------- |
| 1     | Crypto (PQC-Ready)   | Dilithium / ECDSA abstraction |
| 2     | Identity             | Node identity & signing       |
| 3     | Transactions         | Signed metadata transactions  |
| 4     | Blocks               | Merkle-rooted signed blocks   |
| 5     | Ledger               | Chain integrity enforcement   |
| 6     | Validator Governance | Authorization & membership    |
| 7     | Deterministic Leader | Round-robin proposer          |
| 8     | View-Based Failover  | Leader rotation logic         |
| 9     | BFT Voting           | 2f+1 quorum enforcement       |
| 10    | Finality Engine      | Fork prevention & commit lock |

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

### Observations

* Linear scalability
* Merkle hashing dominates compute cost
* Consensus overhead remains minimal
* No fork under tolerated Byzantine conditions

---

# Cryptographic Benchmark Results

AegisQ includes both classical and post-quantum signature implementations for performance comparison and migration readiness.

Benchmarks were executed using Go’s testing framework on:

* **CPU:** 13th Gen Intel® Core™ i7-13620H
* **OS:** Linux amd64
* **Package:** `core/crypto`

---

## Benchmark Command

```bash
go test ./core/crypto -bench=. -benchmem -run=^$
```

---

## Benchmark Environment

```
goos: linux
goarch: amd64
cpu: 13th Gen Intel(R) Core(TM) i7-13620H
pkg: github.com/Sai-shashank-2005/aegisq-protocol/core/crypto
```

---

## Benchmark Results

```
BenchmarkDilithiumKeyGen-16        30282    38682 ns/op    4096 B/op    2 allocs/op
BenchmarkDilithiumSign-16          11571   104051 ns/op    2696 B/op    2 allocs/op
BenchmarkDilithiumVerify-16        31044    37893 ns/op       0 B/op    0 allocs/op

BenchmarkECDSAKeyGen-16            52716    22555 ns/op    1304 B/op   21 allocs/op
BenchmarkECDSASign-16              16744    71969 ns/op    7135 B/op   82 allocs/op
BenchmarkECDSAVerify-16             9475   109940 ns/op    1504 B/op   29 allocs/op
```

---

## Performance Analysis

### Verification (Validator-Critical Path)

| Algorithm             | Verify Time | Allocations |
| --------------------- | ----------- | ----------- |
| Dilithium (ML-DSA-44) | **37.8 µs** | 0           |
| ECDSA P-256           | 109.9 µs    | 29          |

Dilithium verification is approximately **3× faster** than ECDSA under this hardware configuration.

Since validators predominantly perform signature verification rather than signing, this significantly improves consensus throughput.

---

### Signing

| Algorithm   | Sign Time |
| ----------- | --------- |
| ECDSA P-256 | 71.9 µs   |
| Dilithium   | 104 µs    |

ECDSA is faster at signing (~1.4×).
Signing is typically client-side and not validator critical.

---

### Key Generation

| Algorithm | KeyGen Time |
| --------- | ----------- |
| ECDSA     | 22.5 µs     |
| Dilithium | 38.6 µs     |

ECDSA generates keys faster, as expected from classical elliptic curve cryptography.

---

## Memory Behavior

Dilithium verification performs:

```
0 B/op
0 allocs/op
```

This indicates:

* Zero heap pressure
* Predictable GC behavior
* Deterministic runtime characteristics

ECDSA verification involves internal `big.Int` allocations.

---

## Throughput Estimation

Dilithium Verify ≈ 37,893 ns ≈ 0.038 ms

Theoretical single-core ceiling:

```
~26,000 verifications / second
```

On multi-core systems, throughput scales with parallel validation.

---

## Cryptographic Design Implication

These results demonstrate:

* Post-quantum cryptography is production viable.
* Verification-heavy workloads benefit from Dilithium.
* AegisQ’s crypto layer is forward-compatible without sacrificing validator performance.

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

Run all tests:

```bash
go clean -cache -testcache
go test ./... -v
```

Run cryptographic benchmarks:

```bash
go test ./core/crypto -bench=. -benchmem -run=^$
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
