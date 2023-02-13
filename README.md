## URL Sizing CLI Tool
This is a command-line interface (CLI) tool that takes a list of URLs as input, visits each URL, and outputs the list of pairs: URL and response body size. The output is sorted by the size of the response body.

# Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

# Prerequisites
- Go programming language
- Go Modules enabled (Go version 1.11 or later)
# Installation
Clone the repository to your local machine:

```git clone https://github.com/cegielkowski/url-sizing-cli-tool.git```

Navigate to the directory containing the source code:

```cd url-sizing-cli-tool```

Build the executable:

```go build```

# Usage

Run the executable:

```./url-sizing-cli-tool```

You will be prompted to enter a list of URLs, separated by commas. After you have entered the URLs, press the Enter key to submit. The tool will visit each URL and print the list of pairs: URL and response body size, sorted by the size of the response body.

# Example Output
```
Enter the list of URLs (comma-separated): https://www.google.com,https://www.facebook.com,https://www.amazon.com
2023/02/12 22:51:38 https://www.amazon.com: 111111
2023/02/12 22:51:38 https://www.google.com: 123456
2023/02/12 22:51:38 https://www.facebook.com: 654321
```
# Built With
- Go - The programming language used
- net/http - The Go standard library package for HTTP client and server implementations
- sort - The Go standard library package for sorting slices and user-defined collections
# Contributing
If you would like to contribute to this project, please feel free to open a pull request or issue on the GitHub repository.