import java.util.Arrays;
import java.util.TreeSet;

public class Solution {

    public int firstMissingPositive(int[] nums) {
        TreeSet<Integer> numSet = new TreeSet<>();

        Arrays.stream(nums).forEach(n -> numSet.add(n));
        int i = 1;
        while (!numSet.isEmpty()) {
            int current = numSet.pollFirst();
            if (current > i) {
                return i;
            }
            if (current == i) {
                i++;
            }
        }
        return i;
    }

}
