###BigCache Example

##Overview

This project demonstrates a production-ready implementation of a caching system using the BigCache library in Go. It efficiently stores user data in memory, enabling fast reads and writes. The code is designed with robust error handling, structured logging, and graceful shutdown practices.

##Key Features

High-Performance Caching: Uses BigCache to handle large volumes of data with low latency.

Structured Logging: Logs critical operations (add, read, delete) to a file for monitoring.

Robust Error Handling: Gracefully handles cache creation errors and user operations.

Configuration Management: Settings are centralized and easy to adjust.

Graceful Shutdown: Ensures clean closure of resources during application termination.

##Getting Started

Prerequisites

Go 1.18 or higher

BigCache library

Installation

Clone the repository:

git clone https://github.com/yourusername/bigcache-example.git

Change to the project directory:

cd bigcache-example

Install dependencies:

go mod tidy

Running the Application

go run main.go

Log Output

Logs are stored in bigcache.log:

LOG: 2025/05/08 12:34:56 Application started
LOG: 2025/05/08 12:34:57 User updated in cache: 1
LOG: 2025/05/08 12:34:58 User retrieved from cache: 1
LOG: 2025/05/08 12:34:59 User deleted from cache: 1

Usage

The application demonstrates the following:

Adding/Updating a User:

Uses the update method to store user data.

Reading a User:

Uses the read method to retrieve data by ID.

Deleting a User:

Uses the delete method to remove a user from the cache.

Error Handling

If a user is not found, it returns a custom error: errUserNotInCache.

All operations log errors with detailed messages for debugging.

License

This project is licensed under the MIT License.

Contributions

Feel free to fork and submit pull requests for improvements or bug fixes.

Contact

For questions, please open an issue on GitHub or contact me at [your-email@example.com].
