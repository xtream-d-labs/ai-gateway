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
    root.AiGateway.NotebookDetail = factory(root.AiGateway.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The NotebookDetail model module.
   * @module model/NotebookDetail
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>NotebookDetail</code>.
   * Rescale Job information
   * @alias module:model/NotebookDetail
   * @class
   * @param id {String} the container ID
   * @param token {String} Jupyter notebook's login token
   */
  var exports = function(id, token) {
    var _this = this;

    _this['id'] = id;



    _this['token'] = token;



  };

  /**
   * Constructs a <code>NotebookDetail</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/NotebookDetail} obj Optional instance to populate.
   * @return {module:model/NotebookDetail} The populated <code>NotebookDetail</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
      if (data.hasOwnProperty('state')) {
        obj['state'] = ApiClient.convertToType(data['state'], 'String');
      }
      if (data.hasOwnProperty('port')) {
        obj['port'] = ApiClient.convertToType(data['port'], 'Number');
      }
      if (data.hasOwnProperty('token')) {
        obj['token'] = ApiClient.convertToType(data['token'], 'String');
      }
      if (data.hasOwnProperty('mounts')) {
        obj['mounts'] = ApiClient.convertToType(data['mounts'], ['String']);
      }
      if (data.hasOwnProperty('started')) {
        obj['started'] = ApiClient.convertToType(data['started'], 'Date');
      }
      if (data.hasOwnProperty('ended')) {
        obj['ended'] = ApiClient.convertToType(data['ended'], 'Date');
      }
    }
    return obj;
  }

  /**
   * the container ID
   * @member {String} id
   */
  exports.prototype['id'] = undefined;
  /**
   * the container name
   * @member {String} name
   */
  exports.prototype['name'] = undefined;
  /**
   * state of the container
   * @member {String} state
   */
  exports.prototype['state'] = undefined;
  /**
   * the container published port
   * @member {Number} port
   */
  exports.prototype['port'] = undefined;
  /**
   * Jupyter notebook's login token
   * @member {String} token
   */
  exports.prototype['token'] = undefined;
  /**
   * the container labels
   * @member {Array.<String>} mounts
   */
  exports.prototype['mounts'] = undefined;
  /**
   * started unix timestamp
   * @member {Date} started
   */
  exports.prototype['started'] = undefined;
  /**
   * ended unix timestamp
   * @member {Date} ended
   */
  exports.prototype['ended'] = undefined;



  return exports;
}));


