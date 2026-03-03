# AegisQ Protocol

**Post-Quantum Secure • Byzantine-Resilient • Deterministic Blockchain Core**

---

## Overview

**AegisQ** is a permissioned blockchain protocol engineered for:

* Deterministic BFT consensus
* Strict equivocation prevention
* Hardened fork resistance
* Immediate finality
* Post-Quantum signature support

It serves as a secure modular foundation for tamper-proof data anchoring systems.

This repository contains the full consensus, ledger, storage, and cryptographic engine.

---

# Architecture — 10 Layer Deterministic Stack

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

Each layer is independently testable and modular by design.

---

# Current System Capabilities

* 10,000 transactions per block
* Dilithium2 (Post-Quantum) signature support
* Ed25519 classical support (optional)
* Persistent storage via BoltDB
* O(1) transaction lookup (height + index)
* O(1) transaction lookup (hash index)
* Deterministic leader rotation
* Strict equivocation prevention
* Prepare + Commit BFT phases
* Immediate deterministic finality
* Synthetic dataset stress testing
* CLI transaction inspection tools

---

# Security Model

## Byzantine Fault Tolerance

For `n` validators:

```
f = (n - 1) / 3
```

The system tolerates up to `f` Byzantine validators while preserving safety and liveness.

---

## Strict Equivocation Prevention

A validator cannot:

* Vote twice in the same view and phase
* Vote for two different blocks in the same view

All violations are rejected at the voting layer.

---

## Deterministic Finality

Once a block reaches commit quorum:

* The block becomes irreversible
* No probabilistic confirmations
* No chain reorg assumptions

---

# Post-Quantum Cryptography Integration

AegisQ supports both classical and post-quantum signatures.

## Active Production Mode

Currently running with:

* Dilithium2 (ML-DSA-44)

Transactions and blocks are signed using post-quantum cryptography.

---

# Cryptographic Benchmark Results

Environment:

* CPU: 13th Gen Intel® Core™ i7-13620H
* OS: Linux amd64
* Go: native benchmarking

Command used:

```
go test ./core/crypto -bench=. -benchmem -run=^$
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

## Verification Performance (Validator Critical Path)

| Algorithm   | Verify Time | Allocations |
| ----------- | ----------- | ----------- |
| Dilithium2  | ~37.8 µs    | 0           |
| ECDSA P-256 | ~109.9 µs   | 29          |

Dilithium verification is approximately **3× faster** than ECDSA in this environment.

Validators predominantly verify signatures — making Dilithium advantageous.

---

## Signing Performance (Client-Side)

| Algorithm   | Sign Time |
| ----------- | --------- |
| ECDSA P-256 | ~71.9 µs  |
| Dilithium2  | ~104 µs   |

ECDSA signs faster, but signing is client-side and not consensus critical.

---

# Runtime Performance (10,000 Transactions per Block)

Measured using:

```
time go run ./cmd/aegisqd
```

## Dilithium2 Results

* Transaction generation: ~1.2s
* Block finalize: ~117ms
* Total execution: < 1.5s

## Ed25519 Results

* Transaction generation: ~400ms
* Block finalize: ~35ms

Observation:

Dilithium increases signing cost (~3x) but maintains strong verification efficiency.

---

# Transaction Model

Each transaction contains:

* Sender ID
* Public Key
* Algorithm
* DataHash (anchored external data)
* Metadata
* Timestamp
* Signature

Transactions are:

* Individually signed
* Merkle aggregated
* Indexed by height and hash

---

# Storage Architecture

Persistent database:

* BoltDB backend

Buckets:

* blocks
* block_hash_index
* tx_index
* meta

Supports:

* Height-based block retrieval
* Hash-based transaction retrieval
* Constant-time lookup

---

# CLI Commands

---

## Fresh Start (Delete Existing Chain)

If you switch cryptographic algorithms or want to restart from genesis:

```
rm aegisq.db
```

Then run the node again:

```
go run ./cmd/aegisqd
```

This will create a new chain starting at height 1.

---

## Create Block (Default Execution)

Running the node without arguments automatically:

* Initializes validators
* Restores latest height (if DB exists)
* Generates 10,000 synthetic transactions
* Runs BFT Prepare + Commit
* Finalizes and stores the block

Command:

```
go run ./cmd/aegisqd
```

---

## Query Transaction by Height + Index

## Run Node

```
go run ./cmd/aegisqd
```

Creates block with 10,000 synthetic transactions.

---

## Query Transaction by Height + Index

```
go run ./cmd/aegisqd gettx <height> <index>
```

Example:

```
go run ./cmd/aegisqd gettx 1 5000
```

---

## Query Transaction by Hash

```
go run ./cmd/aegisqd gettxhash <data_hash>
```

---

# Stress Testing

Validated against:

* Double vote attack
* Equivocation attempt
* Fork injection exploit
* Unauthorized validator voting
* Malicious block tampering
* High-load 50k transaction simulation

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
```

---

# Current System Status

Architecture completeness: High
PQC Integration: Active
Deterministic Finality: Enforced
Persistent Storage: Enabled
Transaction Indexing: Enabled
Adversarial Test Coverage: Extensive

Next planned enhancements:

* Full chain verification on startup
* Merkle proof generation API
* Multi-node network layer
* Advanced Byzantine live simulation

---

# Version

v1.1.0 — Post-Quantum Persistent BFT Engine

---

<p align="center">
<b>Deterministic Trust. Engineered for Adversarial Environments.</b>
</p>
