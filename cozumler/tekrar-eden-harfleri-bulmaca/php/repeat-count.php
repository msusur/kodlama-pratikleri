<?php
$string      = "Bu cumlede tekrar eden karakterler ve sayilari ekrana basilacaktir.";
$stringCount = [];
for($i = 0; $i < strlen($string); $i++)
{
    $key = $string[$i];
    if(!array_key_exists($key, $stringCount))
    {
        $stringCount[$string[$i]] = 1;
    }
    else
    {
        $stringCount[$string[$i]] = $stringCount[$string[$i]]+1;
    }
}

// Cikti
//Array
//(
//    [B] => 1
//    [u] => 2
//    [ ] => 8
//    [c] => 2
//    [m] => 1
//    [l] => 4
//    [e] => 9
//    [d] => 2
//    [t] => 3
//    [k] => 5
//    [r] => 8
//    [a] => 10
//    [n] => 2
//    [v] => 1
//    [s] => 2
//    [y] => 1
//    [i] => 4
//    [b] => 1
//    [.] => 1
//)
print_r($stringCount);