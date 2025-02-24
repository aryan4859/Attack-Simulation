<?php
session_start();

// Hardcoded username and password
$valid_username = "admin";
$valid_password = "password123";

if (isset($_POST['username']) && isset($_POST['password'])) {
    if ($_POST['username'] === $valid_username && $_POST['password'] === $valid_password) {
        $_SESSION['username'] = $_POST['username'];
        header("Location: dashboard.php");
        exit;
    } else {
        $error_message = "Invalid credentials!";
    }
}

echo "<h1>Login to the Application</h1>";
if (isset($error_message)) {
    echo "<p style='color: red;'>$error_message</p>";
}

echo "<form method='POST'>
        <input type='text' name='username' placeholder='Username' required><br>
        <input type='password' name='password' placeholder='Password' required><br>
        <input type='submit' value='Login'>
      </form>";
?>
