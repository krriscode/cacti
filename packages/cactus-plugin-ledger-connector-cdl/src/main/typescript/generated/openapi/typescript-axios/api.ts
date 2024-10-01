/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cacti Plugin - Connector CDL
 * Can perform basic tasks on Fujitsu CDL service.
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
 * 
 * @export
 * @interface AuthInfoAccessTokenV1
 */
export interface AuthInfoAccessTokenV1 {
    /**
     * 
     * @type {string}
     * @memberof AuthInfoAccessTokenV1
     */
    'accessToken': string;
    /**
     * 
     * @type {string}
     * @memberof AuthInfoAccessTokenV1
     */
    'trustAgentId': string;
}
/**
 * 
 * @export
 * @interface AuthInfoSubscriptionKeyV1
 */
export interface AuthInfoSubscriptionKeyV1 {
    /**
     * 
     * @type {string}
     * @memberof AuthInfoSubscriptionKeyV1
     */
    'subscriptionKey': string;
    /**
     * 
     * @type {string}
     * @memberof AuthInfoSubscriptionKeyV1
     */
    'trustAgentId': string;
    /**
     * 
     * @type {string}
     * @memberof AuthInfoSubscriptionKeyV1
     */
    'trustAgentRole': string;
    /**
     * 
     * @type {string}
     * @memberof AuthInfoSubscriptionKeyV1
     */
    'trustUserId': string;
    /**
     * 
     * @type {string}
     * @memberof AuthInfoSubscriptionKeyV1
     */
    'trustUserRole': string;
}
/**
 * @type AuthInfoV1
 * @export
 */
export type AuthInfoV1 = AuthInfoAccessTokenV1 | AuthInfoSubscriptionKeyV1;

/**
 * 
 * @export
 * @interface CDLCommonResponseV1
 */
export interface CDLCommonResponseV1 {
    /**
     * 
     * @type {any}
     * @memberof CDLCommonResponseV1
     */
    'detail'?: any;
    /**
     * 
     * @type { String}
     * @memberof CDLCommonResponseV1
     */
    'result':  String;
}
/**
 * Error response from the connector.
 * @export
 * @interface ErrorExceptionResponseV1
 */
export interface ErrorExceptionResponseV1 {
    /**
     * Short error description message.
     * @type {string}
     * @memberof ErrorExceptionResponseV1
     */
    'message': string;
    /**
     * Detailed error information.
     * @type {string}
     * @memberof ErrorExceptionResponseV1
     */
    'error': string;
}
/**
 * CDL event linage information (used to identify related events)
 * @export
 * @interface EventLineageV1
 */
export interface EventLineageV1 {
    /**
     * 
     * @type { String}
     * @memberof EventLineageV1
     */
    'cdl:EventId':  String;
    /**
     * 
     * @type { String}
     * @memberof EventLineageV1
     */
    'cdl:LineageId':  String;
    /**
     * 
     * @type { String}
     * @memberof EventLineageV1
     */
    'cdl:DataModelMode':  String;
    /**
     * 
     * @type { String}
     * @memberof EventLineageV1
     */
    'cdl:DataModelVersion':  String;
    /**
     * 
     * @type { String}
     * @memberof EventLineageV1
     */
    'cdl:DataRegistrationTimeStamp':  String;
    /**
     * 
     * @type {Array< String>}
     * @memberof EventLineageV1
     */
    'cdl:NextEventIdList': Array< String>;
    /**
     * 
     * @type {Array< String>}
     * @memberof EventLineageV1
     */
    'cdl:PreviousEventIdList': Array< String>;
}
/**
 * 
 * @export
 * @interface GatewayConfigurationV1
 */
export interface GatewayConfigurationV1 {
    /**
     * Gateway URL
     * @type {string}
     * @memberof GatewayConfigurationV1
     */
    'url': string;
    /**
     * Value of User-Agent header sent to CDL (to identify this client)
     * @type {string}
     * @memberof GatewayConfigurationV1
     */
    'userAgent'?: string;
    /**
     * Set to true to ignore self-signed and other rejected certificates
     * @type {boolean}
     * @memberof GatewayConfigurationV1
     */
    'skipCertCheck'?: boolean;
    /**
     * CA of CDL API gateway server in PEM format to use
     * @type {string}
     * @memberof GatewayConfigurationV1
     */
    'caPath'?: string;
    /**
     * Overwrite server name from cdlApiGateway.url to match one specified in CA
     * @type {string}
     * @memberof GatewayConfigurationV1
     */
    'serverName'?: string;
}
/**
 * 
 * @export
 * @enum {string}
 */

