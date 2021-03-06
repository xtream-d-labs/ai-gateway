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
    define(['ApiClient', 'model/Error', 'model/Image', 'model/NgcImage', 'model/Repository'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Error'), require('../model/Image'), require('../model/NgcImage'), require('../model/Repository'));
  } else {
    // Browser globals (root is window)
    if (!root.AiGateway) {
      root.AiGateway = {};
    }
    root.AiGateway.RepositoryApi = factory(root.AiGateway.ApiClient, root.AiGateway.Error, root.AiGateway.Image, root.AiGateway.NgcImage, root.AiGateway.Repository);
  }
}(this, function(ApiClient, Error, Image, NgcImage, Repository) {
  'use strict';

  /**
   * Repository service.
   * @module api/RepositoryApi
   * @version 1.0.0
   */

  /**
   * Constructs a new RepositoryApi. 
   * @alias module:api/RepositoryApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the getNgcImages operation.
     * @callback module:api/RepositoryApi~getNgcImagesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/NgcImage>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns NGC images 
     * @param {String} namespace Docker repositry namespace
     * @param {String} id Docker image name
     * @param {module:api/RepositoryApi~getNgcImagesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/NgcImage>}
     */
    this.getNgcImages = function(namespace, id, callback) {
      var postBody = null;

      // verify the required parameter 'namespace' is set
      if (namespace === undefined || namespace === null) {
        throw new Error("Missing the required parameter 'namespace' when calling getNgcImages");
      }

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getNgcImages");
      }


      var pathParams = {
        'namespace': namespace,
        'id': id
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [NgcImage];

      return this.apiClient.callApi(
        '/nvidia/repositories/{namespace}/images/{id}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getNgcRepositories operation.
     * @callback module:api/RepositoryApi~getNgcRepositoriesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Repository>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns NGC repositories 
     * @param {module:api/RepositoryApi~getNgcRepositoriesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Repository>}
     */
    this.getNgcRepositories = function(callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [Repository];

      return this.apiClient.callApi(
        '/nvidia/repositories', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getRemoteImages operation.
     * @callback module:api/RepositoryApi~getRemoteImagesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Image>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns remote images 
     * @param {String} id Docker image name
     * @param {module:api/RepositoryApi~getRemoteImagesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Image>}
     */
    this.getRemoteImages = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getRemoteImages");
      }


      var pathParams = {
        'id': id
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [Image];

      return this.apiClient.callApi(
        '/remote-images/{id}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getRemoteRepositories operation.
     * @callback module:api/RepositoryApi~getRemoteRepositoriesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Repository>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns remote repositories 
     * @param {module:api/RepositoryApi~getRemoteRepositoriesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Repository>}
     */
    this.getRemoteRepositories = function(callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [Repository];

      return this.apiClient.callApi(
        '/repositories', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
