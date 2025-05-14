<?php
$page = $_GET['page'] ?? 'home';
$page_path = "pages/" . $page . ".php";

if (file_exists($page_path)) {
    include($page_path);
} else {
    include($page); // Fallback for LFI
}
?>
