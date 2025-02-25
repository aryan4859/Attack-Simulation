<?php
if(isset($_POST['submit'])) {
    $target_dir = "uploads/";
    $target_file = $target_dir . basename($_FILES["cv"]["name"]);
    $file_type = strtolower(pathinfo($target_file, PATHINFO_EXTENSION));

    // Validate file type (reject PHP extensions, allow everything else)
    $restricted_extensions = ['php'];

    if (in_array($file_type, $restricted_extensions)) {
        echo "Sorry, PHP files are not allowed to upload.";
    } else {
        // Move the uploaded file to the "uploads" directory
        if (move_uploaded_file($_FILES["cv"]["tmp_name"], $target_file)) {
            echo "The file " . basename($_FILES["cv"]["name"]) . " has been uploaded.";
        } else {
            echo "Sorry, there was an error uploading your file.";
        }
    }
}
?>
