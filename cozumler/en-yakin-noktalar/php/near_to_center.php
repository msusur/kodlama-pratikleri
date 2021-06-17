<?php

/**
 * @param $a
 * @param $b
 * @return float
 */
function hypotenuse($a, $b)
{
    return sqrt(pow($a, 2) + pow($b, 2));
}

/**
 * @param $array
 * @return mixed
 */
function sortingNearToZero($array)
{
    for ($i = 0; $i < count($array); $i++) {
        $ov = $array[$i];
        $v = hypotenuse($array[$i][0], $array[$i][1]);
        $j = $i - 1;

        while ($j >= 0 && hypotenuse($array[$j][0], $array[$j][1]) > $v) {
            $array[$j + 1] = $array[$j];
            $j--;
        }

        $array[$j + 1] = $ov;
    }

    return $array;
}

/**
 * @param $array
 * @return mixed
 */
function nearToZero($array)
{
    return sortingNearToZero($array);
}

$points = [
    [2, 2],
    [2, 1],
    [3, 1],
    [4, 1],
    [3, 2],
    [3, 3]
];

return var_dump(nearToZero($points));