# Find All Duplicates in an Array Solution

Use the numbers in the array as "index - 1" to go to that place in that array and multiply the number with -1, then you know you've been there. Then while iterating through the array if you encounter a negative number, you know you visited that number and you can save that (index+1) in the result.

Time: O(n), Space: O(1)