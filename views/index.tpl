<!DOCTYPE html>

<html>
  	<head>
		<title>HDMonitor ++</title>

		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

		<link rel="stylesheet" href="/static/css/bootstrap.min.css">
		<script src="/static/js/jquery.min.js"></script>
		<script src="/static/js/bootstrap.min.js"></script>
	</head>

  	<body>

  		<div style="background-color:#A9F16C">
			<div class="container">
				<div class="row">
				  <div>
					<h1>Welcome to HDMonitor ++ !</h1>
					<br />
				  </div>
				</div>
			</div>
		</div>

		<div class='container' style='margin-top:1em;'>
			<div class="row">
				<table class='table table-bordered'>

					<thead>
						<tr>
							<th width=120>ID</th>
							<th>网站</th>
							<th>网址</th>
							<th>方法</th>
							<th>间隔</th>
						</tr>
					</thead>

					<tbody>
						{{range .Sites}}
							<tr>
								<td>{{.HexId}}</td>
								<td>{{.Name}}</td>
								<td>{{.Url}}</td>
								<td>{{.Method}}</td>
								<td>{{.Duration}}</td>
							</tr>
						{{end}}
					</tbody>
				</table>
			</div>
		</div>
	</body>
</html>
