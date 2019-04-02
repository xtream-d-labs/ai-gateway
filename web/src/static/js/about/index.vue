// <script>

var vue = new Vue({
  el: '#data',
  data: {
    current: {
      version: '',
      date: ''
    },
    latest: {
      version: '',
      date: ''
    }
  },
  methods: {
    update: function () {
      API('App').getVersions(function (err, _, res) {
        if (! $.isEmptyObject(err) || ! res || ! res.body) {
          return;
        }
        vue.current = {
          version: res.body.current.version,
          date: res.body.current.build_date
        };
        vue.latest = {
          version: res.body.latest.version,
          date: res.body.latest.build_date
        };
      });
      $('#data').fadeIn();
    }
  }
});

$(document).ready(function () {
   vue.update();
});
// </script>
