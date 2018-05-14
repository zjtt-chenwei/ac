{{define "head"}}
<!DOCTYPE HTML>

<html lang="en">
<head>
 <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Style Css -->
    <link rel="stylesheet" href="../static/bootstrap/css/bootstrap.css" >
    <!-- <link href="../static/css/style3.css" rel="stylesheet" type="text/css">   -->
    <!-- Color Style -->
    <!-- <link href="../static/colors/color1.css" rel="stylesheet" type="text/css"> -->

    <script type="text/javascript" src="../static/bootstrap/js/bootstrap.js"></script>
    <script type="text/javascript" src="js/jquery-1.11.3.min.js"></script>
    
    <script type="text/javascript">
    
        $(document).ready(function(){
          $('ul.nav > li').click(function () {
            $('ul.nav > li').removeClass('active');
            $(this).addClass('active');
          });
        });
       
    </script>
    <style type="text/css">
        html, body {width:100%;height:100%;} /*非常重要的样式让背景图片100%适应整个屏幕*/ 
        .navbar{
            font-family:微软雅黑; 
            margin:0px;
            padding:0px;
        }
        /* .navbar .nav li a{
            color: #000000;
        } */
        .navbar .nav li a:active, .navbar .nav li a:hover {
          color:#000000;
          background-color: #ff8000 !important;
        }
        
        
    </style>
    

{{end}}