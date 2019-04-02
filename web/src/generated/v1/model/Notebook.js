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
    root.ScaleShift.Notebook = factory(root.ScaleShift.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The Notebook model module.
   * @module model/Notebook
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Notebook</code>.
   * Jupyter notebook information
   * @alias module:model/Notebook
   * @class
   * @param id {String} the container ID
   * @param name {String} the container name
   * @param image {String} the image ID
   */
  var exports = function(id, name, image) {
    var _this = this;

    _this['id'] = id;
    _this['name'] = name;
    _this['image'] = image;



  };

  /**
   * Constructs a <code>Notebook</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Notebook} obj Optional instance to populate.
   * @return {module:model/Notebook} The populated <code>Notebook</code> instance.
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
      if (data.hasOwnProperty('image')) {
        obj['image'] = ApiClient.convertToType(data['image'], 'String');
      }
      if (data.hasOwnProperty('state')) {
        obj['state'] = ApiClient.convertToType(data['state'], 'String');
      }
      if (data.hasOwnProperty('port')) {
        obj['port'] = ApiClient.convertToType(data['port'], 'Number');
      }
      if (data.hasOwnProperty('started')) {
        obj['started'] = ApiClient.convertToType(data['started'], 'Date');
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
   * the image ID
   * @member {String} image
   */
  exports.prototype['image'] = undefined;
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
   * started unix timestamp
   * @member {Date} started
   */
  exports.prototype['started'] = undefined;



  return exports;
}));


