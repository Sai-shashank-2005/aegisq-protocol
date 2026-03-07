# AegisQ Explorer

AegisQ Explorer is a web-based blockchain explorer built for the **AegisQ Protocol**, a deterministic BFT blockchain engine with post‑quantum cryptography support.

The explorer provides a real-time interface for inspecting blocks, transactions, and network metrics produced by the AegisQ network.

---

# Features

## Network Dashboard

Displays high level network metrics:

* Latest block height
* Active validator count
* Block size
* Network TPS

Provides a quick overview of the current state of the blockchain.

---

## Block Explorer

Browse all blocks produced by the AegisQ network.

Each block card shows:

* Block height
* Transaction count
* Block hash

Clicking a block opens a detailed block view.

---

## Block Details

Detailed view for each block including:

* Block height
* Number of transactions
* Paginated transaction list

Each transaction entry links to the transaction explorer.

---

## Transaction Explorer

Displays full transaction information:

* Sender
* Signature algorithm
* Data hash
* Metadata
* Timestamp

Transactions can be queried by:

* Block height + index
* Data hash

---

## Search

Global search allows querying:

* Block height
* Transaction hash

---

# Architecture

The explorer is built as a separate module that communicates with the AegisQ node through a REST API.

```
AegisQ Node
   │
   ├── Consensus Engine
   ├── Ledger
   ├── Storage (BoltDB)
   │
   └── Explorer API
          │
          ▼
    Explorer Backend
          │
          ▼
      Web UI
```

---

# Explorer Components

| Component    | Description                    |
| ------------ | ------------------------------ |
| Dashboard    | Real-time network monitoring   |
| Blocks       | Block explorer interface       |
| Transactions | Transaction lookup and details |
| Network      | Validator and network state    |
| Search       | Block and transaction search   |

---

# Example Metrics

Typical runtime values:

| Metric     | Value               |
| ---------- | ------------------- |
| Block Size | 10,000 transactions |
| Validators | 4                   |
| TPS        | ~1200               |

---

# Running the Explorer

Start the AegisQ node:

```
go run ./cmd/aegisqd
```

Start the explorer frontend:

```
npm install
npm run dev
```

Open:

```
http://localhost:3000
```

---

# UI Pages

## Dashboard

Shows overall network statistics.

## Blocks

Displays the full list of blocks in reverse chronological order.

## Transactions

Allows direct transaction lookup.

## Network

Displays validator and network state.

---

# Technology Stack

Frontend:

* React
* Vite
* TailwindCSS
* Recharts

Backend API:

* Go

Storage:

* BoltDB

---

# Use Cases

The explorer enables:

* Blockchain auditing
* Transaction verification
* Network monitoring
* Validator activity inspection

---

# Status

| Component            | Status   |
| -------------------- | -------- |
| Dashboard            | Complete |
| Block Explorer       | Complete |
| Transaction Explorer | Complete |
| Search               | Complete |

---

# Version

Explorer v1.0

---

**AegisQ Explorer — Inspect Deterministic Finality**
