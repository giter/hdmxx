<div class='panel panel-primary' style='width:600px;margin:0 auto;'>

<div class='panel-heading'><b>{{if eq .Site.HexId ""}}创建 - 监控项目{{else}}编辑 - {{.Site.Name}}{{end}}</b></div>
<div class='panel-body'>

	<form id='form-site' action='site.go' method='POST' class='form-horizontal'>

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
			<label class='col-sm-2 control-label'>类型</label> 
			<div class='col-sm-10'>
				<select name='Type' class='form-control required'>
					<option value='HTTP' {{if eq .Site.Type "HTTP"}}selected="selected"{{end}}>HTTP</option>
					<option value='TCP'  {{if eq .Site.Type "TCP"}}selected="selected"{{end}}>TCP</option>
					<option value='UDP'  {{if eq .Site.Type "UDP"}}selected="selected"{{end}}>UDP</option>
				</select>
			</div>
		</div>

		<div class='form-group group-http group'>
			<label class='col-sm-2 control-label'>CheckPoint</label> 
			<div class='col-sm-10'>
				<input class='form-control required' type='text' placeholder='http://example.com/action' name='CheckPoint' value="{{.Site.CheckPoint}}" />
			</div>
		</div>

		<div class='form-group group-http group'>
			<label class='col-sm-2 control-label'>Method</label> 
			<div class='col-sm-10'>
				<select name='Method' class='form-control required'>
					<option value='GET' {{if eq .Site.Method "GET"}}selected="selected"{{end}}>GET</option>
					<option value='POST' {{if eq .Site.Method "POST"}}selected="selected"{{end}}>POST</option>
				</select>
			</div>
		</div>

		<div class='form-group group-tcp group-udp group'>
			<label class='col-sm-2 control-label'>Address</label> 
			<div class='col-sm-10'>

				<input class='form-control required' type='text' placeholder='8.8.8.8' name='Address' value="{{.Site.Address}}" />
			</div>
		</div>

		<div class='form-group group-tcp group-udp group'>
			<label class='col-sm-2 control-label'>Port</label> 
			<div class='col-sm-10'>

				<input class='form-control required' type='text' placeholder='8080' name='Port' value="{{ if gt .Site.Port 0 }}{{.Site.Port}}{{end}}" />
			</div>
		</div>

		<div class='form-group group-http group-tcp group-udp group'>
			<label class='col-sm-2 control-label'>连接超时</label> 
			<div class='col-sm-10'>

				<div class='input-group'>
					<input class='form-control required number' type='text' placeholder='5000' name='CTimeout' value="{{ if gt .Site.CTimeout 0 }}{{.Site.CTimeout}}{{else}}5000{{end}}" />
					<div class='input-group-addon'>ms</div>
				</div>
			</div>
		</div>

		<div class='form-group group-http group-tcp group-udp group'>
			<label class='col-sm-2 control-label'>读取超时</label> 
			<div class='col-sm-10'>

				<div class='input-group'>
					<input class='form-control required number' type='text' placeholder='60000' name='RTimeout' value="{{ if gt .Site.RTimeout 0 }}{{.Site.RTimeout}}{{else}}60000{{end}}" />
					<div class='input-group-addon'>ms</div>
				</div>
			</div>
		</div>

		<div class='form-group group-tcp group-udp group'>
			<label class='col-sm-2 control-label'>输入字节</label> 
			<div class='col-sm-10'>

				<div class='input-group'>
					<input class='form-control' type='text' placeholder='Base64 Encoding' name='Input' value="{{.Site.Input}}" />
					<div class='input-group-addon'>B64</div>
				</div>
			</div>
		</div>


		<div class='form-group group-tcp group-udp group'>
			<label class='col-sm-2 control-label'>输出字节</label> 
			<div class='col-sm-10'>

				<div class='input-group'>
					<input class='form-control' type='text' placeholder='Base64 Encoding' name='Result' value="{{.Site.Result}}" />
					<div class='input-group-addon'>B64</div>
				</div>
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

<script>

	$(function(){

		$("#form-site").find("[name=Type]").change(function(){
			$("#form-site").find(".group").hide().end().find(".group-"+$(this).val().toLowerCase()).show()
		}).trigger('change');
	});
</script>
