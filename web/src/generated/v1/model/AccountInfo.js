/**
 * ScaleShift
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
    if (!root.ScaleShift) {
      root.ScaleShift = {};
    }
    root.ScaleShift.AccountInfo = factory(root.ScaleShift.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The AccountInfo model module.
   * @module model/AccountInfo
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>AccountInfo</code>.
   * @alias module:model/AccountInfo
   * @class
   */
  var exports = function() {
    var _this = this;











  };

  /**
   * Constructs a <code>AccountInfo</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/AccountInfo} obj Optional instance to populate.
   * @return {module:model/AccountInfo} The populated <code>AccountInfo</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

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
   * Password for the private Docker registry
   * @member {String} docker_password
   */
  exports.prototype['docker_password'] = undefined;
  /**
   * E-mail address for NGC console
   * @member {String} ngc_email
   */
  exports.prototype['ngc_email'] = undefined;
  /**
   * Password for NGC console
   * @member {String} ngc_password
   */
  exports.prototype['ngc_password'] = undefined;
  /**
   * NGC - API Key
   * @member {String} ngc_apikey
   */
  exports.prototype['ngc_apikey'] = undefined;
  /**
   * kubecfg
   * @member {String} k8s_config
   */
  exports.prototype['k8s_config'] = undefined;
  /**
   * Rescale platform endopoint
   * @member {module:model/AccountInfo.RescalePlatformEnum} rescale_platform
   */
  exports.prototype['rescale_platform'] = undefined;
  /**
   * Rescale - API Key
   * @member {String} rescale_key
   */
  exports.prototype['rescale_key'] = undefined;


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

