import java.util.Arrays;
import java.util.TreeSet;

public class SolutionWithHelper {

public int firstMissingPositive(int[] nums) {
    int posIdx = segregateNums(nums);
    int n = nums.length, val = 0;
    
    for(int i = 0; i < n-posIdx; i++) {
        val = Math.abs(nums[posIdx+i]);
        if(val <= n-posIdx && nums[posIdx+val-1] > 0) {
            nums[posIdx+val-1] = -nums[posIdx+val-1];
        }
    }
    
    for(int i = posIdx; i < n; i++) {
        if(nums[i] > 0)
            return i-posIdx+1;
    }
    return n-posIdx+1;
}

private static int segregateNums(int[] nums) {
    int i = 0, temp = 0;
    for(int j = 0; j < nums.length; j++) {
        if(nums[j] <= 0) {
            temp = nums[i];
            nums[i++] = nums[j];
            nums[j] = temp;
        }
    }
    return i;
}
}
