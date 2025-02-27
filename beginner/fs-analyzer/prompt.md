Certainly, here's another beginner-friendly Golang problem:

# Problem: Simple File Directory Analyzer

**Task**: Create a Go program that analyzes a given directory and provides statistics about the files it contains.

**Requirements**:

1. The program should accept a directory path as a command-line argument
2. For the specified directory (non-recursive, just the top level):
   - Count the total number of files
   - Calculate the total size of all files in bytes
   - Find the largest and smallest files
   - Group files by extension and count each type

**Example Output**:

```
Directory Analysis: /home/user/documents
Total files: 15
Total size: 1250345 bytes (1.19 MB)

Largest file: presentation.pdf (524288 bytes)
Smallest file: notes.txt (128 bytes)

File types:
  .txt: 5 files
  .pdf: 3 files
  .go: 4 files
  .jpg: 2 files
  [no extension]: 1 file
```

This problem tests your ability to:

- Work with the Go file system package
- Parse command-line arguments
- Handle file operations
- Process collections of data
- Format output appropriately
