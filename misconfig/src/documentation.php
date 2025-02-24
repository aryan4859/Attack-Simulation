<?php
if (isset($_GET['file'])) {
    $file = $_GET['file'];

    // Clean the input (basename ensures no path traversal)
    $safe_file = basename($file);

    // Define the path to the directory where documents are stored
    $dir = '/var/www/html/docs/';

    // Ensure the file exists in that directory
    if (file_exists($dir . $safe_file)) {
        include($dir . $safe_file);
    } else {
        echo "<p>File not found!</p>";
    }
} else {
    echo "<h1>Documentation Page</h1>";
    echo "<p>Welcome to the documentation page. Please read the important information.</p>";
}
?>
