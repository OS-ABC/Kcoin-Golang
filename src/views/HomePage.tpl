<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<title>HomePage</title>
	<link rel="stylesheet" type="text/css" href="../static/css/homepage.css"/>
</head>

<body>
	<div class="header">
		<div class="head">
			<a>登录</a>
			<a href="#container">项目列表</a>
			<a>首页</a>
		</div>
		<div class="title">
			<div class="kcoin">KCOIN</div>
			<h2>基于区块链的项目激励平台</h2>
			<p>您可以在下面的搜索框里</p>
			<p>搜索您感兴趣的项目</p>
			<form action="">
				<input type=" text" class="search" placeholder="搜索">
				<input type="button" class="search_btn">
			</form>
		</div>
	</div>
	<div style="clear: both"></div>
	<div class="container" id="container">
		<div class="container-child">
            {{with .Projects}}
            {{range .Data}}
			<div class="project">
				<img class="project-cover" alt="project" src="{{.ProjectCoverUrl}}"/>
				<div class="project-name">{{.ProjectName}} </div>
				<div class="introduction">introduction</div>
				<div class="CC">CC:12138</div>
				<hr/>
                {{range .MemberList}}
				<img class="head_shot" alt="" src="{{.MemberHeadshotUrl}}"/>
                {{end}}
			</div>
            {{end}}
            {{end}}
		</div>
	</div>
	<div style="clear: both"></div>
	<div class="footer">
		<p>北京大学软件与微电子学院-新工科试验班</p>
	</div>
</body>
</html>
