prints the probability of similarity calculated using the Rabin Karp algorithm

The Rabin–Karp algorithm proceeds by computing, at each position of the text, the hash value of a string starting at that position with the same length as the pattern. If this hash value equals the hash value of the pattern, it performs a full comparison at that position.

The {{aka}} command is a simple tool around that algorithm to compare same name files in subdirectories.
Use the filename (without path) as input.
