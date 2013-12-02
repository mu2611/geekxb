          <div class="container-fluid">
              <div class="row-fluid">
                <div class="span12">
                  <div class="widget-box">
                      <div class="widget-title"> 
                          <span class="icon">
                              <div class="checker" id="uniform-title-checkbox"><span class=""><input type="checkbox" id="title-checkbox" name="title-checkbox" style="opacity: 0; "></span></div>
                          </span>
                          <h5>基本权限配置</h5><div class="buttons"><a href="/Admin_Authority/Edit?editID=-1" class="btn btn-success btn-mini">添加</a></div>
                        </div>
                      <div class="widget-content nopadding">
                        <div id="DataTables_Table_0_wrapper" class="dataTables_wrapper" role="grid">
                        
                        <table class="table table-bordered table-striped with-check">
                                <thead>
                                  <tr>
                                    <th style="width:20px;">ID</th>
									<th style="width:200px;">分组</th>
                                    <th style="width:300px;">名称</th>
                                    <th >Action</th>
                                    <th style="width:150px;">操作</th>
                                  </tr>
                                </thead>
                                <tbody>
                                  {{range .Authitems}}
                                      <tr>
                                        <td><div class="checker" id="uniform-undefined"><span class=""><input type="checkbox" style="opacity: 0; "></span></div></td>
										<td>{{.Authgroup}}</td>
                                        <td>{{.Authname}}</td>
                                        <td>{{.Authaction}}</td>
                                        <td><div class="fr">
                                              <a href="/Admin_Authority/Edit?editID={{.Id}}" class="btn btn-primary btn-mini">编辑</a> 
                                              <a href="/Admin_Authority/Del?delID={{.Id}}" class="btn btn-danger btn-mini">删除</a>
                                            </div>
                                        </td>
                                      </tr>
                                  {{end}}
                                </tbody>
                              </table>
                             {{str2html .PageDiv}}
                        </div>
                      </div>
                  </div>
                </div>
              </div>
            </div>