// <script>
var notebooks = [];
var conditions = {
  firstLoad: true,
  words: '',
  order: 1
};
var vue = new Vue({
  el: '#data',
  data: {
    notebooks: []
  },
  methods: {
    train: function (e) {
      var li = $(e.target).closest('li');
      training.id = li.attr('data-id');
      training.name = li.find('.notebook-name').text();
      training.cmd = '';
      var dialog = $('#training-dialog');
      dialog.find('.form-group.considerable div').css({opacity: 0});
      dialog.find('.btn').addClass('disabled');
      dialog.find('.training-notebook select').html('');
      dialog.modal('open');

      var loadIpynbs = false, loadCoretypes = false;

      API('Notebook').getIPythonNotebooks(training.id, function (err, _, res) {
        if (app.shouldExit(res, err)) {
          return;
        }
        res.body.sort(function (a, b) {
          if (a.name < b.name) return -1;
          if (a.name > b.name) return  1;
          return 0;
        });
        var html = '';
        $.map(res.body, function (ipynb) {
          html += '<option value="'+ipynb.name+'">'+ipynb.name+'</option>';
        });
        if (html == '') {
          html += '<option value="none">There is no notebook in your workspace.</option>';
        }
        var jobType = dialog.find('.training-type select'),
            notebooks = dialog.find('.training-notebook select').html(html);
        if (res.body.length == 0) {
          jobType.val('1').prop("disabled", true);
          notebooks.prop("disabled", true);
          dialog.find('.training-cmds').focus();
          training.cmd = '';
        } else {
          jobType.val('0').removeProp("disabled");
          notebooks.removeProp("disabled");
          training.cmd = 'python <converted-notebook.py>';
        }
        jobType.formSelect();
        notebooks.formSelect();
        loadIpynbs = true;
        loaded()
      });

      API('Rescale').getCoreTypes({appVer:'cpu:cheap'}, function (err, _, res) {
        if (app.shouldExit(res, err) || (res && res.body && res.body.code == 401)) {
          alert('Something went wrong. Check your configurations!')
          window.location.href = '/settings/';
          return;
        }
        var html = '';
        if (res.body.length > 0) {
          $.map(res.body[0].resources, function (resource) {
            var gpus = (resource.gpus > 0) ? ", GPU: " + resource.gpus : "";
            html += '<option value="'+resource.cores+'">' +
              "CPU cores: " + resource.cores + gpus + '</option>';
          });
        }
        dialog.find('.training-cores select').html(html).formSelect();

        loadCoretypes = true;
        loaded()
      });

      function loaded() {
        if (!loadIpynbs || !loadCoretypes) {
          return;
        }
        dialog.find('.form-group.considerable div').animate({opacity: 1}, 500);
        dialog.find('.btn').removeClass('disabled');
      }
    },
    stop: function (e) {
      confirmation.action = 'STOP';
      confirmation.id = app.trim($(e.target).closest('li').attr('data-id'));
      confirmation.name = $(e.target).closest('li').find('.notebook-name').text();
      $('#notebook-modify').modal('open');
    },
    del: function (e) {
      confirmation.action = 'DELETE';
      confirmation.id = app.trim($(e.target).closest('li').attr('data-id'));
      confirmation.name = $(e.target).closest('li').find('.notebook-name').text();
      $('#notebook-modify').modal('open');
    },
    update: function () {
      var filtered = [];
      $.map(notebooks, function (notebook) {
        if (conditions.words != '') {
          if (! app.match([notebook.image, notebook.name], conditions.words)) {
            return;
          }
        }
        if (notebook.state == 'created') {
          return;
        }
        filtered.push(notebook);
      });
      filtered.sort(function (a, b) {
        if (conditions.order == 1) {
          if (a.image < b.image) return -1;
          if (a.image > b.image) return  1;
        }
        var ret = new Date(b.started).getTime() - new Date(a.started).getTime();
        if (ret != 0) return ret;
        if (a.image < b.image) return -1;
        if (a.image > b.image) return  1;
        return 0;
      });
      var formatted = [];
      $.map(filtered, function (notebook) {
        var url = notebook.port ?
          window.location.protocol + "//" + window.location.hostname + ":" + notebook.port :
          '';
        formatted.push({
          id:          notebook.id,
          name:        notebook.name.substring(1),
          image:       notebook.image,
          url:         url,
          state:       notebook.state,
          started:     app.date.format(new Date(notebook.started)),
          classObject: statusClass(notebook.state)
        });
      });
      this.notebooks = formatted;

      if (conditions.firstLoad) {
        conditions.firstLoad = false;
        if (conditions.words != '' && formatted.length > 0) {
          setTimeout(function () {
            // $('#data .collapsible .row-body').eq(0).collapse('show');
            $('#query-words').blur();
          }, 500);
        }
      }
      $('#record-count').text(formatted.length);
    }
  }
});

