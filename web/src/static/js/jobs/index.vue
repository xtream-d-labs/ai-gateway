// <script>
var jobs = [];
var conditions = {
  firstLoad: true,
  words: '',
  order: 1
};
var vue = new Vue({
  el: '#data',
  data: {
    jobs: []
  },
  methods: {
    update: function () {
      var filtered = [];
      $.map(jobs, function (job) {
        if (conditions.words != '') {
          if (! app.match([job.image, job.id, job.ipynb, job.status], conditions.words)) {
            return;
          }
        }
        filtered.push(job);
      });
      filtered.sort(function (a, b) {
        if (conditions.order == 1) {
          var ret = statusOrder(a.status) - statusOrder(b.status);
          if (ret != 0) return ret;
        }
        ret = new Date(b.started).getTime() - new Date(a.started).getTime();
        if (ret != 0) return ret;
        return statusOrder(a.status) - statusOrder(b.status);
      });
      var formatted = [];
      $.map(filtered, function (job) {
        var mounts = (job.mounts.length > 0) ? job.mounts[0] : '';
        if (mounts.length > 0) {
          mounts = '<a href="/workspaces/?q='+encodeURIComponent(mounts)+'">'+mounts+'</a>';
        }
        formatted.push({
          id:          job.id,
          image:       job.image,
          imageHref:   '<a href="/images/?q='+encodeURIComponent(job.image)+'">'+job.image+'</a>',
          commands:    job.commands.join(' '),
          mounts:      mounts,
          status:      statusValue(job.status),
          statusMore:  statusMore(job.status),
          started:     app.date.format(new Date(job.started)),
          classObject: statusClass(job.status)
        });
      });
      this.jobs = formatted;

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
    },
    del: function (e) {
      confirmation.action = 'REMOVE';
      confirmation.id = app.trim($(e.target).closest('li').attr('data-id'));
      $('#job-modify').modal('open');
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
  case 'rescale-start':
    return 'running';
  case 'k8s-job-end':
  case 'rescale-end':
    return 'done';
  }
  return 'unknown';
}

function statusMore(status) {
  switch (status) {
  case 'building-job':
    return 'Building a docker/singularity image for the job';
  case 'pushing-job':
  case 'k8s-job':
  case 'rescale-send':
    return 'Sending files to the cloud';
  case 'k8s-job-start':
  case 'rescale-start':
    return 'Running the job on the cloud';
  case 'k8s-job-end':
  case 'rescale-end':
    return 'Done';
  }
  return 'unknown';
}

function statusOrder(status) {
  switch (status) {
  case 'building-job':
    return 1;
  case 'pushing-job':
  case 'k8s-job':
  case 'rescale-send':
    return 2;
  case 'k8s-job-start':
  case 'rescale-start':
    return 3;
  case 'k8s-job-end':
  case 'rescale-end':
    return 4;
  }
  return 9;
}

function statusClass(state) {
  switch (state) {
  case 'building-job':
    return {
        'label-info':    true,
        'label-warning': false,
        'label-success': false,
        'label-danger':  false
    };
  case 'pushing-job':
  case 'k8s-job':
  case 'rescale-send':
    return {
        'label-info':    false,
        'label-warning': true,
        'label-success': false,
        'label-danger':  false
    };
  case 'k8s-job-start':
  case 'rescale-start':
    return {
        'label-info':    false,
        'label-warning': false,
        'label-success': true,
        'label-danger':  false
    };
  case 'k8s-job-end':
  case 'rescale-end':
    return {
        'label-info':    false,
        'label-warning': false,
        'label-success': false,
        'label-danger':  true
    };
  }
  return {
      'label-info':    false,
      'label-warning': false,
      'label-success': false,
      'label-danger':  true
  };
}

var confirmation = new Vue({
  el: '#job-modify',
  data: {
    action: 'REMOVE',
    id: ''
  },
  methods: {
    exec: function () {
      $('#job-modify').modal('close');
      switch (this.action) {
      case 'REMOVE':
        API('Job').deleteJob(this.id, function (err, _, res) {
          if (! $.isEmptyObject(err)) {
            var message = 'Could not remove the specified job';
            if (res && res.body && res.body.message) {
              message = res.body.message;
            }
            M.toast({html: message, displayLength: 3000});
            return;
          }
          M.toast({html: 'Removed successfully. [ '+confirmation.id+' ]', displayLength: 3000});
          load();
        });
        break;
      }
    },
    close: function () {
      $('#job-modify').modal('close');
    }
  }
});

function update() {
  conditions.words = app.singleSpace($('#query-words').val());
  conditions.order = parseInt($('#query-order-type').val(), 10);
  vue.update();
}

function load(callback) {
  API('Job').getJobs(function (error, data, response) {
    if (! $.isEmptyObject(error) || ! response || ! response.body) {
      return;
    }
    jobs = response.body;
    update();
    callback && callback();
  });
}

$(document).ready(function () {
  $('#menu-jobs').addClass('active');
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
    $('#data').fadeIn();
  });
  setInterval(load, 15 * 1000);
});
// </script>
