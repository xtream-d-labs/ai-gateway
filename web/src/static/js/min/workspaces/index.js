var workspaces=[],conditions={firstLoad:!0,words:"",order:1},vue=new Vue({el:"#data",data:{workspaces:[]},methods:{del:function(e){var t=app.trim($(e.target).closest("li").find(".workspace-path").text()),o=$("#workspace-delete");o.attr("data-path",t),o.find("strong").text(t),o.modal("open")},update:function(){var o=[];$.map(workspaces,function(e){if(""==conditions.words||app.match([e.path],conditions.words)){var t=e.path;e.time=parseInt(t.substring(t.lastIndexOf("-")+1,t.length),10),o.push(e)}}),o.sort(function(e,t){if(1==conditions.order){if(e.path<t.path)return-1;if(e.path>t.path)return 1}var o=t.time-e.time;return 0!=o?o:e.path<t.path?-1:e.path>t.path?1:0});var t=[];$.map(o,function(e){e.notebooks.sort(function(e,t){return t<e?-1:e<t?1:0}),t.push({path:e.path,note:e.notebooks,jobs:e.jobs,full:e.absolute_path,time:app.date.format(new Date(1e3*e.time))})}),this.workspaces=t,conditions.firstLoad&&(conditions.firstLoad=!1,""!=conditions.words&&0<t.length&&setTimeout(function(){$("#data .collapsible .row-body").eq(0).collapse("show"),$("#query-words").blur()},500)),$("#record-count").text(t.length)}}});function load(a){API("Workspace").getWorkspaces(function(e,t,o){$.isEmptyObject(e)&&o&&o.body&&(workspaces=o.body,update(),a&&a())})}function update(){conditions.words=app.singleSpace($("#query-words").val()),conditions.order=parseInt($("#query-order-type").val(),10),vue.update()}$(document).ready(function(){$("#menu-workspaces").addClass("active"),app.query("q")&&$("#query-words").val(app.query("q")).focus(),load(function(){$("#query-words").keyup(function(){update()}),$("#query-order-type").change(function(){update()}),$("#workspace-delete a.cancel").click(function(){$("#workspace-delete").modal("close")}),$("#workspace-delete a.delete").click(function(){$("#workspace-delete").modal("close");var n=$("#workspace-delete").attr("data-path"),e=new models.Workspace(n);API("Workspace").deleteWorkspace(e,function(e,t,o){if(!$.isEmptyObject(e)){var a="Could not delete the specified workspace";return o&&o.body&&o.body.message&&(a=o.body.message),void M.toast({html:a,displayLength:3e3})}M.toast({html:"Deleted successfully. [ "+n+" ]",displayLength:3e3}),load()})}),setTimeout(function(){$(".notebooks").each(function(e,t){$(t).prop("href","/notebooks/?q="+encodeURIComponent($(t).text()))}),$(".jobs").each(function(e,t){$(t).prop("href","/jobs/?q="+encodeURIComponent($(t).text()))})},500),$("#data").fadeIn()})});