// <script>
Vue.use(vuelidate.default);

var images = [];
var conditions = {
  firstLoad: true,
  words: '',
  order: 1
};
var vue = new Vue({
  el: '#data',
  data: {
    images: []
  },
  methods: {
    run: function (e) {
      var li = $(e.target).closest('li');
      runner.tag = app.trim(li.find('.image-tag').text());
      var dialog = $('#run-dialog');
      dialog.find('.btn').addClass('disabled');
      dialog.find('.workspaces select').html('');
      dialog.modal('open');

      API('Workspace').getWorkspaces(function (err, _, res) {
        if (! $.isEmptyObject(err) || ! res || ! res.body) {
          return;
        }
        $.map(res.body, function (workspace) {
          var path = workspace.path;
          workspace.time = parseInt(path.substring(path.lastIndexOf('-')+1, path.length), 10);
        });
        res.body.sort(function (a, b) {
          var ret = b.time - a.time;
          if (ret != 0) return ret;
          if (a.path < b.path) return -1;
          if (a.path > b.path) return  1;
          return 0;
        });
        var html = '';
        $.map(res.body, function (workspace) {
          html += '<option value="'+workspace.path+'">'+workspace.path+'</option>';
        });
        var select = dialog.find('.workspace-type select').val('0');
        if (res.body.length == 0) {
          select.prop("disabled", true);
        }
        select.formSelect();

        select = dialog.find('.workspaces select').html(html);
        select.formSelect();
        M.FormSelect.getInstance(select[0]).destroy();

        dialog.find('.btn').removeClass('disabled');
      });
    },
    update: function () {
      var filtered = [];
      $.map(images, function (image) {
        $.map(image.repoTags, function (tag) {
          if (conditions.words != '') {
            if (! app.match([tag], conditions.words)) {
              return;
            }
          }
          filtered.push({
            id:      image.id,
            tag:     tag,
            size:    image.size,
            created: image.created
          });
        });
      });
      filtered.sort(function (a, b) {
        switch (conditions.order) {
        case 1:
          if (a.tag < b.tag) return -1;
          if (a.tag > b.tag) return  1;
          break;
        case 3:
          return b.size - a.size;
        }
        var ret = new Date(b.created).getTime() - new Date(a.created).getTime();
        if (ret != 0) return ret;
        if (a.tag < b.tag) return -1;
        if (a.tag > b.tag) return  1;
        return 0;
      });
      var formatted = [];
      $.map(filtered, function (image, idx) {
        if (image.id) {
          formatted.push({
            id:      image.id,
            tag:     image.tag,
            size:    app.comma(image.size, 'byte'),
            created: app.date.format(new Date(image.created))
          });
        } else {
          formatted.push({
            id:      'img-'+idx,
            tag:     image.tag,
            size:    '',
            created: 'Now downloading..'
          });
        }
      });
      this.images = formatted;

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
    },
    exists: function (name) {
      var found = false;
      $.map(images, function (image) {
        $.map(image.repoTags, function (tag) {
          found |= (tag == name);
        });
      });
      return found;
    },
    del: function (e) {
      var tag = app.trim($(e.target).closest('li').find('.image-tag').text()),
          dialog = $('#image-delete');
      dialog.attr('data-image-tag', tag);
      dialog.find('strong').text(tag);
      dialog.modal('open');
    },
    delImage: function (tag) {
      $('#image-delete').modal('close');

      var body = new models.ImageName1(tag);
      API('Image').deleteImage(body, function (err, _, res) {
        if (app.shouldExit(res, err)) {
          alert('Something went wrong. Check your configurations!')
          window.location.href = '/settings/';
          return;
        }
        if (! $.isEmptyObject(err)) {
          var message = 'Could not delete the specified image';
          if (res && res.body && res.body.message) {
            message = res.body.message;
          }
          M.toast({html: message, displayLength: 3000});
          return;
        }
        M.toast({html: 'Deleted successfully. [ '+tag+' ]', displayLength: 3000});
        load();
      });
    }
  }
});