function statusClass(state) {
  switch (state) {
  case 'running':
    return {
        'label-success': true,
        'label-warning': false,
        'label-danger':  false
    };
  case 'creating':
    return {
        'label-success': false,
        'label-warning': true,
        'label-danger':  false
    };
  case 'exited':
    return {
        'label-success': false,
        'label-warning': false,
        'label-danger':  true
    };
  }
  // same as exited
  return {
      'label-success': false,
      'label-warning': false,
      'label-danger':  true
  };
}

function update() {
  conditions.words = app.singleSpace($('#query-words').val());
  conditions.order = parseInt($('#query-order-type').val(), 10);
  vue.update();
}

function load(callback) {
  API('Notebook').getNotebooks(function (error, data, response) {
    if (! $.isEmptyObject(error) || ! response || ! response.body) {
      return;
    }
    notebooks = response.body;
    update();
    callback && callback();
  });
}

function loadDetails(el) {
  if ($(el).attr('data-loaded')) {
    return;
  }
  var id = $(el).attr('data-id');
  API('Notebook').getNotebookDetails(id, function (err, _, res) {
    if (app.shouldExit(res, err)) {
      return;
    }
    var a = $('a.endpoint', el);
    if (! res.body.token) {
      setTimeout(function () {loadDetails(el);}, 1000);
      a.prop('href', '');
      return;
    }
    a.prop('href', el.closest('li').attr('data-url')+'?token='+res.body.token);

    $('.notebook-name', el).text(res.body.name.substring(1));
    $('.notebook-started', el).text(app.date.format(new Date(res.body.started)));
    var ended = '-';
    if (res.body.state == 'exited') {
      ended = app.date.format(new Date(res.body.ended));
    }
    $('.notebook-ended', el).text(ended);
    var mounts = (res.body.mounts.length > 0) ? res.body.mounts[0] : '';
    if (mounts.indexOf(':') > 0) {
      mounts = mounts.split(':');
      mounts = '<a href="/workspaces/?q='+encodeURIComponent(mounts[0])+'">'+mounts[0]+'</a>:'+mounts[1];
    }
    $('.notebook-volumes', el).html(mounts);

    $(el).attr('data-loaded', 'done');
  });
}

var training = new Vue({
  el: '#training-dialog',
  data: {
    id: '',
    name: '',
    cmd: ''
  },
  methods: {
    submit: function () {
      setTimeout(function() {$('.btn').blur();}, 500);
      $('form p.errors').show();

      var dialog = $('#training-dialog'),
          commands = app.singleSpace(this.cmd).split(' ');
      if (! commands || commands.length == 0 || commands[0] == '') {
        M.toast({html: 'You have to specify its commands.', displayLength: 3000});
        dialog.find('.training-cmds').focus();
        return;
      }
      dialog.modal('close');

      var entrypoint = dialog.find('.training-notebook select').val();
      if (dialog.find(".training-type select").val() == '1') {
        entrypoint = 'none';
      }
      var body = models.JobAttrs.constructFromObject({
        'notebook_id': this.id,
        'entrypoint_file': entrypoint,
        'commands': commands,
        'coretype': dialog.find(".training-coretype select").val(),
        'cores': parseInt(dialog.find(".training-cores select").val(), 10)
      });
      API('Job').postNewJob(body, function (err, _, res) {
        if (! $.isEmptyObject(err)) {
          if (res && res.body && res.body.code && (res.body.code == '400') && res.body.message) {
            M.toast({html: 'Could not start training with this notebook', displayLength: 3000});
            return;
          }
          if (res && res.body && res.body.code && (res.body.code == '406') && res.body.message) {
            M.toast({html: 'GPU is not allowed yet', displayLength: 3000});
            return;
          }
          if (app.shouldExit(res, err)) {
            alert('Something went wrong. Check your configurations!')
            window.location.href = '/settings/';
            return;
          }
          return;
        }
        location.href = '/jobs/?q=' + encodeURIComponent(res.body.id);
      });
    },
    error: function (message) {
      this.msgCmd = message;
    },
    close: function () {
      $('#training-dialog').modal('close');
    }
  }
});

