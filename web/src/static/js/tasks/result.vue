// <script>
var apiToken = '';

var vue = new Vue({
  el: '#form-job-result',
  data: {
    task: {
      statusMore: '',
      imageHref:  '',
      commands:   '',
      platform:   '',
      link:       '',
      mounts:     '',
      started:    '',
      ended:      ''
    }
  },
  methods: {
    update: function (data) {
      var mounts = (data.mounts.length > 0) ? data.mounts[0] : '';
      if (mounts.length > 0) {
        mounts = '<a href="/workspaces/?q='+encodeURIComponent(mounts)+'">'+mounts+'</a>';
      }
      this.task = {
        statusMore: statusValue(data.status),
        imageHref:  '<a href="/images/?q='+encodeURIComponent(data.image)+'">'+data.image+'</a>',
        commands:   data.commands.join(' '),
        platform:   data.platform,
        link:       data.external_link ? data.external_link : '',
        mounts:     mounts,
        started:    app.date.format(new Date(data.started)),
        ended:      app.date.format(new Date(data.ended))
      }
    }
  }
});
var logs = new Vue({
  el: '#logs',
  data: {
    logs: []
  },
  methods: {
    update: function (logs) {
      var formatted = [];
      $.map(logs, function (log, idx) {
        formatted.push({
          seq:  idx,
          log:  log.log,
          time: app.date.format(new Date(log.time))
        });
      });
      this.logs = formatted;
      if (formatted.length > 0) $('#logs').fadeIn();
    }
  }
});
var files = new Vue({
  el: '#files',
  data: {
    files: []
  },
  methods: {
    update: function (files) {
      var formatted = [];
      $.map(files, function (file, idx) {
        formatted.push({
          seq:  idx,
          name: file.name,
          size: app.comma(file.size, 'byte'),
          url:  file.downloadURL
        });
      });
      this.files = formatted;
      if (formatted.length > 0) $('#files').fadeIn();
    },
    download: function (e) {
      var xhr = new XMLHttpRequest;
      xhr.open("GET", $(e.target).closest('tr').attr('data-url'));
      xhr.addEventListener("load", function () {
        var data = toBinaryString(this.responseText);
        data = "data:application/pdf;base64,"+btoa(data);
        document.location = data;
      }, false);
      xhr.setRequestHeader("Authorization", "Token " + apiToken);
      xhr.overrideMimeType("application/octet-stream; charset=x-user-defined;");
      xhr.send(null);

      function toBinaryString(data) {
        var ret = [], len = data.length, byte;
        for (var i=0; i<len; i++) { 
          byte=( data.charCodeAt(i) & 0xFF )>>> 0;
          ret.push(String.fromCharCode(byte) );
        }
        return ret.join('');
      }
    }
  }
});

function statusValue(status) {
  switch (status) {
  case 'building-job':
    return 'preparing';
  case 'pushing-job':
  case 'k8s-job':
  case 'rescale-send':
    return 'sending';
  case 'k8s-job-start':
  case 'k8s-job-pending':
  case 'k8s-job-runnning':
  case 'rescale-start':
  case 'rescale-runnning':
    return 'running';
  case 'k8s-job-succeeded':
  case 'k8s-job-failed':
  case 'rescale-succeeded':
  case 'rescale-failed':
    return 'done';
  }
  return 'unknown';
}

$(document).ready(function () {
  $('#menu-tasks').addClass('active');
  if (! app.query('id')) {
    window.location.href = '/tasks/';
    return;
  }
  API('Job').getJobDetail(app.query('id'), function (err, _, res) {
    if (! $.isEmptyObject(err) || ! res || ! res.body) {
      return;
    }
    vue.update(res.body);
    logs.update(res.body.logs);
    files.update(res.body.files);
    apiToken = res.body.apiToken;
  
    $('#form-job-result>table').fadeIn();
    $('.wait-icon').hide();
  });
});
// </script>
