<?php  
	include("db_config.php");

	$name = @addslashes($_POST['name']);     // get name
	$ids =  @addslashes($_POST['ids']);      // get student ID
	$message = @addslashes($_POST['message']); // get message/note

	if (@$_FILES["file"]["error"] > 0) {
  		echo "Upload failed!";
  	} else {		
		$filename = @addslashes($_FILES['file']['name']);
    	$target_path  = "./uploads/"; 
    	
    	// Create uploads directory if it doesn't exist
    	if (!is_dir($target_path)){
        	mkdir($target_path);
    	}

    	$rel_file = rand(10,100) . basename($_FILES['file']['name']); 
    	$target_path .= $rel_file;

    	if(@$_POST['submit']) {
			// Only allow .doc files
			if(pathinfo($_FILES['file']['name'], PATHINFO_EXTENSION) != "doc") {
				echo "Only .doc files are allowed. Don't try to get a shell!";
			} else {
				// Move uploaded file
				if(!move_uploaded_file($_FILES['file']['tmp_name'], $target_path)) {
					echo "Unknown upload error";
				} else {
					$rel_file = addslashes($rel_file);
					echo "Submission successful, file path: " . $target_path;
				}

				// Insert submission info into database
				$sql = "INSERT INTO `user_info` (`username`, `ids`, `message`, `filename`) 
				        VALUES ('$name', '$ids', '$message', '$rel_file')";
				$db->query($sql);
			}
		}
  	}
?>
<html>
	<body>
		<meta http-equiv="Content-type" content="text/html; charset=utf-8">
		<center>
			<h1>Submit Assignment</h1>

			<form action="" method="post" enctype="multipart/form-data">
           		<table type="text" align="center" border="1px,solid">
           			<tr>
               			<td>Name</td>
               			<td><input type="text" name="name" id="name"/></td>
           			</tr>
            		<tr>
                		<td>Student ID</td>
                		<td><input type="text" name="ids" id="ids"/> </td>
            		</tr>
            		<tr>
                		<td>Note</td>
                		<td><textarea name="message" id="message" cols="60" role="15"></textarea></td>
           			</tr>
           			<tr>
                		<td>Assignment (Word files only)</td>
                		<td><input type="file" name="file" id="file" /></td>
           			</tr>
           			<tr>
                		<td colspan="2">
                			<input type="submit" name="submit" value="Submit" />
							<input type="reset" name="reset"/>
						</td>
           			</tr>
           		</table>
       		</form>
		</center>
	</body>
</html>
