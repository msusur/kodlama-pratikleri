let rec digital x digit=
  if x <= 9
  then if digit == 1 then x + 9 else x * digit
  else 9 * digit + (digital (x - 9) (digit*10))  

let () =
  Printf.printf "%d" (digital 21 1)
  (* çıktı : 399 *)
  