// <script>
var images = [];
var conditions = {
  firstLoad: true,
  words: '',
  category: 'all'
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
        if (conditions.category != 'all') {
          if (conditions.category != image.namespace) {
            return;
          }
        }
        filtered.push({
          code: image.namespace + '/' + image.name,
          namespace: image.namespace,
          name: image.name,
          description: marked(image.description)
        });
      });
      filtered.sort(function (a, b) {
        if (group(a.namespace) < group(b.namespace)) return -1;
        if (group(a.namespace) > group(b.namespace)) return  1;
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

function group(namespace) {
  switch (namespace) {
  case 'nvidia':        return 1;
  case 'hpc':           return 2;
  case 'nvidia-hpcvis': return 3;
  case 'partners':      return 4;
  }
  return 9;
}

function update() {
  conditions.words = app.singleSpace($('#query-words').val());
  var candidate = $('#categories select').val();
  conditions.category = candidate ? candidate : 'all';
  vue.update();
}

function load(callback) {
  API('Repository').getNgcRepositories(function (err, _, res) {
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
  var ns = $(el).attr('data-ns'),
      nm = $(el).attr('data-nm');
  API('Repository').getNgcImages(ns, nm, function (err, _, res) {
    if (app.shouldExit(res, err)) {
      alert('Something went wrong. Check your configurations!')
      window.location.href = '/settings/';
      return;
    }
    var html = '';
    $.map(res.body, function (image) {
      html += '<tr data-id="' + ns +'/'+ nm +':'+ image.tag +'">';
        html += '<td>'+image.tag+'</td>';
        html += '<td>'+app.comma(image.size, 'byte')+'</td>';
        html += '<td>'+app.date.format(new Date(image.updated))+'</td>';
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
  var body = new models.ImageName('nvcr.io/' + name);
  API('Image').postNewImage(body, function (err, _, res) {
    if (app.shouldExit(res, err)) {
      alert('Something went wrong. Check your configurations!')
      window.location.href = '/settings/';
      return;
    }
    location.href = '/images/?q=' + encodeURIComponent(name);
  });
}

function setCategories() {
  var temp = {};
  $.map(images, function (image) {
    temp[image.namespace] = true;
  });
  var categories = [];
  $.map(temp, function (_, key) {
    categories.push(key);
  });
  categories.sort(function (a, b) {
    if (a < b) return -1;
    if (a > b) return  1;
    return 0;
  });
  var children = '';
  $.map(categories, function (category) {
    children += '<option value="' + category +'" class="hidden-xs"' +
            ((category.toLowerCase() == 'nvidia') ? ' selected="selected"' : '') +
            '>' + category + '</option>';
  });
  children = '<option value="all" class="hidden-xs">all</option>' + children;
  var select = $('#categories select').html(children);
  select.change(function () {
    update();
  });
  select.formSelect();

  conditions.category = select.val();
  update();
}

$(document).ready(function () {
  var c = config.get();
  if (! c.useNGC) {
    window.location.href = '/images/';
    return;
  }
  $('#menu-repositories, #menu-ngc-repo').addClass('active');
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
    setCategories();
    $('#data').fadeIn();
  });
});
// </script>
