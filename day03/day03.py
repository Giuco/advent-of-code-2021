from typing import List, Tuple

TEST_INPUT = """00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010"""


def get_most_common(series: List[int]) -> int:
    size = len(series)
    half = size // 2 if size % 2 == 0 else size // 2 + 1
    return int(sum(series) >= half)


def get_series_pos(series: List[str], pos: int) -> List[int]:
    return [int(x[pos]) for x in series]


def get_gamma_and_epsilon_rate(bit_series: List[str]) -> Tuple[int, int]:
    binary_gamma = ""
    binary_epsilon = ""

    for pos in range(len(bit_series[0])):
        if get_most_common(get_series_pos(bit_series, pos)):
            binary_gamma += "1"
            binary_epsilon += "0"
        else:
            binary_gamma += "0"
            binary_epsilon += "1"

    return int(binary_gamma, 2), int(binary_epsilon, 2)


def get_oxygen_rate(bit_series: List[str]) -> int:
    while True:
        if len(bit_series) == 1:
            break
        for pos in range(len(bit_series[0])):
            if len(bit_series) == 1:
                break
            most_common = get_most_common(get_series_pos(bit_series, pos))
            bit_series = list(filter(lambda x: int(x[pos]) == most_common, bit_series))
    return int(bit_series[0], 2)


def get_co2_rate(bit_series: List[str]) -> int:
    while True:
        if len(bit_series) == 1:
            break
        for pos in range(len(bit_series[0])):
            if len(bit_series) == 1:
                break
            least_common = 1 - get_most_common(get_series_pos(bit_series, pos))
            bit_series = list(filter(lambda x: int(x[pos]) == least_common, bit_series))
    return int(bit_series[0], 2)


if __name__ == "__main__":
    # Test 1
    processed_test_input = TEST_INPUT.splitlines()
    assert get_gamma_and_epsilon_rate(processed_test_input) == (22, 9)

    # Part 1
    with open("data/day03.txt") as f:
        processed_input = f.read().splitlines()
    gamma, epsilon = get_gamma_and_epsilon_rate(processed_input)
    print(f"Part 1. Gamma={gamma}. Epsilon={epsilon}. Energy:{gamma * epsilon}")

    # Test 2
    oxygen, co2 = get_oxygen_rate(processed_test_input), get_co2_rate(processed_test_input)
    assert oxygen == 23, co2 == 10

    # Part 2
    oxygen, co2 = get_oxygen_rate(processed_input), get_co2_rate(processed_input)
    print(f"Part 2. Oxygen={oxygen}. CO2={co2}. LifeSupport={oxygen * co2}")
