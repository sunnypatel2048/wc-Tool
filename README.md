# Write Your Own wc Tool

My own version of Unix command line tool `wc`. It prints the word, line character, and byte count of a given file.

This tool was created as a part of [Johh Crickett's](https://www.linkedin.com/in/johncrickett/) Coding Challenges. You can find the challenge description [here](https://codingchallenges.fyi/challenges/challenge-wc).

## Getting Started

### Prerequisites

- Go (Golang) should be installed on your system. If it's not, you can download it from [here](https://golang.org/dl/).

### Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/sunnypatel2048/wc-Tool.git
   ```
   
2. Switch to the project directory:

    ```bash
    cd wc-Tool
    ```

## Setup

1. Open a terminal and navigate to the project directory.

2. Build the code using following command:
    ```bash
    go build -o ccwc.exe main.go
    ```

3. Copy the `ccwc.exe` file to `~/go/bin`. Note that this directory thoud be present in PATH variable in Environment variables.
    ```bash
    cp ccwc.exe ~/go/bin/
    ```

## Usage

```bash
ccwc [flag] [filepath]
```

## Flags Supported

- `-c` : Prints the number of bytes in the file.

## License

This project is licensed under the [MIT License](LICENSE). You can view the full license text [here](https://opensource.org/licenses/MIT).