var confirmation = new Vue({
  el: '#notebook-modify',
  data: {
    action: 'STOP',
    id: '',
    name: ''
  },
  methods: {
    exec: function () {
      $('#notebook-modify').modal('close');
      switch (this.action) {
      case 'STOP':
        var body = models.NotebookAttrs.constructFromObject({'status': 'stopped'});
        API('Notebook').modifyNotebook(this.id, body, function (err, _, res) {
          if (! $.isEmptyObject(err)) {
            var message = 'Could not stop the specified notebook';
            if (res && res.body && res.body.message) {
              message = res.body.message;
            }
            M.toast({html: message, displayLength: 3000});
            return;
          }
          M.toast({html: 'Stoped successfully. [ '+confirmation.name+' ]', displayLength: 3000});
          load();
        });
        break;
      case 'DELETE':
        API('Notebook').deleteNotebook(this.id, function (err, _, res) {
          if (! $.isEmptyObject(err)) {
            var message = 'Could not delete the specified notebook';
            if (res && res.body && res.body.message) {
              message = res.body.message;
            }
            M.toast({html: message, displayLength: 3000});
            return;
          }
          M.toast({html: 'Deleted successfully. [ '+confirmation.name+' ]', displayLength: 3000});
          load();
        });
        break;
      }
    },
    close: function () {
      $('#notebook-modify').modal('close');
    }
  }
});

$(document).ready(function () {
  $('#menu-notebooks').addClass('active');
  if (app.query('q')) {
    $('#query-words').val(app.query('q')).focus();
  }
  $('.collapsible').on('shown.bs.collapse', function (elem) {
    loadDetails($(elem.target).closest('li'));
  });
  load(function () {
    $('#query-words').keyup(function () {
      update();
    });
    $('#query-order-type').change(function () {
      update();
    });
    $("#training-dialog .training-type select").on('change', function () {
      var dialog = $('#training-dialog');
      if ($(this).val() == '1') {
        dialog.find('.training-notebook select').prop("disabled", true).formSelect();
        dialog.find('.training-cmds').focus();
        training.cmd = '';
      } else {
        dialog.find('.training-notebook select').removeProp("disabled").formSelect();
        training.cmd = 'python <converted-notebook.py>';
      }
    });
    $("#training-dialog .training-coretype select").on('change', function () {
      var dialog = $('#training-dialog'),
          cores = 'cpu:cheap';
      if ($(this).val() == 'dolomite') {
        cores = 'gpu:volta';
      }
      API('Rescale').getCoreTypes({appVer: cores}, function (err, _, res) {
        if (app.shouldExit(res, err) || (res && res.body && res.body.code == 401)) {
          alert('Something went wrong. Check your configurations!')
          window.location.href = '/settings/';
          return;
        }
        var html = '';
        if (res.body.length > 0) {
          $.map(res.body[0].resources, function (resource) {
            var gpus = (resource.gpus > 0) ? ", GPU: " + resource.gpus : "";
            html += '<option value="'+resource.cores+'">' +
              "CPU cores: " + resource.cores + gpus + '</option>';
          });
        }
        dialog.find('.training-cores select').html(html).formSelect();
      });
    });
    $('#data').fadeIn();
  });
  setInterval(load, 3 * 1000);
});
// </script>
