<?php
  session_start();

  // Fake user dashboard
  if(!isset($_SESSION['username'])) {
    header("Location: index.php");
    exit;
  }

  echo "<h1>Welcome, " . htmlspecialchars($_SESSION['username']) . "!</h1>";
  echo "<p>This is your user dashboard.</p>";

  // Simulate a vulnerable link to documentation
  echo "<a href='documentation.php?file=docs/important_readme.txt'>Read Important Documentation</a>";

  // Logout link
  echo "<br><a href='index.php?action=logout'>Logout</a>";
?>
