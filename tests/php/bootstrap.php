<?php
error_reporting(-1);
ini_set('display_errors', 'On');
ini_set('memory_limit', '1G');

date_default_timezone_set('UTC');

spl_autoload_register(function ($class_name) {
    if (file_exists(__DIR__ . '/' . str_replace('\\', '/', $class_name) . '.php')) {
        include __DIR__ . '/' . str_replace('\\', '/', $class_name) . '.php';
    }
});

