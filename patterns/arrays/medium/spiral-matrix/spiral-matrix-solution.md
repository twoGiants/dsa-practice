# Spiral Matrix Solution

Move from left to right, top to bottom, right to left and bottom to top and copy each cell into the result array. The trick is to use 4 pointers and update them:

- create 4 pointers: `left = 0, right = len(matrix[0]), top = 0 and bottom = len(matrix)`
- iterate as long as `left < right and top < bottom`
- go from left to right then cross out top row by incrementing top
- go from top to bottom then cross out right column by decreasing right
- check if `left < right and top < bottom` still holds
- go from right to left then cross out bottom row by decreasing bottom
- go from bottom to top then cross out left column by incrementing left
- don't forget that bottom and right are +1 bigger than the array index

Time: O(m*n), Space: O(1)