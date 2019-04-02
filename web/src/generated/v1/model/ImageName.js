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
    root.ScaleShift.ImageName = factory(root.ScaleShift.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The ImageName model module.
   * @module model/ImageName
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>ImageName</code>.
   * @alias module:model/ImageName
   * @class
   * @param image {String} Docker image name
   */
  var exports = function(image) {
    var _this = this;

    _this['image'] = image;
  };

  /**
   * Constructs a <code>ImageName</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ImageName} obj Optional instance to populate.
   * @return {module:model/ImageName} The populated <code>ImageName</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('image')) {
        obj['image'] = ApiClient.convertToType(data['image'], 'String');
      }
    }
    return obj;
  }

  /**
   * Docker image name
   * @member {String} image
   */
  exports.prototype['image'] = undefined;



  return exports;
}));


