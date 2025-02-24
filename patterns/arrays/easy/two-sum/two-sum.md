# Two Sum

Given an array of unordered integers `nums` and an integer `target`, return the indices `i` and `j` such that `nums[i] + nums[j] == target` and `i != j`.

You may assume that every input has exactly one pair of indices `i` and `j` that satisfy the condition.

Return the answer with the smaller index first. 

## Example 1:

```ts
Input: nums = [3,4,5,6], target = 7

Output: [0,1]
```

## Example 2:

```ts
Input: nums = [4,5,6], target = 10

Output: [0,2]
```

## Example 3:

```ts
Input: nums = [5,5], target = 10

Output: [0,1]
```

## Constrains

- `2 <= nums.length <= 1000`
- `-10,000,000 <= nums[i] <= 10,000,000`
- `-10,000,000 <= target <= 10,000,000`

## Retention Tracking

- [x] retained 24.02.25