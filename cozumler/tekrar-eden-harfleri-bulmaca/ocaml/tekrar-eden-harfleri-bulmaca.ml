module SMap = Map.Make(String)

(* Tüm karakterleri sayar (boşluk ve noktalamalar dahil) . Büyük-Küçük harf ayrımı yapmaz. *)
let rec countChars str =
  if String.length str > 0
  then SMap.merge
    ( fun key a b -> match a, b with
    | (x, None  ) -> x
    | (None  , b) -> b
    | (Some x, Some y) -> Some (x + y)
    )
    ( countChars (String.sub str 1 (String.length str - 1)) )
    ( SMap.add (String.sub str 0 1) 1 SMap.empty )
  else SMap.empty

(* 1 adet olanları filtreler. *)
let countRepeatedChars str =
  SMap.filter (fun key value -> value > 1) (countChars str)

let () =
  let test = "Tekrar eden harfler sorusu icin test dizgesi." in
  SMap.iter
    (fun key count -> Printf.printf "%s : %d\n" key count)
    (countRepeatedChars test)
