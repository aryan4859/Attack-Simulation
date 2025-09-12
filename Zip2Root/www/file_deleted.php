<meta http-equiv="Content-type" content="text/html; charset=utf-8">
<?php
    include("checklogin.php");

    $filename = addslashes($_GET['filename']);
    $filename2 = str_ireplace("./uploads/", "", $filename); // filename used for database deletion
    $filename = str_ireplace("..", "", stripcslashes($filename));

    if (!unlink($filename)) {
        echo $filename . " deletion failed";
    } else {
        $db->query("DELETE FROM user_info WHERE filename='$filename2'");
        echo "Deletion successful";
    }
?>
