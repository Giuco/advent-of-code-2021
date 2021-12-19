from dataclasses import dataclass
from enum import Enum, auto
from typing import List

TEST_RAW = """forward 5
down 5
forward 8
up 3
down 8
forward 2"""


class Direction(Enum):
    FORWARD = auto()
    DOWN = auto()
    UP = auto()


@dataclass
class Command:
    direction: str
    units: int


@dataclass
class Position:
    depth: int = 0
    horizontal: int = 0
    aim: int = 0

    def move(self, command: Command):
        if command.direction == "forward":
            self.horizontal += command.units
        elif command.direction == "up":
            self.depth -= command.units
        elif command.direction == "down":
            self.depth += command.units
        else:
            raise ValueError("Invalid command")

    def move_aim(self, command: Command):
        if command.direction == 'forward':
            self.depth += self.aim * command.units
            self.horizontal += command.units
        elif command.direction == "up":
            self.aim -= command.units
        elif command.direction == "down":
            self.aim += command.units
        else:
            raise ValueError("Invalid command")

    def result(self):
        return self.depth * self.horizontal


def move_ship(commands: List[Command], position: Position = None, aim: bool = False) -> Position:
    if not position:
        position = Position()

    for command in commands:
        position.move_aim(command) if aim else position.move(command)

    return position


if __name__ == "__main__":
    # Test
    processed_input = [Command(x.split()[0], int(x.split()[1])) for x in TEST_RAW.splitlines()]
    assert move_ship(processed_input).result() == 150

    # Part 1
    with open("data/day02.txt") as f:
        processed_input = [Command(x.split()[0], int(x.split()[1])) for x in f.readlines()]

    final_pos = move_ship(processed_input)
    print("Part 1: ", final_pos, final_pos.result())

    # Part 2
    final_pos = move_ship(processed_input, aim=True)
    print("Part 2:", final_pos, final_pos.result())
