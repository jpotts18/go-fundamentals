Here's a different beginner-friendly Golang problem:

# Problem: HTTP Server Status Monitor

**Task**: Create a Go program that periodically checks if specified websites are online and reports their status.

**Requirements**:

1. The program should accept a list of URLs as command-line arguments
2. For each URL, make an HTTP GET request every 30 seconds
3. Track and display the following information for each site:
   - Current status (UP/DOWN)
   - Response time in milliseconds
   - HTTP status code
   - Time since last status change
4. The program should run continuously, updating the console display after each check

**Example Output**:

```
Website Status Monitor - Press Ctrl+C to exit
[2025-02-27 14:30:15]

https://example.com    | UP   | 200 | 87ms   | Uptime: 10m 15s
https://test.org       | DOWN | --- | Timeout| Downtime: 5m 30s
https://golang.org     | UP   | 200 | 132ms  | Uptime: 25m 45s
```

This problem tests your ability to:

- Make HTTP requests in Go
- Handle concurrency with goroutines
- Process command-line arguments
- Format and display time values
- Create a live-updating console application

This exercise involves core Go concepts like goroutines, channels, time handling, and network programming, making it a good test of fundamental skills without being overly complex.
