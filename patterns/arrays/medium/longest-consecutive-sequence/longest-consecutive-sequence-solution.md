# Longest Consecutive Sequence Solution

Draw the numbers in sequence on an x-axis. Then it becomes clearer.

- create a hash set with the numbers which you will iterate through
- create a variable to store the longest sequence
- for every number check if there is a num-1 in the hash set
  - if so, then move on, because it's not the start of a sequence
  - if not, you found the beginning of a sequence
- now init the sequence length with 1 and check in the hash set for the next numbers num+length
  - if found, then increase the length
  - if not, sequence is finished, break the loop
- update the variable

Time: O(n), Space: O(n)