# AegisQ Protocol

<p align="center">
<b>Post-Quantum Secure • Deterministic BFT • Immediate Finality</b>
</p>

---

# Overview

**AegisQ** is a **Post-Quantum Secure Deterministic BFT Blockchain Engine** designed for high-integrity data anchoring and adversarial environments.

It provides:

* Deterministic consensus
* Strict equivocation prevention
* Immediate block finality
* Post-quantum signature support
* Persistent storage
* High-throughput block processing

The protocol is implemented entirely in **Go** and includes a **full blockchain explorer UI** for inspecting blocks and transactions.

---

# What This Project Implements

This repository contains a **complete blockchain stack**, including:

| Component       | Description                                    |
| --------------- | ---------------------------------------------- |
| Protocol Engine | Deterministic BFT blockchain core              |
| Crypto Engine   | Post-Quantum Dilithium2 + Classical signatures |
| Consensus       | Prepare / Commit voting phases                 |
| Ledger          | Chain integrity enforcement                    |
| Storage         | BoltDB persistent state                        |
| Scheduler       | Deterministic leader rotation                  |
| Simulation      | Synthetic transaction generator                |
| Explorer API    | REST interface for querying chain state        |
| Explorer UI     | Real-time blockchain explorer dashboard        |

---

# Architecture — Deterministic 10 Layer Stack

| Layer | Component            | Responsibility                           |
| ----- | -------------------- | ---------------------------------------- |
| 1     | Crypto               | PQC + classical signature abstraction    |
| 2     | Identity             | Validator identity & key management      |
| 3     | Transactions         | Signed data anchoring transactions       |
| 4     | Blocks               | Merkle-rooted block structure            |
| 5     | Ledger               | Chain validation & integrity enforcement |
| 6     | Validator Governance | Membership control                       |
| 7     | Deterministic Leader | Round-robin block proposer               |
| 8     | View Rotation        | Leader failover logic                    |
| 9     | BFT Voting           | 2f+1 quorum enforcement                  |
| 10    | Finality Engine      | Commit lock & fork prevention            |

Each layer is modular and independently testable.
S
---

# Core Features

### Deterministic BFT Consensus

AegisQ implements a deterministic **Prepare → Commit** consensus model.

Properties:

* No probabilistic finality
* No chain reorganizations
* No fork ambiguity

Once a block reaches commit quorum it becomes **permanently finalized**.

---

### Byzantine Fault Tolerance
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

The protocol tolerates up to **f Byzantine validators** while preserving:

* Safety
* Liveness

---

### Strict Equivocation Prevention

Validators are prevented from:

* Voting twice in the same phase
* Voting for two blocks in the same view

Any violation is rejected at the consensus layer.

---

# Post‑Quantum Cryptography

AegisQ supports both classical and post‑quantum signatures.

## Active Mode

Currently running with:

* **Dilithium2 (ML‑DSA‑44)**

All transactions and blocks are signed using post‑quantum cryptography.
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

Environment:

* CPU: Intel Core i7
* OS: Linux amd64
* Language: Go

Command:

```
go test ./core/crypto -bench=. -benchmem -run=^$
```

### Results

```
BenchmarkDilithiumKeyGen        ~38µs
BenchmarkDilithiumSign          ~104µs
BenchmarkDilithiumVerify        ~37µs

BenchmarkECDSAKeyGen            ~22µs
BenchmarkECDSASign              ~72µs
BenchmarkECDSAVerify            ~110µs
```

Observation:

* Dilithium verification is **~3× faster than ECDSA**
* Verification dominates validator workload

Therefore Dilithium improves consensus throughput.

---

# Runtime Performance

Test workload:

* **10,000 transactions per block**

### Dilithium Mode

| Stage                  | Time   |
| ---------------------- | ------ |
| Transaction Generation | ~1.2s  |
| Block Finalization     | ~117ms |
| Total Runtime          | <1.5s  |

### Ed25519 Mode

| Stage                  | Time   |
| ---------------------- | ------ |
| Transaction Generation | ~400ms |
| Block Finalization     | ~35ms  |

Observation:

Post‑quantum signing increases cost but verification remains efficient.

---

# Storage Architecture

Persistent storage uses **BoltDB**.

Buckets:

* blocks
* block_hash_index
* tx_index
* meta

Supports:

* O(1) block retrieval by height
* O(1) transaction lookup by hash

---

# Blockchain Explorer

AegisQ includes a **full explorer UI**.

Explorer capabilities:

* Network dashboard
* Block explorer
* Transaction explorer
* Validator monitoring
* Block search
* Transaction search

### Dashboard Metrics

* Latest block
* Active validators
* Block transaction count
* Network TPS

### Block Explorer

Displays:

* Block height
* Transaction count
* Merkle root
* Validator proposer

### Transaction Explorer

Displays:

* Sender
* Signature algorithm
* Data hash
* Metadata
* Timestamp

---

# CLI Usage

## Run Node

```
go run ./cmd/aegisqd
```

Creates a block with **10,000 synthetic transactions**.

---

## Query Transaction by Height

```
go run ./cmd/aegisqd gettx <height> <index>
```

Example:

```
go run ./cmd/aegisqd gettx 1 5000
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

## Query Transaction by Hash

```
go run ./cmd/aegisqd gettxhash <data_hash>
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
├── block/
├── consensus/
├── crypto/
├── identity/
├── ledger/
├── scheduler/
├── simulation/
├── storage/
├── transaction/

cmd/
└── aegisqd/

explorer/
└── web-ui/
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

# Stress Testing

The system has been tested against:

* Double vote attack
* Validator equivocation
* Fork injection
* Unauthorized validator voting
* Block tampering
* High load transaction simulation

All tests remain within theoretical BFT tolerance bounds.

---

# Current System Status

| Component              | Status      |
| ---------------------- | ----------- |
| Protocol Engine        | Complete    |
| Post‑Quantum Crypto    | Integrated  |
| Deterministic Finality | Enforced    |
| Persistent Storage     | Enabled     |
| Transaction Indexing   | Implemented |
| Blockchain Explorer    | Functional  |
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

**v1.1.0 — Post‑Quantum Deterministic BFT Engine**
`v1.0.0` — Hardened BFT Core Freeze

---

# License

This project is licensed under the **MIT License**.

See the `LICENSE` file for details.

---

<p align="center">
<b>Deterministic Trust — Engineered for Adversarial Environments</b>
<b>Deterministic Trust. Engineered for Adversarial Environments.</b>
</p>
