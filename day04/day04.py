from typing import Tuple, List

RAW_INPUT_TEST = """7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,1,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
 """

Board = List[List[int]]


def parse_input(raw_input: str) -> Tuple[List[int], List[Board]]:
    sections = raw_input.split("\n\n")
    draws = [int(x) for x in sections[0].split(",")]
    boards = sections[1:]

    parsed_boards = []
    for board in boards:
        parsed_board = []
        for line in board.strip().split("\n"):
            parsed_line = [int(x) for x in line.strip().split()]
            parsed_board.append(parsed_line)
        parsed_boards.append(parsed_board)
    return draws, parsed_boards


def mark_board(board: Board, num: int) -> Board:
    for x in range(len(board)):
        for y in range(len(board)):
            if board[x][y] == num:
                board[x][y] = -1
    return board


def is_board_complete(board: Board) -> bool:
    m = len(board)
    rows = [0] * m
    cols = [0] * m
    for x in range(m):
        for y in range(m):
            if board[x][y] == -1:
                rows[x] += 1
                cols[y] += 1

    return any(x == m for x in rows) or any(y == m for y in cols)


def run_bingo(boards: List[Board], draws: List[int]) -> Tuple[Board, int]:
    for num in draws:
        for i in range(len(boards)):
            boards[i] = mark_board(boards[i], num)
            if is_board_complete(boards[i]):
                return boards[i], num
    raise ValueError("No winnable board")


def get_last_board(boards: List[Board], draws: List[int]) -> Tuple[Board, int]:
    for num in draws:
        for i in range(len(boards)):
            boards[i] = mark_board(boards[i], num)
            if is_board_complete(boards[i]) and len(boards) == 1:
                return boards[i], num

        boards = list(filter(lambda x: not is_board_complete(x), boards))

    raise ValueError("No solution")


def count_points(board: Board) -> int:
    count = 0
    m = len(board)
    for x in range(m):
        for y in range(m):
            num = board[x][y]
            if num != -1:
                count += num
    return count


if __name__ == "__main__":
    # Test Case 1
    draws, boards = parse_input(RAW_INPUT_TEST)
    board, last_draw = run_bingo(boards, draws)
    print("Test case", count_points(board) * last_draw, 4512)

    # Part 1
    with open("data/day04.txt") as f:
        RAW_INPUT = f.read()
    draws, boards = parse_input(RAW_INPUT)
    board, last_draw = run_bingo(boards, draws)
    print("Part 1", count_points(board) * last_draw)

    # Test Case 2
    draws, boards = parse_input(RAW_INPUT_TEST)
    board, last_draw = get_last_board(boards, draws)
    print("Test case", count_points(board) * last_draw, 1924)

    # Part 2
    draws, boards = parse_input(RAW_INPUT)
    board, last_draw = get_last_board(boards, draws)
    print("Part 2", count_points(board) * last_draw)
