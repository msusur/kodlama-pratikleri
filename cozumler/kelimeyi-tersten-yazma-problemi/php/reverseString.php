<?php
	
	header('Content-Type: text/html; charset=utf-8');	
	
	$str = "TEST STRING KODLAMA PRATIKLERI";
	$reversedStr = Reverse($str);
	
	echo $reversedStr; // IRELKITARP AMALDOK GNIRTS TSET
	
	function Reverse($str) {
		
		$len = strlen($str);
		
		if ($len == 1) { // Recursion Base Case
			return $str;
		} else {
			
			$len--;
			$firstLetter = substr($str, 0, 1);
			$remaining = substr($str, 1, $len);
			return Reverse($remaining).$firstLetter;
			
		}
	}
?>