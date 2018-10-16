/**
 * n-1 elemanlı tekrar etmeyen pozitif tamsayı dizisinden eksik elemanı getirir
 *
 * Fonksiyonun zaman zorluğu => O(n)
 * ---------------------------------
 * Gauss Yöntemi: n * (n + 1) / 2 => O(1)
 * Array.prototype.reduce() => O(n)
 *
 * @param {Number[]} integerArray n-1 tekrar etmeyen tam sayı içeren dizi
 * @param {Number} n İsteğe bağlı değişken
 * @return {Number} Eksik tamsayı değeri
 */
const findMissingInteger = (integerArray, n = integerArray.length + 1) =>
  (n * (n + 1)) / 2 - integerArray.reduce((total, num) => total + num);

console.assert(findMissingInteger([1, 2, 4, 5, 6]) === 3);
console.assert(findMissingInteger([1, 2, 4, 5, 6], 6) === 3);
