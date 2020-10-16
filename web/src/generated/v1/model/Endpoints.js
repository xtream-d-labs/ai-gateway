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
    root.AiGateway.Endpoints = factory(root.AiGateway.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The Endpoints model module.
   * @module model/Endpoints
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Endpoints</code>.
   * third-party endpoints
   * @alias module:model/Endpoints
   * @class
   */
  var exports = function() {
    var _this = this;





  };

  /**
   * Constructs a <code>Endpoints</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Endpoints} obj Optional instance to populate.
   * @return {module:model/Endpoints} The populated <code>Endpoints</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('docker_registry')) {
        obj['docker_registry'] = ApiClient.convertToType(data['docker_registry'], 'String');
      }
      if (data.hasOwnProperty('ngc_registry')) {
        obj['ngc_registry'] = ApiClient.convertToType(data['ngc_registry'], 'String');
      }
      if (data.hasOwnProperty('kubernetes_api')) {
        obj['kubernetes_api'] = ApiClient.convertToType(data['kubernetes_api'], 'String');
      }
      if (data.hasOwnProperty('rescale_api')) {
        obj['rescale_api'] = ApiClient.convertToType(data['rescale_api'], 'String');
      }
    }
    return obj;
  }

  /**
   * The endpoint for private docker registry
   * @member {String} docker_registry
   */
  exports.prototype['docker_registry'] = undefined;
  /**
   * The endpoint for NGC registry
   * @member {String} ngc_registry
   */
  exports.prototype['ngc_registry'] = undefined;
  /**
   * The endpoint for Kubernetes API
   * @member {String} kubernetes_api
   */
  exports.prototype['kubernetes_api'] = undefined;
  /**
   * The endpoint for Rescale API
   * @member {String} rescale_api
   */
  exports.prototype['rescale_api'] = undefined;



  return exports;
}));


