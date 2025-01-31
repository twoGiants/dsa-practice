# Rotate Image Solution

Move each cell one by one to its desired position. The trick is to use pointers:

- create 4 pointers: left, right, top and bottom
- move cells one by one starting at the top-left position
- to avoid losing cells when overwriting create one extra space variable for top-left
- go backwards and copy the previous corner cell into the current corner cell e.g. bottom-left into top-left etc.
- add iteration counter to corners
- finally increase pointers
- iteration is over when left < right
- after rotating the outermost layer you need to rotate the next inner layer and so on
- each layer has the width: right - left

Time: O(n²), Space: O(1)