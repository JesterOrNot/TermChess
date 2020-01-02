import colorama


def get_move():
    current_x = int(input("What is the current x position?: "))
    current_y = int(input("What is the current y position?: "))
    target_x = int(input("What is the target x position?: "))
    target_y = int(input("What is the target y position?: "))
    return current_y-1, current_x-1, target_y-1, target_x-1


def pretty_print(board: list):
    print("   1  2  3  4  5  6  7  8")
    print(" ╭━━━━━━━━━━━━━━━━━━━━━━━━╮")
    count = 1
    count1 = 0
    for colum in range(8):
        count1 = 0
        print(str(count) + "│", end="")
        count += 1
        for row in range(8):
            if board[colum][row] == "XX":
                print(" x ", end="")
            elif board[colum][row] == "BR":
                print(colorama.Fore.RED + " ♜ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BN":
                print(colorama.Fore.RED + " ♞ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BB":
                print(colorama.Fore.RED + " ♝ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BK":
                print(colorama.Fore.RED + " ♚ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BQ":
                print(colorama.Fore.RED + " ♛ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BP":
                print(colorama.Fore.RED + " ♙ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WP":
                print(colorama.Fore.CYAN + " ♙ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WR":
                print(colorama.Fore.CYAN + " ♖ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WN":
                print(colorama.Fore.CYAN + " ♘ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WB":
                print(colorama.Fore.CYAN + " ♗ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WK":
                print(colorama.Fore.CYAN + " ♔ " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WQ":
                print(colorama.Fore.CYAN + " ♕ " +
                      colorama.Style.RESET_ALL, end="")
            if count1 == 7:
                print("│ " + str(count-1), end="")
                print()
            count1 += 1
    print(" ╰━━━━━━━━━━━━━━━━━━━━━━━━╯")
    print("   1  2  3  4  5  6  7  8")


def new_board():
    return [["BR", "BN", "BB", "BQ", "BK", "BB", "BN", "BR"],
            ["BP", "BP", "BP", "BP", "BP", "BP", "BP", "BP"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["WP", "WP", "WP", "WP", "WP", "WP", "WP", "WP"],
            ["WR", "WN", "WB", "WQ", "WK", "WB", "WN", "WR"]]


def get_bishop_moves(board, current_pos):
    available_moves = []
    count1 = 1
    for i in range(7):
        if current_pos[0]+i > 7 or current_pos[1]+i > 7:
            break
        elif board[current_pos[0]+i][current_pos[1]+i] != "XX":
            break
        else:
            available_moves += [current_pos[0] + i, current_pos[1] + i]
    for i in range(8):
        if current_pos[0]+i > 7 or current_pos[1]-i < 0:
            break
        elif board[current_pos[0]+i][current_pos[1]-i] != "XX":
            break
        else:
            available_moves += [current_pos[0] + i, current_pos[1] - i]
    for i in range(8):
        if current_pos[0]-i < 0 or current_pos[1]-i < 0:
            break
        elif board[current_pos[0]-i][current_pos[1]-i] != "XX":
            break
        else:
            available_moves += [current_pos[0] - i, current_pos[1] - i]
    for i in range(8):
        if current_pos[0]-i < 0 or current_pos[1]+i > 7:
            break
        elif board[current_pos[0]-i][current_pos[1]+i] != "XX":
            break
        else:
            available_moves += [current_pos[0] - i, current_pos[1] + i]
    return available_moves


def make_move(board):
    data = get_move()
    target = board[data[2]][data[3]]
    temp = board[data[0]][data[1]]
    if target == "XX" and temp == "WP":
        if data[2] == data[0]-1 and data[3] == data[1]:
            board[data[2]][data[3]] = temp
            board[data[0]][data[1]] = "XX"
        else:
            print("Invalid move")
    elif target == "XX" and temp == "BP":
        if data[2] == data[0]+1 and data[3] == data[1]:
            board[data[2]][data[3]] = temp
            board[data[0]][data[1]] = "XX"
    elif target == "XX" and (temp == "BK" or temp == "WK"):
        if data[2] == data[0]+1 or data[2] == data[0]+1 and data[3] == data[1]+1 or data[3] == data[1]+1 or data[3] == data[1]-1 or data[2] == data[0]-1 or data[2] == data[0]-1 and data[3] == data[1]+1 or data[2] == data[0]+1 and data[3] == data[1]-1 or data[2] == data[0]-1 and data[3] == data[1]-1:
            board[data[2]][data[3]] = temp
            board[data[0]][data[1]] = "XX"
        else:
            print("Invalid move")
    elif target == "XX" or target.startswith("B") and (temp == "BN" or temp == "WN"):
        if (data[2] == data[0]+2 and data[3] == data[1]+1) or (data[2] == data[0]+2 and data[3] == data[1]-1) or (data[2] == data[0]-2 and data[3] == data[1]-1) or (data[2] == data[0]-2 and data[3] == data[1]+1) or (data[2] == data[0]+1 and data[3] == data[1]-2) or (data[2] == data[0]-1 and data[3] == data[1]-2) or (data[2] == data[0]+1 and data[3] == data[1]+2) or (data[2] == data[0]-1 and data[3] == data[1]+2):
            board[data[2]][data[3]] = temp
            board[data[0]][data[1]] = "XX"
        else:
            print("Invalid move")
    elif temp == "BB" or temp == "WB":
        bishopMoves = get_bishop_moves(board, [data[0], data[1]])
        print(bishopMoves)
        if target in bishopMoves:
            board[data[2]][data[3]] = temp
            board[data[0]][data[1]] = "XX"
    else:
        print("Invalid move")
    return board


if __name__ == "__main__":
    board = new_board()
    # pretty_print(board)
    while True:
        pretty_print(board)
        board = make_move(board)
        print(get_bishop_moves(board, [5, 5]))
    # pretty_print(board)
