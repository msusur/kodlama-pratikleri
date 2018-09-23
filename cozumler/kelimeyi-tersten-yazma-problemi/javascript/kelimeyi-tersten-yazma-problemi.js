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


/* ALTERNATİF YÖNTEM */

function reverse2(str){
  return str.split("").reverse().join("");
}

const test_str2 = "Bu bir reverse test";
console.log(reverse2(test_str2));
// çıktı: tset esrever rib ub
