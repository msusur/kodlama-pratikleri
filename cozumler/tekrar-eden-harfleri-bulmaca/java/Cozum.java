import java.util.HashMap;
import java.util.Map;

public class Cozum {

    public static void main(String[] args) {

        String sentence = "Bir cümle içerisinde tekrar eden harfleri bulmaca";
        Map<Character, Integer> charMap = new HashMap<>();

        // Karmaşıklık: O(n)
        for (int i = 0; i < sentence.length(); i++) {
            char ch = Character.toLowerCase(sentence.charAt(i));
            if (ch != ' ') {
                if (charMap.containsKey(ch)) {
                    charMap.put(ch, charMap.get(ch) + 1);
                } else {
                    charMap.put(ch, 1);
                }
            }
        }
        // Çıktı: [a=4, b=2, c=2, d=2, e=7, f=1, ç=1, h=1, i=5, k=1, l=3, m=2, n=2, r=6, s=1, t=1, u=1, ü=1]
        System.out.println(charMap.entrySet());
    }
}