export const GetLineageOptionDirectionV1 = {
    Backward: 'backward',
    Forward: 'forward',
    Both: 'both'
} as const;

export type GetLineageOptionDirectionV1 = typeof GetLineageOptionDirectionV1[keyof typeof GetLineageOptionDirectionV1];


/**
 * 
 * @export
 * @interface GetLineageRequestV1
 */
export interface GetLineageRequestV1 {
    /**
     * 
     * @type {AuthInfoV1}
     * @memberof GetLineageRequestV1
     */
    'authInfo': AuthInfoV1;
    /**
     * 
     * @type {string}
     * @memberof GetLineageRequestV1
     */
    'eventId': string;
    /**
     * 
     * @type {GetLineageOptionDirectionV1}
     * @memberof GetLineageRequestV1
     */
    'direction'?: GetLineageOptionDirectionV1;
    /**
     * 
     * @type {string}
     * @memberof GetLineageRequestV1
     */
    'depth'?: string;
}


/**
 * 
 * @export
 * @interface GetLineageResponseV1
 */
export interface GetLineageResponseV1 {
    /**
     * 
     * @type {Array<TrailEventDetailsV1>}
     * @memberof GetLineageResponseV1
     */
    'detail': Array<TrailEventDetailsV1>;
    /**
     * 
     * @type { String}
     * @memberof GetLineageResponseV1
     */
    'result':  String;
}
/**
 * 
 * @export
 * @interface RegisterHistoryDataRequestV1
 */
export interface RegisterHistoryDataRequestV1 {
    /**
     * 
     * @type {AuthInfoV1}
     * @memberof RegisterHistoryDataRequestV1
     */
    'authInfo': AuthInfoV1;
    /**
     * 
     * @type {string}
     * @memberof RegisterHistoryDataRequestV1
     */
    'eventId'?: string;
    /**
     * 
     * @type {string}
     * @memberof RegisterHistoryDataRequestV1
     */
    'lineageId'?: string;
    /**
     * 
     * @type {any}
     * @memberof RegisterHistoryDataRequestV1
     */
    'tags'?: any;
    /**
     * 
     * @type {any}
     * @memberof RegisterHistoryDataRequestV1
     */
    'properties'?: any;
}
/**
 * 
 * @export
 * @interface RegisterHistoryDataV1Response
 */
export interface RegisterHistoryDataV1Response {
    /**
     * 
     * @type {TrailEventDetailsV1}
     * @memberof RegisterHistoryDataV1Response
     */
    'detail': TrailEventDetailsV1;
    /**
     * 
     * @type { String}
     * @memberof RegisterHistoryDataV1Response
     */
    'result':  String;
}
/**
 * 
 * @export
 * @interface SearchLineageRequestV1
 */
export interface SearchLineageRequestV1 {
    /**
     * 
     * @type {AuthInfoV1}
     * @memberof SearchLineageRequestV1
     */
    'authInfo': AuthInfoV1;
    /**
     * 
     * @type {SearchLineageTypeV1}
     * @memberof SearchLineageRequestV1
     */
    'searchType': SearchLineageTypeV1;
    /**
     * 
     * @type {any}
     * @memberof SearchLineageRequestV1
     */
    'fields': any;
}


/**
 * 
 * @export
 * @interface SearchLineageResponseV1
 */
export interface SearchLineageResponseV1 {
    /**
     * 
     * @type {Array<TrailEventDetailsV1>}
     * @memberof SearchLineageResponseV1
     */
    'detail': Array<TrailEventDetailsV1>;
    /**
     * 
     * @type { String}
     * @memberof SearchLineageResponseV1
     */
    'result':  String;
}
/**
 * 
 * @export
 * @enum {string}
 */

export const SearchLineageTypeV1 = {
    ExactMatch: 'exactmatch',
    PartialMatch: 'partialmatch',
    RegexMatch: 'regexpmatch'
} as const;

export type SearchLineageTypeV1 = typeof SearchLineageTypeV1[keyof typeof SearchLineageTypeV1];


/**
 * Details of newly created CDL event.
 * @export
 * @interface TrailEventDetailsV1
 */
