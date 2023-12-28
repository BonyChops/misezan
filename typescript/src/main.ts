const pattern = {
  "6": "9",
  "9": "6",
  "2": "5",
  "5": "2",
  ".": ".",
};

export function Calc(a: number, b: number): number {
  if (typeof a !== "number" || typeof b !== "number") {
    throw new Error("a and b must be number");
  }

  if (a === b) {
    return 0;
  }

  const [min, max] = [Math.min(a, b), Math.max(a, b)];

  if (max >= min * 100) {
    return max - min * 17;
  }

  const [minStr, maxStr] = [min.toString(), max.toString()];

  if (maxStr.length !== minStr.length) {
    return max;
  }
  const minStrArr = minStr.split("");

  const isReverse = minStrArr.every(
    (char, index) =>
      (pattern[char as keyof typeof pattern] ?? "") ===
      maxStr[maxStr.length - index - 1]
  );

  if (isReverse) {
    const isSpecial = minStrArr.includes("2") || minStrArr.includes("5");
    if (isSpecial) {
      return 1.1;
    } else {
      return 11;
    }
  }

  return max;
}
