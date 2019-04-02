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
    define(['ApiClient', 'model/Error', 'model/IPythonNotebook', 'model/ImageName2', 'model/Notebook', 'model/NotebookAttrs', 'model/NotebookDetail'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Error'), require('../model/IPythonNotebook'), require('../model/ImageName2'), require('../model/Notebook'), require('../model/NotebookAttrs'), require('../model/NotebookDetail'));
  } else {
    // Browser globals (root is window)
    if (!root.ScaleShift) {
      root.ScaleShift = {};
    }
    root.ScaleShift.NotebookApi = factory(root.ScaleShift.ApiClient, root.ScaleShift.Error, root.ScaleShift.IPythonNotebook, root.ScaleShift.ImageName2, root.ScaleShift.Notebook, root.ScaleShift.NotebookAttrs, root.ScaleShift.NotebookDetail);
  }
}(this, function(ApiClient, Error, IPythonNotebook, ImageName2, Notebook, NotebookAttrs, NotebookDetail) {
  'use strict';

  /**
   * Notebook service.
   * @module api/NotebookApi
   * @version 1.0.0
   */

  /**
   * Constructs a new NotebookApi. 
   * @alias module:api/NotebookApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the deleteNotebook operation.
     * @callback module:api/NotebookApi~deleteNotebookCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * delete a specified notebook 
     * @param {String} id Notebook container ID
     * @param {module:api/NotebookApi~deleteNotebookCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteNotebook = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteNotebook");
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

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/notebooks/{id}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getIPythonNotebooks operation.
     * @callback module:api/NotebookApi~getIPythonNotebooksCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/IPythonNotebook>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns ipynb files on the specified notebook 
     * @param {String} id Notebook container ID
     * @param {module:api/NotebookApi~getIPythonNotebooksCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/IPythonNotebook>}
     */
    this.getIPythonNotebooks = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getIPythonNotebooks");
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

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [IPythonNotebook];

      return this.apiClient.callApi(
        '/notebooks/{id}/ipynbs', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getNotebookDetails operation.
     * @callback module:api/NotebookApi~getNotebookDetailsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/NotebookDetail} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns Jupyter notebook detail information 
     * @param {String} id Notebook container ID
     * @param {module:api/NotebookApi~getNotebookDetailsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/NotebookDetail}
     */
    this.getNotebookDetails = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getNotebookDetails");
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

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = NotebookDetail;

      return this.apiClient.callApi(
        '/notebooks/{id}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getNotebooks operation.
     * @callback module:api/NotebookApi~getNotebooksCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Notebook>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns Jupyter notebook information 
     * @param {module:api/NotebookApi~getNotebooksCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Notebook>}
     */
    this.getNotebooks = function(callback) {
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

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [Notebook];

      return this.apiClient.callApi(
        '/notebooks', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the modifyNotebook operation.
     * @callback module:api/NotebookApi~modifyNotebookCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * modify the notebook status 
     * @param {String} id Notebook container ID
     * @param {module:model/NotebookAttrs} body 
     * @param {module:api/NotebookApi~modifyNotebookCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.modifyNotebook = function(id, body, callback) {
      var postBody = body;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling modifyNotebook");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling modifyNotebook");
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

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/notebooks/{id}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the postNewNotebook operation.
     * @callback module:api/NotebookApi~postNewNotebookCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * creates Jupyter notebook container 
     * @param {module:model/ImageName2} body 
     * @param {module:api/NotebookApi~postNewNotebookCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.postNewNotebook = function(body, callback) {
      var postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling postNewNotebook");
      }


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

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/notebooks', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
