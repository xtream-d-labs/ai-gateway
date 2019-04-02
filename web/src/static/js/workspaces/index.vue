// <script>
var workspaces = [];
var conditions = {
  firstLoad: true,
  words: '',
  order: 1
};
var vue = new Vue({
  el: '#data',
  data: {
    workspaces: []
  },
  methods: {
    del: function (e) {
      var path = app.trim($(e.target).closest('li').find('.workspace-path').text()),
          dialog = $('#workspace-delete');
      dialog.attr('data-path', path);
      dialog.find('strong').text(path);
      dialog.modal('open');
    },
    update: function () {
      var filtered = [];
      $.map(workspaces, function (workspace) {
        if (conditions.words != '') {
          if (! app.match([workspace.path], conditions.words)) {
            return;
          }
        }
        var path = workspace.path;
        workspace.time = parseInt(path.substring(path.lastIndexOf('-')+1, path.length), 10);
        filtered.push(workspace);
      });
      filtered.sort(function (a, b) {
        if (conditions.order == 1) {
          if (a.path < b.path) return -1;
          if (a.path > b.path) return  1;
        }
        var ret = b.time - a.time;
        if (ret != 0) return ret;
        if (a.path < b.path) return -1;
        if (a.path > b.path) return  1;
        return 0;
      });
      var formatted = [];
      $.map(filtered, function (workspace) {
        workspace.notebooks.sort(function (a, b) {
          if (a > b) return -1;
          if (a < b) return  1;
          return 0;
        });
        formatted.push({
          path: workspace.path,
          note: workspace.notebooks,
          jobs: workspace.jobs,
          full: workspace.absolute_path,
          time: app.date.format(new Date(workspace.time*1000))
        });
      });
      this.workspaces = formatted;

      if (conditions.firstLoad) {
        conditions.firstLoad = false;
        if (conditions.words != '' && formatted.length > 0) {
          setTimeout(function () {
            $('#data .collapsible .row-body').eq(0).collapse('show');
            $('#query-words').blur();
          }, 500);
        }
      }
      $('#record-count').text(formatted.length);
    }
  }
});

function load(callback) {
  API('Workspace').getWorkspaces(function (error, data, response) {
    if (! $.isEmptyObject(error) || ! response || ! response.body) {
      return;
    }
    workspaces = response.body;
    update();
    callback && callback();
  });
}

function update() {
  conditions.words = app.singleSpace($('#query-words').val());
  conditions.order = parseInt($('#query-order-type').val(), 10);
  vue.update();
}

$(document).ready(function () {
  $('#menu-workspaces').addClass('active');
  if (app.query('q')) {
    $('#query-words').val(app.query('q')).focus();
  }
  load(function () {
    $('#query-words').keyup(function () {
      update();
    });
    $('#query-order-type').change(function () {
      update();
    });
    $('#workspace-delete a.cancel').click(function () {
      $('#workspace-delete').modal('close');
    });
    $('#workspace-delete a.delete').click(function () {
      $('#workspace-delete').modal('close');

      var path = $('#workspace-delete').attr('data-path'),
          body = new models.Workspace(path);
      API('Workspace').deleteWorkspace(body, function (err, _, res) {
        if (! $.isEmptyObject(err)) {
          var message = 'Could not delete the specified workspace';
          if (res && res.body && res.body.message) {
            message = res.body.message;
          }
          M.toast({html: message, displayLength: 3000});
          return;
        }
        M.toast({html: 'Deleted successfully. [ '+path+' ]', displayLength: 3000});
        load();
      });
    });
    setTimeout(function () {
      $('.notebooks').each(function (_, el) {
        $(el).prop('href', '/notebooks/?q=' + encodeURIComponent($(el).text()));
      });
      $('.jobs').each(function (_, el) {
        $(el).prop('href', '/jobs/?q=' + encodeURIComponent($(el).text()));
      });
    }, 500);
    $('#data').fadeIn();
  });
});
// </script>
