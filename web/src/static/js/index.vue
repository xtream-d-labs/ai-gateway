// <script>
Vue.use(vuelidate.default);

var apiWorks = false;

API('App').getConfigurations(function (err, _, res) {
  if (! $.isEmptyObject(err) || ! res.body) {
    setTimeout(function () {window.location.reload();}, 5000);
    $("#api-server-error").fadeIn();
    return;
  }
  $('#act-signin').removeClass('disabled');
  apiWorks = true;

  config.set(res.body);
  if (! config.get().mustSignedIn) {
    window.location.href = '/images/';
  }
});

API('App').getEndpoints(function (err, _, res) {
  if (! $.isEmptyObject(err) || ! res.body) {
    return;
  }
  var registry = res.body.docker_registry;
  if (registry == "https://registry-1.docker.io") {
    registry = "DockerHub";
  }
  $('#endpoint').text(' on '+registry);
});

var vue = new Vue({
  el: '.form-signin',
  data: {
    username: '',
    password: '',
    msgUsername: '',
    msgPassword: ''
  },
  validations: {
    username: {
      required: validators.required,
      maxLength: validators.maxLength(256)
    },
    password: {
      required: validators.required,
      maxLength: validators.maxLength(256)
    }
  },
  methods: {
    usernameChanged: function () {
      this.msgUsername = '';
    },
    passwordChanged: function () {
      this.msgPassword = '';
    },
    error: function (message) {
      this.msgUsername = message ? message : 'The username or password is incorrect';
      $('#input-username').focus();
    },
    passwordError: function (message) {
      this.msgPassword = message;
      this.password = '';
      $('#input-password').val('').focus();
    },
    submit: function () {
      if (! apiWorks) {
        return;
      }
      setTimeout(function() {$('.btn-primary').blur();}, 500);
      if ($('.btn-primary').attr('disabled')) {
        return;
      }
      $('.btn-primary').attr('disabled', 'disabled');
      $('form p.errors').show();

      if (this.password == '') {
        vue.passwordError('Input your password');
      } else if ($('#input-password').hasClass('invalid')) {
        vue.passwordError('Invalid: password');
      }
      if (this.username == '') {
        vue.error('Input your user name');
      } else if ($('#input-username').hasClass('invalid')) {
        vue.error('Invalid: Username');
      }
      if ((this.username == '') || $('#input-username').hasClass('invalid') ||
          (this.password == '') || $('#input-password').hasClass('invalid')) {
        $('.btn-primary').removeAttr('disabled');
        return;
      }
      auth.clear();

      var body = models.AccountInfo.constructFromObject({
        docker_username: this.username,
        docker_password: this.password,
      });
      API('App').postNewSession(body, function (error, data, response) {
        $('.btn-primary').removeAttr('disabled');
        if (! $.isEmptyObject(error)) {
          vue.error('Username or password is invalid.');
          return;
        }
        if (! response || ! response.body || ! response.body.token) {
          vue.error('Username or password is invalid.');
          return;
        }

        var configs = config.get();
        configs.docker_username = vue.username;
        configs.use_private_registry = "yes";
        configs.usePrivateRegistry = true;
        config.set(configs);

        auth.login(response, vue.error, function () {
          window.location.href = '/images/';
        });
      });
    }
  }
});
$(document).ready(function () {
  $('#input-password').keyup(function (e) {
    if (e.keyCode == 13 && $(this).val() != '') {
      vue.submit();
    }
  });
  $(".form-signin .errors").fadeIn();
  $('#input-username').focus();
});
// </script>
