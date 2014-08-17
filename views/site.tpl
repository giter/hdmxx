<div class='panel panel-primary' style='width:600px;margin:0 auto;'>

<div class='panel-heading'><b>{{if eq .Site.HexId ""}}创建{{else}}编辑{{end}}监控</b></div>
<div class='panel-body'>

	<form action='site.go' method='POST' class='form-horizontal'>

		<input type='hidden' name='Id' value='{{.Site.HexId}}' />

		<div class='form-group'>
			<label class='col-sm-2 control-label'>名称</label> 
			<div class='col-sm-10'>
				<input class='form-control required' type='text' placeholder='项目名称' name='Name' value="{{.Site.Name}}" data-msg-required="请输入项目名称.." />
			</div>
		</div>

		<div class='form-group'>
			<label class='col-sm-2 control-label'>网址</label> 
			<div class='col-sm-10'>
				<input class='form-control required' type='text' placeholder='http://example.com' name='Url' value="{{.Site.Url}}" data-msg-required="请输入网址.."/>
			</div>
		</div>

		<div class='form-group'>
			<label class='col-sm-2 control-label'>CheckPoint</label> 
			<div class='col-sm-10'>
				<input class='form-control required' type='text' placeholder='http://example.com/action' name='CheckPoint' value="{{.Site.CheckPoint}}" />
			</div>
		</div>

		<div class='form-group'>
			<label class='col-sm-2 control-label'>Method</label> 
			<div class='col-sm-10'>
				<select name='Method' class='form-control required'>
					<option value='GET' {{if eq .Site.Method "GET"}}selected="selected"{{end}}>GET</option>
					<option value='POST' {{if eq .Site.Method "POST"}}selected="selected"{{end}}>POST</option>
				</select>
			</div>
		</div>

		<div class='form-group'>
			<label class='col-sm-2 control-label'>监控间隔</label> 
			<div class='col-sm-10'>
				<select name='Duration' class='form-control required'> 
					<option value='1800' {{if eq .Site.Duration 1800}}selected="selected"{{end}}>30分钟</option>
					<option value='3600' {{if eq .Site.Duration 3600}}selected="selected"{{end}}>1小时</option>
					<option value='7200' {{if eq .Site.Duration 7200}}selected="selected"{{end}}>2小时</option>
					<option value='14400' {{if eq .Site.Duration 14400}}selected="selected"{{end}}>4小时</option>
					<option value='28800' {{if eq .Site.Duration 28800}}selected="selected"{{end}}>8小时</option>
					<option value='57600' {{if eq .Site.Duration 57600}}selected="selected"{{end}}>16小时</option>
				</select>
			</div>
		</div>

		<div class='form-group'>
			<label class='col-sm-2 control-label'>禁用</label> 
			<div class='col-sm-10'>
				<input type='checkbox' name='Disabled' value='true' style='margin-top:11px;' {{if eq .Site.Disabled true}}checked='checked'{{end}} />
			</div>
		</div>

		<div class='form-group'>
			<label class='col-sm-2 control-label'>&nbsp;</label> 
			<div class='col-sm-10'>
				<input class='btn btn-primary' type='submit' value='确定' />
			</div>
		</div>

	</form>
</div>
</div>
