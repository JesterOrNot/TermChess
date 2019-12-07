import colorama


def get_move():
    current_x = int(input("What is the current x position?: "))
    current_y = int(input("What is the current y position?: "))
    target_x = int(input("What is the target x position"))
    target_y = int(input("What is the target y position"))
    return current_x-1, current_y-1, target_x-1, target_y-1


def pretty_print(board):
    print("   1  2  3  4  5  6  7  8")
    print(" ╭━━━━━━━━━━━━━━━━━━━━━╮")
    count = 1
    count1 = 0
    for colum in range(7):
        count1 = 0
        print(str(count) + "│", end="")
        count += 1
        for row in range(7):
            if board[colum][row] == "XX":
                print(" x ", end="")
            elif board[colum][row] == "BR":
                print(colorama.Fore.RED + " R " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BN":
                print(colorama.Fore.RED + " N " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BB":
                print(colorama.Fore.RED + " B " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BK":
                print(colorama.Fore.RED + " K " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BQ":
                print(colorama.Fore.RED + " Q " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "BP":
                print(colorama.Fore.RED + " P " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WP":
                print(colorama.Fore.CYAN + " P " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WR":
                print(colorama.Fore.CYAN + " R " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WN":
                print(colorama.Fore.CYAN + " N " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WB":
                print(colorama.Fore.CYAN + " B " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WK":
                print(colorama.Fore.CYAN + " K " +
                      colorama.Style.RESET_ALL, end="")
            elif board[colum][row] == "WQ":
                print(colorama.Fore.CYAN + " Q " +
                      colorama.Style.RESET_ALL, end="")
            if count1 == 6:
                print("│ " + str(count-1), end="")
                print()
            count1 += 1
    print(" ╰━━━━━━━━━━━━━━━━━━━━━╯")

def new_board():
    return [["BR", "BN", "BB", "BQ", "BK", "BB", "BN", "BR"],
            ["BP", "BP", "BP", "BP", "BP", "BP", "BP", "BP"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"],
            ["WP", "WP", "WP", "WP", "WP", "WP", "WP", "WP"],
            ["WR", "WN", "WB", "WQ", "WK", "WB", "WN", "WR"]]


if __name__ == "__main__":
    print(pretty_print(new_board()))
