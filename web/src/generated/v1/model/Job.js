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
    root.AiGateway.Job = factory(root.AiGateway.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The Job model module.
   * @module model/Job
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Job</code>.
   * Rescale Job information
   * @alias module:model/Job
   * @class
   * @param id {String} Job ID
   */
  var exports = function(id) {
    var _this = this;

    _this['id'] = id;








  };

  /**
   * Constructs a <code>Job</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Job} obj Optional instance to populate.
   * @return {module:model/Job} The populated <code>Job</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('platform')) {
        obj['platform'] = ApiClient.convertToType(data['platform'], 'String');
      }
      if (data.hasOwnProperty('status')) {
        obj['status'] = ApiClient.convertToType(data['status'], 'String');
      }
      if (data.hasOwnProperty('image')) {
        obj['image'] = ApiClient.convertToType(data['image'], 'String');
      }
      if (data.hasOwnProperty('mounts')) {
        obj['mounts'] = ApiClient.convertToType(data['mounts'], ['String']);
      }
      if (data.hasOwnProperty('commands')) {
        obj['commands'] = ApiClient.convertToType(data['commands'], ['String']);
      }
      if (data.hasOwnProperty('started')) {
        obj['started'] = ApiClient.convertToType(data['started'], 'Date');
      }
      if (data.hasOwnProperty('ended')) {
        obj['ended'] = ApiClient.convertToType(data['ended'], 'Date');
      }
      if (data.hasOwnProperty('external_link')) {
        obj['external_link'] = ApiClient.convertToType(data['external_link'], 'String');
      }
    }
    return obj;
  }

  /**
   * Job ID
   * @member {String} id
   */
  exports.prototype['id'] = undefined;
  /**
   * platform
   * @member {module:model/Job.PlatformEnum} platform
   */
  exports.prototype['platform'] = undefined;
  /**
   * the status of the job
   * @member {String} status
   */
  exports.prototype['status'] = undefined;
  /**
   * the image ID
   * @member {String} image
   */
  exports.prototype['image'] = undefined;
  /**
   * the container labels
   * @member {Array.<String>} mounts
   */
  exports.prototype['mounts'] = undefined;
  /**
   * the container labels
   * @member {Array.<String>} commands
   */
  exports.prototype['commands'] = undefined;
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
  /**
   * A link to an external status page
   * @member {String} external_link
   */
  exports.prototype['external_link'] = undefined;


  /**
   * Allowed values for the <code>platform</code> property.
   * @enum {String}
   * @readonly
   */
  exports.PlatformEnum = {
    /**
     * value: "kubernetes"
     * @const
     */
    "kubernetes": "kubernetes",
    /**
     * value: "rescale"
     * @const
     */
    "rescale": "rescale"  };


  return exports;
}));


