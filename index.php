#!/usr/bin/env php
<?php

declare(strict_types=1);

array_shift($argv);
$image = array_shift($argv);

if (!$image) {
    fwrite(STDERR, 'No arguments specified.');
    die(1);
}
if (!file_exists($image)) {
    fwrite(STDERR, 'File not found.');
    die(1);
}

$mime = mime_content_type($image);

$supported = ['image/svg', 'image/svg+xml', 'image/gif', 'image/jpeg', 'image/png'];
if (!in_array($mime, $supported)) {
    fwrite(STDERR, 'Image type not supported.');
    die(1);
}

if ($mime === 'image/svg') {
    $mime .= '+xml';
}

$data = file_get_contents($image);

die("data:${mime};base64," . base64_encode($data));
