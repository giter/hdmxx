<!DOCTYPE html>

<html>
  	<head>
		<title>HDMonitor ++</title>

		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

		<link rel="stylesheet" href="/static/css/bootstrap.min.css">
		<link rel="stylesheet" href="/static/css/global.css">

		<script src="/static/js/jquery.min.js"></script>
		<script src="/static/js/bootstrap.min.js"></script>
		
		<!--[if lt IE 9]>
			<script src="/static/js/html5shiv.min.js"></script>
			<script src="/static/js/respond.min.js"></script>
		<![endif]-->

		{{.Heads}}

	</head>

  	<body>

		<header>
			<div class='navbar navbar-inverse navbar-static-top'>
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
							<li><a href='#'>登录</a></li>
						</ul>
					</div>
				</div>
			</div>
		</header>

		<div class='container'>
			<div class='row' style='margin-top:12px;'>
				{{.LayoutContent}}
			</div>
		</div>

		<footer> 
			<div class='container-fluid'> 
				<div class='row text-right' style='padding: 0 12px;'>
					<a target='_blank' href='https://github.com/giter/hdmxx'>@giter</a>
				</div>
			</div>
		</footer>

	</body>
</html>
