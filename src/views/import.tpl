<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<link type="text/css" rel="stylesheet" href="../static/css/import.css">
	<script type="text/javascript" src="../static/js/import.js"></script>
	<title>我的项目</title>
</head>
<body>
	<div class="main">
		<h1>导入项目</h1>
			<ul class="steps">
				<li id="step1" class="active">选择项目</li>
				<li id="step2">初始信息</li>
				<li id="step3">初始分配</li>
			</ul>
			<div class="content">
				<div id="myprojects">
					<form>
						<input class = "url_input" type="text" placeholder="快速导入: 请输入Github项目的URL">
						<input class = "submit" type="submit" value="确&nbsp认">
					</form>
					<div id = "projectslist">
						<p>
							<img class = "profile" src={{.user.Data.HeadShotUrl}} alt="profile" align="top">&nbsp
							<span id="username">{{.user.Data.UserName}}</span>&nbsp
							<span id ="userprojects">({{.memberList_len}}个项目)</span>
						</p>
						
						<ul class="projects">
							{{with .user.Data}}
							{{range .ProjectList}}
								<li>{{.ProjectName}}</li>
							{{end}}
							{{end}}
						</ul>
					</div>
				</div>
				<div id="init_message">
					<p>初始信息页面</p>
				</div>
				<div id="config_project">
					<p>&nbsp</p>
					<p>1. 填写项目名称</p>
					<form>
						<input class = "projectname" type="text">
					</form>
					<p>2. 上传项目封面（选填，但只能添加不大于2M的图片）</p>
					<a href="javascript:;" class="upload">选择文件
						<input class="change" id="upload_image" type="file" accept=".jpg, .png" onchange="showImg(this)"/>
					</a>
				</div>
			</div>
			<div id="next_back">
				<form>
					<button class = "next" id="btnNextStep" type="button" onclick="next_step()">下一步</button>
					<button class = "back" id="btnBackStep" type="button" onclick="back_step()">前一步</button>
				</form>
			</div>
	</div>
</body>
</html>