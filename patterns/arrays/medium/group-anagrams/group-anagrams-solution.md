# Group Anagrams Solution

Use a hash map with count of characters in a string as key and the string itself as value:

- create hash map
- iterate over each string in the array 
- for each string create an array where you will store the count of each character in that string
  - for each character in the string increment the counter in the array
- when done return the values from the hash map in an array

Time: O(m*n), Space: O(m)