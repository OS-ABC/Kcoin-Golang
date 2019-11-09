<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<link type="text/css" rel="stylesheet" href="../static/css/import.css">
	<title>我的项目</title>
</head>
<body>
	<div class="main">
	<h1>导入项目</h1>

		<ul class="steps">
		    <li class="active">选择项目</li>
		    <li class="active">初始信息</li>
		    <li>初始分配</li>
		</ul>
		<div class="content">
		<form>
		<input type="text" placeholder="快速导入: 请输入Github项目的URL">
		</form>

		<p>
			<img class = "profile" src="../static/img/tx2.png" alt="profile" align="top">&nbsp
			<span id="username">{{.user.Data.UserName}}</span>&nbsp
			<span id ="userprojects">(4个子项目)</span>
		</p>
		<ul class="projects">
			<li>-</li>
			<li>-Android-</li>
			<li>HelloWorld</li>
			<li>weatherForcast</li>
		</ul>
		</div>
	</div>
</body>
</html>