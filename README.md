# Local Hack Day: Build - Write Code to Sort a List
To challenge myself, I decided to complete this challenge using a new language. To make it a bit more interesting, I implemented a few of the most common sorting algorithms and tracked the execution time of each. This implementation sorts the hackathon event names scraped from https://mlh.io/seasons/2021/events using goquery.

# Installation
This project requires the installation of [Go](https://golang.org) and the goquery package (used for web scraping).

goquery package installation:
```
$ go get github.com/PuerkitoBio/goquery
```

# Getting Started
### Step 1. Clone the code into a fresh folder
```
$ git clone https://github.com/granaj91/go-sort.git
$ cd go-sort
```

### Step 2. Run the program
```
$ go run sort.go web_scraper.go
```

### Step 3. Run tests
```
$ go test -v
```
