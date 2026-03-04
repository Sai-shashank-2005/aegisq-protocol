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
```

---

## Query Transaction by Hash

```
go run ./cmd/aegisqd gettxhash <data_hash>
```

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

---

# Version

**v1.1.0 — Post‑Quantum Deterministic BFT Engine**

---

<p align="center">
<b>Deterministic Trust — Engineered for Adversarial Environments</b>
</p>
