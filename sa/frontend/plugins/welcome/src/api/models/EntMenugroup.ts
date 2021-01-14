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
    EntMenugroupEdges,
    EntMenugroupEdgesFromJSON,
    EntMenugroupEdgesFromJSONTyped,
    EntMenugroupEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMenugroup
 */
export interface EntMenugroup {
    /**
     * 
     * @type {EntMenugroupEdges}
     * @memberof EntMenugroup
     */
    edges?: EntMenugroupEdges;
    /**
     * Group holds the value of the "group" field.
     * @type {string}
     * @memberof EntMenugroup
     */
    group?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMenugroup
     */
    id?: number;
}

export function EntMenugroupFromJSON(json: any): EntMenugroup {
    return EntMenugroupFromJSONTyped(json, false);
}

export function EntMenugroupFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMenugroup {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntMenugroupEdgesFromJSON(json['edges']),
        'group': !exists(json, 'group') ? undefined : json['group'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntMenugroupToJSON(value?: EntMenugroup | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntMenugroupEdgesToJSON(value.edges),
        'group': value.group,
        'id': value.id,
    };
}


