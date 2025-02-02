# Valid Parentheses

You are given a string `s` consisting of the following characters: `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`.

The input string `s` is valid if and only if:

    Every open bracket is closed by the same type of close bracket.
    Open brackets are closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

Return `true` if `s` is a valid string, and `false` otherwise.

## Example 1:

```ts
Input: s = "[]"

Output: true
```

## Example 2:

```ts
Input: s = "([{}])"

Output: true
```

## Example 3:

```ts
Input: s = "[(])"

Output: false
```

Explanation: The brackets are not closed in the correct order.

## Constrains

- `1 <= s.length <= 1000`

