# Pokédex

A simple **Pokédex CLI** built in **pure Go** (no external packages).  
Run it from your terminal to explore Pokémon.

---

## ✨ Features
- Search Pokémon by name
- Display basic Pokémon data (name, type, height, weight, etc.)
- Minimal dependencies → only Go standard library
- Fast and lightweight command-line interface

---

## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) installed (1.20+ recommended)

Check your Go installation:
```bash
go version

git clone https://github.com/MarianGheorghiu/pokedex.git
cd pokedex

go build -o pokedex

./pokedex

$ ./pokedex pikachu

Name: Pikachu
Type: Electric
Height: 0.4 m
Weight: 6.0 kg
