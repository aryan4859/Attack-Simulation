<meta http-equiv="Content-type" content="text/html; charset=utf-8">
<?php
    include("checklogin.php");

    $filename = filter(@$_GET['filename']);
    $rel_name = "/uploads/" . $filename;
    $down_name = str_ireplace("+", ' ', "./uploads/" . urlencode($filename));

    if (file_exists(dirname(__FILE__) . $rel_name)) {
        $sql = "select message from user_info where filename='$filename'";
        //die($sql);
        $result = $db->query($sql);
        $row = $result->fetch_array();
        echo "<center><h1>$filename</h1></center>";

        echo "<center><label>Note:</label><textarea cols='60' role='20'>" . $row[0] . "</textarea></center><br>";
        echo "<center><a href='$down_name'>Download & View</a></center><br>";
        echo "<center><a href='file_deleted.php?filename=$down_name'>Delete file (if you got the flag, remember to delete the payload file or others may see it)</a></center><br>";
    } else {
        echo "If you want to view the assignment file, the file must exist: " . $rel_name . "  Are you sure it exists???";
    }
?>
