package JavaChess;

public abstract class Piece {

    // Instance Variables

    private int value;

    private boolean isWhite;

    // Abstract Methods

    public abstract void move();

    // Constructors

    public Piece(int value, boolean isWhite) {
        this.value = value;
        this.setIsWhite(isWhite);
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