var imageformat = vuelidate.withParams({type: 'custom'}, function (value) {
  if (value.length <= 0) {
    return false;
  }
  var match = /(.+\/)?([^\/:][^:]+)?(:.*)?/.exec(app.trim(value)); // eslint-disable-line no-useless-escape
  if ((! match) || (match.length < 4)) {
    return false;
  }
  return (match[2] !== undefined);
});
var image = new Vue({
  el: '#image-dialog',
  data: {
    name: '',
    msgName: ''
  },
  validations: {
    name: {custom: imageformat}
  },
  methods: {
    nameChanged: function () {
      this.msgName = '';
    },
    submit: function () {
      setTimeout(function() {$('.btn').blur();}, 500);
      $('form p.errors').show();

      var name = app.trim(this.name);
      if ((name == '') || $('#image-dialog input').hasClass('invalid')) {
        image.error('Invalid image name');
        return;
      }
      if (vue.exists(name)) {
        image.error('Already exists');
        return;
      }
      var body = new models.ImageName(name);
      API('Image').postNewImage(body, function (err, _, res) {
        if (app.shouldExit(res, err)) {
          alert('Something went wrong. Check your configurations!')
          window.location.href = '/settings/';
          return;
        }
        if (! $.isEmptyObject(err)) {
          var message = 'Could not pull the specified image';
          if (res && res.body && res.body.message) {
            message = res.body.message;
          }
          image.error(message);
          return;
        }
        $('#image-dialog').modal('close');
        load();
      });
    },
    error: function (message) {
      this.msgName = message;
    },
    close: function () {
      $('#image-dialog').modal('close');
    }
  }
});

var runner = new Vue({
  el: '#run-dialog',
  data: {
    tag: ''
  },
  methods: {
    submit: function () {
      setTimeout(function() {$('.btn').blur();}, 500);
      $('form p.errors').show();
      $('#run-dialog').modal('close');

      var tag = this.tag,
          wt = $('#run-dialog .workspace-type select').val(),
          work = $('#run-dialog .workspaces select').val(),
          body = models.ImageName2.constructFromObject({
            'image': tag,
            'workspace': (wt == '0') ? '' : work
          });
      API('Notebook').postNewNotebook(body, function (err, _, res) {
        if (app.shouldExit(res, err)) {
          alert('Something went wrong. Check your configurations!')
          window.location.href = '/settings/';
          return;
        }
        location.href = '/notebooks/?q=' + encodeURIComponent(tag);
      });
    },
    close: function () {
      $('#run-dialog').modal('close');
    }
  }
});

function update() {
  conditions.words = app.singleSpace($('#query-words').val());
  conditions.order = parseInt($('#query-order-type').val(), 10);
  vue.update();
}

function load(callback) {
  API('Image').getImages(function (error, data, response) {
    if (! $.isEmptyObject(error) || ! response || ! response.body) {
      return;
    }
    images = response.body;
    update();
    callback && callback();
  });
}

$(document).ready(function () {
  $('#menu-images').addClass('active');
  if (app.query('q')) {
    $('#query-words').val(app.query('q')).focus();
  }
  load(function () {
    $('#act-pull').click(function (e) {
      app.stop(e);
      image.name = '';
      image.msgName = '';
      $('#image-dialog').modal('open');
      $('#image-dialog input').focus();
    });

    $('#query-words').keyup(function () {
      update();
    });
    $('#query-order-type').change(function () {
      update();
    });

    $('#image-delete a.delete').click(function () {
      vue.delImage($('#image-delete').attr('data-image-tag'));
    });
    $('#image-delete a.cancel').click(function () {
      $('#image-delete').modal('close');
    });
    $("#run-dialog .workspace-type select").on('change', function () {
      var select = $('#run-dialog .workspaces select');
      if ($(this).val() == '0') {
        M.FormSelect.getInstance(select[0]).destroy();
      } else {
        select.formSelect();
      }
    });
    $('#data').fadeIn();
  });
  setInterval(load, 10 * 1000);
});
// </script>
