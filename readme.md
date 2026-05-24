# QAE-BAC: Quantifiable Anonymity and Efficiency in Blockchain-Based Access Control

This repository contains the code implementation for the paper:

> **[QAE-BAC: Achieving Quantifiable Anonymity and Efficiency in Blockchain-Based Access Control with Attribute](https://ieeexplore.ieee.org/document/11534385)**  
> Published in *IEEE Internet of Things Journal (IoTJ)* (Accepted)  
> DOI: [10.1109/JIOT.2026.3695861](https://doi.org/10.1109/JIOT.2026.3695861)  
> Full version: arXiv: [2510.21124](https://arxiv.org/abs/2510.21124)

Code implemented by co-author [Zhang Mengke](https://github.com/Mooonk).

---

## Overview

QAE-BAC is an Attribute-Based Access Control (ABAC) system designed for Electronic Health Records (EHR). Its core contribution is an entropy-driven, adaptive policy decision tree that dynamically reorders attribute evaluation order based on historical access patterns, improving both efficiency and quantifiable anonymity.

The experiments in this repository cover:
1. **Access Control Performance** — comparing baseline linear scan (`BaseABAC`), fixed-order trie (`DicABAC`), and entropy-optimized trie (`AnoABAC`)
2. **Anonymity Metrics** — measuring subject-level anonymity across access request distributions

---

## Prerequisites

| Requirement | Version |
|---|---|
| Go | ≥ 1.18 |
| MySQL | ≥ 5.7 |

Install Go dependencies:
```bash
go mod download
```

---

## Repository Structure

```
QAE-BAC/
├── main.go              # Entry point — calls mytest.Simulate()
├── abac/                # Core ABAC engine: trie construction, entropy feature selection
├── anonymity/           # Anonymity metric calculation (CalSubAnonymity)
├── dataGen/             # Synthetic dataset generator (subjects, objects, policies, requests)
├── dataset/             # Generated CSV datasets (subject.csv, object.csv, policy.csv, request.csv)
│   ├── G1/ ~ G15/       # Pre-generated dataset groups for different experiment configurations
│   ├── subject.csv
│   ├── object.csv
│   └── policy.csv
├── model/               # Data models: Sub, Obj, ABACRequest, Policy, MyToken
├── mytest/              # Simulation harness and benchmarking (Simulate, TestCompareAC, TestAnonmity)
├── mytools/             # Utility functions: CSV I/O, hashing, plotting
├── sql/                 # MySQL connector
├── PM/                  # Policy management module
├── RSA/                 # RSA key generation and encryption/decryption
├── UM/                  # User management
├── EM/                  # EHR (Electronic Health Record) management
├── EA/                  # Emergency access delegation
├── zk/                  # Zero-knowledge proof for emergency delegation
├── files/               # RSA key store (public/ and private/ PEM files)
├── result/              # Output directory for benchmark and anonymity results
├── vendor/              # Vendored Go dependencies
│
│   ── Blockchain Integration (see note below) ──
├── network/             # Hyperledger Fabric network scripts and Docker configs
├── fixtures/            # Fabric PKI / MSP crypto material
├── config_monk.yaml     # Fabric SDK connection profile (org1)
├── config_monk1.yaml    # Fabric SDK connection profile (org2)
└── config_monk2.yaml    # Fabric SDK connection profile (orderer)
```

> **Note on Blockchain Components:** The `network/`, `fixtures/`, and `config_monk*.yaml` files implement the Hyperledger Fabric integration (v2.2) used for on-chain policy storage and access logging. These components are part of ongoing work and are **not required** to reproduce the core access control and anonymity experiments described in the paper. See the [Blockchain Integration](#blockchain-integration-hyperledger-fabric) section for guidance on connecting to an existing Fabric network.

---

## Running the Experiments

### Step 1 — Configure the Database

The policy store uses MySQL. Create a database and update the connection string in `sql/sql.go`:

```go
// sql/sql.go
db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/qae_bac")
```

### Step 2 — Configure Experiment Parameters

Open `mytest/mytest.go` and set the desired experiment group by uncommenting the corresponding parameter block. The default (Group 15) is:

```go
// mytest/mytest.go
userNum  = 10000
dataNum  = 10000
reqNum   = 1000000
abac.Sub_Attr  = []string{"A", "B", "C", "D"}
abac.AttrinNum = 2
policyNum      = 100
abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q", "R", "S"}
```

The paper's baseline setup is **4 organizations, 4 attributes (4org-4attr)** with:
- `abac.NewHistoryPool(600)` — history pool size
- `NewTreePools(150)` — tree pool size
- Parallel request counts: 1, 5, 10, 50, 100 × 1000

### Step 3 — Generate Datasets (optional)

Pre-generated datasets for all groups are available under `dataset/G1/` through `dataset/G15/`. To regenerate:

Uncomment the data generation calls in `mytest/mytest.go`:

```go
// mytest/mytest.go  (inside Simulate())
dataGen.CreateData(dataNum, "O1000")
dataGen.CreateUser(userNum, "S1000")
dataGen.CreatePolicy(2, 2)
dataGen.CreateRequest(int(0.2*float64(reqNum)), int(0.8*float64(reqNum)))
```

Generated files are written to `dataset/`:
- `dataset/subject.csv` — subject attribute records
- `dataset/object.csv` — object attribute records
- `dataset/policy.csv` — access policy rules
- `dataset/request.csv` — access request traces (80% deny, 20% permit)

### Step 4 — Run

```bash
go run main.go
```

The entry point calls `mytest.Simulate()`, which sequentially:
1. Generates (or reuses) the request dataset
2. Runs `TestAnonmity()` — computes anonymity metrics and writes `result/anonmityReq.csv`, `result/anonmitySub.csv`, and PNG scatter plots
3. Runs `TestCompareAC()` — benchmarks all three AC strategies and writes `result/introAC.csv`

To run individual tests, edit `main.go`:

```go
// main.go
mytest.TestAnonmity()    // anonymity metrics only
mytest.TestCompareAC()   // AC performance comparison only
```

### Step 5 — Inspect Results

| Output File | Contents |
|---|---|
| `result/anonmityReq.csv` | Per-request anonymity distribution |
| `result/anonmitySub.csv` | Per-subject anonymity distribution |
| `result/introAC.csv` | Throughput (req/s) and latency (ns/req) for BaseABAC, DicABAC, AnoABAC |
| `result/anonmity_sub_pro1.png` | Scatter plot of subject anonymity |
| `result/anonmity_req_pro1.png` | Scatter plot of request anonymity |

> The paper's figures were produced using Excel from the CSV outputs.

---

## Experiment Groups Reference

The following groups correspond to the paper's evaluation scenarios. Switch between them by editing the parameter block in `mytest/mytest.go`:

| Group | userNum | dataNum | Sub_Attr count | AttrinNum | policyNum | Notes |
|---|---|---|---|---|---|---|
| G2 (baseline) | 10000 | 10000 | 4 | 2 | 100 | Baseline |
| G1 | 5000 | 10000 | 4 | 4 | 100 | Fewer users |
| G3 | 15000 | 10000 | 4 | 4 | 100 | More users |
| G4 | 10000 | 5000 | 4 | 4 | 100 | Fewer objects |
| G5 | 10000 | 15000 | 4 | 4 | 100 | More objects |
| G10 | 10000 | 10000 | 4 | 2 | 100 | Lower inValue |
| G11 | 10000 | 10000 | 4 | 6 | 100 | Higher inValue |
| G12 | 10000 | 10000 | 4 | 2 | 50 | Fewer policies |
| G13 | 10000 | 10000 | 4 | 2 | 150 | More policies |
| G15 | 10000 | 10000 | 4 | 2 | 100 | Extended attr list (12 attrs) |

---

## Blockchain Integration (Hyperledger Fabric)

The full QAE-BAC system uses **Hyperledger Fabric v2.2** for decentralized policy storage and tamper-evident access logging. The `network/` directory contains the Docker-based network setup (2 organizations + orderer), and `config_monk*.yaml` are the Fabric Go SDK connection profiles.

**These files are not required to reproduce the paper's core experiments.** The access control and anonymity benchmarks run entirely on local CSV files and MySQL.

If you wish to integrate the blockchain layer using an existing public Fabric deployment, the following adaptations are needed:

1. **Network Setup**: Deploy a Fabric v2.2 network with at least 2 peer organizations. The official [fabric-samples](https://github.com/hyperledger/fabric-samples) `test-network` is a suitable starting point.

2. **Chaincode**: The policy management logic in `PM/pm.go` must be wrapped as a Fabric chaincode. The `PM.GetPolicy()` and `PM.SetPolicy()` calls should be replaced with `contract.EvaluateTransaction()` / `contract.SubmitTransaction()` calls via the [Fabric Go SDK](https://github.com/hyperledger/fabric-sdk-go).

3. **Connection Profile**: Replace `config_monk.yaml` with the connection profile generated by your Fabric network. Update the MSP ID, peer endpoints, orderer endpoints, and TLS certificate paths to match your deployment.

4. **Access Logging**: The access token (`model.MyToken`) issued after a successful `CalTreeABAC` call should be submitted to the ledger for audit. Wire this into the `CalTreeABAC` return path in `abac/ac.go`.

5. **Identity (PKI)**: The `RSA/` and `UM/` modules handle off-chain identity. In a full Fabric deployment, user enrollment should be delegated to the Fabric CA (`fabric-ca-server`), and the RSA keys in `files/` replaced with Fabric-issued x.509 certificates.

> The exact chaincode interface and channel configuration depend on your Fabric network topology and will require project-specific adjustments beyond a drop-in replacement.

---

## Citation

If you use this code, please cite:

```bibtex
@ARTICLE{11534385,
  author={Zhang, Jie and Li, Xiaohong and Zhang, Mengke and Feng, Ruitao and Xu, Shanshan and Hou, Zhe and Bai, Guangdong},
  journal={IEEE Internet of Things Journal}, 
  title={QAE-BAC: Achieving Quantifiable Anonymity and Efficiency in Blockchain-Based Access Control with Attribute}, 
  year={2026},
  doi={10.1109/JIOT.2026.3695861}}

```

---

## License

See [LICENSE](LICENSE).


### Citations

**File:** mytest/mytest.go (L22-190)
```go
func Simulate() {
	// abac.Attr_list=[]string{"A", "B", "E", "C", "D", "X", "Y", "O"}
	// Sub_Attr = []string{"A", "B", "C", "D"}
	// Obj_Attr = []string{"X", "Y"}

	// // group 1: 为保证一致性，policy没有重新生成，而是使用group2的，这样保证只有sub数量是变幻的
	// userNum = 5000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 2---baseline
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 3：只生成sub和req，policy也不动
	// userNum = 15000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 4：只改变obj和req，policy等不动
	// userNum = 10000
	// dataNum = 5000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 5：只改变obj和req，policy等不动
	// userNum = 10000
	// dataNum = 15000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 6: 为保证一致性，不再生成req，直接取其中的500k
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 500000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 7: 为保证一致性，在原有1000k的基础上加500k
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 500000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 8:
	// userNum = 15000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D", "E"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 9:
	// userNum = 15000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C"}
	// abac.AttrinNum = 4
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 10: inValue改变，所有数据集更改
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// group 11: inValue改变，所有数据集更改--
	// 这里对比的是group3，因为6^4=1296===不对，还是2吧
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 6
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 12: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 50
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}

	// // group 13: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 150
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O"}
	// // group 14: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q"}

	// // group 14: 只改变polcy个数，（连带着需要重新生成req）
	// userNum = 10000
	// dataNum = 10000
	// reqNum = 1000000
	// abac.Sub_Attr = []string{"A", "B", "C", "D"}
	// abac.AttrinNum = 2
	// policyNum = 100
	// abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q"}

	// group 15: 只改变polcy个数，（连带着需要重新生成req）
	userNum = 10000
	dataNum = 10000
	reqNum = 1000000
	abac.Sub_Attr = []string{"A", "B", "C", "D"}
	abac.AttrinNum = 2
	policyNum = 100
	abac.Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q", "R", "S"}

	old_attr_list := abac.Attr_list
	abac.PoolNum = reqNum / 10
	for i := 0; i < 1; i++ {
		abac.Attr_list = old_attr_list

		// // fmt.Println("///////////////////////////模拟生成数据集///////////////////////////////////////")
		// dataGen.CreateData(dataNum, "O1000")
		// dataGen.CreateUser(userNum, "S1000")
		// dataGen.CreatePolicy(2, 2)
		dataGen.CreateRequest(int(0.2*float64(reqNum)), int(0.8*float64(reqNum)))
		// // req := mytools.ReadCSV("./dataset/request.csv")
		// // mytools.WriteCSV("./dataset/request.csv", req[:reqNum])
		// fmt.Println("///////////////////////////数据集生成完毕///////////////////////////////////////")

		// 如果散点图不好，可能是策略的原因
		TestAnonmity()
		// TestCompareAC()

	}

```

**File:** main.go (L21-26)
```go
func main() {

	// mytest.TestAnonmity()
	// dataGen.Attr_list = abac.ChooseBestFeature()
	mytest.Simulate()
	// fmt.Println(e, f)
```

**File:** go.mod (L1-22)
```text
module algorithm/mycode

go 1.18

require (
	github.com/go-sql-driver/mysql v1.8.1
	gonum.org/v1/plot v0.15.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	git.sr.ht/~sbinet/gg v0.6.0 // indirect
	github.com/ajstarks/svgo v0.0.0-20211024235047-1546f124cd8b // indirect
	github.com/campoy/embedmd v1.0.0 // indirect
	github.com/go-fonts/liberation v0.3.3 // indirect
	github.com/go-latex/latex v0.0.0-20240709081214-31cef3c7570e // indirect
	github.com/go-pdf/fpdf v0.9.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/image v0.21.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)
```

**File:** PM/pm.go (L1-10)
```go
// policy management
package PM

// import (
// 	"algorithm/mycode/sql"
// 	SQL "database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"strings"
// 	"time"
```
