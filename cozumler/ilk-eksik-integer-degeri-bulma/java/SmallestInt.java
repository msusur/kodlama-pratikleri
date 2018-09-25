public class SmallestInt {

    public static void main(String ... args) {
        System.out.println(solve(new int[]{1, 2, 0}));
        System.out.println(solve(new int[]{3, 4, -1, 1}));
        System.out.println(solve(new int[]{7, 8, 9, 10, 11}));
    }

    static int solve(int[] arr) {
        // extra 2 spaces; 1 for 0 index and the other for 
        // the pigeonhole principle.
        boolean[] flags = new boolean[arr.length+2];
        for (int i = 0; i<arr.length; i++) {
            if (arr[i] > 0 && arr[i] < flags.length) flags[arr[i]] = true;
        }
        for (int i = 1; i<flags.length; i++) {
            if (!flags[i]) return i;
        }

        // not possible
        return -1;
    }
}

