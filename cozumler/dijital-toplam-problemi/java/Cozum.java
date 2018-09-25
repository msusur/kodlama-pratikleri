public class Cozum {
    public static void main(String[] args) {
        System.out.println(solver(5));
        System.out.println(solver(14));
        System.out.println(solver(18));
    }

    private static int solver(int a) {
        int b = Math.floorDiv((a-1), 9);
        return b == 0 ? a + 9 : (a - 9*b + 1) * (int)(Math.pow(10, b)) - 1;
    }
}