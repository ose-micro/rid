# ose-core

> 🚀 Readable ID generator for ose-micro ecosystem.  
Generates **UUID-length (36 char)** IDs that are **human-friendly**, **timestamped**, and **prefix-based**.

[![Go Reference](https://pkg.go.dev/badge/github.com/ose-micro/core.svg)](https://pkg.go.dev/github.com/ose-micro/rid)
[![Go Report Card](https://goreportcard.com/badge/github.com/ose-micro/core)](https://goreportcard.com/report/github.com/ose-micro/rid)
[![License](https://img.shields.io/github/license/ose-micro/rid)](LICENSE)

---

---

## 📌 Features
- Fixed **36 chars** (UUID standard length).
- Includes **prefix** (e.g. PRD, CMP, USR).
- Embeds **date + time** for readability.
- Adds random filler for global uniqueness.
- Output can be **uppercase or lowercase**.

---

## 📦 Install

```bash
go get github.com/ose-micro/rid
