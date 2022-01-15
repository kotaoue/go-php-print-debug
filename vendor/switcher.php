<?php

function convertNiwatori(string $emoji): string
{
    switch ($emoji) {
        case '🐔':
            return 'にわとり';
        case '🐓':
            return 'おんどり';
    }
    return 'にわとり以外';
}
