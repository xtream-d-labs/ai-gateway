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
    root.AiGateway.IPythonNotebook = factory(root.AiGateway.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The IPythonNotebook model module.
   * @module model/IPythonNotebook
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>IPythonNotebook</code>.
   * .ipynb file attributes
   * @alias module:model/IPythonNotebook
   * @class
   * @param name {String} file name
   */
  var exports = function(name) {
    var _this = this;

    _this['name'] = name;
  };

  /**
   * Constructs a <code>IPythonNotebook</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/IPythonNotebook} obj Optional instance to populate.
   * @return {module:model/IPythonNotebook} The populated <code>IPythonNotebook</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
    }
    return obj;
  }

  /**
   * file name
   * @member {String} name
   */
  exports.prototype['name'] = undefined;



  return exports;
}));


