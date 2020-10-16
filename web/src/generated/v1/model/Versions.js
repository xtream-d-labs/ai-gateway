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
    define(['ApiClient', 'model/Version'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Version'));
  } else {
    // Browser globals (root is window)
    if (!root.AiGateway) {
      root.AiGateway = {};
    }
    root.AiGateway.Versions = factory(root.AiGateway.ApiClient, root.AiGateway.Version);
  }
}(this, function(ApiClient, Version) {
  'use strict';




  /**
   * The Versions model module.
   * @module model/Versions
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Versions</code>.
   * application versions
   * @alias module:model/Versions
   * @class
   * @param current {module:model/Version} Current running service version
   * @param latest {module:model/Version} The latest application version which can be installed
   */
  var exports = function(current, latest) {
    var _this = this;

    _this['current'] = current;
    _this['latest'] = latest;
  };

  /**
   * Constructs a <code>Versions</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Versions} obj Optional instance to populate.
   * @return {module:model/Versions} The populated <code>Versions</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('current')) {
        obj['current'] = Version.constructFromObject(data['current']);
      }
      if (data.hasOwnProperty('latest')) {
        obj['latest'] = Version.constructFromObject(data['latest']);
      }
    }
    return obj;
  }

  /**
   * Current running service version
   * @member {module:model/Version} current
   */
  exports.prototype['current'] = undefined;
  /**
   * The latest application version which can be installed
   * @member {module:model/Version} latest
   */
  exports.prototype['latest'] = undefined;



  return exports;
}));


