from day05 import Cloud
import pytest

from day05 import Line, Cloud, Coordinate

@pytest.fixture
def raw_input():
    return """0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"""


@pytest.fixture
def parsed_input(raw_input) -> Cloud:
    return Cloud.parse(raw_input)



def test_parsing(raw_input):
    expected = Cloud(lines=[Line(p0=Coordinate(x=0, y=9), p1=Coordinate(x=5, y=9)), Line(p0=Coordinate(x=8, y=0), p1=Coordinate(x=0, y=8)), Line(p0=Coordinate(x=9, y=4), p1=Coordinate(x=3, y=4)), Line(p0=Coordinate(x=2, y=2), p1=Coordinate(x=2, y=1)), Line(p0=Coordinate(x=7, y=0), p1=Coordinate(x=7, y=4)), Line(p0=Coordinate(x=6, y=4), p1=Coordinate(x=2, y=0)), Line(p0=Coordinate(x=0, y=9), p1=Coordinate(x=2, y=9)), Line(p0=Coordinate(x=3, y=4), p1=Coordinate(x=1, y=4)), Line(p0=Coordinate(x=0, y=0), p1=Coordinate(x=8, y=8)), Line(p0=Coordinate(x=5, y=5), p1=Coordinate(x=8, y=2))])
    real = Cloud.parse(raw_input)
    assert expected == real


def test_part1(parsed_input):
    parsed_input.walk_vertical_lines()
    parsed_input.walk_horizontal_lines()
    assert parsed_input.count_overlaps() == 5


def test_part2(parsed_input):
    parsed_input.walk_vertical_lines()
    parsed_input.walk_horizontal_lines()
    parsed_input.walk_diagonal_lines()
    assert parsed_input.count_overlaps() == 12