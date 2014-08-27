{{define "func-sites" }}
<div class='panel panel-{{.Color}}'>
	<div class='panel-heading clearfix'>
			<b>{{.Name}}</b>
			{{if .Add}}
			<div style='float:right'>

				<a class='btn btn-primary btn-sm' href='/site.go'><span class='glyphicon glyphicon-plus'></span> 监控项目</a>
			</div>
			{{end}}
	</div>
	<table class='table table-bordered table-striped table-hover'>

		<thead>
			<tr>
				<th>类别</th>
				{{/*<th width=120>ID</th>*/}}
				<th>网站</th>
				<th>监控次数</th>
				<th>本次</th>
				<th>耗时</th>
				<th>下次计划</th>
				<th>用户</th>
				<th>操作</th>
			</tr>
		</thead>

		<tbody>

			{{range .Sites}}
				<tr class='type type-{{.Type}} {{if eq .Status 0}}danger{{end}}'>
					<td><b>{{.Type}}</b></td>
					{{/*<td>{{.HexId}}</td>*/}}
					<td><a href='{{.Url}}' target='_blank'>{{.Name}}</a></td>
					<td>{{.Count}}</td>
					<td>{{.TRun}}</td>
					<td>{{.Delay}} ms</td>
					<td>{{.TExpiration}}</td>
					<td>{{range .Users}}<a href='mailto:{{.UserName}}<{{.Email}}>'>{{.UserName}}</a> {{end}}</td>
					<td><a href='/site.go?Id={{.HexId}}'>编辑</a></td>
				</tr>
			{{end}}
		</tbody>
	</table>
</div>
{{end}}

<div class='row'>

	<div class='col-md-2'>
		<ul class='nav nav-stacked nav-pills'>
			<li class='active'><a class='switch-type' href='javascript:void(0)'>全部<span class='badge' style='float:right'>{{if .Stats.ALL}}{{.Stats.ALL}}{{else}}0{{end}}</a></li>
			<li><a class='switch-type' data-type='HTTP' href='javascript:void(0)'>HTTP<span class='badge' style='float:right'>{{if .Stats.HTTP}}{{.Stats.HTTP}}{{else}}0{{end}}</a></li>
			<li><a class='switch-type' data-type='TCP' href='javascript:void(0)'>TCP<span class='badge' style='float:right'>{{if .Stats.TCP}}{{.Stats.TCP}}{{else}}0{{end}}</a></li>
			<li><a class='switch-type' data-type='UDP' href='javascript:void(0)'>UDP<span class='badge' style='float:right'>{{if .Stats.UDP}}{{.Stats.UDP}}{{else}}0{{end}}</a></li>
		</ul>
	</div>

	<div class='page col-md-10'>
			
		<div class='row'>

			{{template "func-sites" .Sites}}

			{{template "func-sites" .Disabled}}
		</div>
	</div>

</div>

<script>
	$(function(){

		$(".switch-type").click(function(){

			$(".switch-type").parent().removeClass("active");
			$(this).parent().addClass("active")


			var type = $(this).data("type")
			if(type){
				$("tr.type").hide();
				$("tr.type.type-"+type).show();
			}else{
				$("tr.type").show();
			}
		});
	});
</script>
