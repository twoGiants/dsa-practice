# Group Anagrams

Given an array of strings `strs`, group all *anagrams* together into sublists. You may return the output in **any order**.

An **anagram** is a string that contains the exact same characters as another string, but the order of the characters can be different.

## Example 1:

```ts
Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
```

## Example 2:

```ts
Input: strs = [""]

Output: [[""]]
```

## Constrains

- `1 <= strs.length <= 1000`
- `0 <= strs[i].length <= 100`
- `strs[i]` is made up of lowercase English letters.


## Solution

[Group anagrams solution.](group-anagrams-solution.md)

## Retention Tracking

- [x] retained 17.02.25