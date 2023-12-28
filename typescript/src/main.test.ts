import testCases from "../../test/test_cases.json";
import { Calc } from "./main";

Object.keys(testCases).forEach((key) => {
  testCases[key as keyof typeof testCases].forEach((testCase) => {
    const { a, b, ans } = testCase;
    test(`[${key}] ${a} 見せ ${b} = ${ans}`, () => {
      expect(Calc(a, b)).toBe(ans);
      expect(Calc(b, a)).toBe(ans);
    });
  });
});

test("a and b must be number", () => {
  expect(() => Calc("a" as any, 1)).toThrow("a and b must be number");
  expect(() => Calc(1, "a" as any)).toThrow("a and b must be number");
});
