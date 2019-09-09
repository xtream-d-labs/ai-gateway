// <script>
var converter = new showdown.Converter();
var images = [];
var conditions = {
  firstLoad: true,
  words: ''
};
var vue = new Vue({
  el: '#data',
  data: {
    images: []
  },
  methods: {
    update: function () {
      var filtered = [];
      $.map(images, function (image) {
        if (conditions.words != '') {
          if (! app.match([image.namespace, image.name], conditions.words)) {
            return;
          }
        }
        filtered.push({
          code: image.namespace + '/' + image.name,
          namespace: image.namespace,
          name: image.name,
          description: image.description ? converter.makeHtml(image.description) : '---'
        });
      });
      filtered.sort(function (a, b) {
        if (a.name < b.name) return -1;
        if (a.name > b.name) return  1;
        return 0;
      });
      this.images = filtered;

      if (conditions.firstLoad) {
        conditions.firstLoad = false;
        if (conditions.words != '' && filtered.length > 0) {
          setTimeout(function () {
            // $('#data .collapsible .row-body').eq(0).collapse('show');
            $('#query-words').blur();
          }, 750);
        }
      }
      $('#record-count').text(filtered.length);
      $('.wait-icon').hide();
    }
  }
});

function update() {
  conditions.words = app.singleSpace($('#query-words').val());
  vue.update();
}

function load(callback) {
  API('Repository').getRemoteRepositories(function (err, _, res) {
    if (app.shouldExit(res, err)) {
      alert('Something went wrong. Check your configurations!')
      window.location.href = '/settings/';
      return;
    }
    images = res.body;
    update();
    callback && callback();
  });
}

function loadDetails(el) {
  if ($(el).attr('data-loaded')) {
    return;
  }
  var ns = $(el).attr('data-ns');
  if (ns) ns += '/';
  var nm = $(el).attr('data-nm');
  API('Repository').getRemoteImages(nm, function (err, _, res) {
    if (app.shouldExit(res, err)) {
      alert('Something went wrong. Check your configurations!')
      window.location.href = '/settings/';
      return;
    }
    res.body.sort(function (a, b) {
      if (a.repoTags[0] < b.repoTags[0]) return  1;
      if (a.repoTags[0] > b.repoTags[0]) return -1;
      return 0;
    });
    var html = '';
    $.map(res.body, function (image) {
      var tag = image.repoTags[0];
      html += '<tr data-id="' + ns + nm +':'+ tag +'">';
        html += '<td>'+tag+'</td>';
        html += '<td><a class="waves-effect waves-light btn blue darken-1">download</a></td>';
      html += '</tr>';
    });
    $(el).find('.progress').hide();

    $('.row-body tbody', el).html(html);
    $('.row-body .btn', el).click(function (e) {
      pullImage($(e.target).closest('tr').attr('data-id'));
    });
    $(el).attr('data-loaded', 'done');
  });
}

function pullImage(name) {
  var body = new models.ImageName(name);
  API('Image').postNewImage(body, function (err, _, res) {
    if (app.shouldExit(res, err)) {
      alert('Something went wrong. Check your configurations!')
      window.location.href = '/settings/';
      return;
    }
    location.href = '/images/?q=' + encodeURIComponent(name);
  });
}

$(document).ready(function () {
  var c = config.get();
  if (! c.usePrivateRegistry) {
    window.location.href = '/images/';
    return;
  }
  $('#menu-repositories, #menu-prv-repo').addClass('active');
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
    $('#data').fadeIn();
  });
});
// </script>
