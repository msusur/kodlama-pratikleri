class WriteSquare {
    public static void main(String[] args) {
        System.out.println(repeatChar(repeatChar("*", 10) + "\n", 10));
    }

    private static String repeatChar(String character, int times) {
        return new String(new char[times]).replace("\0", character);
    }
}