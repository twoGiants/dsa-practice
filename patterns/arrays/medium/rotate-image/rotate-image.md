# Rotate Image

Given a square `n x n` matrix of integers `matrix`, rotate it by 90 degrees *clockwise*.

You must rotate the matrix *in-place*. Do not allocate another 2D matrix and do the rotation.

## Example 1:

![image](mat1.avif)

```ts
Input: matrix = [
  [1,2],
  [3,4],
]

Output: [
  [3,1],
  [4,2],
]
```

## Example 2:

![image](mat2.jpg)

```ts
Input: matrix = [
  [1,2,3],
  [4,5,6],
  [7,8,9],
]

Output: [
  [7,4,1],
  [8,5,2],
  [9,6,3],
]
```

## Example 3:

  ![image](mat3.jpg)

```ts
Input: matrix = [
  [5,1,9,11],
  [2,4,8,10],
  [13,3,6,7],
  [15,14,12,16],
]

Output: [
  [15,13,2,5],
  [14,3,4,1],
  [12,6,8,9],
  [16,7,10,11],
]
```

## Constrains

- `n == matrix.length == matrix[i].length`
- `1 <= n <= 20`
- `-1000 <= matrix[i][j] <= 1000`

