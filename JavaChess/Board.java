package JavaChess;

/**
 * Board
 */
public class Board {

    // Instance Variables

    private String[][] board = new String[7][7];


    public Board() {
        this.board = new String[][]{{"BR", "BN", "BB", "BQ", "BK", "BB", "BN", "BR"},
                                    {"BP", "BP", "BP", "BP", "BP", "BP", "BP", "BP"},
                                    {"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
                                    {"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
                                    {"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
                                    {"XX", "XX", "XX", "XX", "XX", "XX", "XX", "XX"},
                                    {"WP", "WP", "WP", "WP", "WP", "WP", "WP", "WP"},
                                    {"WR", "WN", "WB", "WQ", "WK", "WB", "WN", "WR"}};
    }

    // Getters

    public String[][] getBoard() {
        return this.board;
    }

    // Setters

    public void setBoard(String[][] board) {
        this.board = board;
    }

}