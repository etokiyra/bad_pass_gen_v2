# bad_pass_gen_v2 💀 (Go edition)

**The world's worst password generator — now rewritten in Go.**

Same terrible idea as the original: instead of generating a real random password, this tool picks a random line straight out of `rockyou.txt`, one of the most infamous leaked password lists on the internet. Every password it hands you is already known to every password cracker on Earth.

This is satire, built as a learning project while porting from Rust to Go. **Do not use anything this tool outputs as an actual password.**

## ⚠️ Disclaimer

This project is provided "as is", for educational and humorous purposes only, with absolutely no warranty of any kind — see the [LICENSE](LICENSE).

- This tool does **not** generate secure passwords. It selects an existing line from a public, well-known breach corpus (`rockyou.txt`).
- Any password it outputs is, by definition, already compromised and trivially crackable.
- If you use output from this tool as an actual password for an actual account and something bad happens as a result, that is entirely on you — the tool tells you exactly what it's doing, right here, in the source code, and in its own output.
- The author assumes no liability for any use, misuse, account compromise, data loss, or bad life decisions resulting from running this software.

## What it does

1. Loads `rockyou.txt` (~14 million real leaked passwords)
2. Picks one at random
3. Hands it to you, unearned confidence included

## Requirements

- [Go](https://go.dev/doc/install)
- Your own copy of `rockyou.txt` — not included in this repo (it's ~140MB and, again, a breach corpus). It's a widely available wordlist; find your own copy responsibly.

## Installation & running

```bash
git clone https://github.com/etokiyra/bad_pass_gen_v2.git
cd bad_pass_gen_v2
```

Place your own `rockyou.txt` file directly in the `bad_pass_gen_v2/` project folder (same level as `go.mod`).

Then run:

```bash
go run main.go
```

Each run prints one (bad, insecure, thoroughly cracked) password.

## Also available in Rust

The original version of this joke lives at [bad_pass_gen](https://github.com/etokiyra/bad_pass_gen), if you want to compare the Rust and Go implementations side by side.

## License

See [LICENSE](LICENSE).
