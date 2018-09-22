# Sierpinski Üçgeni

Zorluk: Zor

Sierpinski Üçgeni kendi kendini tekrar eden üçgenlerden oluşan bir fraktaldır. İsimini matematikçi Waclaw Sierpinski'den alır. Aşağıdaki resimde örnekleri görünmektedir.

![Alt text](../images/sierpinski.png)

Algoritma;
1. Eş kenar üçgen çiz.
2. Her kenarın orta noktalarını belirleyerek üçgenin içerisinde 4 tane daha eş kenar üçgen çiz.
3. Her küçük üçgen için ikinci adımı tekrarla.

Problem 1.
  Yukarıdaki algoritmayı çalıştırarak yakınlaş/uzaklaş özellikleri de olacak şekilde ekrana çizdiren uygulamayı yazdırın.

Problem 2.
  Yukarıdaki algoritmayı göz önüne alarak kaçıncı seviyede kaç tane üçgen olacağını ekrana yazdıran uygulamayı yazın.

Opsiyonel.
  Yukarıdaki algoritma ve problemleri en küçük kodu yazarak tamamlayın.

*Bu problem ARM'nin yazılımcı mülakatından alıntıdır*

[Çözüm](../cozumler/sierpinski-ucgeni)