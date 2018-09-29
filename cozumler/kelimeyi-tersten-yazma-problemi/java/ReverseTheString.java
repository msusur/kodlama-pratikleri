package freePackage;

import org.apache.commons.lang3.ArrayUtils;

public class ReverseTheString {

    private static String reverseString = "";

    public static void main(String[] args) {
        String example = "Muvaffakiyetsizleştiricileştiriveremeyebileceklerimizdenmişsinizcesine";
        makeReverse(example.toCharArray());
        System.out.println(reverseString);
    }

    public static void makeReverse(char[] word) {
        if (word.length > 0) {
            reverseString += word[word.length - 1];
            word = ArrayUtils.remove(word, word.length - 1);
            makeReverse(word);
        }
    }
    
}
