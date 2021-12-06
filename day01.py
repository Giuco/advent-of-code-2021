from typing import List

RAW_INPUT = """199
200
208
210
200
207
240
269
260
263"""


def depth_increase_measure(ys: List[int], window_size: int = 1) -> int:
    count = 0
    for i in range(1, len(ys)):
        if ys[i] > ys[i - window_size]:
            count += 1
    return count


if __name__ == '__main__':
    # Test
    processed_input = [int(x) for x in RAW_INPUT.splitlines()]
    assert depth_increase_measure(processed_input) == 7

    # Part 1
    with open("data/day01.txt") as f:
        processed_input = [int(x) for x in f.readlines()]
    print("Part 1: ", depth_increase_measure(processed_input))

    # Part 2
    print("Part 2: ", depth_increase_measure(processed_input, window_size=3))
