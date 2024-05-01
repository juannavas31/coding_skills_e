# Rolling Hash Algorithm

This application implements a rolling hash based file diffing algorithm to compare an original and updated version of a file. It returns a list of _deltas_ describing: 
- The type of operation ( _Insert_, _Delete_, _Replace_ )
- The start position in the original file
- The end position in the original file
- The data changed (inserted data, deleted data or new data replacing the old ones)

The implementation is based on the [Rabin-Karp hash function](https://www.geeksforgeeks.org/rabin-karp-algorithm-for-pattern-searching/) which efficiently calculates the hash of an initial data window and subsequent hashes of the sliding window.

A high level description. A table of hashes is built for each file, each hash corresponding to a sliding window. They are traversed to find differences. When a difference is found, a delta is created and the search for a new delta is resumed from a point where the hashes are again equal. 

The real-world use case would be to compare files the way git does at pull requests, for instance.

# Usage

To build the applicaion, just use the `go build -o diffing` command

To run unit tests, just use the command `go test -v ./...`

To run the application, use the following command syntax

`./diffing --window=<n> <original_file> <updated_file>`

where: 

* _window=n_ optional (default=6). Specifies the window size, recommended value is the size of a word, for instance 6
* *original_file* is the base file to compare with
* *updated_file* is the modified file

For example: 

`./diffing --window=6 original.txt updated.txt`


Happy diffing!!
