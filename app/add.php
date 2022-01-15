<?php

function add(int $a, int $b) : int
{
    var_dump("i intentionally injected this var_dump");
    return $a + $b;
}
