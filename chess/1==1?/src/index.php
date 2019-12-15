<?php

highlight_file(__FILE__);

include('flag.php');

if(isset($_GET['input']))
{
	$in = $_GET['input'];
	$ans = 1;

	if(preg_match('/^1(\.0*$|e0*)|^1$|true/is', (string)$in)) echo 'GG';
	else if($in == $ans) echo $flag;
	else echo 'GG';
}

?>