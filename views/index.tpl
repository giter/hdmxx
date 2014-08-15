<div class='text-right col-12' style='margin-bottom:12px'>

	<a class='btn btn-primary btn-sm' href='/site.go'>
			<span class='glyphicon glyphicon-plus'></span> 监控网站
	</a>
</div>

{{define "func-sites" }}
<div class='panel panel-{{.Color}}'>
	<div class='panel-heading clearfix'>
			<b>{{.Name}}({{len .Sites}})</b>
	</div>
	<table class='table table-bordered table-striped table-hover'>

		<thead>
			<tr>
				{{/*<th width=120>ID</th>*/}}
				<th>网站</th>
				<th>监控次数</th>
				<th>上次监控</th>
				<th>下次计划</th>
				<th>用户</th>
				{{/*<th>操作</th>*/}}
			</tr>
		</thead>

		<tbody>

			{{range .Sites}}
				<tr>
					{{/*<td>{{.HexId}}</td>*/}}
					<td><a href='{{.Url}}' target='_blank'>{{.Name}}</a></td>
					<td>{{.Count}}</td>
					<td>{{.TRun}}</td>
					<td>{{.TExpiration}}</td>
					<td>{{range .Users}}<a href='mailto:{{.Email}}'>{{.Account}}</a>{{end}}</td>
					{{/*<td><a href='/site.go?Id={{.HexId}}'>编辑</a></td>*/}}
				</tr>
			{{end}}
		</tbody>
	</table>
</div>
{{end}}

{{template "func-sites" .Sites}}

{{template "func-sites" .Disabled}}

