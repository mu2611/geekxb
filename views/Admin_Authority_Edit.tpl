	<div class="container-fluid">
       <div class="row-fluid">
       <div class="span12">
		<div class="widget-box">
	        <div class="widget-title"> <span class="icon"> <i class="icon-align-justify"></i> </span>
	          <h5>添加基本权限数据</h5>
	        </div>
	        <div class="widget-content nopadding">
	          <form action="Update" method="post" class="form-horizontal">
				<input type="hidden" name="authid" value={{.editID}}>
				<div class="control-group">
	              <label class="control-label">分 组 :</label>
	              <div class="controls">
	                <input type="text" class="span11" name="authgroup" value={{.authGroup}}>
	              </div>
	            </div>
	            <div class="control-group">
	              <label class="control-label">名 称 :</label>
	              <div class="controls">
	                <input type="text" class="span11" name="authname" value={{.authName}}>
	              </div>
	            </div>
	            <div class="control-group">
	              <label class="control-label">Action Url :</label>
	              <div class="controls">
	                <input type="text" class="span11" name="authaction" value={{.authAction}}>
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