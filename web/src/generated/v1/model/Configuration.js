/**
 * AI Gateway
 * A platform for machine learning & high performance computing 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.AiGateway) {
      root.AiGateway = {};
    }
    root.AiGateway.Configuration = factory(root.AiGateway.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The Configuration model module.
   * @module model/Configuration
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Configuration</code>.
   * app&#39;s configurations
   * @alias module:model/Configuration
   * @class
   * @param mustSignedIn {module:model/Configuration.MustSignedInEnum} Users should be signed in
   */
  var exports = function(mustSignedIn) {
    var _this = this;

    _this['must_signed_in'] = mustSignedIn;














  };

  /**
   * Constructs a <code>Configuration</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Configuration} obj Optional instance to populate.
   * @return {module:model/Configuration} The populated <code>Configuration</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('must_signed_in')) {
        obj['must_signed_in'] = ApiClient.convertToType(data['must_signed_in'], 'String');
      }
      if (data.hasOwnProperty('use_private_registry')) {
        obj['use_private_registry'] = ApiClient.convertToType(data['use_private_registry'], 'String');
      }
      if (data.hasOwnProperty('use_ngc')) {
        obj['use_ngc'] = ApiClient.convertToType(data['use_ngc'], 'String');
      }
      if (data.hasOwnProperty('use_k8s')) {
        obj['use_k8s'] = ApiClient.convertToType(data['use_k8s'], 'String');
      }
      if (data.hasOwnProperty('use_rescale')) {
        obj['use_rescale'] = ApiClient.convertToType(data['use_rescale'], 'String');
      }
      if (data.hasOwnProperty('docker_registry')) {
        obj['docker_registry'] = ApiClient.convertToType(data['docker_registry'], 'String');
      }
      if (data.hasOwnProperty('docker_hostname')) {
        obj['docker_hostname'] = ApiClient.convertToType(data['docker_hostname'], 'String');
      }
      if (data.hasOwnProperty('docker_username')) {
        obj['docker_username'] = ApiClient.convertToType(data['docker_username'], 'String');
      }
      if (data.hasOwnProperty('docker_password')) {
        obj['docker_password'] = ApiClient.convertToType(data['docker_password'], 'String');
      }
      if (data.hasOwnProperty('ngc_email')) {
        obj['ngc_email'] = ApiClient.convertToType(data['ngc_email'], 'String');
      }
      if (data.hasOwnProperty('ngc_password')) {
        obj['ngc_password'] = ApiClient.convertToType(data['ngc_password'], 'String');
      }
      if (data.hasOwnProperty('ngc_apikey')) {
        obj['ngc_apikey'] = ApiClient.convertToType(data['ngc_apikey'], 'String');
      }
      if (data.hasOwnProperty('k8s_config')) {
        obj['k8s_config'] = ApiClient.convertToType(data['k8s_config'], 'String');
      }
      if (data.hasOwnProperty('rescale_platform')) {
        obj['rescale_platform'] = ApiClient.convertToType(data['rescale_platform'], 'String');
      }
      if (data.hasOwnProperty('rescale_key')) {
        obj['rescale_key'] = ApiClient.convertToType(data['rescale_key'], 'String');
      }
    }
    return obj;
  }

  /**
   * Users should be signed in
   * @member {module:model/Configuration.MustSignedInEnum} must_signed_in
   */
  exports.prototype['must_signed_in'] = undefined;
  /**
   * Private registry will be used or not
   * @member {module:model/Configuration.UsePrivateRegistryEnum} use_private_registry
   */
  exports.prototype['use_private_registry'] = undefined;
  /**
   * NGC will be used or not
   * @member {module:model/Configuration.UseNgcEnum} use_ngc
   */
  exports.prototype['use_ngc'] = undefined;
  /**
   * Kubernetes will be used or not
   * @member {module:model/Configuration.UseK8sEnum} use_k8s
   */
  exports.prototype['use_k8s'] = undefined;
  /**
   * Rescale will be used or not
   * @member {module:model/Configuration.UseRescaleEnum} use_rescale
   */
  exports.prototype['use_rescale'] = undefined;
  /**
   * Docker Registry endpoint
   * @member {String} docker_registry
   */
  exports.prototype['docker_registry'] = undefined;
  /**
   * Hostname for the private Docker registry
   * @member {String} docker_hostname
   */
  exports.prototype['docker_hostname'] = undefined;
  /**
   * Username for the private Docker registry
   * @member {String} docker_username
   */
  exports.prototype['docker_username'] = undefined;
  /**
   * Fist 3 chars of the password of the private Docker registry
   * @member {String} docker_password
   */
  exports.prototype['docker_password'] = undefined;
  /**
   * E-mail address for NGC console
   * @member {String} ngc_email
   */
  exports.prototype['ngc_email'] = undefined;
  /**
   * Fist 3 chars of the password for NGC console
   * @member {String} ngc_password
   */
  exports.prototype['ngc_password'] = undefined;
  /**
   * Fist 5 chars of NGC API Key
   * @member {String} ngc_apikey
   */
  exports.prototype['ngc_apikey'] = undefined;
  /**
   * kubecfg
   * @member {String} k8s_config
   */
  exports.prototype['k8s_config'] = undefined;
  /**
   * @member {module:model/Configuration.RescalePlatformEnum} rescale_platform
   */
  exports.prototype['rescale_platform'] = undefined;
  /**
   * Fist 5 chars of Rescal API Key
   * @member {String} rescale_key
   */
  exports.prototype['rescale_key'] = undefined;


  /**
   * Allowed values for the <code>must_signed_in</code> property.
   * @enum {String}
   * @readonly
   */
  exports.MustSignedInEnum = {
    /**
     * value: "yes"
     * @const
     */
    "yes": "yes",
    /**
     * value: "no"
     * @const
     */
    "no": "no"  };

  /**
   * Allowed values for the <code>use_private_registry</code> property.
   * @enum {String}
   * @readonly
   */
  exports.UsePrivateRegistryEnum = {
    /**
     * value: "yes"
     * @const
     */
    "yes": "yes",
    /**
     * value: "no"
     * @const
     */
    "no": "no"  };

  /**
   * Allowed values for the <code>use_ngc</code> property.
   * @enum {String}
   * @readonly
   */
  exports.UseNgcEnum = {
    /**
     * value: "yes"
     * @const
     */
    "yes": "yes",
    /**
     * value: "no"
     * @const
     */
    "no": "no"  };

  /**
   * Allowed values for the <code>use_k8s</code> property.
   * @enum {String}
   * @readonly
   */
  exports.UseK8sEnum = {
    /**
     * value: "yes"
     * @const
     */
    "yes": "yes",
    /**
     * value: "no"
     * @const
     */
    "no": "no"  };

  /**
   * Allowed values for the <code>use_rescale</code> property.
   * @enum {String}
   * @readonly
   */
  exports.UseRescaleEnum = {
    /**
     * value: "yes"
     * @const
     */
    "yes": "yes",
    /**
     * value: "no"
     * @const
     */
    "no": "no"  };

  /**
   * Allowed values for the <code>rescale_platform</code> property.
   * @enum {String}
   * @readonly
   */
  exports.RescalePlatformEnum = {
    /**
     * value: "https://platform.rescale.com"
     * @const
     */
    "platform.rescale.com": "https://platform.rescale.com",
    /**
     * value: "https://platform.rescale.jp"
     * @const
     */
    "platform.rescale.jp": "https://platform.rescale.jp",
    /**
     * value: "https://kr.rescale.com"
     * @const
     */
    "kr.rescale.com": "https://kr.rescale.com",
    /**
     * value: "https://eu.rescale.com"
     * @const
     */
    "eu.rescale.com": "https://eu.rescale.com"  };


  return exports;
}));


