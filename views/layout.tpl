<!DOCTYPE html>

<html>
  	<head>
		<title>HDMonitor ++</title>

		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

		<link rel="stylesheet" href="/static/css/bootstrap.min.css">
		<link rel="stylesheet" href="/static/css/global.css">

		<script src="/static/js/jquery.min.js"></script>
		<script src="/static/js/jquery.validate.min.js"></script>
		<script src="/static/js/bootstrap.min.js"></script>
		<script src="/static/js/global.js"></script>
		
		{{ str2html "<!--[if lt IE 9]>"}}
			<script src="/static/js/html5shiv.min.js"></script>
			<script src="/static/js/respond.min.js"></script>
		{{ str2html "<![endif]-->"}}

		{{.Heads}}

	</head>

  	<body>

		<header>
			<div class='navbar navbar-default navbar-static-top'>
				<div class='container-fluid'> 
					<div style='padding:0 12px;'>
						<div class='navbar-header'>
							<a class="navbar-brand" href="/">HDM++</a>
						</div>
						<ul class='nav navbar-nav'>
							<li class='active'><a href='/'`>概览</a></li>
							<li ><a href='#'>监控中心</a></li>
						</ul>
						<ul class='nav navbar-nav navbar-right'>
							{{if .GUser.Account}}
								<li><a href='javascript:void(0)'>{{.GUser.UserName}}</a></li>
								<li><a href='/logout.go'>退出</a></li>
							{{else}}
								<li><a href='/login.go'>登录{{.GUser.Account}}</a></li>
							{{end}}
						</ul>
					</div>
				</div>
			</div>
		</header>

		<article class='container'>
				{{.LayoutContent}}
		</article>

		<footer> 
			<div class='container-fluid'> 
				<div class='row text-right' style='padding: 0 12px;'>
					<small>COPY RIGHT 2014-2099 &copy; giter&lt;nubix at qq dot com&gt;</small>
					<!--<a target='_blank' href='https://github.com/giter/hdmxx'>@giter</a>-->
				</div>
			</div>
		</footer>

	</body>
</html>
