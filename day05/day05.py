from dataclasses import dataclass, field
from typing import Coroutine, List


@dataclass
class Coordinate:
    x: int
    y: int

    @staticmethod
    def parse(raw: str) -> 'Coordinate':
        x, y = raw.split(",")
        return Coordinate(int(x), int(y))


@dataclass
class Line:
    p0: Coordinate
    p1: Coordinate

    def is_vertical(self) -> bool:
        return self.p0.x == self.p1.x

    def is_horizontal(self) -> bool:
        return self.p0.y == self.p1.y

    def walk(self) -> List[Coordinate]:
        miny = min(self.p0.y, self.p1.y)
        maxy = max(self.p0.y, self.p1.y)
        minx = min(self.p0.x, self.p1.x)
        maxx = max(self.p0.x, self.p1.x)
        
        if self.is_vertical():
            return [Coordinate(self.p0.x, y) for y in range(miny, maxy+1)]        
        elif self.is_horizontal():
            return [Coordinate(x, self.p0.y) for x in range(minx, maxx+1)]
        else:
            left_to_right = 1 if self.p1.y > self.p0.y else -1
            up_to_down = 1 if self.p1.x > self.p0.x else -1
            dist = abs(self.p1.y - self.p0.y)
            return [Coordinate(self.p0.x + i * up_to_down, self.p0.y + i * left_to_right) for i in range(dist+1)]

    @staticmethod
    def parse(raw: str) -> 'Line':
        p0, p1 = raw.split(" -> ")
        return Line(Coordinate.parse(p0), Coordinate.parse(p1))

@dataclass
class Cloud:
    lines: List[Coordinate]
    diagram: List[List[int]] = field(init=False)
    
    def __post_init__(self):
        self.max_coords = self.get_max_coordinates()
        self.diagram = [[0]*(self.max_coords.y+1) for _ in range(self.max_coords.x+1)]

    @staticmethod
    def parse(raw: str) -> 'Cloud':
        return Cloud(lines=[Line.parse(l) for l in raw.splitlines()])
    
    def get_max_coordinates(self) -> Coordinate:
        max_x, max_y = 0, 0

        for l in self.lines:
            max_x = max(max_x, l.p0.x, l.p1.x)
            max_y = max(max_y, l.p0.y, l.p1.y)

        return Coordinate(max_x, max_y)

    def walk_vertical_lines(self) -> None:
        lines = list(filter(lambda x: x.is_vertical(), self.lines))
        for l in lines:
            for c in l.walk():
                self.diagram[c.x][c.y] += 1

    def walk_horizontal_lines(self) -> None:
        lines = list(filter(lambda x: x.is_horizontal(), self.lines))
        for l in lines:
            for c in l.walk():
                self.diagram[c.x][c.y] += 1

    def walk_diagonal_lines(self) -> None:
        lines = list(filter(lambda x: not (x.is_horizontal() or x.is_vertical()), self.lines))
        for l in lines:
            print(l)
            for c in l.walk():
                print(c)
                self.diagram[c.x][c.y] += 1
            print("----------------")


    def count_overlaps(self) -> int:
        count = 0
        for row in self.diagram:
            for cell in row:
                if cell > 1:
                    count += 1
        return count


if __name__ == "__main__":
    with open("../data/day05.txt") as f:
        raw_input = f.read()
    cloud = Cloud.parse(raw_input)
    cloud.walk_vertical_lines()
    cloud.walk_horizontal_lines()
    overlap_count = cloud.count_overlaps()
    print(f"Part1 - Overlap Count: {overlap_count}")
    cloud.walk_diagonal_lines()
    overlap_count = cloud.count_overlaps()
    print(f"Part2 - Overlap Count: {overlap_count}")