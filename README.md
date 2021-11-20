# files
Convenience file functions for Go.

## What
This is a small collection of file functions to make development simpler.

- NewCloser() creates a deferrable Closer with no return value, which holds a list of files to close. Any errors from each close operation will be printed to stderr with a timestamp.
- LoadJSON() is a convenience function to load JSON from a file.
- SaveJSON() is the saving counterpart to LoadJSON() which uses a Closer for consistent error output.
- Exosts() tests if a file or directory exists.
- DirExists() tests if a directory exists, and returns false if the path is a file.
- FileExists() tests if a file exists, and returns false if the path is a directory.
- WriteFile() writes a file to disk, logging both open and close errors to stderr with a timestamp. There is no return value.
