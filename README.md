# Nuchain-Txn: Local Nuchain Transaction

This project demonstrates how to interact with a **nuchain** blockchain locally and send transactions using Go.

## Requirements
- Go 1.17+
- Nuchain node running locally
- Ethereum-compatible RPC URL for connecting to your Nuchain node

## Setup

1. **Install Go**:
    - Download and install Go from the official website: https://golang.org/dl/
    - Ensure that Go is correctly installed by running: go version go1.19 linux/amd64

2. **Start the Nuchain Node**:
    - Ensure your Nuchain node is running locally. You can use the following command to run it:
    

3. **Clone the Repository**:
    - Clone this repository to your local machine:
    

4. **Install Dependencies**:
    - In the project directory, install any Go dependencies:
    

## Modifying the  File

The  file is responsible for interacting with the blockchain. Modify the following code to set your RPC URL and the address you want to send tokens to:

### In :



## Running the Code

1. **Run the Go Application**:
    - To run the application, execute the following command:
    

2. **Check Transactions**:
    - Monitor the output to ensure that the transactions are being processed correctly.
    - You can use the Nuchain block explorer or logs to track the transaction status.

## License

This project is licensed under the MIT `License`. See  for details.
