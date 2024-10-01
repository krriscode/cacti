/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Example - Carbon Accounting App
 * Demonstrates how a business use case can be satisfied with Cactus when multiple distinct ledgers are involved.
 *
 * The version of the OpenAPI document: 2.0.0-rc.7
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError } from './base';

/**
 * Stores global constants related to the authorization of the application. Specifically enumerates the claims to validate for as per RFC 7519, section 4.1. See: https://tools.ietf.org/html/rfc7519#section-4.1
 * @export
 * @enum {string}
 */

export const AuthzJwtClaim = {
    /**
    * The &quot;iss&quot; (issuer) claim identifies the principal that issued the JWT.  The processing of this claim is generally application specific. The &quot;iss&quot; value is a case-sensitive string containing a StringOrURI value.  Use of this claim is OPTIONAL.
    */
    iss: 'Hyperledger Labs - Carbon Accounting Tool'
} as const;

export type AuthzJwtClaim = typeof AuthzJwtClaim[keyof typeof AuthzJwtClaim];


/**
 * 
 * @export
 * @enum {string}
 */

export const AuthzScope = {
    /**
    * Identities with the group:admin scope are administrators of the system.
    */
    GroupAdmin: 'group:admin',
    /**
    * Identities with the group:user scope are end users of the system who only have authorization to perform a limited set of actions.
    */
    GroupUser: 'group:user'
} as const;

export type AuthzScope = typeof AuthzScope[keyof typeof AuthzScope];


/**
 * 
 * @export
 * @interface Checkpoint
 */
export interface Checkpoint {
    /**
     * 
     * @type {number}
     * @memberof Checkpoint
     */
    'fromBlock': number;
    /**
     * 
     * @type {string}
     * @memberof Checkpoint
     */
    'votes': string;
}
/**
 * 
 * @export
 * @interface DaoTokenGetAllowanceNotFoundResponse
 */
export interface DaoTokenGetAllowanceNotFoundResponse {
    /**
     * Indicates whether the account was found or not.
     * @type {boolean}
     * @memberof DaoTokenGetAllowanceNotFoundResponse
     */
    'accountFound': boolean;
    /**
     * Indicates whether the spender was found or not.
     * @type {boolean}
     * @memberof DaoTokenGetAllowanceNotFoundResponse
     */
    'spenderFound': boolean;
}
/**
 * 
 * @export
 * @interface DaoTokenGetAllowanceRequest
 */
export interface DaoTokenGetAllowanceRequest {
    /**
     * The address of the account holding the funds
     * @type {string}
     * @memberof DaoTokenGetAllowanceRequest
     */
    'account': string;
    /**
     * The address of the account spending the funds
     * @type {string}
     * @memberof DaoTokenGetAllowanceRequest
     */
    'spender': string;
}
/**
 * 
 * @export
 * @interface DaoTokenGetAllowanceResponse
 */
export interface DaoTokenGetAllowanceResponse {
    /**
     * The number of tokens approved
     * @type {number}
     * @memberof DaoTokenGetAllowanceResponse
     */
    'allowance': number;
}
/**
 * 
 * @export
 * @enum {string}
 */

export const EnrollAdminInfo = {
    SUCCESSFULLY_ENROLLED_ADMIN: 'Successfully enrolled admin user and imported it into the wallet',
    ORG_ADMIN_REGISTERED: 'ORG ADMIN REGISTERED'
} as const;

export type EnrollAdminInfo = typeof EnrollAdminInfo[keyof typeof EnrollAdminInfo];


/**
 * 
 * @export
 * @interface EnrollAdminV1Request
 */
export interface EnrollAdminV1Request {
    /**
     * 
     * @type {string}
     * @memberof EnrollAdminV1Request
     */
    'orgName': string;
}
/**
 * 
 * @export
 * @interface EnrollAdminV1Response
 */
export interface EnrollAdminV1Response {
    /**
     * 
     * @type {string}
     * @memberof EnrollAdminV1Response
     */
    'info': string;
    /**
     * 
     * @type {string}
     * @memberof EnrollAdminV1Response
     */
    'orgName': string;
    /**
     * 
     * @type {string}
     * @memberof EnrollAdminV1Response
     */
    'msp': string;
    /**
     * 
     * @type {string}
     * @memberof EnrollAdminV1Response
     */
    'caName': string;
    /**
     * The key under which the identity created will be persisted to the keychain for later retrieval.
     * @type {string}
     * @memberof EnrollAdminV1Response
     */
    'identityId'?: string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Get the number of tokens `spender` is approved to spend on behalf of `account`
         * @param {DaoTokenGetAllowanceRequest} [daoTokenGetAllowanceRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        daoTokenGetAllowanceV1: async (daoTokenGetAllowanceRequest?: DaoTokenGetAllowanceRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-example-carbon-accounting-backend/dao-token/get-allowance`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(daoTokenGetAllowanceRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Registers an admin account within the Fabric organization specified.
         * @param {EnrollAdminV1Request} [enrollAdminV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        enrollAdminV1: async (enrollAdminV1Request?: EnrollAdminV1Request, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/utilityemissionchannel/registerEnroll/admin`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(enrollAdminV1Request, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Get the number of tokens `spender` is approved to spend on behalf of `account`
         * @param {DaoTokenGetAllowanceRequest} [daoTokenGetAllowanceRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async daoTokenGetAllowanceV1(daoTokenGetAllowanceRequest?: DaoTokenGetAllowanceRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<DaoTokenGetAllowanceResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.daoTokenGetAllowanceV1(daoTokenGetAllowanceRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Registers an admin account within the Fabric organization specified.
         * @param {EnrollAdminV1Request} [enrollAdminV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async enrollAdminV1(enrollAdminV1Request?: EnrollAdminV1Request, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<EnrollAdminV1Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.enrollAdminV1(enrollAdminV1Request, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @summary Get the number of tokens `spender` is approved to spend on behalf of `account`
         * @param {DaoTokenGetAllowanceRequest} [daoTokenGetAllowanceRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        daoTokenGetAllowanceV1(daoTokenGetAllowanceRequest?: DaoTokenGetAllowanceRequest, options?: any): AxiosPromise<DaoTokenGetAllowanceResponse> {
            return localVarFp.daoTokenGetAllowanceV1(daoTokenGetAllowanceRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Registers an admin account within the Fabric organization specified.
         * @param {EnrollAdminV1Request} [enrollAdminV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        enrollAdminV1(enrollAdminV1Request?: EnrollAdminV1Request, options?: any): AxiosPromise<EnrollAdminV1Response> {
            return localVarFp.enrollAdminV1(enrollAdminV1Request, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @summary Get the number of tokens `spender` is approved to spend on behalf of `account`
     * @param {DaoTokenGetAllowanceRequest} [daoTokenGetAllowanceRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public daoTokenGetAllowanceV1(daoTokenGetAllowanceRequest?: DaoTokenGetAllowanceRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).daoTokenGetAllowanceV1(daoTokenGetAllowanceRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Registers an admin account within the Fabric organization specified.
     * @param {EnrollAdminV1Request} [enrollAdminV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public enrollAdminV1(enrollAdminV1Request?: EnrollAdminV1Request, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).enrollAdminV1(enrollAdminV1Request, options).then((request) => request(this.axios, this.basePath));
    }
}


