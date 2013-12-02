	<div class="container-fluid">
       <div class="row-fluid">
       <div class="span12">
		<div class="widget-box">
	        <div class="widget-title"> <span class="icon"> <i class="icon-align-justify"></i> </span>
	          <h5>编辑用户组</h5>
	        </div>
	        <div class="widget-content nopadding">
	          <form action="Update" method="post" class="form-horizontal">
				<input type="hidden" name="editid" value={{.editID}}>
	            <div class="control-group">
	              <label class="control-label">名 称 :</label>
	              <div class="controls">
	                <input type="text" class="span11" name="Ugname" value={{.Ugname}}>
	              </div>
	            </div>
	            <div class="control-group">
	              <label class="control-label">Action Url :</label>
	              <div class="controls">
	                <input type="text" class="span11" name="Ugmanageauthority" value={{.Ugmanageauthority}}>
	              </div>
	            </div>
				<div class="form-actions">
	              <button type="submit" class="btn btn-success">提 交</button>
	            </div>
	          </form>
	        </div>
	      </div>
	</div>
	</div>
</div>