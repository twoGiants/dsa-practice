# Top K Frequent Elements Solution

The trick is to use an array where the index is the frequency and the value at that index is an array of all the numbers which appear that many times.

- use a hash map to count how often each number appears
- the trick: use a 2D array where the index is the frequency and the value at that index all the numbers with that frequency
- iterate the frequency array backwards and grab k numbers from it

Time: O(n), Space: O(n)