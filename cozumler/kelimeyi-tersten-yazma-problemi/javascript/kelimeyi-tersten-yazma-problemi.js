function reverse(str, index) {
  if ( str[index] ) {
    return reverse(str, index + 1) + str[index];
  } else {
    return "";
  }
}

const test_str = "TEST STRING KODLAMA PRATİKLERİ";
console.log( reverse(test_str , 0) );
// çıktı : İRELKİTARP AMALDOK GNIRTS TSET