export interface TrailEventDetailsV1 {
    /**
     * 
     * @type {any}
     * @memberof TrailEventDetailsV1
     */
    'cdl:Event'?: any;
    /**
     * 
     * @type {EventLineageV1}
     * @memberof TrailEventDetailsV1
     */
    'cdl:Lineage': EventLineageV1;
    /**
     * 
     * @type {any}
     * @memberof TrailEventDetailsV1
     */
    'cdl:Tags': any;
    /**
     * 
     * @type {any}
     * @memberof TrailEventDetailsV1
     */
    'cdl:Verification': any;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Get lineage trail from CDL.
         * @param {GetLineageRequestV1} [getLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getLineageV1: async (getLineageRequestV1?: GetLineageRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-cdl/get-lineage`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(getLineageRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Register new data trail on CDL
         * @param {RegisterHistoryDataRequestV1} [registerHistoryDataRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        registerHistoryDataV1: async (registerHistoryDataRequestV1?: RegisterHistoryDataRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-cdl/register-history-data`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(registerHistoryDataRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Search lineage using global data fields.
         * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        searchLineageByGlobalDataV1: async (searchLineageRequestV1?: SearchLineageRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-cdl/search-lineage-by-globaldata`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(searchLineageRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Search lineage using header fields.
         * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        searchLineageByHeaderV1: async (searchLineageRequestV1?: SearchLineageRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-cdl/search-lineage-by-header`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(searchLineageRequestV1, localVarRequestOptions, configuration)

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
         * @summary Get lineage trail from CDL.
         * @param {GetLineageRequestV1} [getLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getLineageV1(getLineageRequestV1?: GetLineageRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetLineageResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getLineageV1(getLineageRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Register new data trail on CDL
         * @param {RegisterHistoryDataRequestV1} [registerHistoryDataRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async registerHistoryDataV1(registerHistoryDataRequestV1?: RegisterHistoryDataRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RegisterHistoryDataV1Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.registerHistoryDataV1(registerHistoryDataRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Search lineage using global data fields.
         * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async searchLineageByGlobalDataV1(searchLineageRequestV1?: SearchLineageRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SearchLineageResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.searchLineageByGlobalDataV1(searchLineageRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Search lineage using header fields.
         * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async searchLineageByHeaderV1(searchLineageRequestV1?: SearchLineageRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SearchLineageResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.searchLineageByHeaderV1(searchLineageRequestV1, options);
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
         * @summary Get lineage trail from CDL.
         * @param {GetLineageRequestV1} [getLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getLineageV1(getLineageRequestV1?: GetLineageRequestV1, options?: any): AxiosPromise<GetLineageResponseV1> {
            return localVarFp.getLineageV1(getLineageRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Register new data trail on CDL
         * @param {RegisterHistoryDataRequestV1} [registerHistoryDataRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        registerHistoryDataV1(registerHistoryDataRequestV1?: RegisterHistoryDataRequestV1, options?: any): AxiosPromise<RegisterHistoryDataV1Response> {
            return localVarFp.registerHistoryDataV1(registerHistoryDataRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Search lineage using global data fields.
         * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        searchLineageByGlobalDataV1(searchLineageRequestV1?: SearchLineageRequestV1, options?: any): AxiosPromise<SearchLineageResponseV1> {
            return localVarFp.searchLineageByGlobalDataV1(searchLineageRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Search lineage using header fields.
         * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        searchLineageByHeaderV1(searchLineageRequestV1?: SearchLineageRequestV1, options?: any): AxiosPromise<SearchLineageResponseV1> {
            return localVarFp.searchLineageByHeaderV1(searchLineageRequestV1, options).then((request) => request(axios, basePath));
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
     * @summary Get lineage trail from CDL.
     * @param {GetLineageRequestV1} [getLineageRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getLineageV1(getLineageRequestV1?: GetLineageRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getLineageV1(getLineageRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Register new data trail on CDL
     * @param {RegisterHistoryDataRequestV1} [registerHistoryDataRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public registerHistoryDataV1(registerHistoryDataRequestV1?: RegisterHistoryDataRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).registerHistoryDataV1(registerHistoryDataRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Search lineage using global data fields.
     * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public searchLineageByGlobalDataV1(searchLineageRequestV1?: SearchLineageRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).searchLineageByGlobalDataV1(searchLineageRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Search lineage using header fields.
     * @param {SearchLineageRequestV1} [searchLineageRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public searchLineageByHeaderV1(searchLineageRequestV1?: SearchLineageRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).searchLineageByHeaderV1(searchLineageRequestV1, options).then((request) => request(this.axios, this.basePath));
    }
}


