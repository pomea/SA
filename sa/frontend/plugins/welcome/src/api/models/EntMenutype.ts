/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Menu
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntMenutypeEdges,
    EntMenutypeEdgesFromJSON,
    EntMenutypeEdgesFromJSONTyped,
    EntMenutypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMenutype
 */
export interface EntMenutype {
    /**
     * 
     * @type {EntMenutypeEdges}
     * @memberof EntMenutype
     */
    edges?: EntMenutypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMenutype
     */
    id?: number;
    /**
     * Type holds the value of the "type" field.
     * @type {string}
     * @memberof EntMenutype
     */
    type?: string;
}

export function EntMenutypeFromJSON(json: any): EntMenutype {
    return EntMenutypeFromJSONTyped(json, false);
}

export function EntMenutypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMenutype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntMenutypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function EntMenutypeToJSON(value?: EntMenutype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntMenutypeEdgesToJSON(value.edges),
        'id': value.id,
        'type': value.type,
    };
}


