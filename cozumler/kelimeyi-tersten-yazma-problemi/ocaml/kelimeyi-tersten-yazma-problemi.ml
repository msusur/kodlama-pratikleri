let rec reverse str =
  if String.length str > 0 
  then 
  reverse (String.sub str 1 (String.length str - 1))
    ^
  String.sub str 0 1
  else str

let () =
  let test_string = "TEST STRING KODLAMA PRATIKLERI" in
  Printf.printf "%s" (reverse test_string)
  (* çıktı : IRELKITARP AMALDOK GNIRTS TSET *)
