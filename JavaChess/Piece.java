package JavaChess;

public abstract class Piece {

    // Instance Variables

    private int value;

    private boolean isWhite;

    private boolean isAlive;

    // Abstract Methods

    public abstract Board makeMove(Board theBoard);

    public boolean isAlive() {
        return isAlive;
    }

    public void setAlive(boolean isAlive) {
        this.isAlive = isAlive;
    }

    public abstract int[][] getLegalMoves(Board theBoard);

    // Constructors

    public Piece(int value, boolean isWhite) {
        this.value = value;
        this.setIsWhite(isWhite);
        this.setAlive(true);
    }

    // Getters

    public int getValue() {
        return this.value;
    }

    public boolean getIsWhite() {
        return isWhite;
    }

    // Setters

    public void setValue(int value) {
        this.value = value;
    }

    public void setIsWhite(boolean isWhite) {
        this.isWhite = isWhite;
    }
}
