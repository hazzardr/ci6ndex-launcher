import { describe, it, expect } from "vitest";
import { useCounter } from "~/composables/useCounter";

describe("useCounter", () => {
  it("increments", () => {
    const { count, increment } = useCounter();
    expect(count.value).toBe(0);
    increment();
    expect(count.value).toBe(1);
  });

  it("doubles", () => {
    const { doubled, increment } = useCounter(3);
    expect(doubled.value).toBe(6);
    increment();
    expect(doubled.value).toBe(8);
  });
});
