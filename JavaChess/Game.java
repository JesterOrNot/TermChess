package JavaChess;

/**
 * Game
 */
public class Game {

    private int moveNumber;

    private boolean isWhiteTurn;

    private Board theBoard;

    public int getMoveNumber() {
        return this.moveNumber;
    }

    public boolean getIsWhiteTurn() {
        return this.isWhiteTurn;
    }

    public void setWhiteTurn(boolean isWhiteTurn) {
        this.isWhiteTurn = isWhiteTurn;
    }

    public void setMoveNumber(int moveNumber) {
        this.moveNumber = moveNumber;
    }
}