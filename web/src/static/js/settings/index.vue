// <script>
Vue.use(vuelidate.default);

API('App').getConfigurations(function (err, _, res) {
  if (! $.isEmptyObject(err) || ! res || ! res.body) {
    return;
  }
  vue.registry    = res.body.docker_registry;
  vue.hostname    = res.body.docker_hostname;
  vue.username    = res.body.docker_username;
  vue.password    = res.body.docker_password;
  vue.ngcEmail    = res.body.ngc_email;
  vue.ngcPassword = res.body.ngc_password;
  vue.ngcKey      = res.body.ngc_apikey;
  vue.ngcKey      = res.body.ngc_apikey;
  vue.k8sConfig   = res.body.k8s_config;
  vue.rescaleKey  = res.body.rescale_key;

  config.set(res.body);
  app.refreshMenus(res.body);

  var select = $('#input-rescale-platform').hide().val(res.body.rescale_platform);
  select.formSelect();
  select.closest('.form-group').fadeIn();

  $('.errors').show();
});

var vue = new Vue({
  el: '.form-signin',
  data: {
    registry: '',
    hostname: '',
    username: '',
    password: '',
    ngcEmail: '',
    ngcPassword: '',
    ngcKey: '',
    k8sConfig: '',
    rescaleKey: '',
    message: ''
  },
  methods: {
    error: function (message) {
      this.message = message ? message : 'The username or password is incorrect';
      $('#input-username').focus();
    },
    submit: function () {
      setTimeout(function() {$('.btn-primary').blur();}, 500);
      if ($('.btn-primary').attr('disabled')) {
        return;
      }
      $('.btn-primary').attr('disabled', 'disabled');
      $('form p.errors').show();

      var body = models.AccountInfo.constructFromObject({
        docker_registry:  this.registry,
        docker_hostname:  this.hostname,
        docker_username:  this.username,
        docker_password:  this.password,
        ngc_email:        this.ngcEmail,
        ngc_password:     this.ngcPassword,
        ngc_apikey:       this.ngcKey,
        k8s_config:       this.k8sConfig,
        rescale_platform: $('#input-rescale-platform option:selected').val(),
        rescale_key:      this.rescaleKey
      });
      API('App').postConfigurations(body, function (err, data, res) {
        $('.btn-primary').removeAttr('disabled');

        if (! $.isEmptyObject(err)) {
          vue.error(err.response.body.message);
          if (err.response.body.message.indexOf('Docker registry') > -1) {
            $('#input-username').focus();
          }
          if (err.response.body.message.indexOf('NGC Email') > -1) {
            $('#input-ngc-email').focus();
          }
          if (err.response.body.message.indexOf('NGC API Key') > -1) {
            $('#input-ngc-key').focus();
          }
          return;
        }
        if (! res || ! res.body || ! res.body.token) {
          vue.error('Invalid parameters');
          return;
        }
        auth.clear();
        auth.login(res, vue.error, function () {
          alert('Saved successfully');
          window.location.reload();
        });
      });
    }
  }
});
$(document).ready(function () {
  $('#input-username').focus();
});
// </script>
