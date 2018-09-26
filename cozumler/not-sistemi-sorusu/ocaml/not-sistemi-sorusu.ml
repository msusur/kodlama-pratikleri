(*
  fintIndex Verilen koşulu sağlayan ilk indeksi yolluyor.
  Native olarak OCaml'da olmadığı için yazmam gerekti.
*)
let rec findIndex conditionFn arr start =
  match (conditionFn (Array.get arr start)) with
    | true  -> start
    | false -> (findIndex conditionFn arr (start + 1))
 
let scale score =
  let grades = [| "FF"; "FD"; "DD"; "CC"; "CB"; "BB"; "BA"; "AA" |] in
  let scores = [|   29;   49;   59;   64;   69;   79;   89;  100 |] in
  findIndex (fun value -> score <= value) scores 0
  |> Array.get grades

let () =
  assert ( scale 29 = "FF" ) ;
  assert ( scale 45 = "FD" ) ;
  assert ( scale 55 = "DD" ) ;
  assert ( scale 62 = "CC" ) ;
  assert ( scale 65 = "CB" ) ;
  assert ( scale 75 = "BB" ) ;
  assert ( scale 85 = "BA" ) ;
  assert ( scale 90 = "AA" ) ;

