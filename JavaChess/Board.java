package JavaChess;

/**
 * Board
 */
public class Board {

    // Instance Variables

    private Piece[][] board = new Piece[7][7];

    public Board() {
        this.board = new Piece[][] {
                    {
                        new Rook(5, false), new Knight(3, false), new Bishop(3, false), new Queen(9, false),
                        new King(-1, false), new Bishop(3, false), new Knight(3, false), new Rook(5, false)
                    },
                    {
                        new Pawn(1, false), new Pawn(1, false), new Pawn(1, false), new Pawn(1, false),
                        new Pawn(1, false), new Pawn(1, false), new Pawn(1, false), new Pawn(1, false)
                    },
                    { null, null, null, null, null, null, null, null }, 
                    { null, null, null, null, null, null, null, null },
                    { null, null, null, null, null, null, null, null }, 
                    { null, null, null, null, null, null, null, null },
                    { 
                        new Pawn(1, true), new Pawn(1, true), new Pawn(1, true), new Pawn(1, true),
                        new Pawn(1, true), new Pawn(1, true), new Pawn(1, true), new Pawn(1, true)
                    },
                    { 
                        new Rook(5, true), new Knight(3, true), new Bishop(3, true), new Queen(9, true),
                        new King(-1, true), new Bishop(3, true), new Knight(3, true), new Rook(5, true)
                    }
                };
    }

    // Getters

    public Piece[][] getBoard() {
        return this.board;
    }

    // Setters

    public void setBoard(Piece[][] board) {
        this.board = board;
    }

}