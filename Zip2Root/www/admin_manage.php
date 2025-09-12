<meta http-equiv="Content-type" content="text/html; charset=utf-8">
<?php

include("checklogin.php");
$sql = "select * from user_info";  
$result = $db->query($sql);  
  
if ($result)   
{  
    if ($result->num_rows > 0)  
    {   
        echo '<table type="text" align="center" border="1px,solid">';
		echo "<center><h1>All submitted assignments are listed here</h1></center>";
		echo "<tr><td>Name</td>";
        echo "<td>Student ID</td>";
        echo "<td>Assignment Document (Click to view)</td><BR></tr>";
        
        while ($rows = $result->fetch_array()) 
        { 
            echo "<tr><td>".$rows['username']."</td>";
            echo "<td>".$rows['ids']."</td>";
            echo "<td><a href='file_checked.php?filename=".urlencode($rows['filename'])."'>".$rows['filename']."</a></td>";
        }	
 		
        echo "<BR></table>";
    }
    else
    {  
        echo "<BR>No results found!";     
    }
}
else
{  
    echo "<BR>Query failed!";   
}
?>
