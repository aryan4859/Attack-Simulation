<?php
  session_start();
  // Simple "login" (not functional but looks realistic)
  if(isset($_GET['action']) && $_GET['action'] == 'logout') {
      session_destroy();
      echo "You have been logged out.";
      exit;
  }



  // Fake authentication system
  if(isset($_POST['username']) && isset($_POST['password'])) {
      $_SESSION['username'] = $_POST['username'];
      header("Location: dashboard.php");
      exit;
  }

  echo "<h1>Welcome to the Secure Application</h1>";
  echo "<p>Please login to continue:</p>";
  echo "<form method='POST'>
          <input type='text' name='username' placeholder='Username' required><br>
          <input type='password' name='password' placeholder='Password' required><br>
          <input type='submit' value='Login'>
        </form>";
?>
