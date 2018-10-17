import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class EnYakinNoktalar {
   public static void main(String[] args) {
      Double[][] noktalar = new Double[][] { { 1.3, 2.1 }, { 3.9, 1.1 }, { 8.0, 9.0 }, { 2.0, 3.0 }, { 2.0, 2.0 } };

      List<Double[]> enYakinNoktalar = enYakinNoktaBul(noktalar, 2);
      List<Double[]> beklenenSonuc = Arrays.asList(noktalar[0], noktalar[2]);

      assert enYakinNoktalar.equals(beklenenSonuc);
   }

   public static List<Double[]> enYakinNoktaBul(Double[][] noktalar, int adet) {
      return Arrays.stream(noktalar).sorted((o1, o2) -> uzaklik(o1).compareTo(uzaklik(o2)))
            .collect(Collectors.toList()).subList(0, adet);
   }

   public static Double uzaklik(Double[] item) {
      return Math.sqrt(Math.pow(item[0], 2) + Math.pow(item[1], 2));
   }
}
