(function () {
var generated  = require('./v1/index');
var jwtDecode  = require('jwt-decode');

window.models = generated;

window.API = function (name) {
  var api = new generated[name+'Api']();

  // base path
  var base = document.getElementById("api-base");
  if (base) api.apiClient.basePath = base.value;

  // api token
  var auth = localStorage.getItem('auth');
  if (auth) api.apiClient.defaultHeaders = {'Authorization': auth};

  api.apiClient.timeout = 1000 * 15; // 15sec
  api.apiClient.cache = false;
  return api;
}

function login(response, fail, success) {
  if (! response || ! response.body || ! response.body.token) {
    clear();
    fail();
    return;
  }
  localStorage.setItem('auth', response.body.token);
  success();
}

function jwt() {
  var auth = localStorage.getItem('auth');
  var jwt = auth ? jwtDecode(auth) : {claims: ''};
  if (jwt.claims == '') {
    return empty();
  }
  var sess = jwt.claims;
  return {
    signedIn: true,
    username: sess.docker_username
  };
  function empty() {
    return {
      signedIn: false,
      username: '---'
    };
  }
}

function setConfig(value) {
  var json = JSON.stringify(value);
  localStorage.setItem('configs', json);
}

function getConfig() {
  var configs = JSON.parse(localStorage.getItem('configs')) || {};
  configs.mustSignedIn = (configs && configs.must_signed_in == "yes");
  configs.usePrivateRegistry = (configs && configs.use_private_registry == "yes");
  configs.useNGC = (configs && configs.use_ngc == "yes");
  configs.useRescale = (configs && configs.use_rescale == "yes");
  configs.useKubernetes = (configs && configs.use_k8s == "yes");
  return configs;
}

function clear() {
  localStorage.removeItem('auth');
}

function exit() {
  clear();
  if (config.get().mustSignedIn) {
    window.location.href = '/';
    return;
  }
  window.location.href = '/settings/';
}

window.auth = {
  login: login,
  session: jwt,
  clear: clear,
  exit: exit
}

window.config = {
  set: setConfig,
  get: getConfig
}
})();
