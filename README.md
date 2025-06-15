# GolangCards

A simple card deck management app written in Go. This project was built while following the **[Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/)** Udemy course by Stephen Grider. It demonstrates core language features such as slices, functions, methods, file I/O, and randomization.

## 🃏 Features

- Generate a new deck of cards
- Print all cards to the console
- Deal hands from the deck
- Save a deck to a file
- Load a deck from a file
- Shuffle the deck randomly

## 📁 Project Structure

- `deck.go`: Contains the core logic for the `deck` type and its methods.
- `my_cards`: A saved deck of cards as a plain text file.
- `go.mod`: Go module file.

## 🔧 How to Run

```bash
go run deck.go
Ensure you have Go installed. If not, download and install Go.

📦 Functions Overview
newDeck() deck: Creates a full deck.

print(): Prints each card with its index.

deal(deck, int) (deck, deck): Deals a hand and returns the remainder.

saveToFile(string) error: Saves the deck to a file.

newDeckFromFile(string) deck: Loads a deck from a file.

shuffle(): Shuffles the cards in the deck.

💡 Notes
The deck currently includes a subset of cards: Ace through Four of each suit.

File operations use os.WriteFile and os.ReadFile, so they overwrite and read plain-text files like my_cards.

🧠 Lessons Learned
This project explores:

Custom types in Go (type deck []string)

Slices and range-based loops

Method receivers

File I/O and error handling

Random number generation with custom seeds

📚 Based On
Go: The Complete Developer's Guide (Golang) by Stephen Grider
