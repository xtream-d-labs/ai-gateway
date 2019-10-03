// <script>
var errors = [];
var conditions = {
  firstLoad: true,
  words: ''
};
var vue = new Vue({
  el: '#data',
  data: {
    errors: []
  },
  methods: {
    update: function () {
      var filtered = [];
      $.map(errors, function (e) {
        if (conditions.words != '') {
          if (! app.match([e.caption, e.detail, e.owner], conditions.words)) {
            return;
          }
        }
        filtered.push(e);
      });
      filtered.sort(function (a, b) {
        return new Date(b.occursAt).getTime() - new Date(a.occursAt).getTime();
      });
      var formatted = [];
      $.map(filtered, function (e) {
        formatted.push({
          id:       formatted.length,
          caption:  e.caption,
          detail:   e.detail,
          owner:    (e.owner == 'anonymous') ? '-' : e.owner,
          occursAt: app.date.format(new Date(e.occursAt))
        });
      });
      this.errors = formatted;

      if (conditions.firstLoad) {
        conditions.firstLoad = false;
        if (conditions.words != '' && formatted.length > 0) {
          setTimeout(function () {
            $('#data .collapsible .row-body').eq(0).collapse('show');
            $('#query-words').focus();
          }, 750);
        }
      }
      $('#record-count').text(formatted.length);
    }
  }
});

function update() {
  conditions.words = app.singleSpace($('#query-words').val());
  vue.update();
}

function load(callback) {
  API('AppErrors').getAppErrors(function (error, data, response) {
    if (! $.isEmptyObject(error) || ! response || ! response.body) {
      return;
    }
    errors = response.body;
    update();
    callback && callback();
  });
}

$(document).ready(function () {
  $('#menu-system, #menu-system-error').addClass('active');
  if (app.query('q')) {
    $('#query-words').val(app.query('q')).focus();
  }
  load(function () {
    $('#query-words').keyup(function () {
      update();
    });
    $('#data').fadeIn();
  });
  setInterval(load, 30 * 1000);
});
// </script>
