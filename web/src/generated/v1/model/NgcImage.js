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
    root.AiGateway.NgcImage = factory(root.AiGateway.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The NgcImage model module.
   * @module model/NgcImage
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>NgcImage</code>.
   * NGC docker image information
   * @alias module:model/NgcImage
   * @class
   * @param tag {String} the image tag
   * @param size {Number} the image size
   * @param updated {Date} updated unix timestamp
   */
  var exports = function(tag, size, updated) {
    var _this = this;

    _this['tag'] = tag;
    _this['size'] = size;
    _this['updated'] = updated;
  };

  /**
   * Constructs a <code>NgcImage</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/NgcImage} obj Optional instance to populate.
   * @return {module:model/NgcImage} The populated <code>NgcImage</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('tag')) {
        obj['tag'] = ApiClient.convertToType(data['tag'], 'String');
      }
      if (data.hasOwnProperty('size')) {
        obj['size'] = ApiClient.convertToType(data['size'], 'Number');
      }
      if (data.hasOwnProperty('updated')) {
        obj['updated'] = ApiClient.convertToType(data['updated'], 'Date');
      }
    }
    return obj;
  }

  /**
   * the image tag
   * @member {String} tag
   */
  exports.prototype['tag'] = undefined;
  /**
   * the image size
   * @member {Number} size
   */
  exports.prototype['size'] = undefined;
  /**
   * updated unix timestamp
   * @member {Date} updated
   */
  exports.prototype['updated'] = undefined;



  return exports;
}));


